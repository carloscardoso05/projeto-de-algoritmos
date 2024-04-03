package numero_complexo

import "fmt"

type NumeroComplexo struct {
	Real int
	Imag int
}

func Criar(real, imag int) NumeroComplexo {
	return NumeroComplexo{
		Real: real,
		Imag: imag,
	}
}
func (n NumeroComplexo) Imprimir() {
	fmt.Printf("%d + %di\n", n.Real, n.Imag)
}
func (n NumeroComplexo) EhReal() bool {
	return n.Imag == 0
}
func Somar(n1, n2 NumeroComplexo) NumeroComplexo {
	return NumeroComplexo{
		Real: n1.Real + n2.Real,
		Imag: n1.Imag + n2.Imag,
	}
}
func (n1 NumeroComplexo) CopiarPara(n2 *NumeroComplexo) {
	n2.Real = n1.Real
	n2.Imag = n1.Imag
}
