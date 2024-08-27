package cmd

import (
	"github.com/kostikovk/gohooks/lib"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall GoHooks CLI",
	Long:  `Uninstall GoHooks CLI...`,
	Run:   lib.RunUninstall,
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
