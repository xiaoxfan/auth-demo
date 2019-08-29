package config

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T)  {
	config := loadConfig()
	fmt.Println(config)
	fmt.Println(Config.DBUrl())
}