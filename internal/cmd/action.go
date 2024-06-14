package cmd

import (
	"aquarium/internal/domein/logic"
	"github.com/gookit/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func ActionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ActionCmd",
		Short: "Action Command",
		RunE: func(cmd *cobra.Command, args []string) error {
			tank := logic.NewTankManager().Load()
			color.Cyanln("TankName: " + tank.Name)

			prompt := promptui.Select{
				Label: "There is an aquarium in front of you.",
				Items: []string{
					"Init Tank",
					"Add Fish",
				},
			}
			_, result, _ := prompt.Run()

			switch result {
			case "Init Tank":
				color.Cyanln("Init Tank")
				_ = InitTankCmd().Execute()
			case "Add Fish":
				color.Cyanln("Add Fish")
				_ = AddFishCmd().Execute()
			default:
				color.Cyanln("Exception")
			}

			return nil
		},
	}

	return cmd
}
