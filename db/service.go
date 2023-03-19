package db

import (
	"context"
)

func (p *Postgres) GetService(ctx context.Context, id int) (*Service, error) {

	var svc Service

	err := p.conn.QueryRow(ctx, "SELECT Id, Name, Csi, Type, Status, Repository FROM SERVICE WHERE ID=$1", id).Scan(&svc.Id, &svc.Name, &svc.Csi, &svc.Type, &svc.Status, &svc.Repository)
	if err != nil {
		return nil, err
	}

	return &svc, nil
}

func (p *Postgres) CreateService(ctx context.Context, s *Service) (*Service, error) {
	_, err := p.conn.Exec(ctx, `
		INSERT INTO SERVICE (NAME, CSI, TYPE, STATUS, REPOSITORY)
			VALUES ($1, $2, $3, $4, $5);
	`, s.Name, s.Csi, s.Type, s.Status, s.Repository)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (p *Postgres) ListService(ctx context.Context, csi string) ([]Service, error) {
	rows, err := p.conn.Query(ctx, "SELECT Id, Name, Csi, Type, Status, Repository FROM SERVICE WHERE CSI=$1", csi)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var services []Service
	for rows.Next() {
		var s Service
		err := rows.Scan(&s.Id, &s.Name, &s.Csi, &s.Type, &s.Status, &s.Repository)
		if err != nil {
			return nil, err
		}

		services = append(services, s)
	}

	return services, nil

}

func (p *Postgres) RegisterSbom(ctx context.Context, serviceId int, sbom string) error {

	_, err := p.conn.Exec(ctx, `
		UPDATE SERVICE
			SET SBOM=$1
			WHERE ID=$2;
	`, sbom, serviceId)

	if err != nil {
		return err
	}

	return nil
}

type Service struct {
	Id         int
	Name       string
	Csi        string
	Type       string
	Status     string
	Repository string
}
