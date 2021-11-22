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

const (
	fileNameNonYaml = "config-non-yaml.txt"
	fileDataNonYaml = `This is a text file`
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

	// create temporary file
	f1, err := os.Create(fileNameNonYaml)
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	// write data to file
	f1.Write([]byte(fileDataNonYaml))

	testRunCode := m.Run()
	os.Remove(f.Name())
	os.Remove(f1.Name())

	os.Exit(testRunCode)
}

func TestLoadConfigFromFile(t *testing.T) {
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
			name: "non-yaml file",
			give: fileNameNonYaml,
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
