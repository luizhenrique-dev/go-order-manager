## Guia de Execução do Projeto

Este guia fornece as instruções para executar o projeto que envolve o uso do SQLite, Docker com RabbitMQ e uma aplicação Go para consumir e salvar mensagens no banco de dados.

### Pré-requisitos

Antes de começar, certifique-se de ter instalado em seu computador as seguintes ferramentas:

- Docker (com Docker Compose)
- SQLite3
- Go (Golang)

### Passo 1: Criar a tabela no banco de dados SQLite

No terminal, execute o comando abaixo para acessar o cliente SQLite no banco de dados **_db.sqlite3_**:

```
sqlite3 db.sqlite3
```

### Passo 2: Criar a tabela "orders"

Dentro do cliente SQLite, cole e execute o seguinte comando para criar a tabela **_orders_**:

```
CREATE TABLE orders (id varchar(255) not null, price float not null, tax float not null, final_price float not null, primary key (id));
```

### Passo 3: Executar o Docker Compose com o RabbitMQ

Certifique-se de que o Docker esteja instalado e em execução. No mesmo terminal, navegue até o diretório onde está localizado o arquivo **_docker-compose.yml_** e execute o seguinte comando para iniciar o RabbitMQ em um container:

```
docker-compose up -d
```

### Passo 4: Configurar a queue "order" no RabbitMQ

Abra um navegador e acesse a interface web do RabbitMQ em `http://localhost:15672` (login: guest, senha: guest). Crie uma nova fila chamada **_order_**.

### Passo 5: Publicar uma mensagem na queue "order"

Utilizando a interface web do RabbitMQ, acesse a aba **_Publish message_** da fila **_order_** e publique a seguinte mensagem no formato _JSON_:

```
{"id":"3","price":13.0,"tax":0.3}
```

### Passo 6: Executar a aplicação Go

No mesmo terminal, navegue até o diretório onde se encontra o arquivo "main.go" da aplicação Go. Em seguida, execute o seguinte comando para rodar a aplicação:

```
go run cmd/order/main.go
```

### Passo 7: Consumir e salvar a mensagem no banco de dados

A aplicação Go consumirá a mensagem publicada na fila **_order_** no RabbitMQ e salvará os dados na tabela **_orders_** do banco de dados SQLite.