package config 

import (
  "os"
  "flag"
  "log"
  "github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
  Env string `yaml:"env" env-required:"true" env-default:"production"`
  StoragePath string `yaml:"storage" env-required:"true"`
  HttpServer `yaml:"http_server"`
}

type HttpServer struct  {
  Address string `yaml:"address" env-required:"true" env-default:"localhost:3000"`
}

func MustLoad() *Config {
  configPath := os.Getenv("CONFIG_PATH")

  if configPath == "" {
    cfg := flag.String("config", "", "path to config file")
    flag.Parse()

    configPath = *cfg

    if configPath == "" {
      log.Fatal("config path is not set")
    }

  }

  if _, err := os.Stat(configPath); os.IsNotExist(err) {
    log.Fatalf("config file does not exist: %s", configPath)
  }

  var cfg Config 

  err := cleanenv.ReadConfig(configPath, &cfg)
  if err != nil {
    log.Fatalf("cannot read config: %s", err)
  }

  return &cfg
}

