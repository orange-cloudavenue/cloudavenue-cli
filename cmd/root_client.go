package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/orange-cloudavenue/cloudavenue-sdk-go"
	clientcloudavenue "github.com/orange-cloudavenue/cloudavenue-sdk-go/pkg/clients/cloudavenue"
	"github.com/spf13/viper"

	"github.com/orange-cloudavenue/cloudavenue-cli/pkg/customerrors"
)

// function to initialize the configuration file
func initConfig() error {
	// Find the home directory
	home, err := homedir.Dir()
	if err != nil {
		return fmt.Errorf("home directory %s is: %w, %w", home, customerrors.ErrNoHomeDirectory, err)
	}
	// Set default file configuration
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(home + configPath)
	if home == "" {
		return fmt.Errorf("error home directory: %w", errors.New(customerrors.ErrNoHomeDirectory.Error()))
	}
	// Create configuration file if not exist
	if _, err = os.Stat(home + fileConfigPath); os.IsNotExist(err) {
		if err = os.MkdirAll(home+configPath, 0o755); err != nil {
			return fmt.Errorf("error canot create directory %s: %w, %w", home+configPath, customerrors.ErrNoHomeDirectory, err)
		}
		// Set default configuration
		cloudavenueConfig := cloudavenueConfig{}
		// set struct to viper
		v.Set("cloudavenue", cloudavenueConfig)

		// Write configuration file
		if err = v.SafeWriteConfig(); err != nil {
			return fmt.Errorf("error write config: %w", err)
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
		return fmt.Errorf("unable to read config file: %w, %w", customerrors.ErrConfigFile, err)
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
		return fmt.Errorf("unable to init cloudavenue client: %w, %w", customerrors.ErrClient, err)
	}
	return nil
}
