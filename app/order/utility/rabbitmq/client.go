package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	amqp "github.com/rabbitmq/amqp091-go"
)

// OrderRabbitMQClient 订单服务专用的RabbitMQ客户端
type OrderRabbitMQClient struct {
	conn    *amqp.Connection //连接对象
	channel *amqp.Channel    //通道对象
	mutex   sync.Mutex       //保证并发操作线程安全
}

var (
	instance *OrderRabbitMQClient //单例实例
	once     sync.Once            //确保初始化只执行一次
)

// GetOrderRabbitMQClient 获取订单服务专用的RabbitMQ客户端单例
func GetOrderRabbitMQClient(ctx context.Context) (*OrderRabbitMQClient, error) {
	var err error
	once.Do(func() { //保证只有一个goruntine执行，后续调用Do的goroutint会阻塞，知道首次执行完成
		instance, err = newOrderRabbitMQClient(ctx)
	})

	if err != nil {
		return nil, err
	}

	// 检查连接是否已关闭，如果已关闭则重新创建
	if instance.conn.IsClosed() {
		instance.mutex.Lock()
		defer instance.mutex.Unlock()

		// 双重检查，防止重新创建后，其他goroutine得到锁进入此逻辑重复创建
		if instance.conn.IsClosed() {
			instance, err = newOrderRabbitMQClient(ctx)
			if err != nil {
				return nil, fmt.Errorf("重新创建RabbitMQ连接失败: %v", err)
			}
		}
	}

	return instance, nil
}

// newOrderRabbitMQClient 创建新的RabbitMQ客户端实例（内部方法）
func newOrderRabbitMQClient(ctx context.Context) (*OrderRabbitMQClient, error) {
	// 从配置获取RabbitMQ连接地址
	address := g.Cfg().MustGet(ctx, "rabbitmq.default.address").String()
	if address == "" {
		return nil, fmt.Errorf("RabbitMQ连接地址不能为空，请检查配置文件")
	}

	g.Log().Info(ctx, "订单服务正在连接RabbitMQ服务器: "+address)

	// 建立连接
	conn, err := amqp.Dial(address)
	if err != nil {
		return nil, fmt.Errorf("连接RabbitMQ失败: %v", err)
	}

	// 创建通道
	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("创建RabbitMQ通道失败: %v", err)
	}

	// 设置QoS（服务质量），空值消费者预取消息数量
	prefetchCount := g.Cfg().MustGet(ctx, "rabbitmq.default.consumer.prefetchCount").Int()
	if prefetchCount <= 0 {
		prefetchCount = 1
	}
	err = channel.Qos(
		prefetchCount, // 预取计数，消费者最多可以预取多少条未确认的消息
		0,             // 预取大小 ，0表示不限制
		false,         // 是否全局设置，如果为false则消费者之间独立受限服务质量，如果为true则所有消费者共享服务质量
	)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("设置QoS失败: %v", err)
	}

	client := &OrderRabbitMQClient{
		conn:    conn,
		channel: channel,
	}

	// 初始化交换机和队列
	err = client.initExchangeAndQueue(ctx)
	if err != nil {
		client.Close()
		return nil, err
	}

	g.Log().Info(ctx, "订单服务RabbitMQ客户端初始化成功")
	return client, nil
}

// Close 关闭RabbitMQ连接
func (r *OrderRabbitMQClient) Close() error {
	var err error
	if r.channel != nil {
		err = r.channel.Close()
		if err != nil {
			g.Log().Error(context.Background(), "关闭RabbitMQ通道失败: "+err.Error())
		}
	}
	if r.conn != nil {
		if closeErr := r.conn.Close(); closeErr != nil && err == nil {
			err = closeErr
			g.Log().Error(context.Background(), "关闭RabbitMQ连接失败: "+closeErr.Error())
		}
	}
	return err
}

// initExchangeAndQueue 初始化延迟交换机和队列
func (r *OrderRabbitMQClient) initExchangeAndQueue(ctx context.Context) error {
	// 获取交换机名称
	exchangeName := g.Cfg().MustGet(ctx, "rabbitmq.default.exchange.orderDelayExchange").String()
	if exchangeName == "" {
		return fmt.Errorf("交换机名称不能为空，请检查配置文件")
	}

	g.Log().Info(ctx, "正在创建延迟交换机: "+exchangeName)

	// 声明延迟交换机
	err := r.channel.ExchangeDeclare(
		exchangeName,        // 交换机名称
		"x-delayed-message", // 交换机类型
		true,                // 持久化
		false,               // 自动删除
		false,               // 内部使用
		false,               // 不等待服务器响应
		amqp.Table{
			"x-delayed-type": "direct", // 指定延迟交换机的底层类型
		},
	)
	if err != nil {
		return fmt.Errorf("创建延迟交换机失败: %v", err)
	}

	// 获取队列名称
	queueName := g.Cfg().MustGet(ctx, "rabbitmq.default.queue.orderTimeoutQueue").String()
	if queueName == "" {
		return fmt.Errorf("队列名称不能为空，请检查配置文件")
	}

	g.Log().Info(ctx, "正在创建队列: "+queueName)

	// 声明队列
	_, err = r.channel.QueueDeclare(
		queueName, // 队列名称
		true,      // 持久化
		false,     // 自动删除
		false,     // 独占
		false,     // 不等待服务器响应
		nil,       // 参数
	)
	if err != nil {
		return fmt.Errorf("创建队列失败: %v", err)
	}

	g.Log().Info(ctx, "正在绑定队列到交换机")

	// 绑定队列到交换机
	err = r.channel.QueueBind(
		queueName,       // 队列名称
		"order.timeout", // 路由键
		exchangeName,    // 交换机名称
		false,           // 不等待服务器响应
		nil,             // 参数
	)
	if err != nil {
		return fmt.Errorf("绑定队列到交换机失败: %v", err)
	}

	g.Log().Info(ctx, "RabbitMQ交换机和队列初始化完成")
	return nil
}

// SendOrderTimeoutMessage 发送订单超时消息
func (r *OrderRabbitMQClient) SendOrderTimeoutMessage(ctx context.Context, orderId int32, delayMs int) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// 获取交换机名称
	exchangeName := g.Cfg().MustGet(ctx, "rabbitmq.default.exchange.orderDelayExchange").String()
	if exchangeName == "" {
		return fmt.Errorf("交换机名称不能为空，请检查配置文件")
	}

	//消息体构造
	message := map[string]interface{}{
		"order_id":  orderId,
		"type":      "order_timeout",
		"timestamp": time.Now().Format(time.RFC3339),
	}

	body, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("序列化消息失败: %v", err)
	}

	//消息发布
	err = r.channel.Publish(
		exchangeName,    // 交换机
		"order.timeout", // 路由键
		false,           // 不强制路由
		false,           // 不立即返回
		amqp.Publishing{ //消息内容 amqp.Publishing类型
			ContentType: "application/json",
			Body:        body, //实际内容 字节切片类型
			Headers: amqp.Table{
				"x-delay": delayMs, // 延迟时间，单位毫秒
			},
			DeliveryMode: amqp.Persistent, // 持久化消息
		})
	if err != nil {
		return fmt.Errorf("发送消息失败: %v", err)
	}

	g.Log().Infof(ctx, "订单超时消息发送成功, 订单ID: %d, 延迟: %d毫秒", orderId, delayMs)
	return nil
}

// GetOrderTimeoutDelay 从配置获取订单超时延迟时间
func GetOrderTimeoutDelay(ctx context.Context) int {
	// 从配置获取订单超时时间，默认30分钟
	timeout := g.Cfg().MustGet(ctx, "business.orderTimeout").Int()
	if timeout <= 0 {
		timeout = 30 * 60 * 1000 // 30分钟
	}
	return timeout
}

// SendOrderTimeoutMessageStatic 静态方法发送订单超时消息
func SendOrderTimeoutMessageStatic(ctx context.Context, orderId int32, delayMs int) error {
	client, err := GetOrderRabbitMQClient(ctx)
	if err != nil {
		return err
	}

	return client.SendOrderTimeoutMessage(ctx, orderId, delayMs)
}
