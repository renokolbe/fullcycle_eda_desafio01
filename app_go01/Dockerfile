## IMAGEM INICIAL - base-go

# Imagem Base
FROM golang:1.20 as base-go

RUN apt-get update && apt-get install -y librdkafka-dev 

WORKDIR /usr/src/app
COPY . .

RUN go build -v -o /usr/local/bin/walletcore ./cmd/walletcore/main.go
CMD [ "/usr/local/bin/walletcore" ]

## -----------
## IMAGEM FINAL
## ------------
# Imagem Base - scratch
##FROM scratch

# Copia os dados do Estagio Incial para a Imagem Final
##COPY --from=base-go /usr/local/bin/walletcore .
##CMD [ "./walletcore" ]