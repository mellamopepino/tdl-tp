package main

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

type Warehouse struct {
	addOp chan addResources
	useOp chan useResources
	getOp chan getResources
	done  bool
}

func MakeWarehouse() *Warehouse {
	warehouse := Warehouse{
		addOp: make(chan addResources),
		useOp: make(chan useResources),
		getOp: make(chan getResources),
		done:  false,
	}
	return &warehouse
}

// Goroutine dueña del estado del warehouse
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
			}
		}
	}()
}

// Function receiver para agregar cosas al warehouse
func (warehouse *Warehouse) Add(material string, amount int) {
	addOp := addResources{
		material: material,
		amount:   amount,
	}
	warehouse.addOp <- addOp
}

// Function receiver para usar recursos del warehouse
func (warehouse *Warehouse) Use(materials map[string]int) bool {
	useOp := useResources{
		materials: materials,
		ok:        make(chan bool),
	}
	warehouse.useOp <- useOp
	return <-useOp.ok
}
