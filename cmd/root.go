package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/fx"

	"github.com/cristiano-pacheco/goflix/internal/identity"
	shared_modules "github.com/cristiano-pacheco/goflix/internal/shared/modules"
)

var rootCmd = &cobra.Command{
	Use:   "goflix",
	Short: "GoFlix API",
	Long:  `GoFlix API is a RESTful API for managing a video streaming platform.`,
	Run: func(cmd *cobra.Command, args []string) {
		app := fx.New(
			shared_modules.Module,
			identity.Module,
		)
		app.Run()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
