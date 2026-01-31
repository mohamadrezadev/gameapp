package config

import (
	"strings"

	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func Load(configPath string) Config {
	var k = koanf.New(".")
	
	k.Load(confmap.Provider(defaultConfig,"."), nil)

	k.Load(file.Provider(configPath),yaml.Parser())

	k.Load(env.Provider("GAMAPP_",".",func (s string) string {
		return strings.Replace(strings.ToLower(strings.TrimPrefix(s,"GAMAPP_")),"_",".",-1)
	}),nil)

	var cfg Config
	if err:=k.Unmarshal("",&cfg);err !=nil{
		panic(err)
	}
	return cfg
}