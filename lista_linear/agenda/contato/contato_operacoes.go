package contato

import "agenda/data"

const (
	MAX_TAM_NOME     = 40
	MAX_TAM_CELULAR  = 15
	MAX_TAM_TELEFONE = 15
	MAX_TAM_EMAIL    = 40
)

func (c *Contato) Nome() string {
	return c.nome
}

func (c *Contato) AtribuirNome(nome string) {
	max_len := min(MAX_TAM_NOME, len(nome))
	c.nome = nome[:max_len]
}

func (c *Contato) Telefone() string {
	return c.telefone
}

func (c *Contato) AtribuirTelefone(telefone string) {
	max_len := min(MAX_TAM_TELEFONE, len(telefone))
	c.telefone = telefone[:max_len]
}

func (c *Contato) Celular() string {
	return c.celular
}

func (c *Contato) AtribuirCelular(celular string) {
	max_len := min(MAX_TAM_CELULAR, len(celular))
	c.celular = celular[:max_len]
}

func (c *Contato) Email() string {
	return c.email
}

func (c *Contato) AtribuirEmail(email string) {
	max_len := min(MAX_TAM_EMAIL, len(email))
	c.email = email[:max_len]
}

func (c *Contato) DataAniversario() data.Data {
	return c.dataAniversario
}

func (c *Contato) AtribuirDataAniversario(data data.Data) {
	c.dataAniversario = data
}
