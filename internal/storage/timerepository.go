package storage

import (
	"context"
	"fmt"

	"github.com/shegai01/timer/internal/model"
)

func (storage *Storage) CreateListTimer(ctx context.Context) error {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	_, err := storage.conn.Exec(ctx, createTable)
	if err != nil {
		return fmt.Errorf("storage.conn.Exec: %w", err)
	}

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
		return nil, fmt.Errorf("storage.conn.QueryRow: %w", err)
	}

	return &time, nil
}

func (storage *Storage) GetTimerByID(ctx context.Context, id int) (*model.Timer, error) {
	var time model.Timer
	storage.mu.Lock()
	defer storage.mu.Unlock()

	err := storage.conn.QueryRow(ctx, getTimer, id).Scan(
		&time.ID,
		&time.Title,
		&time.StartTime,
		&time.StopTime)

	if err != nil {
		return nil, fmt.Errorf("storage.conn.QueryRow: %w", err)
	}

	return &time, nil
}

func (storage *Storage) ShowAllTimers(ctx context.Context) ([]*model.Timer, error) {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	rows, err := storage.conn.Query(ctx, showAllTimers)
	if err != nil {
		return nil, fmt.Errorf("storage.conn.Query: %w", err)
	}
	defer rows.Close()

	var allTimers []*model.Timer
	for rows.Next() {
		timer := &model.Timer{}
		err := rows.Scan(&timer.ID, &timer.Title, &timer.StartTime, &timer.StopTime)
		if err != nil {
			return nil, fmt.Errorf("rows.Scan: %w", err)
		}

		allTimers = append(allTimers, timer)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("rows.Err: %w", err)
	}

	return allTimers, nil
}

func (storage *Storage) Delete(ctx context.Context, id int) error {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	_, err := storage.conn.Exec(ctx, delete, id)
	if err != nil {
		return fmt.Errorf("storage.conn.Exec: %w", err)
	}

	return nil
}

func (storage *Storage) StopTimer(ctx context.Context, id int) (*model.Timer, error) {
	var timer model.Timer
	storage.mu.Lock()
	defer storage.mu.Unlock()

	_, err := storage.conn.Exec(ctx, stopTimer, id)
	if err != nil {
		return nil, fmt.Errorf("storage.conn.Exec: %w", err)
	}

	return &timer, nil
}
