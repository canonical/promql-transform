package main

import (
	cmd "github.com/canonical/promql-transform/cmd/root"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
