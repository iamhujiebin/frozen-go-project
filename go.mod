module frozen-go-project

go 1.14

require (
	github.com/Shopify/sarama v1.27.2
	github.com/aws/aws-sdk-go v1.35.10 // indirect
	github.com/bitly/go-simplejson v0.5.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/golang/protobuf v1.4.3
	github.com/golang/snappy v0.0.2 // indirect
	github.com/google/uuid v1.1.2
	github.com/jinzhu/copier v0.0.0-20190924061706-b57f9002281a
	github.com/klauspost/compress v1.11.1 // indirect
	github.com/tal-tech/go-zero v1.0.22
	go.mongodb.org/mongo-driver v1.4.2
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	golang.org/x/net v0.0.0-20201021035429-f5854403a974 // indirect
	golang.org/x/sync v0.0.0-20201008141435-b3e1573b7520 // indirect
	google.golang.org/grpc v1.33.0
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc v1.33.0 => google.golang.org/grpc v1.29.1 //grpc 和etcd 不兼容，需要指定版本才行
