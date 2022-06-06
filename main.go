package main

import (
	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
	"log"
	"os"
	"path/filepath"
)

var (
	repo  = "addozhang/alfred-safari-toolkit"
	query string
	wf    *aw.Workflow
)

func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))
	//init icon cache folder
	iconPath := filepath.Join(wf.CacheDir(), "/favicons")
	if _, err := os.Stat(iconPath); os.IsNotExist(err) {
		os.Mkdir(iconPath, os.ModePerm)
	}
}

func main() {
	wf.Run(run)
}

func run() {
	if err := checkForUpdate(); err != nil {
		log.Printf("Error starting update check: %s", err)
	}
	args := wf.Args()
	switch args[0] {
	case "history":
		searchHistory()
		//TODO: more action such as tab searching
	}
}
