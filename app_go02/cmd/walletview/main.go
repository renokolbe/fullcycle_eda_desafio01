package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"

	"github.com/renokolbe/fc-ms-wallet-balance/internal/database"
	"github.com/renokolbe/fc-ms-wallet-balance/internal/dto"
	"github.com/renokolbe/fc-ms-wallet-balance/internal/entity"
	"github.com/renokolbe/fc-ms-wallet-balance/internal/infra/kafka"
	"github.com/renokolbe/fc-ms-wallet-balance/internal/transformer"
	get_accountbalance "github.com/renokolbe/fc-ms-wallet-balance/internal/usecase/get_account-balance"
	"github.com/renokolbe/fc-ms-wallet-balance/internal/web"
	"github.com/renokolbe/fc-ms-wallet-balance/internal/web/webserver"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	strConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local", "root", "root", "mysql", "3306", "wallet")
	db, err := sql.Open("mysql", strConnection)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	accountBalances := make(chan *entity.AccountBalance)
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	kafkaMsgChan := make(chan *ckafka.Message)

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}

	kafka := kafka.NewConsumer(&configMap, []string{"balances"})

	go kafka.Consume(kafkaMsgChan)

	fmt.Println("Monitor Consumo Kafka - Topico balances is running...")

	accountBalanceDB := database.NewAccountBalanceDB(db)

	// Recebe os Dados do Kafka, Transforma e Joga no Canal de Entrada
	go func() {
		for msg := range kafkaMsgChan {
			wg.Add(1)
			fmt.Println(string(msg.Value))
			transactionInput := dto.TransactionInputDTO{}
			err := json.Unmarshal(msg.Value, &transactionInput)
			if err != nil {
				panic(err)
			}
			accountFrom := transformer.TransformInputAccountFrom(transactionInput)
			accountBalances <- accountFrom
			accountTo := transformer.TransformInputAccountTo(transactionInput)
			accountBalances <- accountTo
		}
	}()

	aBalances := entity.NewAccountBalances(accountBalances, wg)
	go aBalances.Registro(accountBalanceDB)

	fmt.Println("Monitor Canal de Transações is running...")

	getAccountBalanceUsecase := get_accountbalance.NewGetAccountBalanceUseCase(accountBalanceDB)
	accountBalanceHandler := web.NewWebAccountBalanceHandler(*getAccountBalanceUsecase)
	webserver := webserver.NewWebServer(":3003")

	webserver.AddHandler("/account-balance/{id}", accountBalanceHandler.GetAccountBalance)

	fmt.Println("Server is running on port 3003")

	webserver.Start()
}
