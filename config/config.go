package config

type Config struct {
	Provider string
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
