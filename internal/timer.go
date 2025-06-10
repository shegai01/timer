package internal

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
)

const (
	createTable string = `create table timer(
		id bigserial primary key,
		start_time datetime,
		finish_time datetime, 
	);`
	insertTimer string = `insert into timer( 
		start_time, finish_time
	) values (
		now(),
		now()
	);`
	selectTimer string = `select (
		id, start_time, finish_time 
		) from timer;`
	delete string = `DELETE FROM timer WHERE id=$1,name=$2`
)

type Timer struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	StartTime  time.Time `json:"start_time"`
	FinishTime time.Time `json:"finish_time"`
	// Finished   bool          `json:"finish"`
}

func (t *Timer) TimeDuration() time.Duration {
	return t.FinishTime.Sub(t.StartTime)
}

// storage db
type TimerDB struct {
	conn *pgx.Conn
}

func NewTimerDB(cfg string) (*TimerDB, error) {
	db, err := pgx.Connect(context.Background(), cfg)
	if err != nil {
		log.Println("connection failed")
		return nil, err
	}
	_, err = db.Exec(context.Background(), createTable)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &TimerDB{
		conn: db,
	}, nil
}
func (t *TimerDB) CreateTimer() (*Timer, error) {
	var time Timer
	err := t.conn.QueryRow(context.Background(), insertTimer).Scan(
		&time.ID,
		&time.Name,
		&time.StartTime,
		&time.FinishTime,
	)
	if err != nil {
		log.Println("creating failed", err)
		return nil, err
	}
	return &time, nil
}
func (t *TimerDB) ShowAllTimers() ([]Timer, error) {
	rows, err := t.conn.Query(context.Background(), selectTimer)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var alltimers []Timer
	for rows.Next() {
		var timer Timer
		err := rows.Scan(&timer.ID, &timer.Name, &timer.StartTime, &timer.FinishTime)
		if err != nil {
			return nil, err
		}
		alltimers = append(alltimers, timer)
	}
	return alltimers, nil
}

func (t *TimerDB) DeletebyID(id int, name string) error {
	_, err := t.conn.Exec(context.Background(), delete, id, name)
	if err != nil {
		fmt.Println("delete failed", err)
		return err
	}
	return nil
}
func (t *TimerDB) UpdateTimer(start, finish time.Time, name string) (*Timer, error) {
	_, err := t.conn.Exec(context.Background(), "UPDATE timer set name=$1, start_time=$2, finish_time=$3", start, finish, name)
	if err != nil {
		log.Println("updating failed")
		return nil, err
	}
	return nil, nil
}
