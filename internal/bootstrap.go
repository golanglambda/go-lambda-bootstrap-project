package internal

func BootstrapHandler(cfg *Config) (*Handler, error) {
	//Load secure parameters
	//Create database connections

	h := NewHandler(cfg)
	return h, nil
}
