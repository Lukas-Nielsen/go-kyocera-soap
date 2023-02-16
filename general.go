package kyocera_soap

const (
	PORT_INSECURE uint32 = 9090
	PORT_SECURE   uint32 = 9091
)

type Config struct {
	host     string
	username string
	password string
	port     uint32
	debug    bool
}

func NewConfig(host string) *Config {
	return &Config{
		host:     host,
		username: "Admin",
		password: "Admin",
		port:     PORT_SECURE,
		debug:    false,
	}
}

func (c *Config) SetPort(port uint32) *Config {
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

func (c *Config) SetDebug() *Config {
	c.debug = true
	return c
}
