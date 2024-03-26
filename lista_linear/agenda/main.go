package main

import (
	"agenda/agenda"
	"agenda/contato"

)

func main() {
	agenda := agenda.CriarAgenda()
	agenda.Imprimir()

	c1 := contato.Contato{}
	c1.AtribuirNome("Carlos")
	c2 := contato.Contato{}
	c2.AtribuirNome("Pedro")
	c3 := contato.Contato{}
	c3.AtribuirNome("Renata")

	agenda.InserirContato(c1)
	agenda.InserirContato(c2)
	agenda.InserirContato(c3)

	agenda.Imprimir()
}