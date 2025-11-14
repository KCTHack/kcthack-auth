package config

type Config struct {
	App struct {
		Port string
	}
}

func Init() (*Config, error) {
	var cfg Config
	loadFromEnv()

	return &cfg, nil
}

func loadFromEnv() {

}
