package main

import (
	"encoding/json"
	"os"
)

// ReadResourcesConfig lee el archivo de configuracion y devuelve dos slices:
// uno de recursos y otro de cantidad de consumers (gatherers) por recurso
func ReadResourcesConfig(filePath string) (resources []Resource) {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	err = json.NewDecoder(file).Decode(&resources)
	check(err)

	return
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// ReadWeaponsConfig para leer la configuracion de las armas
func ReadWeaponsConfig(filePath string) (weapons []Weapon) {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	err = json.NewDecoder(file).Decode(&weapons)
	check(err)

	return
}
