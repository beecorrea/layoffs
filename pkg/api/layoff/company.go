package layoff

type Company struct {
	Name    string      `json:"name"`
	Country CountryCode `json:"country"`
	Market  MarketCode  `json:"market"`
}

type CountryCode string

const (
	Brazil CountryCode = "BR"
)

type MarketCode string

const (
	LatinAmerica MarketCode = "LATAM"
)
