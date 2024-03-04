## Guia de Execução do Projeto

Este guia fornece as instruções para executar o projeto que envolve o uso do MySQL, Docker com RabbitMQ e uma aplicação Go para consumir e salvar mensagens no banco de dados.

### Pré-requisitos

Antes de começar, certifique-se de ter instalado em seu computador as seguintes ferramentas:

- Docker (com Docker Compose)
- Go (Golang) 1.21+

### Passo 1: Criar o diretório para o banco de dados MySQL

No terminal, execute o comando abaixo para criar um diretório onde será armazenado os dados do MySQL:

```
mkdir /tmp/mysql-data
```

### Passo 2: Executar o Docker Compose com o RabbitMQ

Certifique-se de que o Docker esteja instalado e em execução. No mesmo terminal, navegue até o diretório onde está localizado o arquivo **_docker-compose.yml_** e execute o seguinte comando para iniciar o RabbitMQ em um container:

```
docker-compose up -d
```

### Passo 3: Configurar a queue "order" no RabbitMQ

Abra um navegador e acesse a interface web do RabbitMQ em `http://localhost:15672` (login: guest, senha: guest). Crie uma nova fila chamada **_order_**.

### Passo 4: Publicar uma mensagem na queue "order"

Utilizando a interface web do RabbitMQ, acesse a aba **_Publish message_** da fila **_order_** e publique a seguinte mensagem no formato _JSON_:

```
{"id":"3","price":13.0,"tax":0.3}
```

### Passo 5: Executar a aplicação Go

No mesmo terminal, execute:

```
make run
```

- A aplicação estará disponível em [http://localhost:8000](http://localhost:8000).
- As _migrations_ serão executadas automaticamente no startup da aplicação.
- O GraphQL Playground estará disponível em [http://localhost:8080](http://localhost:8080).
- O gRPC server estará disponível em [http://localhost:50051](http://localhost:50051).

### Passo 6: Consumir e salvar a mensagem no banco de dados

A aplicação Go consumirá a mensagem publicada na fila **_order_** no RabbitMQ e salvará os dados na tabela **_orders_** do banco de dados SQLite.

## Disponibilizar API do projeto

### Passo 1: Gerar imagem Docker
No terminal no diretório da aplicação, execute o seguinte comando para gerar a imagem Docker:
```
make docker-build
```

### Passo 2: Executar aplicação em container
No terminal no diretório da aplicação, execute o seguinte comando para rodar aplicação em container Docker:
```
make docker-run
```
