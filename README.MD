## Family tree challenge

### Descrição

Este repositório contém uma API escrita em Golang que criar e gerenciar uma árvore genealógica.

### Pré-requisitos

- Docker (https://docs.docker.com/get-docker/)
- Docker Compose (https://docs.docker.com/compose/install/)

### Como rodar a aplicação

1. Clone o repositório para sua máquina local:
   ```shell
   git clone https://github.com/rodriguezbarreto/family-tree-challenge
   cd family-tree-challenge
   ```
   
2. Construa a imagem Docker da API:

    ```shell
    docker build -t family-tree-challenge .
    ```

3. Inicie a aplicação e o banco de dados com o Docker Compose:

    ```shell
   docker-compose up -d
    ```
    Este comando iniciará dois contêineres, um para a API e outro para o banco de dados PostgreSQL.

4. Acesse a API no seu navegador ou com um cliente HTTP, como o curl:

    ```shell
    curl http://localhost:8080
    ```

5. Se você quiser interromper e remover os contêineres, use o seguinte comando:

    ```shell
   docker-compose down
    ```

6. Há um arquivo para preencher o banco de dados com algumas informações. Siga as instruções a seguir para semear o banco de dados:

    ```shell
   cd cmd/seed
   go run seed.go
    ```

### Estrutura de pastas

    ├──📁 cmd
    │   ├──📁configs
    │   │   ├── seed-database.go
    │   │   └── setup-middlewares.go
    │   ├──📁 enpoints
    │   │   └── setup-endpoints.go
    │   ├──📁 seed
    |   |   └── seed.go
    │   └── main.go
    ├──📁 internal
    │   ├──📁 domain
    │   │   ├──📁 errors
    │   │   │   └── validator.go
    │   │   ├──📁 tests
    │   │   |   ├── person_test.go
    │   │   |   └── relationship_test.go
    │   │   ├──relationship.go
    │   │   └──person.go
    │   ├──📁 infra
    │   │   ├──📁 controllers
    │   │   │   ├── create-person-controller.go
    │   │   │   ├── create-relationship-controller.go
    │   │   │   ├── delete-person-controller.go
    │   │   │   ├── delete-relationship-controller.go
    │   │   │   ├── get-bacon-number-controller.go
    │   │   │   ├── get-genealogy-controller.go
    │   │   │   ├── list-person-controller.go
    │   │   │   ├── list-relationship-controller.go
    │   │   │   └── update-person controller.go
    │   │   ├──📁 database
    │   │   │   ├── postgres-connection.go
    │   │   │   └──📁 repositories
    │   │   │       ├── person-repository.go
    │   │   │       └── relationship-repository.go
    │   │   └──📁 factories
    │   │       ├── bacon-number-facotry.go
    │   │       ├── create-person-factory.go
    │   │       ├── create-relationship-factory.go
    │   │       ├── delete-person-factory.go
    │   │       ├── delete-relationship.go
    │   │       ├── get-genealogy-factory.go
    │   │       ├── list-person-factory.go
    │   │       ├── list-relationship-factory.go
    │   │       ├── person-repository-factory.go
    │   │       ├── relationship-repository-factory.go
    │   │       └── update-person-factory.go
    │   └──📁 use-cases
    │       ├──📁 dto
    │       │   ├── genealogy-dto.go
    │       │   ├── person-dto.go
    │       │   └── relationship-dto.go
    │       ├──📁 repositories
    │       │   ├── person-repository.go
    │       │   └── relationship-repository.go
    |       ├──📁 tests
    │       │   ├── create_person_test.go
    │       │   ├── create_relationship_test.go
    │       │   ├── delete_person_test.go
    │       │   ├── delete_relationship_test.go
    │       │   ├── list_person_test.go
    │       │   ├── list_relationship_test.go
    │       │   ├── mock_test.go
    │       │   └── update_person_test.go
    │       ├── create-person.go
    │       ├── create-relationship.go
    │       ├── delete-person.go
    │       ├── delete-relationship.go
    │       ├── get-bacon-number.go
    │       ├── get-genealogy.go
    │       ├── list-person.go
    │       ├── list-relationship.go
    │       └── update-person.go
    |
    ├── docker-compose.yml
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    └── README.MD

<br>

### Endpoints


---
>### Persons
#### POST `/persons`

Create Person

- **Request Body:** JSON
    ```Json
    {
        "name": "Daniel"
    }
    ```

- **Response:**
  - **HTTP 201**: Pessoa criada com sucesso
  - **HTTP 400**: Requisição inválida.
  - **HTTP 500**: Erro interno do servidor.

<br>

---

#### GET `/persons/{id}`

List Persons by ID

- **Parameters:** 
    + **`id`** (int, path, obrigatório): ID da `Person`.
- **Response Body:** JSON
    ```Json
    {
        "id": "7f272309-bd0b-4cf3-8c09-81391ba2d2df",
        "name": "Sonny"
    }
    ```
- **Status Response:**
    - **HTTP 200**: Retorna JSON com informações da pessoa.
    - **HTTP 404**: Pessoa não encontrada.
    - **HTTP 400**: Requisição inválida.
    - **HTTP 500**: Erro interno do servidor.

<br>

---

#### GET `/persons`

List all Persons

- **Response Body:** JSON
    ```Json
    {
        "id": "3c8f6f29-9dab-4dfc-9fee-a93279865e89",
        "name": "Ann"
    },
    {
        "id": "d3ab4771-e066-4e26-b13e-986e92241601",
        "name": "Dunny"
    },
    {
        "id": "b1c6dd69-0abf-4aeb-ad48-9e838fd76d34",
        "name": "Bruce"
    },
    {
        "id": "bc1afa20-a896-4085-9700-44f7eb74106a",
        "name": "Advick"
    }
    ```

- **Status Response:**
    - **HTTP 200**: Retorna JSON com informações das pessoas.
    - **HTTP 500**: Erro interno do servidor.

<br>

---

#### PUT `/persons/{id}`

Update Person

- **Parameters:**
  - `id` (int, path, obrigatório): ID da pessoa.
- **Request Body:** JSON
    ```Json
    {
        "name": "Sonny"
    }
    ```
- **Status Response:**
    - **HTTP 201**: Pessoa atualizada com sucesso
    - **HTTP 404**: Pessoa não encontrada.
    - **HTTP 400**: Requisição inválida.
    - **HTTP 500**: Erro interno do servidor

<br>

---

#### DELETE `/persons/{id}`

Delete Person

- **Parameters:** 
  - `id` (int, path, obrigatório): ID da pessoa.
- **Status Response:**
    - **HTTP 204**: Pessoa Deletada com sucesso.
    - **HTTP 400**: Requisição inválida.
    - **HTTP 500**: Erro interno do servidor

<br>

---
>### Reelationships

Create Relationship

- **Request Body:** JSON
    ```Json
    {
        "child": "da2e1e9a-0d12-4506-93e1-833d6887a83f",
        "parent": "23dc076d-8eed-4349-bffa-e0fcce80d5f1"
    }
    ```
- **Response:**
  - **HTTP 201**: Relacionamento criado com sucesso.
  - **HTTP 400**: Requisição inválida.
  - **HTTP 500**: Erro interno do servidor.

<br>

---
#### GET `/relationships`

List all Relationship.

- **Parameters:** 
  - `relId` (query string): Filtra relacionamento por ID
  - `childId` (query string): Filtra todos relacionamentos na condição de filho
  - `parentId` (query string): Filtra todos relacionamentos na condição de pai

  <br>

- **Response Body:** JSON
    ```Json
    [
       {
            "id": "4a82227b-c64b-462e-87de-6b3a49521231",
            "child_id": "3c8f6f29-9dab-4dfc-9fee-a93279865e89",
            "parent_id": "7f272309-bd0b-4cf3-8c09-81391ba2d2df",
            "child": {
                "id": "3c8f6f29-9dab-4dfc-9fee-a93279865e89",
                "name": "Ann"
            },
            "parent": {
                "id": "7f272309-bd0b-4cf3-8c09-81391ba2d2df",
                "name": "Sonny"
            }
        },
        {
            "id": "93ca5dc7-c45f-4c37-951d-56bd8e3be415",
            "child_id": "6ea48a6b-283d-4832-a60c-70c6d834b79e",
            "parent_id": "da6e7212-4f6a-4c3e-adf8-8aa443006df4",
            "child": {
                "id": "6ea48a6b-283d-4832-a60c-70c6d834b79e",
                "name": "Phoebe"
            },
            "parent": {
                "id": "da6e7212-4f6a-4c3e-adf8-8aa443006df4",
                "name": "Martin"
            }
        },
    ]
    ```

- **Status Response:**
  - **HTTP 200**: Retorna JSON com uma lista de relacionamentos.
  - **HTTP 500**: Erro interno do servidor.

<br>

---
>### Genealogy

#### GET `/genealogy/{id}`

Get Genealogy.

- **Parameters:** 
  - `id` (int, path, obrigatório): ID da pessoa.

- **Response Body:** JSON
    ```Json
   {
        "members": [
            {
                "name": "Phoebe",
                "relationships": [
                    {
                        "name": "Martin",
                        "relationship": "parent"
                    },
                    {
                        "name": "Anastasia",
                        "relationship": "parent"
                    }
                ]
            },
            {
                "name": "Anastasia",
                "relationships": []
            },
            {
                "name": "Martin",
                "relationships": []
            },
            {
                "name": "Bruce",
                "relationships": [
                    {
                        "name": "Phoebe",
                        "relationship": "parent"
                    }
                ]
            },
            {
                "name": "Dunny",
                "relationships": [
                    {
                        "name": "Phoebe",
                        "relationship": "parent"
                    }
                ]
            },
            {
                "name": "Clark",
                "relationships": [
                    {
                        "name": "Anastasia",
                        "relationship": "parent"
                    }
                ]
            }
        ]
    }
    ```
- **Status Response:**
  - **HTTP 200**: Retorna JSON com a árvore genealógica.
  - **HTTP 404**: Pessoa não encontrada.
  - **HTTP 500**: Erro interno do servidor.

<br>

---
#### GET `/genealogy/bacon-number`

Get Bacon Number.

- **Request Body:** JSON
    ```Json
    {
        "sourceId": "b1c6dd69-0abf-4aeb-ad48-9e838fd76d34",
        "targetId": "401244a0-cc21-46e9-a1b8-27067d4da7c0"
    }
    ```

- **Response Body:** JSON
    ```Json
    {
        "baconNumber": "4"
    }   
    ```

- **Status Response:**
  - **HTTP 200**: Retorna JSON com o Bacon Number.
  - **HTTP 500**: Erro interno do servidor.