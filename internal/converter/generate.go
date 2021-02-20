package converter

import (
	"encoding/json"
	"fmt"
	"github.com/lorenzoaiello/flightsim-rdv2volanta/internal/rdv"
	"github.com/lorenzoaiello/flightsim-rdv2volanta/internal/volanta"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

func GenerateVolantaFlights() {
	aircraftMappings := loadAircraftMappings()

	for _, reservation := range loadReservations() {
		// rdv: load flight reports
		var report rdv.Report
		for _, ri := range LoadReports() {
			if ri.FlightreservationID == reservation.ID {
				report = ri
				break
			}
		}

		// rdv: load flight data
		data, err := ioutil.ReadFile(fmt.Sprintf("./source/%s/rawdata.json", reservation.ID))
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] Skipped - Error loading RDV rawdata file: %s", reservation.ID, err))
			continue
		}
		var rdvpos []rdv.Position
		err = json.Unmarshal(data, &rdvpos)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] Skipped - Error parsing RDV rawdata file: %s", reservation.ID, err))
			continue
		}

		// reformat position data
		var positions []volanta.Position
		var lastAircraft string
		var startTime int
		var endTime int
		var touchdownPosition rdv.Position
		for i, pos := range rdvpos {
			vsfpm, err := strconv.ParseFloat(pos.VsFpm, 0)
			if err != nil {
				fmt.Println(fmt.Sprintf("[%s] Warning - Unable to convert Vertical Speed data: %s", reservation.ID, err))
			}

			gs, err := strconv.ParseFloat(pos.Groundspeed, 0)
			if err != nil {
				fmt.Println(fmt.Sprintf("[%s] Warning - Unable to convert Groundspeed data: %s", reservation.ID, err))
			}

			hdg, _ := strconv.ParseFloat(pos.HeadingMagnetic, 0)
			if err != nil {
				fmt.Println(fmt.Sprintf("[%s] Warning - Unable to convert Heading data: %s", reservation.ID, err))
			}

			timeTaken, err := strconv.Atoi(pos.TakenTimestamp)
			if err != nil {
				fmt.Println(fmt.Sprintf("[%s] Warning - Unable to convert Timestamp data: %s", reservation.ID, err))
			}

			altitude, err := strconv.ParseFloat(pos.AltitudeMsl, 0)
			if err != nil {
				fmt.Println(fmt.Sprintf("[%s] Warning - Unable to convert Altitude data: %s", reservation.ID, err))
			}

			if i == 0 {
				startTime = timeTaken
			}

			altitudeParts := strings.Split(fmt.Sprintf("%f", math.Round(altitude)), ".")

			positions = append(positions, volanta.Position{
				Latitude:      pos.Lat,
				Longitude:     pos.Lon,
				VerticalSpeed: int(vsfpm),
				Speed:         int(gs),
				Altitude:      altitudeParts[0],
				Heading:       int(hdg),
				Time:          timeTaken,
			})
			lastAircraft = pos.AircraftString
			endTime = timeTaken

			if i > 1 && pos.OnGround == "1" && rdvpos[i-1].OnGround == "0" {
				touchdownPosition = pos
			}
		}

		// check if aircraft mapping for aircraft exists
		var aircraftType string
		if val, ok := aircraftMappings[lastAircraft]; ok {
			aircraftType = val
		} else {
			fmt.Println(fmt.Sprintf("[%s] Warning - Unable to find aircraft mapping for: '%s'", reservation.ID, lastAircraft))
		}

		// rdv: convert general flight report data
		flightId, err := strconv.Atoi(reservation.ID)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] Warning - Unable to convert Flight ID: %s", reservation.ID, err))
		}

		distanceNm, err := strconv.Atoi(reservation.DistanceNm)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] Warning - Unable to convert Flight Distance: %s", reservation.ID, err))
		}

		paxCount, err := strconv.Atoi(report.PaxCount)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] Warning - Unable to convert Passenger Total: %s", reservation.ID, err))
		}

		fuelUsed, err := strconv.Atoi(report.FuelUsed)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] Warning - Unable to convert Fuel Usage: %s", reservation.ID, err))
		}

		landingRate, err := strconv.Atoi(report.LandingRate)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] Warning - Unable to convert Landing Rate: %s", reservation.ID, err))
		}

		minutesEnroute, err := strconv.Atoi(report.MinutesEnroute)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] Warning - Unable to convert Enroute Time: %s", reservation.ID, err))
		}

		// volanta: build landing event
		touchdownTime, err := strconv.Atoi(touchdownPosition.TakenTimestamp)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] Warning - Unable to convert touchdown Timestamp data: %s", reservation.ID, err))
		}

		rate, err := strconv.ParseFloat(touchdownPosition.VsFpm, 0)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] Warning - Unable to convert touchdown Vertical Speed data: %s", reservation.ID, err))
		}

		force, err := strconv.ParseFloat(touchdownPosition.GForce, 0)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] Warning - Unable to convert touchdown G-Force data: %s", reservation.ID, err))
		}

		speed, err := strconv.ParseFloat(touchdownPosition.Groundspeed, 0)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] Warning - Unable to convert touchdown Groundspeed data: %s", reservation.ID, err))
		}

		pitch, err := strconv.ParseFloat(touchdownPosition.PitchDeg, 0)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] Warning - Unable to convert touchdown Pitch data: %s", reservation.ID, err))
		}

		roll, err := strconv.ParseFloat(touchdownPosition.BankDeg, 0)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] Warning - Unable to convert touchdown Bank data: %s", reservation.ID, err))
		}

		latitude, err := strconv.ParseFloat(touchdownPosition.Lat, 0)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] Warning - Unable to convert touchdown Latitude data: %s", reservation.ID, err))
		}

		longitude, err := strconv.ParseFloat(touchdownPosition.Lon, 0)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] Warning - Unable to convert touchdown Longitude data: %s", reservation.ID, err))
		}

		touchdownEvent := volanta.TouchdownEvent{
			Rate:      int(rate),
			Force:     force,
			Pitch:     int(pitch),
			Roll:      int(roll),
			Speed:     int(speed),
			Latitude:  latitude,
			Longitude: longitude,
		}

		landingData, err := json.Marshal(touchdownEvent)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] Warning - Unable to convert landing event: %s", reservation.ID, err))
		}

		landingEvent := volanta.Event{
			Name: "landing",
			Data: string(landingData),
			Time: touchdownTime,
		}

		// volanta: build flight data
		flight := volanta.Flight{
			ID:           flightId,
			FlightNumber: fmt.Sprintf("%s/%s", reservation.Callsign, reservation.FlightNum),
			Route:        fmt.Sprintf("%s - %s", reservation.Depicao, reservation.Arricao),
			FlightRules:  "IFR",
			FlightType:   "Scheduled",
			Network:      "OFFLINE",
			Note:         report.PilotRemarks,
			StatusCode:   "202",
			Distance:     distanceNm,
			Aircraft: volanta.Aircraft{
				Registration: volanta.AircraftRegistration{
					Registration: "PH-RDV",
					Airline: volanta.Airline{
						Logo: "https://storage.googleapis.com/pf-assets/img/airlines/logos/571-klm-royal-dutch-airlines.png?v=5c3335fa",
						Name: "KLM Royal Dutch Airlines",
						Icao: "KLM",
						Iata: "KL",
					},
					Aircraft: volanta.Fleet{
						Type: aircraftType,
					},
				},
			},
			Departure: volanta.Airport{
				Icao: reservation.Depicao,
			},
			Arrival: volanta.Airport{
				Icao: reservation.Arricao,
			},
			Flight: volanta.FlightBlock{
				ID:          flightId,
				Diverted:    reservation.Diverticao,
				FuelBurned:  fuelUsed,
				BlockFuel:   fuelUsed,
				LandingFpm:  landingRate,
				BlockTime:   fmtDuration(time.Minute * time.Duration(minutesEnroute)),
				Pax:         paxCount,
				TimeStarted: time.Unix(int64(startTime), 0).Format("2006-01-02 15:04:05"),
				TimeEnded:   time.Unix(int64(endTime), 0).Format("2006-01-02 15:04:05"),
				Path:        positions,
				Events:      []volanta.Event{landingEvent},
			},
		}

		// volanta: export
		output, err := json.Marshal(flight)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] Error - Unable to generate Volanta export: %s", reservation.ID, err))
			continue
		}

		err = ioutil.WriteFile("./output/Flight "+reservation.ID+".json", output, 0755)
		if err != nil {
			fmt.Println(fmt.Sprintf("[%s] Error - Unable to write Volanta flight file: %s", reservation.ID, err))
			continue
		}

		fmt.Println(fmt.Sprintf("[%s] Converted Successfully", reservation.ID))
	}
}
