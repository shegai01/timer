package storage

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v4"
)

const (
	createTable string = `create table if not exists time_tracker(
		id bigserial primary key,
		title varchar(255) not null,
		start_time timestamp not null,
		stop_time timestamp
	);`

	insertTimer string = `insert into time_tracker (title, start_time)
		values ($1, now())
		returning id, title, start_time, stop_time
	;`

	stopTimer string = `update time_tracker set stop_time = now() where id = $1;`

	// TODO: pagination
	showAllTimers string = `select id, title, start_time, stop_time from time_tracker;`

	getTimer string = `select id, title, start_time, stop_time from time_tracker where id=$1;`

	delete string = `delete from time_tracker where id=$1;`
)

type Storage struct {
	mu   sync.Mutex
	conn *pgx.Conn
}

func NewStorage() *Storage {
	return &Storage{}
}

func (t *Storage) Connect(connString string) error {
	db, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return fmt.Errorf("pgx.Connect: %w", err)
	}

	t.conn = db
	_, err = t.conn.Exec(context.Background(), createTable)
	if err != nil {
		return fmt.Errorf("t.conn.Exec: %w", err)
	}

	return nil
}

func (storage *Storage) Close() error {
	defer func() { storage.conn = nil }()

	err := storage.conn.Close(context.Background())
	if err != nil {
		return fmt.Errorf("storage.conn.Close: %w", err)
	}

	return nil
}
