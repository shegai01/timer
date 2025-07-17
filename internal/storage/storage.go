package storage

import (
	"context"
	"log"
	"sync"

	"github.com/jackc/pgx/v4"
)

const (
	createTable string = `create table if not exists time_tracker(
		id bigserial primary key,
		title varchar(255) not null,
		start_time timestamp,
		stop_time timestamp
	);`

	insertTimer string = `insert into time_tracker (title, start_time)
		values (
		$1, now())
		returning id, title, start_time, stop_time
	;`

	stopTimer string = `update time_tracker set stop_time = $2 where id = $1;`

	showALLtimer string = `select id, title, start_time, stop_time from time_tracker;`

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
func (t *Storage) Connect(cfg string) error {
	db, err := pgx.Connect(context.Background(), cfg)
	if err != nil {
		log.Println(err)
		return err
	}
	t.conn = db
	_, err = t.conn.Exec(context.Background(), createTable)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("it's ok")
	return nil
}
func (storage *Storage) Close() error {
	storage.conn = nil
	return storage.conn.Close(context.Background())

}
