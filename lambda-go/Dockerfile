# Use Amazon Linux 2 image
FROM amazonlinux:2

# Instale o Go
RUN yum install -y \
    glibc-static \
    gcc \
    go \
    zip

# Configure o Go
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# Crie o diretório de trabalho
WORKDIR /go/src/app

# Crie um módulo Go
RUN go mod init app

# Instale o pacote AWS Lambda Go
RUN go get -u github.com/aws/aws-lambda-go/lambda

# Copie o código da função Lambda para o contêiner
COPY fizzbuzz.go .

# Compile o código da função Lambda
RUN go build -o bootstrap

# Crie o arquivo zip
RUN zip main.zip bootstrap

# Comando para exportar o arquivo zip
CMD ["cp", "main.zip", "/output/main.zip"]
