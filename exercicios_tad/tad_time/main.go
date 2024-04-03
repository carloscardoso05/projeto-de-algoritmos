package main

import (
	"fmt"
	"time_futebol/jogador_de_futebol"
	"time_futebol/time"
)

func main() {
	var jogadores []jogador_de_futebol.JogadorDeFutebol

	for i := 1; i <= 5; i++ {
		nome := fmt.Sprintf("Jogador %d", i)
		jogadores = append(jogadores, jogador_de_futebol.Criar(nome, 0, 0, 0))
	}

	time, err := time.Criar("Remo", jogadores)
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Imprimir()
}
