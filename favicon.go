package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	url2 "net/url"
	"os"
	"path/filepath"
)

var (
	IconDBPath = "/Library/Safari/Favicon Cache/favicons.db"
	iconDB     *sql.DB
	IconQuery  = `
SELECT i.uuid, i.url
FROM page_url p, icon_info i
WHERE p.uuid = i.uuid and
      p.url LIKE ?
LIMIT 1
`
)

//look up site favicon with url
func lookupFavicon(url string) string {
	if iconDB == nil {
		home, _ := os.UserHomeDir()
		iconDBPath := filepath.Join(home, IconDBPath)
		rt, err := sql.Open("sqlite3", iconDBPath)
		if err != nil {
			log.Println(err.Error())
		}
		iconDB = rt
	}

	if iconDB == nil {
		return ""
	}
	parsedURL, err := url2.Parse(url)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	query := fmt.Sprintf("%%%s%%", parsedURL.Host)
	rows, err := iconDB.Query(IconQuery, query)
	if err != nil {
		log.Println(err.Error())
	}

	for rows.Next() {
		var iconUrl sql.NullString
		var iconUUID sql.NullString
		var result string
		err := rows.Scan(&iconUUID, &iconUrl)
		if err != nil {
			return ""
		}
		if !iconUUID.Valid || !iconUrl.Valid {
			return ""
		}
		result = wf.CacheDir() + "/favicons/" + iconUUID.String
		_, err = os.Stat(result)
		if err != nil {
			if os.IsNotExist(err) {
				go downloadIcon(iconUrl.String, result)
			}
			return ""
		}
		return result
	}
	return ""
}

func downloadIcon(url, dest string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	file, err := os.Create(dest)
	if err != nil {
		log.Println(err.Error())
		return
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
