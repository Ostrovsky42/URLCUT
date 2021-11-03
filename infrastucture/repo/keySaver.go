package repo

import (
	"database/sql"
	"sync"
)

type  keysaver struct {
	repo sql.DB
	mu sync.RWMutex
}

func (k keysaver) Save(url string, key string) error {
	k.mu.RLock()
	return nil
}

func (k keysaver) Get(s string) (string, error) {
	panic("implement me")
}

func NewKeySaver(repo *sql.DB)*keysaver  {

}