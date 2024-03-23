package time

import (
	"errors"
	"fmt"
	"time_futebol/jogador_de_futebol"
)

type Time struct {
	Nome      string
	Treinador string
	Vitorias  uint
	Empates   uint
	Derrotas  uint
	Jogadores []jogador_de_futebol.JogadorDeFutebol
}

func CriarTime(nome string, jogadores []jogador_de_futebol.JogadorDeFutebol) (*Time, error) {
	if len(jogadores) < 5 {
		return nil, errors.New("a quantidade de jogadores não pode ser menor que 5")
	}
	if len(nome) == 0 {
		return nil, errors.New("o nome do time não pode ser vazio")
	}

	time := &Time{Nome: nome, Jogadores: jogadores}
	return time, nil
}

func (t *Time) AtribuirJogadores(jogadores []jogador_de_futebol.JogadorDeFutebol) error {
	if len(jogadores) < 5 {
		return errors.New("a quantidade de jogadores não pode ser menor que 5")
	}
	t.Jogadores = jogadores
	return nil
}

func (t *Time) Pontuacao() uint {
	return t.Vitorias*3 + t.Empates
}

func (t *Time) Imprimir() {
	fmt.Printf("\n----- Time -----\n")
	fmt.Printf("Nome: %s\n", t.Nome)
	fmt.Printf("Treinador: %s\n", t.Treinador)
	fmt.Printf("Vitorias: %d\n", t.Vitorias)
	fmt.Printf("Empates: %d\n", t.Empates)
	fmt.Printf("Derrotas: %d\n", t.Derrotas)
	fmt.Print("Jogadores: \n")
	for _, jogador := range t.Jogadores {
		jogador.Imprimir()
	}
}
