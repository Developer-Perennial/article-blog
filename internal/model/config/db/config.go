package db

import (
	"fmt"
	"os"

	"github.com/DevPer/article-blog/internal/constants"
	"github.com/DevPer/article-blog/internal/util"
)

type Config struct {
	Username   string `yaml:"username" validate:"required"`
	Password   string `yaml:"password" validate:"required"`
	Host       string `yaml:"host" validate:"required"`
	Port       string `yaml:"port" validate:"required"`
	Database   string `yaml:"dbname" validate:"required"`
	MaxIdleCon *int   `yaml:"max_idle_con,omitempty"`
	MaxOpenCon *int   `yaml:"max_open_con,omitempty"`
	Debug      bool   `yaml:"debug"`

	Dsn string `yaml:"-" json:"-"`
}

func (d *Config) GenDsn() error {
	if err := d.Validate(); err != nil {
		return err
	}
	d.Dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", d.Username, d.Password, d.Host, d.Port, d.Database)
	return nil
}

func (d Config) Validate() error {
	return util.ValidateRequiredFields(d)
}

func (d *Config) InjectEnvValues() {
	if v := os.Getenv(constants.DB_CONFIG_HOST); v != "" {
		d.Host = v
	}
	if v := os.Getenv(constants.DB_CONFIG_PORT); v != "" {
		d.Port = v
	}
}
