package storage

import (
	"context"
	"log"
	"time"

	"github.com/shegai01/timer/internal/model"
)

func (storage *Storage) CreateListTimer() error {
	_, err := storage.conn.Exec(context.Background(), createTable)
	if err != nil {
		log.Println("create table failed")
		return err
	}
	log.Println("database created", createTable)

	return nil
}
func (storage *Storage) CreateTimer(tittle string) (*model.Timer, error) {
	var time model.Timer
	err := storage.conn.QueryRow(context.Background(), insertTimer, tittle).Scan(
		&time.ID,
		&time.Tittle,
		&time.StartTime,
		&time.StopTime)
	if err != nil {
		log.Println("can't creating without tittle")
		return nil, err
	}

	return &time, nil
}
func (storage *Storage) GetTimerbyTittle(tittle string) (*model.Timer, error) {
	var time model.Timer
	err := storage.conn.QueryRow(context.Background(), selectTimer, tittle).Scan(
		&time.ID,
		&time.Tittle,
		&time.StartTime,
		&time.StopTime)

	if err != nil {
		log.Println(err)
		log.Println("get failed")
		return nil, err
	}
	return &time, nil

}

func (storage *Storage) ShowAllTimers() ([]*model.Timer, error) {
	rows, err := storage.conn.Query(context.Background(), selectALLtimer)
	if err != nil {
		log.Println(err)

		return nil, err
	}
	defer rows.Close()

	var alltimers []*model.Timer
	for rows.Next() {
		timer := &model.Timer{}
		err := rows.Scan(&timer.ID, &timer.Tittle, &timer.StartTime, &timer.StopTime)
		if err != nil {
			log.Println(err) // error
			return nil, err
		}
		timer.Duration = timer.StopTime.Sub(timer.StartTime)
		alltimers = append(alltimers, timer)
	}
	return alltimers, nil
}

func (storage *Storage) DeletebyID(tittle string) error {
	_, err := storage.conn.Exec(context.Background(), delete, tittle)
	if err != nil {
		log.Println("delete failed", err)

		return err
	}
	return nil
}
func (storage *Storage) UpdateTimer(tittle string, stop time.Time) (*model.Timer, error) {
	var timer model.Timer
	// duration := timer.StopTime.Sub(timer.StartTime)
	_, err := storage.conn.Exec(context.Background(), updateTimer, tittle, stop)
	if err != nil {
		log.Println("update failed")
		return nil, err
	}
	log.Println("timer updated")
	return &timer, nil
}

// белые и серые айпи адреса
