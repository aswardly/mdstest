package config

//AppConfig represents application main config
type AppConfig struct {
	Server		ServerConfig
	Database	DbConfig
}

//ServerConfig represents web server specific config
type ServerConfig struct {
	ListeningPort int	//port where application is listening
	ShutdownTimeout string //duration to wait for graceful shutdown
}

//DbConfig represents database related config
type DbConfig struct {
	DriverType	string		//database driver type
	DSN 		string		//datasource string for db connection
}
