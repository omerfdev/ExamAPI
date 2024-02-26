package configs

var ac AppConfig

type AppConfig struct {
	URI     string
	Database         string
	Collection string
}

func SetConfigs(env string) {
	ac = AppConfig{
		URI:     Configs[env].URI,
		Database:         Configs[env].Database,
		Collection: Configs[env].Collection,
	}
}

func GetConfigs() *AppConfig {
	return &ac
}
