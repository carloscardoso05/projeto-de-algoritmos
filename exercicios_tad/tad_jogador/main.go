package main

import "jogador_de_futebol/jogador_de_futebol"

func main() {
	jogador1 := jogador_de_futebol.JogadorDeFutebol{
		Nome:         "Carlos",
		Jogos:        10,
		Gols:         2,
		Assistencias: 3,
	}

	jogador1.Imprime()

	jogador2 := jogador_de_futebol.JogadorDeFutebol{
		Nome:         "Pedro",
		Jogos:        15,
		Gols:         10,
		Assistencias: 4,
	}

	jogador2.Imprime()
}
