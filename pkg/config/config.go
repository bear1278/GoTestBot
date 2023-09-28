package config

import "github.com/spf13/viper"

type Config struct {
	TelegramToken string
	DbPtah        string
	Messages      Messages
}

type Messages struct {
	Errors
	Responses
}

type Responses struct {
	Start          string `mapstructure:"strat"`
	UnvalidCommand string `mapstructure:"unvalid_command"`
	TaskAdded      string `mapstructure:"task_add"`
	TaskDelete     string `mapstructure:"task_del"`
	TaskMark       string `mapstructure:"task_mark"`
	EnterTask      string `mapstructure:"task_enter"`
}

type Errors struct {
	IndenticalTask string `mapstructure:"IndenticalTask"`
	NoTasks        string `mapstructure:"NoTasks"`
	UnvalidText    string `mapstructure:"UnvalidText"`
	TaskNotFound   string `mapstructure:"TaskNotFound"`
	AlreadyMark    string `mapstructure:"AlreadyMark"`
	UnvalidError   string `mapstructure:"unvalid_err"`
}

func Init() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg Config

	err = viper.UnmarshalKey("messages.responses", &cfg.Messages.Responses)
	if err != nil {
		return nil, err
	}

	err = viper.UnmarshalKey("messages.errors", &cfg.Messages.Errors)
	if err != nil {
		return nil, err
	}

	if err = ParseEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func ParseEnv(cfg *Config) error {

	if err := viper.BindEnv("token"); err != nil {
		return err
	}

	cfg.TelegramToken = viper.GetString("token")

	if err := viper.BindEnv("mysql"); err != nil {
		return err
	}

	cfg.DbPtah = viper.GetString("mysql")
	return nil

}
