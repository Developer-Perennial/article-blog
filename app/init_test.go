package app

import (
	"context"
	"github.com/DevPer/article-blog/internal/constants"
	"os"
	"testing"
	"time"
)

const (
	fileName = "config.yaml"
	fileData = `
env: dev

server_config:
  host: 0.0.0.0
  port: 9090

db_config:
  username: root
  password: 123456
  host: 0.0.0.0
  port: 3306
  dbname: blog_system
  max_idle_con: 2
  max_open_con: 5
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

	testRunCode := m.Run()
	os.Remove(f.Name())

	os.Exit(testRunCode)
}

func TestInit(t *testing.T){
	tests := []struct{
		name string
		give string
		want error
		setupFunc func()
		cleanupFunc func()
	}{
		{
			name: "file not present",
			give: "no-file.yml",
			want: nil,
		},
		{
			name: "happy path",
			give: fileName,
			want: nil,
			setupFunc: func() {
				_ = os.Setenv(constants.DB_CONFIG_HOST, "localhost")
				_ = os.Setenv(constants.DB_CONFIG_PORT, "3306")
			},
			cleanupFunc: func() {
				_ = os.Unsetenv(constants.DB_CONFIG_HOST)
				_ = os.Unsetenv(constants.DB_CONFIG_PORT)
			},
		},
	}
	for _, tt := range  tests{
		t.Run(tt.name, func(t *testing.T) {
			defer func(want error) {
				if r := recover(); r != nil {
					if tt.want != nil {
						t.Error("Got unexpected panic", r)
					}
				}
			}(tt.want)
			if tt.setupFunc != nil{
				tt.setupFunc()
			}
			_, err := Init(tt.give)
			if err != tt.want{
				t.Errorf("Want: %#v, Got: %#v", tt.want, err)
			}
			if tt.cleanupFunc != nil{
				tt.cleanupFunc()
			}
		})
	}
}

func TestApp(t *testing.T) {
	app, _ := Init(fileName)

	sigChan := make(chan struct{})
	time.AfterFunc(2 * time.Second, func() {
		sigChan <- struct{}{}
	})

	go func() {
		_ = app.Run(context.Background())
	}()
	<-sigChan

	err := app.ShutDown(nil)
	if err != nil{
		t.Errorf("Unexpected error")
	}
}