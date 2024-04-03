package contato

import (
	"fmt"
)

type Contato struct {
	nome     string
	endereco string
	numero   string
}

func Criar(nome string, numero string, endereco string) Contato {
	contato := Contato{}
	contato.AlterarNome(nome)
	contato.AlterarNumero(numero)
	contato.AlterarEndereco(endereco)
	return contato
}

func (c *Contato) Nome() string {
	return c.nome
}

func (c *Contato) Endereco() string {
	return c.endereco
}

func (c *Contato) Numero() string {
	return c.numero
}

func (c *Contato) AlterarNumero(numero string) {
	c.numero = numero
}

func (c *Contato) AlterarEndereco(endereco string) {
	c.endereco = endereco
}

func (c *Contato) AlterarNome(nome string) {
	c.nome = nome
}

func (c *Contato) Imprimir() {
	fmt.Printf("\n----- Contato -----\n")
	fmt.Printf("Nome: %s\n", c.nome)
	fmt.Printf("Número: %s\n", c.numero)
	fmt.Printf("Endereço: %s\n\n", c.endereco)
}
