package jogador_de_futebol

import "testing"

func TestJogadorEhBom(t *testing.T) {
	j := &JogadorDeFutebol{
		nome:         "Pedro",
		jogos:        10,
		gols:         10,
		assistencias: 5,
	}

	j.Imprime()
	
	if !j.EhBom() {
		t.Errorf("Esperava jogador é bom, recebeu não é bom")
	}
	
	j = &JogadorDeFutebol{
		nome:         "Pedro",
		jogos:        10,
		gols:         0,
		assistencias: 6,
	}

	j.Imprime()
	
	if !j.EhBom() {
		t.Errorf("Esperava jogador é bom, recebeu não é bom")
	}
	
	j = &JogadorDeFutebol{
		nome:         "Pedro",
		jogos:        10,
		gols:         2,
		assistencias: 2,
	}

	j.Imprime()
	
	if j.EhBom() {
		t.Errorf("Esperava jogador não é bom, recebeu é bom")
	}
	
	j = &JogadorDeFutebol{
		nome:         "Pedro",
		jogos:        10,
		gols:         0,
		assistencias: 0,
	}
	
	j.Imprime()

	if j.EhBom() {
		t.Errorf("Esperava jogador não é bom, recebeu é bom")
	}
}
