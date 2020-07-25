package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// ReadConfig lee el archivo de configuracion y devuelve dos slices:
// uno de recursos y otro de cantidad de consumers (gatherers) por recurso
func ReadConfig(filePath string) (resources []string, gatherers []int, fileError bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("File error:", r)
			fileError = true
		}
	}()
	fileError = false
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	check(err)

	for _, line := range lines {
		var resourceName string = line[0]
		var resourceGatherers string = line[1]
		parsedGatherers, err := strconv.Atoi(resourceGatherers)
		check(err)
		resources = append(resources, resourceName)
		gatherers = append(gatherers, parsedGatherers)
	}
	return
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
