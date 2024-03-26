package agenda

import (
	"agenda/contato"
	"fmt"
)

type Agenda struct {
	Inicio *NóContato
}

type NóContato struct {
	Valor   contato.Contato
	Proximo *NóContato
}

func CriarAgenda() Agenda {
	return Agenda{}
}

func (a *Agenda) InserirContato(contato contato.Contato) {
	if a.Inicio == nil {
		nó := NóContato{
			Valor: contato,
		}
		a.Inicio = &nó
		return
	}

	cursor := a.Inicio
	for cursor.Proximo != nil {
		cursor = cursor.Proximo
	}

	nó := NóContato{
		Valor: contato,
	}
	cursor.Proximo = &nó
}

func (a *Agenda) Imprimir() {
	if a.Inicio == nil {
		fmt.Printf("Agenda vazia\n")
		return
	}
	
	fmt.Printf("\n%+v\n", a.Inicio.Valor)
	
	cursor := a.Inicio
	for cursor.Proximo != nil {
		cursor = cursor.Proximo
		fmt.Printf("\n%+v\n", cursor.Valor)
	}
}
