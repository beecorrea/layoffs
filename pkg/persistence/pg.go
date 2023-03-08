package persistence

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type DatabaseConnectionParams struct {
	Host         string
	Port         string
	Password     string
	Username     string
	DatabaseName string
}

func (params *DatabaseConnectionParams) ToDSN() string {
	return ""
}

func (params *DatabaseConnectionParams) ToPgxConf() *pgx.ConnConfig {
	return nil
}

func (params *DatabaseConnectionParams) Open(ctx context.Context) (*DatabaseConnection, error) {
	conf := params.ToPgxConf()
	c, err := pgx.ConnectConfig(ctx, conf)
	if err != nil {
		return nil, err
	}

	return &DatabaseConnection{dsn: params.ToDSN(), config: conf, conn: c}, nil
}

type DatabaseConnection struct {
	dsn    string
	config *pgx.ConnConfig
	conn   *pgx.Conn
}

func (dc *DatabaseConnection) GetConn() *pgx.Conn {
	return dc.conn
}
