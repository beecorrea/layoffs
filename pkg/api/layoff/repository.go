package layoff

import (
	"context"
	"time"

	"github.com/beecorrea/layoffs/pkg/persistence"
	"github.com/jackc/pgx/v5"
)

type LayoffRepository struct {
	Database *persistence.DatabaseConnection
}
type LayoffTable struct {
	CompanyName    string     `json:"name"`
	CompanyCountry string     `json:"country"`
	CompanyMarket  string     `json:"market"`
	HappenedAt     *time.Time `json:"happened_at"`
	ReportedAt     *time.Time `json:"reported_at"`
	ConfirmedAt    *time.Time `json:"confirmed_at"`
	AmountAffected int        `json:"amount_affected"`
}

const persistLayoffStatement = "INSERT INTO layoffs VALUES($1, $2, $3, $4, $5, $6)"

func (lr *LayoffRepository) PersistLayoff(ctx context.Context, layoff Layoff) error {
	c := lr.Database.GetConn()
	tx, err := c.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, persistLayoffStatement, layoff.Company.Name, layoff.Company.Country, layoff.HappenedAt, layoff.ReportedAt, layoff.ConfirmedAt, layoff.AmountAffected)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

const getLayoffByCompanyStatement = "SELECT * FROM layoffs WHERE companyName = $1"

func (lr *LayoffRepository) GetLayoffByCompany(ctx context.Context, company string) ([]LayoffTable, error) {
	c := lr.Database.GetConn()

	r, err := c.Query(ctx, getLayoffByCompanyStatement, company)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	rows, err := pgx.CollectRows(r, pgx.RowTo[LayoffTable])
	if err != nil {
		return nil, err
	}

	return rows, nil
}
