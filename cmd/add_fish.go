package cmd

import (
	"aquarium/usecase"
	"github.com/spf13/cobra"
)

func AddFishCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "InitTankCmd",
		Short: "Init Tank Command",
		RunE: func(cmd *cobra.Command, args []string) error {
			usecase.NewAddFishUseCase().AddFish()

			return nil
		},
	}

	return cmd
}
