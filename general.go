package kyocera_soap

type port uint32

const (
	PORT_INSECURE port = 9090
	PORT_SECURE   port = 9091
)

type Config struct {
	host     string
	username string
	password string
	port     port
}

func NewConfig(host string) *Config {
	return &Config{
		host:     host,
		username: "Admin",
		password: "Admin",
		port:     PORT_SECURE,
	}
}

func (c *Config) SetPort(port port) *Config {
	if port == PORT_SECURE || port == PORT_INSECURE {
		c.port = port
	}
	return c
}

func (c *Config) SetPassword(password string) *Config {
	c.password = password
	return c
}

func (c *Config) SetUsername(username string) *Config {
	c.username = username
	return c
}
