package main

import "jogador_de_futebol/jogador_de_futebol"

func main() {
	j1 := jogador_de_futebol.Criar("Carlos", 10, 2, 4)
	j1.Imprimir()

	j2 := jogador_de_futebol.Criar("Pedro", 15, 10, 4)
	j2.Imprimir()

	j3 := jogador_de_futebol.Criar("João", 14, 11, 2)
	j3.AtribuirNome("João Pedro")
	j3.AtribuirJogos(17)
	j3.AtribuirGols(15)
	j3.AtribuirAssistencias(3)
	j3.Imprimir()
}
