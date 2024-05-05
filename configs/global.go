package configs

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Config struct {
	Debug bool
	Local bool
}

func GetGlobalConfig() *Config {
	return &Config{
		Debug: viper.GetBool(FlagDebug),
		Local: viper.GetBool(FlagLocal),
	}
}

func BindFlags(cmd *cobra.Command) error {
	errs := make([]error, 0)
	cmd.Flags().VisitAll(
		func(f *pflag.Flag) {
			err := viper.BindPFlag(f.Name, f)
			if err != nil {
				errs = append(
					errs,
					fmt.Errorf("error binding %s via viper: %s", f.Name, err.Error()),
				)
			}
			// cobra commands don't support environment variables so manually
			// tell cobra a flag was set if viper got a value via alternate method
			if viper.IsSet(f.Name) && viper.GetString(f.Name) != "" {
				setErr := cmd.Flags().Set(f.Name, viper.GetString(f.Name))
				if setErr != nil {
					errs = append(
						errs,
						fmt.Errorf(
							"setting flag with viper value for %s: %s",
							f.Name,
							setErr.Error(),
						),
					)
				}
			}
		},
	)

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}
