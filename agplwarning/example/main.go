package main

import (
	"fmt"

	"github.com/osu-datenshi/lib/agplwarning"
)

func main() {
	err := agplwarning.Warn("agplwarning", "AGPLWarning")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("License agreed")
}
