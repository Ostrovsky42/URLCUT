package repo

import (
	"database/sql"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

type keysaver struct {
	repo *sql.DB
}

func (k keysaver) Save(url string, key string) error {
	insert := `insert into "CutUrl"("LongURL","Key") values($1, $2)`
	_, err := k.repo.Exec(insert, url, key)
	if err != nil {
		log.Print("Exec error", err)
		return err
	}
	return nil
}

func (k keysaver) Get(key string) (string, error) {
	query := `SELECT "LongURL" FROM "CutUrl" where "Key"=$1`
	row, err := k.repo.Query(query, key)
	defer row.Close()
	if err != nil {
		log.Print("Query error", err)
		return "", err
	}
	var longUrl string

	for row.Next() {
		err = row.Scan(&longUrl)
		if err != nil {
			log.Print("Scan row error", err)
			return "", err
		}
	}
	return longUrl, nil
}

func NewKeySaver(repo *sql.DB) *keysaver {
	return &keysaver{repo: repo}
}
