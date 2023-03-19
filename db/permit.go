package db

import (
	"context"
	"log"
)

func (pg *Postgres) RegisterPermit(ctx context.Context, pmt *Permit) (int, error) {

	var id int
	err := pg.conn.QueryRow(ctx, `
		INSERT INTO PERMIT (CSI, TYPE, STATUS)
			VALUES ($1, $2, $3)
			RETURNING ID;
	`, pmt.Csi, pmt.Type, pmt.Status).Scan(&id)
	if err != nil {
		return -1, err
	}

	for _, svcID := range pmt.Services {
		_, err := pg.conn.Exec(ctx, `
		INSERT INTO PERMIT_SERVICE (PERMIT_ID, SERVICE_ID)
			VALUES ($1, $2)
		`, id, svcID)
		if err != nil {
			return -1, err
		}
	}

	return id, nil
}

func (pg *Postgres) ListPermits(ctx context.Context, csi string) ([]Permit, error) {
	rows, err := pg.conn.Query(ctx, "SELECT Id, Csi, Type, Status FROM PERMIT WHERE CSI=$1", csi)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var permits []Permit
	for rows.Next() {
		var p Permit
		err := rows.Scan(&p.Id, &p.Csi, &p.Type, &p.Status)
		if err != nil {
			return nil, err
		}

		jrows, err := pg.conn.Query(ctx, "SELECT service_id FROM PERMIT_SERVICE WHERE PERMIT_ID=$1", p.Id)
		if err != nil {
			return nil, err
		}
		defer jrows.Close()
		for jrows.Next() {
			var svc int
			err := jrows.Scan(&svc)
			if err != nil {
				return nil, err
			}

			p.Services = append(p.Services, svc)
		}

		permits = append(permits, p)
	}

	return permits, nil
}

// GetServices for permit fetches all services for a given permit
func (pg *Postgres) GetServicesForPermit(ctx context.Context, permitId int) ([]Service, error) {
	rows, err := pg.conn.Query(ctx, "SELECT service_id FROM PERMIT_SERVICE WHERE PERMIT_ID=$1", permitId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	log.Println("rows", rows)

	var services []Service
	for rows.Next() {
		var svcId int
		err := rows.Scan(&svcId)
		if err != nil {
			return nil, err
		}

		svc, err := pg.GetService(ctx, svcId)
		if err != nil {
			log.Printf("error getting service %d: %v", svcId, err)
			return nil, err
		}

		services = append(services, *svc)
	}

	return services, nil
}

type Permit struct {
	Id       int
	Type     string
	Status   string
	Csi      string
	Services []int
}
