package main

import (
	"encoding/json"
	"os"
)

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

func ReadWeaponsConfig(filePath string) (weapons []Weapon) {
	file, err := os.Open(filePath)
	check(err)
	defer file.Close()

	err = json.NewDecoder(file).Decode(&weapons)
	check(err)

	return
}
