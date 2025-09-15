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

安装Elasticsearch  
docker run -d --name elasticsearch `
-p 9200:9200 -p 9300:9300 `
-e "discovery.type=single-node" `
-e "xpack.security.enabled=false" `
-e "ES_JAVA_OPTS=-Xms512m -Xmx512m" `
-v es-data:/usr/share/elasticsearch/data `
docker.elastic.co/elasticsearch/elasticsearch:8.11.0  
安装Kibana  
docker run -d --name kibana ^
-p 5601:5601 ^
-e "ELASTICSEARCH_HOSTS=http://elasticsearch:9200" ^
-e "SERVER_NAME=kibana" ^
-e "XPACK_SECURITY_ENABLED=false" ^
-e "XPACK_MONITORING_ENABLED=false" ^
--link elasticsearch:elasticsearch ^
docker.elastic.co/kibana/kibana:8.11.0  
IK分词器下载elasticsearch-analysis-ik-8.11.0.zip    
https://release.infinilabs.com/analysis-ik/stable/  
docker cp D:\MyPlugins\elasticsearch-analysis-ik-8.11.0.zip elasticsearch:/tmp/ik.zip  
安装IK分词器  
# 进入容器内部
docker exec -it elasticsearch /bin/bash

# 在容器内部执行安装命令（针对复制的ZIP文件）
./bin/elasticsearch-plugin install file:///tmp/ik.zip

# 安装完成后，退出容器
exit  

# 重启elasticsearch  
docker restart elasticsearch

微服务编写流程  
初始化文件 gf init app/admin -a   
修改hack/config.yaml配置文件  如dao和pbentity的配置
生成dao模型和pbentity模型  gf gen dao  gf gen pbentity
创建接口proto  创建目录manifest/protobuf/admin_info/v1/admin_info.proto  
生成控制器  gf gen pb
