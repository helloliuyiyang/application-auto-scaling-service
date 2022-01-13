package config

import (
	"log"

	"github.com/pkg/errors"
	"gopkg.in/ini.v1"
)

const (
	SceneMPC    = "mpc"
	SceneSearch = "search"
)

type Config struct {
	// CCE 集群元信息
	ClusterId          string `ini:"cluster_id"`
	SyncInstanceToVega bool   `ini:"sync_instance_to_vega"`

	LogConf LogConf `ini:"log"`
	// 运行场景，默认转码MPC
	Scene string `ini:"scene"`

	// 转码MPC场景
	ObsConf      ObsConf      `ini:"obs"`
	StrategyConf StrategyConf `ini:"strategy"`
	K8sConf      K8sConf      `ini:"k8s"`

	// 搜索场景
	HttpServerConf HttpServerConf `ini:"http_server"`
	AppGatewayConf AppGatewayConf `ini:"gateway"`
}

// HttpServerConf http服务端配置
type HttpServerConf struct {
	HttpAddr string `ini:"addr"`
	HttpPort int    `ini:"port"`
}

// AppGatewayConf gateway grpc server相关配置
type AppGatewayConf struct {
	GrpcHost string `ini:"host"`
	GrpcPort int    `ini:"port"`
}

// LogConf log相关配置
type LogConf struct {
	Level      string `ini:"level"`
	Path       string `ini:"path"`
	MaxSize    int    `ini:"max_size"`
	MaxBackups int    `ini:"max_backups"`
	MaxDays    int    `ini:"max_days"`
	Compress   bool   `ini:"compress"`
}

// ObsConf obs相关配置
type ObsConf struct {
	Endpoint                  string `ini:"endpoint"`
	BucketName                string `ini:"bucket_name"`
	SourceFileNodeIdsTemplate string `ini:"source_file_node_ids_template"`
	// nodeIds文件路径
	ObjectKeyNodeIdsTemplate string `ini:"object_key_node_ids_template"`
	// 伸缩策略文件路径
	ObjectKeyStrategiesTemplate    string `ini:"object_key_strategies_template"`
	SyncNodeIdsToOBSIntervalMinute int    `ini:"sync_node_ids_to_obs_interval_minute"`
}

// StrategyConf 扩缩策略相关配置
type StrategyConf struct {
	// 策略来源，enum："local"/"GTM"
	Source string `ini:"source"`
	// 本地策略文件路径，只有在 Source 为 "local" 时需要
	LocalPath string `ini:"local_path"`
}

// K8sConf k8s相关配置
type K8sConf struct {
	// kebeconfig文件路径，只有在k8s体外运行时需要
	Kubeconfig string `ini:"kubeconfig"`
}

// LoadConfig 加载配置文件
func LoadConfig(configFile string) (*Config, error) {
	config := GetDefaultConfig()
	if err := readConfig(configFile, config); err != nil {
		return nil, err
	}
	return config, nil
}

func readConfig(configFile string, config *Config) error {
	log.Printf("Reading config file: %s", configFile)
	conf, err := ini.LoadSources(ini.LoadOptions{
		SpaceBeforeInlineComment: true,
	}, configFile)
	if err != nil {
		return errors.Wrapf(err, "read conf file[%s] err", configFile)
	}
	if err = conf.MapTo(config); err != nil {
		return errors.Wrapf(err, "invalid config from file[%s]", configFile)
	}
	return nil
}

// GetDefaultConfig 获取默认配置
func GetDefaultConfig() *Config {
	return &Config{
		ClusterId: "default_cluster_id",
		//ClusterName: "default_cluster_name",
		LogConf: LogConf{
			Level:      "info",
			Path:       "/opt/cloud/logs/application-auto-scaling-service/application-auto-scaling-service.log",
			MaxSize:    20,
			MaxBackups: 50,
			MaxDays:    90,
			Compress:   true,
		},
	}
}
