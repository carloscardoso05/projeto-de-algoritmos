package agenda

import "agenda/contato"

type Agenda struct {
	contatos []contato.Contato
}

func (a *Agenda) BuscarPorNome(nome string) *contato.Contato {
	for i, contato := range a.contatos {
		if contato.Nome() == nome {
			return &a.contatos[i]
		}
	}
	return nil
}

func (a *Agenda) Adicionar(contato contato.Contato) {
	a.contatos = append(a.contatos, contato)
}

func (a *Agenda) Listar() {
	for _, contato := range a.contatos {
		contato.Imprimir()
	}
}
