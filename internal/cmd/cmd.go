package cmd

import (
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

func init() {
	cobra.OnInitialize()
	Root.AddCommand(
		ActionCmd(),
		InitTankCmd(),
		AddFishCmd(),
	)
	Root.PersistentFlags().IntP(
		"nest",
		"n",
		5,
		"Specify the depth of the directory",
	)
	Root.PersistentFlags().StringP(
		"char",
		"c",
		"utf-8",
		"Select charcter code(utf-8, shift-jis)",
	)
}

var Root = &cobra.Command{
	Use:   "aquarium",
	Short: "Aquarium Simulator",
	Run: func(cmd *cobra.Command, args []string) {
		color.Cyanln("🐟Welcome to Aquarium")
		ActionCmd().Execute()
	},
}
