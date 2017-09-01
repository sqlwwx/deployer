package main

import (
	"log"
	"os/exec"
)

func runScript(item *Item) (err error) {
	script := item.Script
	// out, err := exec.Command("bash", "-c", script).Output()
	cmd := exec.Command("bash", "-c", script)
	cmd.Dir = item.ScriptDir
	out, err := cmd.Output()

	if err != nil {
		log.Printf("Exec command failed: %s\n", err)
	}

	log.Printf("Run %s output: %s\n", script, string(out))
	return
}
