package agenda

import "agenda/contato"

type Agenda struct {
	Inicio *NóContato
}

type NóContato struct {
	Valor   contato.Contato
	Proximo *NóContato
}