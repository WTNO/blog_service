package setting

import "github.com/spf13/viper"

type Setting struct {
	vp *viper.Viper
}

// 用于初始化本项目的配置的基础属性
func NewSetting() (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Setting{vp}, nil
}
