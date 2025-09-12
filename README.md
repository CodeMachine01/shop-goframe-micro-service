启动etcd  
docker run -d --name etcd -p 2379:2379 -e ALLOW_NONE_AUTHENTICATION=yes quay.io/coreos/etcd:v3.5.0

启动mysql  
docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql:latest

安装微服务组件  
go get -u github.com/gogf/gf/contrib/rpc/grpcx/v2

安装grpc-go插件  
go install google.golang.org/protobuf/cmd/protoc-gen-go@l
atest  
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

安装数据库驱动  
go get -u github.com/gogf/gf/contrib/drivers/mysql/v2

安装etcd组件  
go get -u github.com/gogf/gf/contrib/registry/etcd/v2

微服务编写流程  
初始化文件 gf init app/admin -a   
修改hack/config.yaml配置文件  如dao和pbentity的配置
生成dao模型和pbentity模型  gf gen dao  gf gen pbentity
创建接口proto  创建目录manifest/protobuf/admin_info/v1/admin_info.proto  
生成控制器  gf gen pb
