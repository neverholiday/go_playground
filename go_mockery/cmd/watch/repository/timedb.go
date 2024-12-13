package repository

import "time"

type WatchRepo struct {
}

func NewWatchRepo() *WatchRepo {
	return &WatchRepo{}
}

func (w *WatchRepo) GetCurrentTime() (string, error) {
	t := time.Now()
	return t.Format(time.TimeOnly), nil
}
