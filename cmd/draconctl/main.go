package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/ocurity/dracon/cmd/draconctl/components"
	"github.com/ocurity/dracon/cmd/draconctl/migrations"
	"github.com/ocurity/dracon/cmd/draconctl/pipelines"
)

var rootCmd = &cobra.Command{
	Use:   "draconctl",
	Short: "A CLI to manage all things related to Dracon",
	PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
		// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
		return initializeConfig(cmd)
	},
}

func main() {
	rootCmd.AddGroup(&cobra.Group{
		ID:    "top-level",
		Title: "Top-level Commands:",
	})
	pipelines.RegisterPipelinesSubcommands(rootCmd)
	migrations.RegisterMigrationsSubcommands(rootCmd)
	components.RegisterComponentsSubcommands(rootCmd)
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{})))
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initializeConfig(cmd *cobra.Command) error {
	v := viper.New()

	// When we bind flags to environment variables expect that the
	// environment variables are prefixed, e.g. a flag like --number
	// binds to an environment variable DRACONCTL_NUMBER. This helps
	// avoid conflicts.
	v.SetEnvPrefix("draconctl")

	// Environment variables can't have dashes in them, so bind them to their equivalent
	// keys with underscores, e.g. --favorite-color to DRACONCTL_FAVORITE_COLOR
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	// Bind to environment variables
	// Works great for simple config names, but needs help for names
	// like --favorite-color which we fix in the bindFlags function
	v.AutomaticEnv()

	// Bind the current command's flags to viper
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			if err := cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val)); err != nil {
				log.Println(err)
			}
		}
	})

	return nil
}
