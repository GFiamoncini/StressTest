
# Load Tester CLI

Este é um sistema CLI (Command-Line Interface) desenvolvido em Go para realizar testes de carga em serviços web. Ele permite configurar o número total de requisições, o nível de concorrência e gerar relatórios detalhados após a execução dos testes.

## Funcionalidades

- Envia múltiplas requisições HTTP para um serviço web.
- Controla o número de requisições simultâneas (concorrência).
- Gera um relatório com:
  - Tempo total do teste.
  - Total de requisições realizadas.
  - Número de respostas HTTP 200.
  - Distribuição de outros códigos de status HTTP.

---

## Pré-requisitos

- [Docker](https://www.docker.com) instalado.
- (Opcional) [Go](https://golang.org) para compilar e testar localmente.

---

## Configuração

### Construindo a Imagem Docker

1. Clone este repositório ou copie os arquivos necessários para o diretório local:
   ```bash
   git clone <URL_DO_REPOSITORIO>
   cd <NOME_DO_DIRETORIO>
   ```

2. Certifique-se de que o arquivo `Dockerfile` está no diretório raiz.

3. Construa a imagem Docker:
   ```bash
   docker build -t load-tester .
   ```

---

## Uso

### Executando o Teste de Carga

Após construir a imagem Docker, execute o container com os seguintes parâmetros:

```bash
docker run load-tester --url=<URL> --requests=<TOTAL_REQUESTS> --concurrency=<CONCURRENCY_LEVEL>
```

#### Parâmetros

- `--url`: URL do serviço que será testado (exemplo: `http://example.com`).
- `--requests`: Total de requisições que serão enviadas.
- `--concurrency`: Número de requisições simultâneas.

#### Exemplo

```bash
docker run load-tester --url=http://example.com --requests=1000 --concurrency=10
```

### Resultado do Relatório

Após a execução, a aplicação exibirá um relatório no terminal com informações como:

- **Tempo total gasto**: Tempo total da execução do teste.
- **Total de requests**: Total de requisições realizadas.
- **Requests com status 200**: Quantidade de respostas bem-sucedidas.
- **Distribuição de status HTTP**: Frequência de outros códigos de status (404, 500, etc.).

---

## Desenvolvimento Local (Opcional)

Se preferir testar a aplicação localmente sem Docker:

1. Instale o Go e execute o programa:
   ```bash
   go run main.go --url=http://example.com --requests=1000 --concurrency=10
   ```

2. Para criar um binário executável:
   ```bash
   go build -o load-tester
   ./load-tester --url=http://example.com --requests=1000 --concurrency=10
   ```

---

## Problemas Conhecidos

- Certifique-se de que o serviço fornecido pela URL aceita conexões HTTP.
- Verifique os limites do servidor para evitar bloqueios causados por alta carga.

---
