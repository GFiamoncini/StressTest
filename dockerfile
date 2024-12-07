# Usar imagem base Golang
FROM golang:1.23.3

# Criar diretório de trabalho
WORKDIR /app

# Copiar código
COPY . .

# Construir aplicação
RUN go build -o load-tester

# Comando de execução
ENTRYPOINT ["./load-tester"]
