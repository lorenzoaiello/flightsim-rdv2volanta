package cmd

import (
	"github.com/spf13/cobra"

	"github.com/lorenzoaiello/flightsim-rdv2volanta/internal/mapping"
)

func init() {
	rootCmd.AddCommand(mappingCmd)
}

var mappingCmd = &cobra.Command{
	Use:   "mappings",
	Short: "Generates the aircraft mapping script",
	Run: func(cmd *cobra.Command, args []string) {
		mapping.GenerateAircraftMappings()
	},
}
