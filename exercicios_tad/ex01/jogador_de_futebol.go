package jogador_de_futebol

import (
	"errors"
	"fmt"
)

type JogadorDeFutebol struct {
	nome         string
	jogos        uint
	gols         uint
	assistencias uint
}

func (j *JogadorDeFutebol) Imprime() {
	fmt.Printf("Jogador de Futebol\n")
	fmt.Printf("	Nome: %s\n", j.nome)
	fmt.Printf("	Jogos: %d\n", j.jogos)
	fmt.Printf("	Gols: %d\n", j.gols)
	fmt.Printf("	Assistências: %d\n", j.assistencias)
	fmt.Printf("	É bom: %t\n\n", j.EhBom())
}

func (j *JogadorDeFutebol) EhBom() bool {
	return (float64(j.gols)+float64(j.assistencias))/float64(j.jogos) >= 0.6
}

func (j *JogadorDeFutebol) AtribuirNome(nome string) error {
	if len(nome) < 3 {
		msg := fmt.Sprintf("o nome \"%s\" deve conter pelo menos 3 caracteres", nome)
		return errors.New(msg)
	}
	j.nome = nome
	return nil
}

func (j *JogadorDeFutebol) AtribuirJogos(jogos uint) {
	j.jogos = jogos

}

func (j *JogadorDeFutebol) AtribuirPartidas(gols uint) {
	j.gols = gols
}

func (j *JogadorDeFutebol) AtribuirAssistencias(assistencias uint) {
	j.assistencias = assistencias
}
