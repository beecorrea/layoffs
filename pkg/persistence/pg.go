package persistence

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Open(ctx context.Context) (*DatabaseConnection, error) {
	conf, _ := pgx.ParseConfig("")
	c, err := pgx.ConnectConfig(ctx, conf)
	if err != nil {
		return nil, err
	}

	return &DatabaseConnection{conn: c, config: conf}, nil
}

type DatabaseConnection struct {
	config *pgx.ConnConfig
	conn   *pgx.Conn
}

func (dc *DatabaseConnection) GetConn() *pgx.Conn {
	return dc.conn
}
