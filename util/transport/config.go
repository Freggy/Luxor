package transport

import "strconv"

type Config struct {
	Host      string
	Port      int
	User      string
	Password  string
	KeyPath   string
	CertPath  string
	UseTLS    bool
	ExtraData map[string]string
}

func (c Config) String() string {
	return "amqp://" + c.User + ":" + c.Password + "@" + c.Host + strconv.Itoa(c.Port) + "/"
}
