package server

// Config structure to hold configuration settings.
type Config struct {
	Port            string `yaml:"port"`
	Region          string `yaml:"region"`
	AccessKeyID     string `yaml:"access_key_id"`
	SecretAccessKey string `yaml:"secret_access_key"`
	ModelID         string `yaml:"model_id"`
}
