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
		finish_time timestamp
	);`
	insertTimer string = `insert into time_tracker (tittle, start_time)
		values (
	 	tittle,
		now(),
		now()
	);`
	updateTimer string = `insert into time_tracker (title, start_time, finish_time)
	values (
	tittle, start_time, finish_time
	);`
	selectTimer string = `select id, start_time, finish_time from time_tracker;`
	delete      string = `DELETE FROM time_tracker WHERE id=$1;`
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
	return nil
}
func (storage *Storage) CloseDB() error {
	return storage.conn.Close(context.Background())

}

// func (s *Storage) TrackerTimer() *TimeTrackerRepo {
// 	if s.trackertimeRepo != nil {
// 		return s.trackertimeRepo
// 	}
// 	s.trackertimeRepo = &TimeTrackerRepo{
// 		store: s,
// 	}
// 	return nil
// }
