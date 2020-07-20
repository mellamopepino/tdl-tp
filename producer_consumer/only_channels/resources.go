package main

type Resources map[string]int

type addResources struct {
	material  string
	amount int
	ok chan bool
}

type buildResources struct {
	material  string
	build string
	amount  int
	ok chan bool
}

type getResources struct {
	resp chan Resources
}

type ResourcesHandler struct {
	addOp chan addResources
	buildOp chan buildResources
	getOp chan getResources
}


func MakeResourcesHandler() ResourcesHandler {
	return ResourcesHandler{
		addOp: make(chan addResources),
		buildOp: make(chan buildResources),
		getOp: make(chan getResources),
	}
}

func (r ResourcesHandler) Listen() {
	go func() {
		var resources = make(Resources)
		for {
			select {
			case a := <-r.addOp:
				resources[a.material] += a.amount
				a.ok <- true
			case b := <-r.buildOp:
				if(resources[b.material] - b.amount >= 0) {
					resources[b.material] -= b.amount
					resources[b.build]++
					b.ok <- true
				} else {
					b.ok <- false
				}
			case g := <-r.getOp:
				g.resp <- resources
			}
		}
	}()
}

func (r ResourcesHandler) Add(material string, amount int) (Resources, bool) {
	addOp := addResources{
		material: material,
		amount: amount,
		ok: make(chan bool),
	}
	r.addOp <- addOp

	/*
	* No podemos hacer `return r.GetAll(), <-addOp.ok`
	* Porque el `r.GetAll()` va a intentar mandar un mensaje
	* a el handler a través del channel `getOp`.
	* Esto mientras la función `Listen` del handler está
	* bloqueada esperando que alguien tome el mensaje que mandó
	* por el channel que para nosotros es `addOp.ok`.
	* Entonces el `Listen` no puede tomar el request de get hasta
	* que se desbloquee, pero solo se va a desbloquear cuando
	* hagamos `<-addOp.ok`. Que, esto último, se ejecutaría después
	* de llamar al `r.GetAll()`. Cuestión muere todo. :)
	*/
	ok := <-addOp.ok
	return r.GetAll(), ok
}

func (r ResourcesHandler) Build(build string, material string, amount int) (Resources, bool) {
	buildOp := buildResources{
		build: build,
		material: material,
		amount: amount,
		ok: make(chan bool),
	}
	r.buildOp <- buildOp

	ok := <-buildOp.ok
	return r.GetAll(), ok
}

func (r ResourcesHandler) GetAll() Resources {
	getOp := getResources{ resp: make(chan Resources) }
	r.getOp <- getOp
	return <-getOp.resp
}
