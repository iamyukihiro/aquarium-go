package cmd

import (
	"aquarium/internal/usecase"
	"github.com/gookit/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func InitTankCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "InitTankCmd",
		Short: "Init Tank Command",
		RunE: func(cmd *cobra.Command, args []string) error {
			prompt := promptui.Select{
				Label: "⚠ Do you really want to init?",
				Items: []string{
					"Yes",
					"No",
				},
			}

			_, result, _ := prompt.Run()

			switch result {
			case "Yes":
				color.Cyanln("bye.")
				usecase.NewInitTankUseCase().InitTank()
			default:
				_ = ActionCmd().Execute()
			}

			return nil
		},
	}

	return cmd
}
