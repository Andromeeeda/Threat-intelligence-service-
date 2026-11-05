package repository_postgres_rep

import (
	domain_threatDomain "Threat-intelligence-service/internal/core/domain"
	common_errors "Threat-intelligence-service/internal/core/errors/common"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
)

func (r *ThreatRepository) AddThreat(ctx context.Context, threat *domain_threatDomain.Threat) error {
	query := `
	INSERT INTO app.threats (md5, name) VALUES ($1, $2)`

	_, err := r.conn.Exec(ctx, query, threat.Md5, threat.Name)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return fmt.Errorf("repo Save: %w", common_errors.ErrConflict)
			}
		}


		return fmt.Errorf("repo Save: %w", err)
	}

	return nil 
}
