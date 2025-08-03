package cmd

import (
	"github.com/bad3r/antidot/internal/tui"
	"github.com/bad3r/antidot/internal/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var rulesSource = "https://raw.githubusercontent.com/bad3r/antidot/master/rules.yaml"

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update rules file",
	Run: func(cmd *cobra.Command, args []string) {
		tui.Debug("Updating rules...")
		err := utils.Download(rulesSource, rulesFilePath)
		tui.FatalIfError("Failed to update rules", err)

		tui.Print("Rules updated successfully")
	},
}
