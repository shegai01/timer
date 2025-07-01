package storage

import (
	"context"

	"github.com/jackc/pgx/v4"
)

const (
	dbURI       string = "postgres"
	createTable string = `create table if not exists time_tracker(
		id bigserial primary key,
		title varchar(255) not null,
		start_time timestamp,
		finish_time timestamp
	);`
	insertTimer string = `insert into timerdb (tittle, start_time)
		values (
	 	tittle,
		now(),
		now()
	);`
	selectTimer string = `select id, start_time, finish_time from timer;`
	delete      string = `DELETE FROM timer WHERE id=$1;`
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
