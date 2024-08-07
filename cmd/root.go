/*
Copyright © 2024 NAME HERE yusufakdenizxd@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scriptmaster",
	Short: "Script Manager for Bash and Zsh Scripts",

	Long: `Script Master is a cross-platform and easy-to-use script manager
designed to simplify the process of installing and managing Bash and Zsh scripts.

Whether you're working on a single script or managing a collection of scripts,
Script Master makes it easy to keep everything organized and accessible.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
