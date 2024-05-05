package cmd

import (
	"fmt"
	"github.com/actvehigh/go-cli-project-template/configs"
	"github.com/actvehigh/go-cli-project-template/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"slices"
	"strings"
)

func init() {
	cobra.OnInitialize(
		func() {
			viper.SetEnvPrefix(configs.PrefixENVVars)
			viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
			viper.AutomaticEnv()

			logger.InitializeLogger()
		},
	)
}

func GetRootCommand() *cobra.Command {
	// get version
	var rootCommand = &cobra.Command{
		Use:     "",
		Short:   "",
		Example: "...",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Validate the log level arg
			logLevel := strings.ToLower(viper.GetString("log-level"))
			validLogLevels := []string{"debug", "info", "warn", "error"}
			if !slices.Contains(validLogLevels, logLevel) {
				return fmt.Errorf(
					"invalid log level provided: '%s'. Must be one of %v",
					logLevel,
					validLogLevels,
				)
			}

			return configs.BindFlags(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// Print help screen if no args provided
			return cmd.Help()
		},
	}

	return rootCommand
}
