package layoff

import (
	"context"
	"time"
)

type Layoff struct {
	Id             int
	CompanyName    string     `json:"name"`
	CompanyCountry string     `json:"country"`
	CompanyMarket  string     `json:"market"`
	HappenedAt     *time.Time `json:"happened_at"`
	ReportedAt     *time.Time `json:"reported_at"`
	ConfirmedAt    *time.Time `json:"confirmed_at"`
	AmountAffected int        `json:"amount_affected"`
}

type LayoffService struct {
	repo LayoffRepository
}

func NewLayoffService(repo LayoffRepository) LayoffService {
	return LayoffService{repo}
}
func (ls LayoffService) GetLayoffs(ctx context.Context, company string) ([]Layoff, error) {
	return ls.repo.GetLayoffByCompany(ctx, company)
}

func (ls LayoffService) PersistLayoff(ctx context.Context, layoff Layoff) (*Layoff, error) {
	return ls.repo.PersistLayoff(ctx, layoff)
}

// func getMockLayoff(company string) []Layoff {
// 	layoffs := make([]Layoff, 0)
// 	meta := getMockMetadata()
// 	mock := Layoff{
// 		Company:        Company{"Twitter", Brazil, LatinAmerica},
// 		HappenedAt:     meta[0],
// 		ReportedAt:     meta[1],
// 		ConfirmedAt:    meta[2],
// 		AmountAffected: 100,
// 	}
// 	layoffs = append(layoffs, mock)
// 	return layoffs
// }

// func getMockMetadata() []*time.Time {
// 	yt := time.Now().Add(-24 * time.Hour)
// 	eightH := time.Now().Add(-8 * time.Hour)
// 	now := time.Now()
// 	return []*time.Time{&yt, &eightH, &now}
// }
