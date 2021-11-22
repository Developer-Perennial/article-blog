package config

import (
	"github.com/DevPer/article-blog/internal/model/config/server"
	"os"
	"reflect"
	"testing"
)

const (
	fileName = "config.yaml"
	fileData = `
env: dev

server_config:
  host: 0.0.0.0
  port: 8080
`
)

func TestMain(m *testing.M) {
	// create temporary file
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// write data to file
	f.Write([]byte(fileData))

	os.Exit(m.Run())
}

func TestLoadConfigFromFile(t *testing.T) {
	defer func() {
		os.Remove(fileName)
	}()
	tests := []struct {
		name string
		give string
		want *Config
	}{
		{
			name: "file not present",
			give: "no-file.yaml",
			want: nil,
		},
		{
			name: "happy path",
			give: fileName,
			want: &Config{
				Env: "dev",
				ServerConfig: &server.Config{
					Host: "0.0.0.0",
					Port: "8080",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func(want *Config) {
				if r := recover(); r != nil {
					if tt.want != nil {
						t.Error("Got unexpected panic", r)
					}
				}
			}(tt.want)
			r := LoadConfigFromFile(tt.give)
			if !reflect.DeepEqual(r, tt.want) {
				t.Errorf("Want: %#v, Got: %#v", tt.want, r)
			}
		})
	}
}
