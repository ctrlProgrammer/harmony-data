package types

type Seller struct {
	PDV         int     `json:"pdv"`
	VendorCode  string  `json:"vendor_code"`
	Address     string  `json:"address"`
	SalesLiters float32 `json:"sales_liters"`
	SalesUSD    float32 `json:"sales_usd"`
	SalesUnits  float32 `json:"sales_units"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	City        string  `json:"city"`
	ProductName string  `json:"product_name"`
	VendorName  string  `json:"vendor_name"`
	Country     string  `json:"country"`
	Continent   string  `json:"continent"`
	Category    string  `json:"category"`
	Brand       string  `json:"brand"`
	SubBrand    string  `json:"sub_brand"`
	Item        string  `json:"item"`
	Year        string  `json:"year"`
	Month       string  `json:"month"`
	Quarter     string  `json:"quarter"`
	NSE         string  `json:"NSE"`
}
