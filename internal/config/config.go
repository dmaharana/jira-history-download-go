package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type Config struct {
	URL      string
	Username string
	Token    string
	JQL      string
}

func Load(filename string) (*Config, error) {
	cfg, err := ini.Load(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to load config file: %v", err)
	}

	jiraSec := cfg.Section("jira")
	return &Config{
		URL:      jiraSec.Key("url").String(),
		Username: jiraSec.Key("username").String(),
		Token:    jiraSec.Key("token").String(),
		JQL:      jiraSec.Key("jql").String(),
	}, nil
}
