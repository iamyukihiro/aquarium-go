package main

import (
	"aquarium/internal/cmd"
)

func main() {
	err := cmd.Root.Execute()
	if err != nil {
		return
	}
}
