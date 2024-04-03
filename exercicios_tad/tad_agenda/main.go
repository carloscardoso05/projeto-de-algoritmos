package main

import (
	"agenda/agenda"
	"agenda/contato"
	"fmt"
)

func main() {
	agenda := &agenda.Agenda{}

	agenda.Adicionar(contato.Criar("Carlos", "1111-1111", "Rua 10"))
	agenda.Adicionar(contato.Criar("Pedro", "2222-2222", "Rua 20"))
	agenda.Adicionar(contato.Criar("Marcos", "3333-3333", "Rua 30"))

	nomeBusca := "Carlos"
	contato := agenda.BuscarPorNome(nomeBusca)
	contato.Imprimir()
	contato.AlterarNome("Carlos Vitor")
	contato.AlterarNumero("5555-5555")
	contato.AlterarEndereco("Avenida 10")
	contato.Imprimir()

	contato = agenda.BuscarPorNome(nomeBusca)
	if contato == nil {
		fmt.Printf("\nNenhum contato com o nome \"%s\" encontrado\n", nomeBusca)
	} else {
		contato.Imprimir()
	}
}
