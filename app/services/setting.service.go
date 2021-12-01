package services

import (
	//	"errors"
	"todo/core"
	//	"gorm.io/gorm"
)

type settingServiceLayer struct {
	config *core.Config
}

func newSettingServiceLayer(c *core.Config) *settingServiceLayer {
	return &settingServiceLayer{
		config: c,
	}
}

func (s *settingServiceLayer) GetHealth() core.Response {

	return core.Success(&map[string]interface{}{
		"status": "Up",
	}, core.String("Status checked"))
}
