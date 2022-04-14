package client

type Config struct {
	Durable    bool `koanf:"durable"`
	AutoDelete bool `koanf:"auto_delete"`
	Exclusive  bool `koanf:"exclusive"`
	Wait       bool `koanf:"wait"`
	Mandatory  bool `koanf:"mandatory"`
	Local      bool `koanf:"local"`
	AutoAck    bool `koanf:"auto_ack"`
}
