## IMAGEM INICIAL - base-go

# Imagem Base
FROM golang:1.20 as base-go

RUN apt-get update && apt-get install -y librdkafka-dev 

WORKDIR /usr/src/app
COPY . .

RUN go build -v -o /usr/local/bin/walletview ./cmd/walletview/main.go
CMD [ "/usr/local/bin/walletview" ]

## -----------
## IMAGEM FINAL
## ------------
# Imagem Base - scratch
##FROM scratch

# Copia os dados do Estagio Incial para a Imagem Final
##COPY --from=base-go /usr/local/bin/walletview .
##CMD [ "./walletview" ]