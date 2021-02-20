package converter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/lorenzoaiello/flightsim-rdv2volanta/internal/rdv"
)

func loadAircraftMappings() map[string]string {
	var mappings map[string]string

	data, err := ioutil.ReadFile("./mappings.json")
	if err != nil {
		fmt.Println("Error loading aircraft mappings file: ", err)
		return mappings
	}

	err = json.Unmarshal(data, &mappings)
	if err != nil {
		fmt.Println("Error parsing aircraft mappings file: ", err)
		return mappings
	}

	return mappings
}

func loadReservations() []rdv.Reservation {
	var reservations []rdv.Reservation
	data, err := ioutil.ReadFile("./source/reservations.json")
	if err != nil {
		fmt.Println("Error loading RDV reservation file: ", err)
		return reservations
	}

	err = json.Unmarshal(data, &reservations)
	if err != nil {
		fmt.Println("Error parsing RDV reservation file: ", err)
		return reservations
	}

	return reservations
}

func LoadReports() []rdv.Report {
	var reports []rdv.Report

	data, err := ioutil.ReadFile("./source/reports.json")
	if err != nil {
		fmt.Println("Error loading RDV report file: ", err)
		return reports
	}

	err = json.Unmarshal(data, &reports)
	if err != nil {
		fmt.Println("Error parsing RDV report file: ", err)
		return reports
	}

	return reports
}
