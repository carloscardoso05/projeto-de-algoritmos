package jogador_de_futebol

import "fmt"

type JogadorDeFutebol struct {
	nome         string
	jogos        uint
	gols         uint
	assistencias uint
}

func Criar(nome string, jogos, gols, assistencias uint) JogadorDeFutebol {
	j := JogadorDeFutebol{}
	j.AtribuirNome(nome)
	j.AtribuirJogos(jogos)
	j.AtribuirGols(gols)
	j.AtribuirAssistencias(assistencias)
	return j
}

func (j JogadorDeFutebol) Imprimir() {
	fmt.Printf("----- Jogador de Futebol -----\n")
	fmt.Printf("Nome: %s\n", j.nome)
	fmt.Printf("Jogos: %d\n", j.jogos)
	fmt.Printf("Gols: %d\n", j.gols)
	fmt.Printf("Assistências: %d\n", j.assistencias)
	if j.EhBom() {
		fmt.Printf("É bom?: Sim\n")
	} else {
		fmt.Printf("É bom?: Não\n")
	}
	fmt.Println()
}

// Getters
func (j *JogadorDeFutebol) Nome() string {
	return j.nome
}
func (j *JogadorDeFutebol) Jogos() uint {
	return j.jogos
}
func (j *JogadorDeFutebol) Gols() uint {
	return j.gols
}
func (j *JogadorDeFutebol) Assistencias() uint {
	return j.assistencias
}
func (j *JogadorDeFutebol) EhBom() bool {
	return (float64(j.gols+j.assistencias))/float64(j.jogos) >= 0.6
}

// Setters
func (j *JogadorDeFutebol) AtribuirNome(nome string) {
	j.nome = nome
}
func (j *JogadorDeFutebol) AtribuirJogos(jogos uint) {
	j.jogos = jogos
}
func (j *JogadorDeFutebol) AtribuirGols(gols uint) {
	j.gols = gols
}
func (j *JogadorDeFutebol) AtribuirAssistencias(assistencias uint) {
	j.assistencias = assistencias
}
