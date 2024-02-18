package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/customErrors"
	"github.com/orange-cloudavenue/cloudavenue-sdk-go"
	clientcloudavenue "github.com/orange-cloudavenue/cloudavenue-sdk-go/pkg/clients/cloudavenue"
	"github.com/spf13/viper"
)

// function to initialize the configuration file
func initConfig() error {
	// Find the home directory
	home, err := homedir.Dir()
	if err != nil {
		return fmt.Errorf("Home Directory %s is: %w, %w", home, customErrors.ErrNoHomeDirectory, err)
	}
	// Set default file configuration
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(home + configPath)
	if home == "" {
		return fmt.Errorf("Error Home Directory: %w", errors.New(customErrors.ErrNoHomeDirectory.Error()))
	}
	// Create configuration file if not exist
	if _, err = os.Stat(home + fileConfigPath); os.IsNotExist(err) {
		if err = os.MkdirAll(home+configPath, 0755); err != nil {
			return fmt.Errorf("Error canot create directory %s: %w, %w", home+configPath, customErrors.ErrNoHomeDirectory, err)
		}
		// Set default configuration
		cloudavenueConfig := cloudavenueConfig{}
		// set struct to viper
		v.Set("cloudavenue", cloudavenueConfig)

		// Write configuration file
		if err = v.SafeWriteConfig(); err != nil {
			return fmt.Errorf("Error write config: %w", err)
		}
		s.FinalMSG = `
					***
					Configuration file is created in your Home Directory under .cav/config.yaml
					Please fill it with your credentials and re-run the command.
					***`
		s.Stop()
		os.Exit(0)
	}
	// Read configuration file
	err = v.ReadInConfig()
	if err != nil {
		return fmt.Errorf("Unable to read config file: %w, %w", customErrors.ErrConfigFile, err)
	}
	return initClient(v)
}

func initClient(v *viper.Viper) (err error) {
	// Set client CloudAvenue
	c, err = cloudavenue.New(&cloudavenue.ClientOpts{
		CloudAvenue: &clientcloudavenue.Opts{
			Username: v.GetString("cloudavenue.username"),
			Password: v.GetString("cloudavenue.password"),
			Org:      v.GetString("cloudavenue.org"),
			URL:      v.GetString("cloudavenue.url"),
			Debug:    v.GetBool("cloudavenue.debug"),
		},
	})
	if err != nil {
		return fmt.Errorf("Unable to init cloudavenue client: %w, %w", customErrors.ErrClient, err)
	}
	return nil
}
