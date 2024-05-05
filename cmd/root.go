package cmd

import (
	"fmt"
	"slices"
	"strings"
)

func init() {
	cobra.OnInitialize(
		func() {
			viper.SetEnvPrefix(config.PrefixEnvVars)
			viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
			viper.AutomaticEnv()

			logger.InitialiseLogger()
			config.BindDefaultS3BucketSettings()
		},
	)
}

func GetRootCommand() *cobra.Command {
	// get version
	versionStr := viper.GetString(config.DKBinaryVersion)

	var rootCommand = &cobra.Command{
		Use: "dk",
		Short: fmt.Sprintf(
			"deploy-kit: standardised deployments at Rokt, version: %v",
			versionStr,
		),
		Example: "dk ...",
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

			return config.BindFlags(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// Print help screen if no args provided
			return cmd.Help()
		},
	}

	return rootCommand
}
