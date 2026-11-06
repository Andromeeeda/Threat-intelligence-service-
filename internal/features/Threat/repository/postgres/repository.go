package repository_postgres_rep

import "github.com/jackc/pgx/v5"

type ThreatRepository struct {
	conn *pgx.Conn
}

func NewThreatRepository(conn *pgx.Conn) *ThreatRepository {
	return &ThreatRepository{
		conn: conn,
	}
}
