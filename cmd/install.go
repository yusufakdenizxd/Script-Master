/*
Copyright © 2024 NAME HERE yusufakdenizxd@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Deneme",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("install command bruuh")
		dat, err := os.ReadDir("/Users/yusufakdeniz/dev/dot-config/zsh/script")
		if err != nil {
			panic(err)
		}

		for _, e := range dat {
			fmt.Println(e.Name())
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
