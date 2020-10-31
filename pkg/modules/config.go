package modules

//TODO: Read config from external source

func DefaultConfig() *Config {
	return &Config{
		GorrorFilePath:       "config/",
		Port:                 "8080",
		SessionSecret:        "S3cReT",
		SessionName:          "defaultSession",
		DBDialect:            "postgres",
		DBAddress:            "localhost",
		DBPort:               "5432",
		Database:             "test",
		DBUser:               "test",
		DBPassword:           "t3sT",
		DBMaxIdleConnections: 10,
		DBLogging:            true,
		DBAudit:              false,
		DBShouldLog: 		  false,
		DBSSLMode: 			  "disable",
		LogFile:              "../service.log",
		LogLevel:             "info",
		LogTimestampFormat:   "02-01-2006_15:04:05",
	}
}

func TestConfig() *Config {
	return &Config{
		GorrorFilePath:       "../../../config/",
		Port:                 "8080",
		SessionSecret:        "S3cReT",
		SessionName:          "defaultSession",
		DBDialect:            "postgres",
		DBAddress:            "localhost",
		DBPort:               "5432",
		Database:             "test",
		DBUser:               "test",
		DBPassword:           "t3sT",
		DBMaxIdleConnections: 10,
		DBLogging:            true,
		DBAudit:              false,
		DBShouldLog: 		  false,
		DBSSLMode: 			  "disable",
		LogFile:              "../service.log",
		LogLevel:             "info",
		LogTimestampFormat:   "02-01-2006_15:04:05",
	}
}

type Config struct {
	GorrorFilePath 		 string
	Port                 string
	SessionSecret        string
	SessionName          string
	DBDialect            string
	DBAddress            string
	DBPort               string
	Database             string
	DBUser               string
	DBPassword           string
	DBMaxIdleConnections int
	DBLogging            bool
	DBAudit              bool
	DBShouldLog 		 bool
	DBSSLMode 			 string
	LogFile              string
	LogLevel             string
	LogTimestampFormat   string
}
