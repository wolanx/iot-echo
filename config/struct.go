package config

type Model struct {
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
