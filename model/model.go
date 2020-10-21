package model

type Venue struct {
	Id         int      `json:"id"`
	Name       string   `json:"name"`
	Abbr       string   `json:"abbr"`
	Address    string   `json:"address"`
	PhoneNo    string   `json:"phoneNo"`
	AvatarUrl  string   `json:"avatarUrl"`
	CoverPhoto string   `json:"coverPhoto,omitempty"`
	InfoUrl    string   `json:"infoUrl,omitempty"`
	Notes      []string `json:"notes,omitempty"`
}

type Info struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"paragraphs,omitempty"`
	Notes      []string `json:"notes,omitempty"`
	Expanded   bool     `json:"expanded"`
}

type Attraction struct {
	Date    string   `json:"date"`
	VenueId int      `json:"venueId"`
	Notes   []string `json:"notes,omitempty"`
}

type HotelStay struct {
	CheckInDate    string   `json:"checkInDate"`
	CheckOutDate   string   `json:"checkOutDate"`
	ConfirmationNo string   `json:"confirmationNo"`
	VenueId        int      `json:"venueId"`
	Notes          []string `json:"notes,omitempty"`
}

type DiningReservation struct {
	Date           string   `json:"date"`
	ConfirmationNo string   `json:"confirmationNo"`
	VenueId        int      `json:"venueId"`
	Notes          []string `json:"notes,omitempty"`
}

type Flight struct {
	Airline          string   `json:"airline"`
	FlightNumber     string   `json:"flightNumber"`
	DepartureDate    string   `json:"departureDate"`
	DepartureVenueId int      `json:"departureVenueId,omitempty"`
	ArrivalDate      string   `json:"arrivalDate,omitempty"`
	ArrivalVenueId   int      `json:"arrivalVenueId"`
	Terminal         string   `json:"terminal,omitempty"`
	BaggageClaim     string   `json:"baggageClaim,omitempty"`
	RecordLocator    string   `json:"recordLocator,omitempty"`
	Notes            []string `json:"notes,omitempty"`
}

type Itinerary struct {
	Id                  string              `json:"id"`
	StartDate           string              `json:"startDate"`
	EndDate             string              `json:"endDate"`
	SalutationName      string              `json:"salutationName"`
	SalutationText      string              `json:"salutationText"`
	SalutationSignature string              `json:"salutationSignature"`
	Infos               []Info              `json:"infos,omitempty"`
	Attractions         []Attraction        `json:"attractions,omitempty"`
	HotelStays          []HotelStay         `json:"hotelStays,omitempty"`
	DiningReservations  []DiningReservation `json:"diningReservations,omitempty"`
	Flights             []Flight            `json:"flights,omitempty"`
}
