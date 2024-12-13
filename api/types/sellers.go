package types

type Seller struct {
	PDV          int     `json:"pdv"`
	Vendor_code  string  `json:"vendor_code"`
	Address      string  `json:"address"`
	Sales_liters float64 `json:"sales_liters"`
	Sales_usd    float64 `json:"sales_usd"`
	Sales_units  float64 `json:"sales_units"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	City         string  `json:"city"`
	Product_name string  `json:"product_name"`
	Vendor_name  string  `json:"vendor_name"`
	Country      string  `json:"country"`
	Continent    string  `json:"continent"`
	Category     string  `json:"category"`
	Brand        string  `json:"brand"`
	Sub_brand    string  `json:"sub_brand"`
	Item         string  `json:"item"`
	Year         int     `json:"year"`
	Month        string  `json:"month"`
	Quarter      string  `json:"quarter"`
	NSE          int     `json:"NSE"`
	District     string  `json:"district"`
}
