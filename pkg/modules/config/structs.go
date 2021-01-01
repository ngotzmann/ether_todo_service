package config

type Service struct {
	Address 			 string
	Port                 string
	SessionSecret        string
	SessionName          string
	DefaultLang 		 string
}

type Database struct {
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
}

type Log struct {
	LogFile              string
	LogLevel             string
	LogTimestampFormat   string
}