package contato

import "agenda/data"

type Contato struct {
	nome            string
	telefone        string
	celular         string
	email           string
	dataAniversario data.Data
}