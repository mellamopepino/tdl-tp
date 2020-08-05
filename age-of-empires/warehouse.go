package main

import (
	"sync"
)

// Resources representa los recursos guardados. Usamos un mapa que mapea strings a ints
type Resources map[string]int

// Warehouse estructura que guarda los recursos
type Warehouse struct {
	resources Resources
	m         sync.Mutex
	done      bool
}

// MakeWarehouse es una función auxiliar que crea warehouses
func MakeWarehouse() *Warehouse {
	warehouse := Warehouse{Resources{}, sync.Mutex{}, false}
	return &warehouse
}

// Add agrega recursos al warehouse
func (warehouse *Warehouse) Add(material string, amount int) {
	defer warehouse.m.Unlock()
	warehouse.m.Lock()
	warehouse.resources[material] += amount
}

// Use resta recursos al warehouse (si están disponibles)
func (warehouse *Warehouse) Use(materials map[string]int) bool {
	defer warehouse.m.Unlock()
	warehouse.m.Lock()
	ok := true
	for material, amount := range materials {
		if warehouse.resources[material] < amount {
			ok = false
			break
		}
	}
	if ok {
		for material, amount := range materials {
			warehouse.resources[material] -= amount
		}
	}
	return ok
}

// GetAll devuelve los recursos del warehouse
func (warehouse *Warehouse) GetAll() Resources {
	return warehouse.resources
}
