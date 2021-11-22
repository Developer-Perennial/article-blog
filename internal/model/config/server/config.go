package server

import "github.com/DevPer/article-blog/internal/util"

type Config struct {
	Host string `yaml:"host" validate:"required"`
	Port string `yaml:"port" validate:"required"`
}

func (d Config) Validate() error {
	return util.ValidateRequiredFields(d)
}

func (d *Config) InjectEnvValues() {
}
