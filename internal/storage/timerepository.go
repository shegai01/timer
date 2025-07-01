package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/shegai01/timer/internal/model"
)

func (storage *Storage) CreateTimer(tittle string) (*model.Timer, error) {
	var time model.Timer
	err := storage.conn.QueryRow(context.Background(), insertTimer, tittle).Scan(
		&time.ID,
		&time.Tittle,
		&time.StartTime,
		&time.FinishTime)
	if err != nil {
		log.Println("can't creating without name")
		return nil, err
	}
	return &time, nil
}
func (storage *Storage) GetTimerbyID(tittle string) (*model.Timer, error) {
	var time model.Timer
	err := storage.conn.QueryRow(context.Background(), selectTimer, tittle).Scan(
		&time.ID,
		&time.Tittle,
		&time.StartTime,
		&time.FinishTime)

	if err != nil {
		return nil, err
	}
	return &time, nil

}

func (storage *Storage) ShowAllTimers() ([]model.Timer, error) {
	rows, err := storage.conn.Query(context.Background(), selectTimer)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var alltimers []model.Timer
	for rows.Next() {
		var timer model.Timer
		err := rows.Scan(&timer.ID, &timer.Tittle, &timer.StartTime, &timer.FinishTime)
		if err != nil {
			return nil, err
		}
		alltimers = append(alltimers, timer)
	}
	return alltimers, nil
}

func (storage *Storage) DeletebyID(id int) error {
	_, err := storage.conn.Exec(context.Background(), delete)
	if err != nil {
		fmt.Println("delete failed", err)
		return err
	}
	return nil
}
func (storage *Storage) UpdateTimer(start, finish time.Time, name string) (*model.Timer, error) {
	_, err := storage.conn.Exec(context.Background(), updateTimer, start, finish, name)
	if err != nil {
		log.Println("updating failed")
		return nil, err
	}
	return nil, nil
}

//белые и серые айпи адреса
