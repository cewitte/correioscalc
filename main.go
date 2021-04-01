// Copyright 2021 Carlos Eduardo Witte (@cewitte)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"

	"github.com/cewitte/correioscalc/correios"
	"github.com/cewitte/correioscalc/maptricks"
)

//  Teste é um objeto que representa o seguinte request HTTP GET': http://ws.correios.com.br/calculador/CalcPrecoPrazo.aspx?nCdEmpresa=08082650&sDsSenha=564321&sCepOrigem=70002900&sCepDestino=04547000&nVlPeso=1&nCdFormato=1&nVlComprimento=20&nVlAltura=20&nVlLargura=20&sCdMaoPropria=N&nVlValorDeclarado=0&sCdAvisoRecebimento=n&nCdServico=04510&nVlDiametro=0&StrRetorno=xml&nIndicaCalculo=3
var Teste = correios.CalcPrecoPrazo{
	NCdEmpresa:          "08082650",
	SDsSenha:            "564321",
	NCdServico:          "04014",
	SCepOrigem:          "70002900",
	SCepDestino:         "04547000",
	NVlPeso:             "1",
	NCdFormato:          1,
	NVlComprimento:      20,
	NVlAltura:           20,
	NVlLargura:          20,
	SCdMaoPropria:       "N",
	NVlValorDeclarado:   0,
	SCdAvisoRecebimento: "N",
}

// 1, 2, 3 Testando: som, som, som.
func main() {
	srv, err := Teste.PrecoPrazo()
	if err != nil {
		log.Panic(err)
	}

	// Vamos inverter o mapa de códigos para localizar o nome do serviço (Sedex, PAC, etc.) pelo código.
	codNome := maptricks.ReverseMap(correios.CodigosVigentes)

	// Os dados verdadeiramente úteis vêm dentro do slice srv.CServico.
	for _, v := range srv.CServico {
		fmt.Printf("Código: %s (%s)\n", v.Codigo, codNome[v.Codigo])
		fmt.Println("Valor: ", v.Valor)
		fmt.Println("PrazoEntrega: ", v.PrazoEntrega)
		fmt.Println("ValorSemAdicionais: ", v.ValorSemAdicionais)
		fmt.Println("ValorMaoPropria: ", v.ValorMaoPropria)
		fmt.Println("ValorAvisoRecebimento: ", v.ValorAvisoRecebimento)
		fmt.Println("ValorValorDeclarado: ", v.ValorValorDeclarado)
		fmt.Println("EntregaDomiciliar: ", v.EntregaDomiciliar)
		fmt.Println("EntregaSabado: ", v.EntregaSabado)
		fmt.Println("ObsFim: ", v.ObsFim)
		fmt.Println("Erro: ", v.MsgErro)
	}
}
