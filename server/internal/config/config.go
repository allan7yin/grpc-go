package config

type ServerConfig struct {
	Host         string
	Port         string
	GrpcProtocol string
	GrpcPort     string
}

type config struct {
	serverConfig *ServerConfig
}

// we define this interface, viper will implement this for us
type Config interface {
	GetConfig() *config
	GetServerConfig() *ServerConfig
}
