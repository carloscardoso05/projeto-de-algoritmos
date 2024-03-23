package jogador_de_futebol

import "fmt"

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
