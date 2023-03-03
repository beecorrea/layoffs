package layoff

import "time"

type Layoff struct {
	Company Company `json:"company"`

	HappenedAt  *time.Time `json:"happened_at"`
	ReportedAt  *time.Time `json:"reported_at"`
	ConfirmedAt *time.Time `json:"confirmed_at"`

	AmountAffected int `json:"amount_affected"`
}

func GetLayoffs(company string) []Layoff {
	return getMockLayoff(company)
}

func getMockLayoff(company string) []Layoff {
	layoffs := make([]Layoff, 0)
	meta := getMockMetadata()
	mock := Layoff{
		Company:        Company{"Twitter", Brazil, LatinAmerica},
		HappenedAt:     meta[0],
		ReportedAt:     meta[1],
		ConfirmedAt:    meta[2],
		AmountAffected: 100,
	}
	layoffs = append(layoffs, mock)
	return layoffs
}

func getMockMetadata() []*time.Time {
	yt := time.Now().Add(-24 * time.Hour)
	eightH := time.Now().Add(-8 * time.Hour)
	now := time.Now()
	return []*time.Time{&yt, &eightH, &now}
}
