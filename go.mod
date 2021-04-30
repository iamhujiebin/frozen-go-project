module frozen-go-project

go 1.14

require (
	cloud.google.com/go/firestore v1.5.0 // indirect
	firebase.google.com/go v3.13.0+incompatible // indirect
	github.com/BurntSushi/toml v0.3.1
	github.com/Shopify/sarama v1.27.2
	github.com/alicebob/miniredis v2.5.0+incompatible // indirect
	github.com/antlr/antlr4 v0.0.0-20210216171041-fde0b28dfbd6 // indirect
	github.com/aws/aws-sdk-go v1.35.10 // indirect
	github.com/bitly/go-simplejson v0.5.0
	github.com/bluele/gcache v0.0.2 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/dchest/siphash v1.2.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/extrame/ole2 v0.0.0-20160812065207-d69429661ad7 // indirect
	github.com/extrame/xls v0.0.1 // indirect
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.6.3
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/gogf/gf v1.15.5 // indirect
	github.com/golang/protobuf v1.4.3
	github.com/golang/snappy v0.0.2 // indirect
	github.com/google/gops v0.3.7 // indirect
	github.com/google/uuid v1.1.2
	github.com/iancoleman/strcase v0.1.3 // indirect
	github.com/jinzhu/copier v0.0.0-20190924061706-b57f9002281a
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/klauspost/compress v1.11.1 // indirect
	github.com/nicksnyder/go-i18n v1.10.1 // indirect
	github.com/nicksnyder/go-i18n/v2 v2.1.2
	github.com/ouqiang/timewheel v1.0.1 // indirect
	github.com/prometheus/common v0.9.1
	github.com/stretchr/testify v1.7.0
	github.com/tal-tech/go-zero v1.1.6
	github.com/tealeg/xlsx v1.0.5 // indirect
	go.mongodb.org/mongo-driver v1.4.2
	go.uber.org/automaxprocs v1.4.0 // indirect
	golang.org/x/blog v0.0.0-20210219171517-8bdb56a492da // indirect
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b
	golang.org/x/oauth2 v0.0.0-20210218202405-ba52d332ba99
	golang.org/x/text v0.3.4
	google.golang.org/grpc v1.35.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.5
)

replace google.golang.org/grpc v1.33.0 => google.golang.org/grpc v1.29.1 //grpc 和etcd 不兼容，需要指定版本才行
