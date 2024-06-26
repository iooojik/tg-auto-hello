package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/iooojik/tg-auto-hello/internal/bot"
)

type Config struct {
	BotConfig *bot.Config `yaml:"bot"`
}

func ReadCfg(p string) (*Config, error) {
	content, err := os.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("read config from %s: %w", p, err)
	}

	cfg := &Config{}

	err = yaml.Unmarshal(content, cfg)
	if err != nil {
		return nil, fmt.Errorf("unmarshal config %s: %w", p, err)
	}

	return cfg, nil
}
