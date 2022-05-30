package main

import (
	aw "github.com/deanishe/awgo"
	"log"
	"os"
	"os/exec"
)

var iconAvailable = &aw.Icon{Value: "update-available.png"}

func doUpdate() error {
	log.Println("Checking for update...")
	return wf.CheckForUpdate()
}

func checkForUpdate() error {
	if !wf.UpdateCheckDue() || wf.IsRunning("update") {
		return nil
	}
	cmd := exec.Command(os.Args[0], "update")
	return wf.RunInBackground("update", cmd)
}

func showUpdateStatus() {
	if query != "" {
		return
	}

	if wf.UpdateAvailable() {
		wf.Configure(aw.SuppressUIDs(true))
		log.Println("Update available!")
		wf.NewItem("An update is available!").
			Subtitle("⇥ or ↩ to install update").
			Valid(false).
			Autocomplete("workflow:update").
			Icon(iconAvailable)
	}
}
