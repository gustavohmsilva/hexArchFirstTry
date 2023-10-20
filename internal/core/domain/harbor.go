package domain

// Harbor is the main business domain of this microservice
type Harbor struct {
	InternalID  string    `json:"uuid"`        // internal ID in uuid format, used for internal keying
	ExternalID  string    `json:"id"`          // external ID (probably from other systems), preserved
	Name        string    `json:"name"`        // readable name of the harbor
	City        string    `json:"city"`        // city in which the harbor is located
	Country     string    `json:"country"`     // country in which the harbor is located
	Alias       []string  `json:"alias"`       // Alias/nicknames for the Alias
	Regions     []string  `json:"regions"`     // destination regions from this harbor (maybe?)
	Coordinates []float32 `json:"coordinates"` // Geo coordinates (lat/long) for geolocation purposes
	Province    string    `json:"province"`    // province where the harbor is located
	Timezone    string    `json:"timezone"`    // Timezone in which the harbor is located
	Unlocs      []string  `json:"unlocs"`      // no idea
	Code        string    `json:"code"`        // external numerical harbor code
}
