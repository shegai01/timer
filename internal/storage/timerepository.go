package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/shegai01/timer/internal/model"
)

type TimeTracker interface {
	CreateList() error
	CreateTimer() (*model.Timer, error)
	ShowAllTimers() ([]model.Timer, error)
	DeletebyID(id int) error
	UpdateTimer(start, finish time.Time, name string)
}

func (tdb *Storage) CreateList() error {
	createTebale, err := tdb.conn.Exec(context.Background(), createTable)
	if err != nil {
		log.Println("create table failed", err)
		return err
	}
	log.Println("database created", createTebale)
	return nil

}

func (tdb *Storage) CreateTimer(tittle string) (*model.Timer, error) {
	var time model.Timer
	err := tdb.conn.QueryRow(context.Background(), insertTimer, tittle).Scan(
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
func (tdb *Storage) ShowAllTimers() ([]model.Timer, error) {
	rows, err := tdb.conn.Query(context.Background(), selectTimer)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()
	var alltimers []model.Timer
	for rows.Next() {
		var timer model.Timer
		err := rows.Scan(&timer.ID, &timer.StartTime, &timer.FinishTime)
		if err != nil {
			return nil, err
		}
		alltimers = append(alltimers, timer)
	}
	return alltimers, nil
}

func (tdb *Storage) DeletebyID(id int) error {
	_, err := tdb.conn.Exec(context.Background(), delete)
	if err != nil {
		fmt.Println("delete failed", err)
		return err
	}
	return nil
}
func (tdb *Storage) UpdateTimer(start, finish time.Time, name string) (*model.Timer, error) {
	_, err := tdb.conn.Exec(context.Background(), "UPDATE timer set name=$1, start_time=$2, finish_time=$3", start, finish, name)
	if err != nil {
		log.Println("updating failed")
		return nil, err
	}
	return nil, nil
}

//белые и серые айпи адреса
