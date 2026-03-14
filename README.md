# finance-notification

Aplicação em Go que consulta cotações de ações da bolsa brasileira (B3) via API [brapi.dev](https://brapi.dev), projetada para ser executada como uma função AWS Lambda.

## O que faz

- Busca a lista completa de ações disponíveis na B3
- Para cada ação, consulta a cotação atual e o dado histórico do dia anterior
- Serializa os resultados em JSON e os retorna como resposta da Lambda

## Estrutura do projeto

```
finance-notification/
├── main.go           # Handler da Lambda e ponto de entrada
├── client/
│   └── brapi.go      # Cliente HTTP da API brapi.dev
└── model/
    └── quote.go      # Structs de domínio e conversão de dados
```

## Pré-requisitos

- Go 1.21+
- Conta na [brapi.dev](https://brapi.dev) para obter um token de acesso
- AWS CLI configurado (para deploy)

## Configuração

A aplicação utiliza a variável de ambiente abaixo:

| Variável      | Descrição                          | Obrigatória |
|---------------|------------------------------------|-------------|
| `BRAPI_TOKEN` | Token de autenticação da brapi.dev | Sim         |

## Executando localmente

```bash
export BRAPI_TOKEN=seu_token_aqui
go run main.go
```

## Deploy na AWS Lambda

1. Compile o binário para Linux:

```bash
GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
```

2. Empacote em zip:

```bash
zip function.zip bootstrap
```

3. Faça o deploy via AWS CLI:

```bash
aws lambda create-function \
  --function-name finance-notification \
  --runtime provided.al2 \
  --handler bootstrap \
  --zip-file fileb://function.zip \
  --role arn:aws:iam::<account-id>:role/<execution-role> \
  --environment Variables={BRAPI_TOKEN=seu_token_aqui}
```

Para atualizar uma função já existente:

```bash
aws lambda update-function-code \
  --function-name finance-notification \
  --zip-file fileb://function.zip
```

### Configurações da Lambda

| Configuração | Valor            |
|--------------|------------------|
| Runtime      | `provided.al2`   |
| Handler      | `bootstrap`      |
| Arquitetura  | `x86_64`         |

## Exemplo de resposta

```json
[
  {
    "symbol": "PETR4",
    "shortName": "PETROBRAS PN",
    "regularMarketPrice": 30.65,
    "regularMarketChange": 0.22,
    "regularMarketChangePercent": 0.72,
    "regularMarketVolume": 21075300,
    "historicalDataPrice": {
      "date": "24/01/2026",
      "open": 30.47,
      "close": 30.65,
      "volume": 21075300,
      "adjustedClose": 30.65
    }
  }
]
```

## API utilizada

[brapi.dev](https://brapi.dev) é uma API brasileira que fornece dados do mercado financeiro, incluindo cotações de ações, FIIs e índices da B3.

| Endpoint             | Descrição                            |
|----------------------|--------------------------------------|
| `GET /api/quote/list`| Lista todas as ações disponíveis     |
| `GET /api/quote/:ticker?interval=1d&range=1d` | Cotação e histórico de uma ação |