package volanta

type Flight struct {
	ID           int         `json:"id"`
	FlightNumber string      `json:"flight_number"`
	Route        string      `json:"route"`
	FlightRules  string      `json:"flight_rules"`
	FlightType   string      `json:"flight_type"`
	Network      string      `json:"network"`
	Note         interface{} `json:"note"`
	StatusCode   string      `json:"status_code"`
	Distance     int         `json:"distance"`
	Directory    interface{} `json:"directory"`
	Aircraft     Aircraft    `json:"aircraft"`
	Departure    Airport     `json:"departure"`
	Arrival      Airport     `json:"arrival"`
	Flight       FlightBlock `json:"flight"`
	Username     string      `json:"username"`
}

type Aircraft struct {
	Registration AircraftRegistration `json:"registration"`
}

type AircraftRegistration struct {
	Registration string  `json:"registration"`
	Airline      Airline `json:"airline"`
	Aircraft     Fleet   `json:"aircraft"`
}

type Airline struct {
	Logo string      `json:"logo"`
	Name interface{} `json:"name"`
	Icao string      `json:"icao"`
	Iata string      `json:"iata"`
}

type Fleet struct {
	Type string `json:"type"`
}

type Airport struct {
	Name      string `json:"name"`
	Icao      string `json:"icao"`
	Iata      string `json:"iata"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type FlightBlock struct {
	ID          int         `json:"id"`
	Diverted    interface{} `json:"diverted"`
	FuelBurned  int         `json:"fuel_burned"`
	BlockFuel   int         `json:"block_fuel"`
	BlockTime   string      `json:"block_time"`
	LandingFpm  int         `json:"landing_fpm"`
	Pax         interface{} `json:"pax"`
	Path        []Position  `json:"path"`
	Events      []Event     `json:"events"`
	TimeStarted string      `json:"time_started"`
	TimeEnded   string      `json:"time_ended"`
}

type Position struct {
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	VerticalSpeed int    `json:"vertical_speed"`
	Speed         int    `json:"speed"`
	Altitude      string `json:"altitude"`
	Heading       int    `json:"heading"`
	Time          int    `json:"time"`
}

type Event struct {
	Name string `json:"name"`
	Data string `json:"data"`
	Time int    `json:"time"`
}

type TouchdownEvent struct {
	Rate      int     `json:"Rate"`
	Force     float64 `json:"Force"`
	Pitch     int     `json:"Pitch"`
	Roll      int     `json:"Roll"`
	Speed     int     `json:"Speed"`
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
}
