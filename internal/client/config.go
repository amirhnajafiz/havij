package client

type Config struct {
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	Wait       bool
	Mandatory  bool
	Local      bool
	AutoAck    bool
}
