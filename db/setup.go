package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var connectionString string = os.Getenv("DATABASE_URL")

func NewPostgres(ctx context.Context) (*Postgres, error) {
	conn, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		return nil, err
	}

	_, err = conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS service (
			ID SERIAL PRIMARY KEY, 
			name VARCHAR(255),
			csi VARCHAR(10),
			type VARCHAR(10),
			status VARCHAR(10) DEFAULT 'Draft',
			sbom JSON,
			repository VARCHAR(255)
		);

		CREATE TABLE IF NOT EXISTS permit (
			ID SERIAL PRIMARY KEY, 
			csi VARCHAR(10),
			type VARCHAR(10),
			status VARCHAR(10) DEFAULT 'Requested'
		);

		CREATE TABLE IF NOT EXISTS permit_service (
			ID SERIAL PRIMARY KEY, 
			permit_id INT,
			CONSTRAINT fk_permit
      			FOREIGN KEY(permit_id) 
	  			REFERENCES permit(ID),
			service_id INT,
			CONSTRAINT fk_service
      			FOREIGN KEY(service_id) 
	  			REFERENCES service(ID)
		);
	`)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		conn: conn,
	}, err
}

type Postgres struct {
	conn *pgxpool.Pool
}
