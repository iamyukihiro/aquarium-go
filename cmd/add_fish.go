package cmd

import (
	"aquarium/usecase"
	"github.com/spf13/cobra"
)

func AddFishCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "InitTankCmd",
		Short: "Init Tank Cmd",
		RunE: func(cmd *cobra.Command, args []string) error {
			usecase.AddFish()

			return nil
		},
	}

	return cmd
}
