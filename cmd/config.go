package cmd

import (
	"github.com/mitchellh/mapstructure"
	alice "github.com/ohhfishal/alice/api/v1"
	"github.com/spf13/viper"
)

func NewConfig(vConfig *viper.Viper) (*alice.Config, error) {
	var config alice.Config
	err := vConfig.Unmarshal(&config, UseJson)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func UseJson(c *mapstructure.DecoderConfig) {
	c.TagName = "json"
}
