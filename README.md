<div align="center">
    <img alt="Golang Logo" title="Golang Logo" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" width="300" height="300">
</div>

# Virtual Wallet - REST API - GOLANG

## Techs utilizadas

- [Golang](https://go.dev)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)

## Variáveis de Ambiente

Para rodar esse projeto, você vai precisar adicionar as seguintes variáveis de ambiente no seu .env

`API_PORT`

`DATABASE_URL`

## Requisitos

- [x] Crie um enpoint que recebe o dois ids de usuários, e um valor monetário representando a transferência entre eles.
- [x] Crie um endpoint que recebe um id de usuário e retorna o saldo dele.
- [x] Valide se o usuário de origem tem saldo suficiente antes da transferência.
- [x] É preciso pensar na possibilidade de concorrência de transferências onde duas pessoas tranferem dinheiro ao mesmo tempo para uma terceira.
- [x] Se uma transferência falhar, o saldo do usuário de origem deve ser restaurado.
- [x] Não é necessário endpoints para criar usuários, popule o banco de forma com que os dois usuários existam e que transferências possam ser feitas entre eles.

## Endpoints

#### Criar uma transferência

```http
  POST /transfer/
```

| Parâmetro        | Tipo     | Descrição                               |
| :--------------- | :------- | :-------------------------------------- |
| `amount`         | `int64`  | **Obrigatório**. Valor da transferência |
| `debtor_id`      | `string` | **Obrigatório**. ID do Debitante        |
| `beneficiary_id` | `string` | **Obrigatório**. ID do Beneficiário     |

#### Buscar usuário pelo id

```http
  GET /users/{id}
```

- Exemplo de Payload:

```json
{
  "amount": 100,
  "debtor_id": "089557bc-ddf2-4ec5-8077-d8bf09fe3ddc",
  "beneficiary_id": "f2f0e0d1-e37e-45c3-ad06-e6c2a66544fc"
}
```

- Exemplo de Requisição:

```sh
curl -X POST \
  http://localhost:8080/opportunity \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d '{
	"amount": 10,
	"debtor_id": "089557bc-ddf2-4ec5-8077-d8bf09fe3ddc",
	"beneficiary_id": "f2f0e0d1-e37e-45c3-ad06-e6c2a66544fc"
}'
```
