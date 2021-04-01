# correios-calc

### Para clonar o repositório e rodar em sua máquina local (to clone the repo and run in your local machine):

1. Clone o repositório em sua máquina local (clone the repo to your local machine):
```
git clone https://github.com/cewitte/correios-calc.git
```

2. Execute o código (run the code):
```
go run main.go
```

3. Siga o seu coração e faça o Bem (follow your heart and do Good).


### Para usar o código em um projeto Go existente (to use the code in an existing Go project):
```
go get -u github.com/cewitte/correios-calc
```


### Notas em Português do Brasil
Esta é uma implementação Go (Golang) do serviço web dos Correios que permite obter o preço e o prazo de envio de um pacote com base em alguns parâmetros.

Essa implementação foi baseada no **"Manual técnico de integração web services ao Sistema Calculador de Preço e Prazo - SCPP.pdf"** na versão 2.2 de 25/09/2019. Meu último teste com sucesso foi realizado na noite de 31 de março de 2021, portanto acredito que a documentação continua (e continuará) válida.

O manual SCPP está disponível aqui no diretório "manual-correios", apenas em português do Brasil.

Por fim, peço desculpas por manter todo o código em um único arquivo (`main.go`) em vez de empacotá-lo corretamente. Ainda assim, o código é tão curto que não faria muito sentido fazê-lo. Você notará que metade do código são comentários ou a função `main()` de teste.

Mais uma coisa: todos os comentários e a maioria das variáveis estão em português brasileiro. Normalmente prefiro fazer tudo em inglês, mas dada a especificidade desta implementação, a decisão me parece justificada.

#### Contribuições
Suas contribuições são bem vindas. Que tal um Pull Request?

#### Licenciamento
The Apache 2.0 licence. Por favor, veja o arquivo da licença.


### Notes in U.S. English
This is a Go (Golang) implementation of the Brazilian Postal Service's (aka "Correios") web service which allows you to retrieve the price and ETA of a package shipment based on some parameters.

This implementation was based on the **"Technical manual for integrating web services into the Price and Term Calculator System - PTCS.pdf"** (title loosely translated) version 2.2 dated 09/25/2019. My last successful test was performed on the night of March 31, 2021, so I believe the documentation remains (and shall continue to be) valid.

The PTCS manual is available under the "manual-correios" folder, only in Brazilian Portuguese though.

At last, I apologize for keeping all the code in a single file (`main.go`) instead of correctly packaging it. Still, IMHO the code is so short that it wouln't make too much sense to do so. You will notice that half of it are comments or the `func main()` for tests.

One more thing: all comments and most variables are in Brazilian Portuguese. I usually prefer doing everything in English, but given the specificity of this implementation, the decision is justified.

#### Contributions
Your contributions are welcome. What about a Pull Request?

#### Licensing
The Apache 2.0 License. Please see the License File for more.
