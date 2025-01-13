package config

import (
	"github.com/spf13/viper"
	"strings"
)

var (
	DbUser     = "postgres"
	DbPassword = "postgres"
	DbName     = "maker_checker"
	DbHost     = "localhost"
	DbPort     = "5432"

	MigrationsLocation = "file://store/psql/migrations"
	SkipMigrations     = false // Skip DB migrations at the service launch

	LogLevel = "info"

	Profile = "prod"
	Port    = "8080"
)

func setDBDefaults(v *viper.Viper) {
	v.SetDefault("DB-USER", DbUser)
	v.SetDefault("DB-PASSWORD", DbPassword)
	v.SetDefault("DB-NAME", DbName)
	v.SetDefault("DB-HOST", DbHost)
	v.SetDefault("DB-PORT", DbPort)
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("migrations-location", MigrationsLocation)
	v.SetDefault("skip-migrations", SkipMigrations)
	v.SetDefault("profile", Profile)
	v.SetDefault("port", Port)
	v.SetDefault("log-level", LogLevel)
	setDBDefaults(v)
}

func addConfigFile(v *viper.Viper, profile string) error {
	if profile == "prod" {
		return nil
	}

	v.SetConfigName(profile)
	v.SetConfigType("yaml")
	v.AddConfigPath("./cmd/maker-checker/config")
	return v.ReadInConfig()
}

func setDBEnvVars(v *viper.Viper) error {
	envs := []string{
		"DB-USER",
		"DB-PASSWORD",
		"DB-NAME",
		"DB-HOST",
		"DB-PORT",
	}

	var err error
	for _, env := range envs {
		if err = v.BindEnv(env); err != nil {
			return err
		}
	}

	return nil
}

func addEnvVars(v *viper.Viper) error {
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	envs := []string{
		"migrations-location",
		"skip-migrations",
		"port",
	}

	var err error
	for _, env := range envs {
		if err = v.BindEnv(env); err != nil {
			return err
		}
	}

	if err := setDBEnvVars(v); err != nil {
		return err
	}

	v.AutomaticEnv()
	return err
}

func Init(v *viper.Viper) error {
	setDefaults(v)
	if err := addEnvVars(v); err != nil {
		return err
	}
	return nil
}

func Load(v *viper.Viper) error {
	tempProfile := v.GetString("profile")
	switch tempProfile {
	case "dev", "prod":
		Profile = tempProfile
	}

	if err := addConfigFile(v, Profile); err != nil {
		return err
	}

	DbUser = v.GetString("DB-USER")
	DbPassword = v.GetString("DB-PASSWORD")
	DbName = v.GetString("DB-NAME")
	DbHost = v.GetString("DB-HOST")
	DbPort = v.GetString("DB-PORT")

	LogLevel = v.GetString("LOG-LEVEL")
	MigrationsLocation = v.GetString("migrations-location")
	SkipMigrations = v.GetBool("skip-migrations")

	Port = v.GetString("PORT")

	return nil
}
