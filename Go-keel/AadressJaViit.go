package main

import "fmt"

type Isik struct {
	eesnimi  string
	perenimi string
}

func main() {
	var i Isik
	i = Isik{"Priit", "Parmakson"}

	v := &i
	vahetaNimed(v)

	fmt.Println(i.eesnimi, " ", i.perenimi)
}

func vahetaNimed(v *Isik) {
	v.eesnimi, v.perenimi = v.perenimi, v.eesnimi
}
