package time

import (
	"errors"
	"fmt"
	"time_futebol/jogador_de_futebol"
)

type Time struct {
	nome      string
	treinador string
	vitorias  uint
	empates   uint
	derrotas  uint
	jogadores []jogador_de_futebol.JogadorDeFutebol
}

func Criar(nome string, jogadores []jogador_de_futebol.JogadorDeFutebol) (*Time, error) {
	if len(jogadores) < 5 {
		return nil, errors.New("a quantidade de jogadores não pode ser menor que 5")
	}

	time := &Time{nome: nome, jogadores: jogadores}
	return time, nil
}

func (t *Time) Imprimir() {
	fmt.Printf("\n----- Time -----\n")
	fmt.Printf("Nome: %s\n", t.nome)
	fmt.Printf("Treinador: %s\n", t.treinador)
	fmt.Printf("Vitorias: %d\n", t.vitorias)
	fmt.Printf("Empates: %d\n", t.empates)
	fmt.Printf("Derrotas: %d\n", t.derrotas)
	fmt.Print("Jogadores: \n")
	for _, jogador := range t.jogadores {
		jogador.Imprimir()
	}
}

// Getters
func (t *Time) Nome() string {
	return t.nome
}
func (t *Time) Treinador() string {
	return t.treinador
}
func (t *Time) Vitorias() uint {
	return t.vitorias
}
func (t *Time) Empates() uint {
	return t.empates
}
func (t *Time) Derrotas() uint {
	return t.derrotas
}
func (t *Time) Jogadores() []jogador_de_futebol.JogadorDeFutebol {
	return t.jogadores
}
func (t *Time) Pontuacao() uint {
	return t.vitorias*3 + t.empates
}

// Setters
func (t *Time) AtribuirNome(nome string) {
	t.nome = nome
}
func (t *Time) AtribuirTreinador(treinador string) {
	t.treinador = treinador
}
func (t *Time) AtribuirVitorias(vitorias uint) {
	t.vitorias = vitorias
}
func (t *Time) AtribuirEmpates(empates uint) {
	t.empates = empates
}
func (t *Time) AtribuirDerrotas(derrotas uint) {
	t.derrotas = derrotas
}
func (t *Time) AtribuirJogadores(jogadores []jogador_de_futebol.JogadorDeFutebol) error {
	if len(jogadores) < 5 {
		return errors.New("a quantidade de jogadores não pode ser menor que 5")
	}
	t.jogadores = jogadores
	return nil
}