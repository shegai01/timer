package storage

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/shegai01/timer/internal/model"
)

func (storage *Storage) CreateListTimer(ctx context.Context) error {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	_, err := storage.conn.Exec(context.Background(), createTable)
	if err != nil {
		log.Printf("CreateTavle failed : %v\n", err)

		return fmt.Errorf("CreateListTimer failed: %w", err)
	}
	log.Println("database created", createTable)

	return nil
}
func (storage *Storage) CreateTimer(ctx context.Context, title string) (*model.Timer, error) {
	var time model.Timer
	storage.mu.Lock()
	defer storage.mu.Unlock()
	err := storage.conn.QueryRow(ctx, insertTimer, title).Scan(
		&time.ID,
		&time.Title,
		&time.StartTime,
		&time.StopTime)

	if err != nil {
		log.Printf("CreateTimer failed : %v\n", err)

		return nil, fmt.Errorf("CreateTimer failed %w", err)
	}

	return &time, nil
}
func (storage *Storage) GetTimerbyID(ctx context.Context, id int) (*model.Timer, error) {
	var time model.Timer
	storage.mu.Lock()
	defer storage.mu.Unlock()
	err := storage.conn.QueryRow(ctx, getTimer, id).Scan(
		&time.ID,
		&time.Title,
		&time.StartTime,
		&time.StopTime)

	if err != nil {
		log.Printf("GetTimerByID failed : %v\n", err)

		return nil, fmt.Errorf("GetTimer failed %w", err)
	}
	return &time, nil

}

func (storage *Storage) ShowAllTimers(ctx context.Context) ([]*model.Timer, error) {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	rows, err := storage.conn.Query(ctx, showALLtimer)
	if err != nil {
		log.Printf("ShowAllTimers failed : %v\n", err)

		return nil, fmt.Errorf("ShowAlltimers failed %w", err)
	}
	defer rows.Close()

	var alltimers []*model.Timer
	for rows.Next() {
		timer := &model.Timer{}
		err := rows.Scan(&timer.ID, &timer.Title, &timer.StartTime, &timer.StopTime)
		if err != nil {
			return nil, fmt.Errorf("ShowAlltimers failed %w", err)
		}
		alltimers = append(alltimers, timer)
	}
	return alltimers, nil
}

func (storage *Storage) Delete(ctx context.Context, id int) error {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	_, err := storage.conn.Exec(ctx, delete, id)
	if err != nil {
		log.Printf("Delete failed : %v\n", err)

		return fmt.Errorf("ShowAlltimers failed %w", err)

	}
	return nil
}
func (storage *Storage) StopTimer(ctx context.Context, id, stop time.Time) (*model.Timer, error) {
	var timer model.Timer
	storage.mu.Lock()
	defer storage.mu.Unlock()
	_, err := storage.conn.Exec(ctx, stopTimer, id, stop)
	if err != nil {
		return nil, fmt.Errorf("StopTimer failed %w", err)
	}

	log.Println("timer stoped")
	return &timer, nil
}
