<p align="center">
  <img src="https://cdn-icons-png.flaticon.com/512/6218/6218295.png" width="100" />
</p>
<p align="center">
    <h1 align="center">GO-EXPERT TEMP BY CEP</h1>
</p>
<p align="center">
    <em>Desafio da pós em GO Expert</em>
</p>
<p align="center">
	<img src="https://img.shields.io/github/last-commit/mhayk/GO-Expert-rate-limiter?style=flat&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/mhayk/GO-Expert-rate-limiter?style=flat&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/mhayk/GO-Expert-rate-limiter?style=flat&color=0080ff" alt="repo-language-count">
<p>
<p align="center">
    <em>Developed with ❤️ by Mhayk Whandson</em>
</p>
<p align="center">
		<em>Developed with the language, software and tools below.</em>
</p>
<p align="center">
	<img src="https://img.shields.io/badge/YAML-CB171E.svg?style=flat&logo=YAML&logoColor=white" alt="YAML">
	<img src="https://img.shields.io/badge/V8-4B8BF5.svg?style=flat&logo=V8&logoColor=white" alt="V8">
	<img src="https://img.shields.io/badge/Docker-2496ED.svg?style=flat&logo=Docker&logoColor=white" alt="Docker">
	<img src="https://img.shields.io/badge/Go-00ADD8.svg?style=flat&logo=Go&logoColor=white" alt="Go">
</p>
<hr>


# Lab GO: Sistema de Temperatura por CEP

## Descrição do Desafio

Este projeto faz parte de um desafio Full Cycle, onde o objetivo é criar uma API capaz de calcular a temperatura com base no CEP informado.

### API Externa Utilizada:

- **API**:  https://temp-by-cep-1035948664484.us-central1.run.app/temp-by-cep?zipcode=SEU_CEP_AQUI
- **Exemplo**:   https://temp-by-cep-1035948664484.us-central1.run.app/temp-by-cep?zipcode=69304350

#### Retorno Esperado:
```json
{
    "temp_C": 28.7,
    "temp_F": 83.6,
    "temp_K": 301.7
}
```

### Como Rodar em Ambiente de Desenvolvimento
Siga os passos abaixo para configurar e rodar o projeto em sua máquina local:

1. Clone o Projeto
Baixe o repositório para a sua máquina:

```bash
git clone https://github.com/mhayk/GO-Expert-temp-by-cep
cd GO-Expert-temp-by-cep
```
2. Instale as Dependências
Certifique-se de ter o Go instalado, e execute o comando:

```bash
go mod tidy
```

3. Crie uma Conta na WeatherAPI
Acesse WeatherAPI e crie uma conta. Esta API será utilizada para consultar informações de temperatura.

4. Configure as Variáveis de Ambiente
Faça uma cópia do arquivo .env.example e renomeie para .env:
```bash
cp .env.example .env
```
Dentro do arquivo .env, adicione sua chave da WeatherAPI na variável WEATHER_API_KEY:
```text
WEATHER_API_KEY=YOUR_API_KEY
```
5. Rodar com Docker Compose
Execute o comando abaixo para levantar os serviços com Docker Compose:

```bash
docker compose up
```
6. Acessar o Container e Rodar a Aplicação
Acesse o container e inicie a aplicação com o comando:

```bash
docker exec -it <nome_do_container> /bin/sh
go run main.go
```

7. Executar os Testes
Para rodar os testes do projeto, utilize o comando:

``` bash
go test ./...
```

### Tecnologias Utilizadas
* Go
* Docker
* WeatherAPI
* Google Cloud Run
