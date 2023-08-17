## Family tree challenge

### Descrição

Este repositório contém uma API escrita em Golang que criar e gerenciar uma árvore genealógica,

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

6. A uma arquivo para popular o banco de dados com algumas informações as instruções para semear o banco de dados é:

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

