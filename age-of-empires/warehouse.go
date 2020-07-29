package main

// Resources representa los recursos guardados. Usamos un mapa que mapea strings a ints
type Resources map[string]int

type addResources struct {
	material string
	amount   int
}

type useResources struct {
	materials map[string]int
	ok        chan bool
}

type getResources struct {
	resp chan Resources
}

// Warehouse es un struct con métodos asociados que facilitan compartir recursos
type Warehouse struct {
	addOp chan addResources
	useOp chan useResources
	getOp chan getResources
	done  bool
}

// MakeWarehouse es una función auxiliar que crea warehouses
func MakeWarehouse() *Warehouse {
	warehouse := Warehouse{
		addOp: make(chan addResources),
		useOp: make(chan useResources),
		getOp: make(chan getResources),
		done:  false,
	}
	return &warehouse
}

// Listen es la goroutine stateful que efectivamente contiene los recursos compartidos
func (warehouse *Warehouse) Listen() {
	go func() {
		var resources = make(Resources)
		for {
			select {
			case addOp := <-warehouse.addOp:
				resources[addOp.material] += addOp.amount
			case useOp := <-warehouse.useOp:
				ok := true
				for material, amount := range useOp.materials {
					if resources[material] < amount {
						ok = false
						break
					}
				}
				if ok {
					for material, amount := range useOp.materials {
						resources[material] -= amount
					}
				}
				useOp.ok <- ok
			case getOp := <-warehouse.getOp:
				getOp.resp <- resources
			}
		}
	}()
}

// Add agrega recursos al warehouse
func (warehouse *Warehouse) Add(material string, amount int) {
	addOp := addResources{
		material: material,
		amount:   amount,
	}
	warehouse.addOp <- addOp
}

// Use resta recursos al warehouse (si están disponibles)
func (warehouse *Warehouse) Use(materials map[string]int) bool {
	useOp := useResources{
		materials: materials,
		ok:        make(chan bool),
	}
	warehouse.useOp <- useOp
	return <-useOp.ok
}

// GetAll devuelve los recursos del warehouse
func (warehouse *Warehouse) GetAll() Resources {
	getOp := getResources{resp: make(chan Resources)}
	warehouse.getOp <- getOp
	return <-getOp.resp
}
