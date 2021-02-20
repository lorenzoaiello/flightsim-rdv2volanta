package mapping

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/lorenzoaiello/flightsim-rdv2volanta/internal/converter"
)

func GenerateAircraftMappings() {
	uniqueAircraft := map[string]int{}

	// rdv: load flight reports
	for _, report := range converter.LoadReports() {
		var aircraftUsed []string

		if report.AircraftUsed == "" {
			fmt.Println(fmt.Sprintf("[%s] Error parsing aircraft", report.FlightreservationID))
			continue
		}

		err := json.Unmarshal([]byte(report.AircraftUsed), &aircraftUsed)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] Error parsing aircraft: %s", report.FlightreservationID, err))
			continue
		}

		for _, ac := range aircraftUsed {
			if _, ok := uniqueAircraft[ac]; !ok {
				uniqueAircraft[ac] = 1
			}
		}
	}

	mapping := map[string]string{}
	for ac := range uniqueAircraft {
		if ac == "" {
			continue
		}

		mapping[ac] = ""
	}

	output, err := json.Marshal(mapping)
	if err != nil {
		fmt.Println("Unable to generate aircraft mapping: ", err)
		os.Exit(1)
	}

	err = ioutil.WriteFile("./mappings.json", output, 0755)
	if err != nil {
		fmt.Println("Unable to save aircraft mapping: ", err)
		os.Exit(1)
	}

	fmt.Println("Aircraft mappings file generated successfully")
}
