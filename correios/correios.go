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

package correios

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/cewitte/correioscalc/maptricks"
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

var Formatos = map[int]string{
	1: "Formato caixa/pacote",
	2: "Formato rolo/prisma",
	3: "Envelope",
}

var CodigosVigentesPorCodigo = maptricks.ReverseMap(CodigosVigentes)

// CalcPrecoPrazo encapsula os parâmetros de uma requisição para o WebService dos Correios que informa o preço, prazo e outras opção de uma encomenda a ser entregue pelos Correios.
type CalcPrecoPrazo struct {
	NCdEmpresa          string  // Sem contrato, enviar vazio.
	SDsSenha            string  // Sem contrato, enviar vazio.
	NCdServico          string  // Consultar códigos vigentes.
	SCepOrigem          string  // Sem hífen.
	SCepDestino         string  // Sem hífen.
	NVlPeso             string  // Quilogramas.
	NCdFormato          int     // 1-Caixa/Pcte; 2-Rolo/prisma. 3-Envel.
	NVlComprimento      float64 // Cm.
	NVlAltura           float64 // Cm. Se envelope, informar 0.
	NVlLargura          float64 // Cm
	NVlDiametro         float64 // Cm
	SCdMaoPropria       string  // S ou N (Sim ou Não).
	NVlValorDeclarado   float64 // Se não quiser, informar zero.
	SCdAvisoRecebimento string  // S ou N (Sim ou Não).
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
	vals.Add("nCDEmpresa", params.NCdEmpresa)
	vals.Add("sDsSenha", params.SDsSenha)
	vals.Add("nCdServico", params.NCdServico)
	vals.Add("sCepOrigem", params.SCepOrigem)
	vals.Add("sCepDestino", params.SCepDestino)
	vals.Add("nVlPeso", params.NVlPeso)
	vals.Add("nCdFormato", strconv.Itoa(params.NCdFormato))
	vals.Add("nVlComprimento", fmt.Sprintf("%.2f", params.NVlComprimento))
	vals.Add("nVlAltura", fmt.Sprintf("%.2f", params.NVlAltura))
	vals.Add("nVlLargura", fmt.Sprintf("%.2f", params.NVlLargura))
	vals.Add("nVlDiametro", fmt.Sprintf("%.2f", params.NVlDiametro))
	vals.Add("sCdMaoPropria", params.SCdMaoPropria)
	vals.Add("nVlValorDeclarado", fmt.Sprintf("%.2f", params.NVlValorDeclarado))
	vals.Add("sCdAvisoRecebimento", params.SCdAvisoRecebimento)
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
	if err != nil {
		return nil, err
	}
	reader := bytes.NewReader(xmlData)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&srv)
	if err != nil {
		return nil, err
	}

	return &srv, nil
}
