package main

import (
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "wl",}

	var cmd = &cobra.Command{
		Use: "update"
		Short: "Atualiza aplicacao"
		Run: func(cmd *cobra.PositionalArgs)
	}

}
