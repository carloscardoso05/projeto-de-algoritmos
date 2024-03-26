package data

import "fmt"

func CriarData(dia uint, mes uint) Data {
	data := Data{}
	data.AtribuirDia(dia)
	data.AtribuirMes(mes)
	return data
}

func (d *Data) Dia() uint {
	return d.dia
}

func (d *Data) AtribuirDia(dia uint) {
	if dia == 0 || dia > 31 {
		fmt.Printf("Dia \"%d\" é inválido. Dia não foi atribuído\n", dia)
		return
	}
	d.dia = dia
}

func (d *Data) Mes() uint {
	return d.mes
}

func (d *Data) AtribuirMes(mes uint) {
	if mes == 0 || mes > 12 {
		fmt.Printf("Mês \"%d\" é inválido. Mês não foi atribuído\n", mes)
		return
	}
	d.mes = mes
}
