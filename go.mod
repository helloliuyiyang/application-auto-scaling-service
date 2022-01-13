module nanto.io/application-auto-scaling-service

go 1.16

require (
	github.com/emicklei/go-restful v2.9.5+incompatible
	github.com/huaweicloud/huaweicloud-sdk-go-obs v3.21.8+incompatible
	github.com/huaweicloud/huaweicloud-sdk-go-v3 v0.0.74
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0
	github.com/prometheus/common v0.26.0
	github.com/robfig/cron/v3 v3.0.0
	github.com/sirupsen/logrus v1.8.1
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/ini.v1 v1.64.0
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/apimachinery v0.0.0-20211001003147-df63df3af3fc
	k8s.io/client-go v0.0.0-20211001003700-dbfa30b9d908
	k8s.io/code-generator v0.23.1
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20211001003357-dd4141958dfc
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20211001003147-df63df3af3fc
	k8s.io/client-go => k8s.io/client-go v0.0.0-20211001003700-dbfa30b9d908
)
