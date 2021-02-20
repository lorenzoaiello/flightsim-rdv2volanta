package cmd

import (
	"github.com/spf13/cobra"

	"github.com/lorenzoaiello/flightsim-rdv2volanta/internal/converter"
)

func init() {
	rootCmd.AddCommand(convertCmd)
}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Runs the conversion script",
	Run: func(cmd *cobra.Command, args []string) {
		converter.GenerateVolantaFlights()
	},
}
