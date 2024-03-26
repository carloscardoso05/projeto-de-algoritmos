package jogador_de_futebol

import "fmt"

type JogadorDeFutebol struct {
	Nome         string
	Jogos        uint
	Gols         uint
	Assistencias uint
}

func (j *JogadorDeFutebol) Imprime() {
	fmt.Printf("\n----- Jogador de Futebol -----\n")
	fmt.Printf("Nome: %s\n", j.Nome)
	fmt.Printf("Jogos: %d\n", j.Jogos)
	fmt.Printf("Gols: %d\n", j.Gols)
	fmt.Printf("Assistências: %d\n", j.Assistencias)
	if j.EhBom() {
		fmt.Printf("É bom: Sim\n")
	} else {
		fmt.Printf("É bom: Não\n")
	}
}

func (j *JogadorDeFutebol) EhBom() bool {
	return (float64(j.Gols)+float64(j.Assistencias))/float64(j.Jogos) >= 0.6
}
