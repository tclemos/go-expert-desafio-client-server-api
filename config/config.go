package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

type ClientConfig struct {
	DolarProvider DolarProviderConfig `json:"dolar_provider"`
	Output        string              `json:"output"`
}

type ServerConfig struct {
	Host          string              `json:"host"`
	Port          int                 `json:"port"`
	DolarProvider DolarProviderConfig `json:"dolar_provider"`
	DB            SQLiteConfig        `json:"database"`
}

type DolarProviderConfig struct {
	Endpoint     string   `json:"endpoint"`
	ReadTimeout  Duration `json:"read_timeout"`
	WriteTimeout Duration `json:"write_timeout"`
}

type SQLiteConfig struct {
	Path string `json:"path"`
}

func MustLoadClientConfig(path string) ClientConfig {
	var cfg ClientConfig
	err := loadJsonFile(path, &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func MustLoadServerConfig(path string) ServerConfig {
	var cfg ServerConfig
	err := loadJsonFile(path, &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

func loadJsonFile[T any](path string, instance *T) error {
	fmt.Printf("Loading config from: %v\n", path)
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	fileContent, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(fileContent, &instance); err != nil {
		return err
	}
	fmt.Printf("Config loaded: %v\n", string(fileContent))

	return nil
}

type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Duration(d).String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		*d = Duration(time.Duration(value))
		return nil
	case string:
		tmp, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		*d = Duration(tmp)
		return nil
	default:
		return errors.New("invalid duration")
	}
}
