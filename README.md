# correioscalc 

Implementação Go (Golang) de um web service dos Correios, que permite obter o preço e o prazo de remessas (A Go/Golang implementation of a Correios web service, which allows you to obtain the price and delivery term).

### Para clonar o repositório e rodar em sua máquina local (to clone the repo and run in your local machine):

1. Clone o repositório em sua máquina local (clone the repo to your local machine):
```
git clone https://github.com/cewitte/correioscalc.git
```

2. Execute o código (run the code):
```
go run main.go
```

3. Siga o seu coração e faça o Bem (follow your heart and do Good).


### Para usar o código em seu projeto Go existente (using the code in your existing Go project):
```
go get -u github.com/cewitte/correioscalc
```


### Notas em Português do Brasil

Esta é uma implementação Go (Golang) do serviço web dos Correios que permite obter o preço e o prazo de envio de um pacote com base em alguns parâmetros.

Essa implementação foi baseada no **"Manual técnico de integração web services ao Sistema Calculador de Preço e Prazo - SCPP.pdf"** na versão 2.2 de 25/09/2019. Meu último teste com sucesso foi realizado na noite de 31 de março de 2021, portanto acredito que a documentação continua (e continuará) válida.

O manual SCPP está disponível aqui no diretório "manual-correios", apenas em português do Brasil.

Todos os comentários e a maioria dos nomes de variáveis estão em português brasileiro. Normalmente prefiro fazer tudo em inglês, mas dada a especificidade desta implementação, a decisão me parece justificada.

#### Contribuições

Suas contribuições são bem vindas. Que tal um Pull Request?

#### Licenciamento

The Apache 2.0 licence. Por favor, veja o arquivo da licença.


### Notes in U.S. English

This is a Go (Golang) implementation of the Brazilian Postal Service's (aka "Correios") web service which allows you to retrieve the price and ETA of a package shipment based on some parameters.

This implementation was based on the **"Technical manual for integrating web services into the Price and Term Calculator System - PTCS.pdf"** (title loosely translated) version 2.2 dated 09/25/2019. My last successful test was performed on the night of March 31, 2021, so I believe the documentation remains (and shall continue to be) valid.

The PTCS manual is available under the "manual-correios" folder, only in Brazilian Portuguese though.

All comments and most variables names are in Brazilian Portuguese. I usually prefer doing everything in English, but given the specificity of this implementation, the decision is justified.

#### Contributions

Your contributions are welcome. What about a Pull Request?

#### Licensing

The Apache 2.0 License. Please see the License File for more.
