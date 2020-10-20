module frozen-go-project

go 1.14

require (
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2
	github.com/jinzhu/copier v0.0.0-20190924061706-b57f9002281a
	github.com/tal-tech/go-zero v1.0.21
	google.golang.org/grpc v1.33.0
	google.golang.org/protobuf v1.25.0
)

replace (
	google.golang.org/grpc v1.33.0 => google.golang.org/grpc v1.29.1 //grpc 和etcd 不兼容，需要指定版本才行
)