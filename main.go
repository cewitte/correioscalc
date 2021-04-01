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
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/net/html/charset"
)

var (
	URL string = "http://ws.correios.com.br/calculador/CalcPrecoPrazo.aspx"
)

// CodigosVigentes implementa um map[string]string
// Atenção! Os códigos abaixo podem não estar mais válidos. Dados obtidos do "Manual técnico de integração web services ao Sistema Calculador de Preço e Prazo - SCPP.pdf" na versão 2.2 de 25/09/2019.
// Clientes com contrato devem consultar os códigos vigentes no contrato.
var CodigosVigentes = map[string]string{
	"SEDEX à vista":      "04014",
	"PAC à vista":        "04510",
	"SEDEX 12 (à vista)": "04782",
	"SEDEX 10 (à vista)": "04790",
	"SEDEX Hoje à vista": "04804",
}

// ReverseMap permite inverter um map[string]string, isto é, chave-valor é invertido para valor-chave. Use a função abaixo para reverter o mapa acima para chave: código e valor: nome. Pode ser útil para interpretar o retorno dos Correios.
func ReverseMap(m map[string]string) map[string]string {
	n := make(map[string]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}

// CalcPrecoPrazo encapsula os parâmetros de uma requisição para o WebService dos Correios que informa o preço, prazo e outras opção de uma encomenda a ser entregue pelos Correios.
type CalcPrecoPrazo struct {
	nCdEmpresa          string  // Sem contrato, enviar vazio.
	sDsSenha            string  // Sem contrato, enviar vazio.
	nCdServico          string  // Consultar códigos vigentes.
	sCepOrigem          string  // Sem hífen.
	sCepDestino         string  // Sem hífen.
	nVlPeso             string  // Quilogramas.
	nCdFormato          int     // 1-Caixa/Pcte; 2-Rolo/prisma. 3-Envel.
	nVlComprimento      float64 // Cm.
	nVlAltura           float64 // Cm. Se envelope, informar 0.
	nVlLargura          float64 // Cm
	nVlDiametro         float64 // Cm
	sCdMaoPropria       string  // S ou N (Sim ou Não).
	nVlValorDeclarado   float64 // Se não quiser, informar zero.
	sCdAvisoRecebimento string  // Se não quiser, informar zero.
}

//  Teste é um objeto que representa o seguinte request HTTP GET': http://ws.correios.com.br/calculador/CalcPrecoPrazo.aspx?nCdEmpresa=08082650&sDsSenha=564321&sCepOrigem=70002900&sCepDestino=04547000&nVlPeso=1&nCdFormato=1&nVlComprimento=20&nVlAltura=20&nVlLargura=20&sCdMaoPropria=N&nVlValorDeclarado=0&sCdAvisoRecebimento=n&nCdServico=04510&nVlDiametro=0&StrRetorno=xml&nIndicaCalculo=3
var Teste = CalcPrecoPrazo{
	nCdEmpresa:          "08082650",
	sDsSenha:            "564321",
	nCdServico:          "04014",
	sCepOrigem:          "70002900",
	sCepDestino:         "04547000",
	nVlPeso:             "1",
	nCdFormato:          1,
	nVlComprimento:      20,
	nVlAltura:           20,
	nVlLargura:          20,
	sCdMaoPropria:       "N",
	nVlValorDeclarado:   0,
	sCdAvisoRecebimento: "N",
}

// Servicos representa o XML de retorno dos correios.
type Servicos struct {
	XMLName  xml.Name `xml:"Servicos"`
	Text     string   `xml:",chardata"`
	CServico []struct {
		Text                  string `xml:",chardata"`
		Codigo                string `xml:"Codigo"`
		Valor                 string `xml:"Valor"`
		PrazoEntrega          string `xml:"PrazoEntrega"`
		ValorSemAdicionais    string `xml:"ValorSemAdicionais"`
		ValorMaoPropria       string `xml:"ValorMaoPropria"`
		ValorAvisoRecebimento string `xml:"ValorAvisoRecebimento"`
		ValorValorDeclarado   string `xml:"ValorValorDeclarado"`
		EntregaDomiciliar     string `xml:"EntregaDomiciliar"`
		EntregaSabado         string `xml:"EntregaSabado"`
		ObsFim                string `xml:"obsFim"`
		Erro                  string `xml:"Erro"`
		MsgErro               string `xml:"MsgErro"`
	} `xml:"cServico"`
}

// PrecoPrazo é o método que monta o HTTP GET Request, conecta a URL dos Correios, baixa o XML de resposta e faz o unmarshalling do XML de retorno para um objeto Servicos.
func (params *CalcPrecoPrazo) PrecoPrazo() (*Servicos, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	vals := req.URL.Query()
	vals.Add("nCDEmpresa", params.nCdEmpresa)
	vals.Add("sDsSenha", params.sDsSenha)
	vals.Add("nCdServico", params.nCdServico)
	vals.Add("sCepOrigem", params.sCepOrigem)
	vals.Add("sCepDestino", params.sCepDestino)
	vals.Add("nVlPeso", params.nVlPeso)
	vals.Add("nCdFormato", strconv.Itoa(params.nCdFormato))
	vals.Add("nVlComprimento", fmt.Sprintf("%.2f", params.nVlComprimento))
	vals.Add("nVlAltura", fmt.Sprintf("%.2f", params.nVlAltura))
	vals.Add("nVlLargura", fmt.Sprintf("%.2f", params.nVlLargura))
	vals.Add("nVlDiametro", fmt.Sprintf("%.2f", params.nVlDiametro))
	vals.Add("sCdMaoPropria", params.sCdMaoPropria)
	vals.Add("nVlValorDeclarado", fmt.Sprintf("%.2f", params.nVlValorDeclarado))
	vals.Add("sCdAvisoRecebimento", params.sCdAvisoRecebimento)
	vals.Add("StrRetorno", "xml")
	vals.Add("nIndicaCalculo", "3")

	req.URL.RawQuery = vals.Encode()
	req.Header.Add("Accept", "application/xml")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var srv Servicos
	xmlData, err := ioutil.ReadAll(resp.Body)
	reader := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&srv)
	if err != nil {
		return nil, err
	}

	return &srv, nil
}

// 1, 2, 3 Testando: som, som, som.
func main() {
	srv, err := Teste.PrecoPrazo()
	if err != nil {
		log.Panic(err)
	}

	// Vamos inverter o mapa de códigos para localizar o nome do serviço (Sedex, PAC, etc.) pelo código.
	codNome := ReverseMap(CodigosVigentes)

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
