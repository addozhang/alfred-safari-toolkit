package main

import (
	aw "github.com/deanishe/awgo"
)

var (
	//repo  = "addozhang/alfred-safari-toolkit"
	query string
	wf    *aw.Workflow
)

func init() {
	wf = aw.New()
}

func main() {
	wf.Run(run)
}

func run() {
	search()
}
