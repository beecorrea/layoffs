package layoff

import (
	"context"
	"log"

	"github.com/beecorrea/layoffs/pkg/persistence"
	"github.com/jackc/pgx/v5"
)

type LayoffRepository struct {
	Database *persistence.DatabaseConnection
}

const persistLayoffStatement = "INSERT INTO layoffs(company_name, company_country,company_market,happened_at,reported_at,confirmed_at,amount_affected) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING *;"

func (lr *LayoffRepository) PersistLayoff(ctx context.Context, layoff Layoff) (*Layoff, error) {
	c := lr.Database.GetConn()
	tx, err := c.Begin(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	rows, _ := tx.Query(ctx, persistLayoffStatement, layoff.CompanyName, layoff.CompanyCountry, layoff.CompanyMarket, layoff.HappenedAt, layoff.ReportedAt, layoff.ConfirmedAt, layoff.AmountAffected)
	row, err := pgx.CollectOneRow(rows, pgx.RowToStructByPos[Layoff])
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &row, nil
}

const getLayoffByCompanyStatement = "SELECT * FROM layoffs WHERE company_name = $1;"

func (lr *LayoffRepository) GetLayoffByCompany(ctx context.Context, company string) ([]Layoff, error) {
	c := lr.Database.GetConn()

	r, err := c.Query(ctx, getLayoffByCompanyStatement, company)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	rows, err := pgx.CollectRows(r, pgx.RowToStructByPos[Layoff])
	if err != nil {
		return nil, err
	}

	return rows, nil
}
