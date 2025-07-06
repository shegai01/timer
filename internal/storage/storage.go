package storage

import (
	"context"

	"github.com/jackc/pgx/v4"
)

const (
	createTable string = `create table if not exists time_tracker(
		id bigserial primary key,
		tittle varchar(255) not null,
		start_time timestamp,
		stop_time timestamp
	);`

	insertTimer string = `insert into time_tracker (tittle, start_time, stop_time)
		values (
		$1, now(), now())
		returning
		id,
		tittle,
		start_time,
		stop_time
	;`

	updateTimer string = `update time_tracker (tittle, stop_time)
	set tittle = $1, stop_time = $2;`

	selectALLtimer string = `select id, tittle, start_time, stop_time from time_tracker;`

	selectTimer string = `select id, tittle, start_time, stop_time from time_tracker where tittle=$1;`

	delete string = `delete  from time_tracker where tittle=$1;`
)

// storage db
type Storage struct {
	conn *pgx.Conn
}

func NewStorage() *Storage {
	return &Storage{}
}
func (t *Storage) ConnectDB(cfg string) error {
	db, err := pgx.Connect(context.Background(), cfg)
	if err != nil {
		return err
	}
	t.conn = db
	_, err = t.conn.Exec(context.Background(), createTable)
	if err != nil {
		return err
	}
	return nil
}
func (storage *Storage) CloseDB() error {
	return storage.conn.Close(context.Background())

}
