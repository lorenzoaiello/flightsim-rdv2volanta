package rdv

type Reservation struct {
	ID             string      `json:"id"`
	ReservedTime   string      `json:"reserved_time"`
	ExpirationTime string      `json:"expiration_time"`
	FlightNum      string      `json:"flight_num"`
	Callsign       string      `json:"callsign"`
	Depicao        string      `json:"depicao"`
	Arricao        string      `json:"arricao"`
	Diverticao     interface{} `json:"diverticao"`
	DistanceNm     string      `json:"distance_nm"`
}

type Report struct {
	FlightreportID      string      `json:"flightreport_id"`
	FlightreservationID string      `json:"flightreservation_id"`
	FuelUsed            string      `json:"fuel_used"`
	CargoKg             string      `json:"cargo_kg"`
	PaxCount            string      `json:"pax_count"`
	LandingRate         string      `json:"landing_rate"`
	MinutesEnroute      string      `json:"minutes_enroute"`
	PilotRemarks        interface{} `json:"pilot_remarks"`
	AircraftUsed        string      `json:"aircraft_used"`
}

type Position struct {
	Lat             string `json:"lat"`
	Lon             string `json:"lon"`
	HeadingMagnetic string `json:"heading_magnetic"`
	AircraftString  string `json:"aircraft_string"`
	VsFpm           string `json:"vs_fpm"`
	GForce          string `json:"gforce"`
	PitchDeg        string `json:"pitch_deg"`
	BankDeg         string `json:"bank_deg"`
	Groundspeed     string `json:"groundspeed"`
	Gear            string `json:"gear"`
	AltitudeMsl     string `json:"altitude_msl"`
	OnGround        string `json:"onground"`
	AirPath         string `json:"air_path"`
	RealUtcTime     string `json:"real_utc_time"`
	TakenTimestamp  string `json:"taken_timestamp"`
}
