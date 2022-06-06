package main

import (
	"database/sql"
	"fmt"
	aw "github.com/deanishe/awgo"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
LIMIT 50
`
)

func searchHistory() error {
	showUpdateStatus()

	home, _ := os.UserHomeDir()
	dbPath := filepath.Join(home, DBPath)
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	query := fmt.Sprintf("%%%s%%", strings.Join(strings.Split(os.Args[3], " "), "%"))
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
		if !title.Valid || len(title.String) == 0 {
			title.String = url.String
		}
		item := wf.NewItem(title.String).
			Valid(true).
			UID(strconv.Itoa(id)).
			Subtitle(url.String).
			Arg(url.String)
		iconPath := lookupFavicon(url.String)
		if iconPath != "" {
			item.Icon(&aw.Icon{
				Value: iconPath,
			})
		}
	}
	wf.WarnEmpty("No matching history found", "Try another?")
	wf.SendFeedback()
	return nil
}
