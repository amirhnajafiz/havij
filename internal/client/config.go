package client

type Config struct {
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	Wait       bool
	Mandatory  bool
	Immediate  bool
	Local      bool
}
