package lista_encadeada

type ListEncadeada[T any] struct {
	Inicio *Nó[T]
}

type Nó[T any] struct {
	Valor   T
	Proximo *Nó[T]
}
