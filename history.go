package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DBPath = "/Library/Safari/History.db"
	QUERY  = `
SELECT history_items.id, title, url
FROM history_items
INNER JOIN history_visits
ON history_visits.history_item = history_items.id
WHERE url LIKE ? OR title LIKE ?
GROUP BY url
ORDER BY visit_time DESC
`
)

func search() error {
	home, _ := os.UserHomeDir()
	dbPath := filepath.Join(home, DBPath)
	cachePath := filepath.Join(wf.CacheDir(), "history.db")
	if err := cache(dbPath, cachePath); err != nil {
		return err
	}

	db, err := sql.Open("sqlite3", cachePath)
	if err != nil {
		return err
	}
	defer db.Close()
	query := fmt.Sprintf("%%%s%%", os.Args[3])
	rows, err := db.Query(QUERY, query, query)
	if err != nil {
		return err
	}

	for rows.Next() {
		var id int
		var title, url sql.NullString
		err := rows.Scan(&id, &title, &url)
		if err != nil {
			return err
		}
		if !title.Valid {
			title.String = url.String
		}
		wf.NewItem(title.String).
			Valid(true).
			UID(strconv.Itoa(id)).
			Subtitle(url.String).
			Arg(url.String)
	}
	if os.Args[3] != "" {
		wf.Filter(os.Args[3])
	}
	wf.WarnEmpty("No matching history found", "Try another?")
	wf.SendFeedback()
	return nil
}

func cache(src, dst string) error {
	if file, err := os.Stat(dst); err != nil {
		return err
	} else if time.Now().After(file.ModTime().Add(time.Second * 60)) {
		return nil
	}
	source, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dst, source, 0644)
	return err
}
