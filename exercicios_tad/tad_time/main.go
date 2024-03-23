package main

import (
	"fmt"
	"time_futebol/jogador_de_futebol"
	"time_futebol/time"
)

func main() {
	var jogadores []jogador_de_futebol.JogadorDeFutebol

	for i := 0; i < 5; i++ {
		jogadores = append(jogadores, jogador_de_futebol.JogadorDeFutebol{Nome: fmt.Sprintf("Jogador%d", i+1)})
	}

	time, err := time.CriarTime("Remo", jogadores)
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Imprimir()
}
