package main

import (
	"aquarium/cmd"
)

func main() {
	err := cmd.Root.Execute()
	if err != nil {
		return
	}
}
