package config

type Config struct {
	Provider string // 暂时没用，准备兼容各个平台 [echo,aliyun,aws,azure]
	Server   struct {
		Host string
		Tls  bool
	}
	Device struct {
		ProductKey   string
		DeviceName   string
		DeviceSecret string
	}
}
