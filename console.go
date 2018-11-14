package main

import (
	"os"

	"github.com/AdilHoumadi/background/command/email"
	c "github.com/AdilHoumadi/background/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{Use: c.ConsoleExec}

func init() {

	// Some config for cobra
	cobra.EnableCommandSorting = false
	RootCmd.LocalFlags().SortFlags = false
	RootCmd.Flags().SortFlags = false
	RootCmd.PersistentFlags().SortFlags = false

	// flags
	RootCmd.PersistentFlags().StringP(
		c.EnvFlag,
		c.EnvShortcut,
		os.Getenv(c.Env),
		c.EnvDescription,
	)
	RootCmd.PersistentFlags().StringP(
		c.VerboseFlag,
		c.VerboseShortcut,
		os.Getenv(c.VerboseFlag),
		c.UnSecureDescription,
	)

	// Bind persistent flags
	viper.BindPFlag(
		c.EnvFlag,
		RootCmd.PersistentFlags().Lookup(c.EnvFlag),
	)
	viper.BindPFlag(
		c.VerboseFlag,
		RootCmd.PersistentFlags().Lookup(c.VerboseFlag),
	)
}

func main() {
	RootCmd.AddCommand(email.Cmd)
	RootCmd.Execute()
}
