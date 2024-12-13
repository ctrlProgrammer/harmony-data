package types

type District struct {
	Name      string `json:"name"`
	City      string `json:"city"`
	UUID      string `json:"uuid"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Radius    int    `json:"radius"`
}
