package storage

type ConfigDB struct {
	Dsn struct {
		Dsn string `json:"dsn"`
	} `json:"db"`
}

func NewConfigStorage() *ConfigDB {
	return &ConfigDB{
	}
}
