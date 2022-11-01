package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bxcodec/faker"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
	_ "github/go-sql-driver/mysql"
)

type Transaction struct {
	TransactionId      int     `json:"TransactionId"`
	RequestId          int     `json:"RequestId"`
	TerminalId         int     `json:"TerminalId"`
	ParterObjectId     int     `json:"ParterObjectId"`
	AmountTotal        int     `json:"AmountTotal"`
	CommissionPS       float32 `json:"CommissionPS"`
	CommissionClient   int     `json:"CommissionClient"`
	CommissionProvider float32 `json:"CommissionProvider"`
	DateInput          string  `json:"DateInput"`
	DatePost           string  `json:"DatePost"`
	Status             string  `json:"Status"`
	PaymentType        string  `json:"PaymentType"`
	PaymentNumber      string  `json:"PaymentNumber"`
	ServiceId          int     `json:"ServiceId"`
	Service            string  `json:"Service"`
	PayeeId            int     `json:"PayeeId"`
	PayeeName          string  `json:"PayeeName"`
	PayeeBankMfo       int     `json:"PayeeBankMfo"`
	PayeeBankAccount   string  `json:"PayeeBankAccount"`
	PaymentNarrative   string  `json:"PaymentNarrative"`
}

func getMySQLDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@(127.0.0.1:3306)/test&?parseTime=true")
	if err!= nil {
		log.Fatal(err)
	}
	return db
}
func main() {
	
	var db = getMySQLDB()

	app := fiber.New()
    
	app.Get("/", func(c *fiber.Ctx) error {
	transaction := []Transaction{}

	db.Find(&transaction)

		return c.JSON(transaction)
	})
    transaction := []Transaction{}
	csvFile, err := os.Open("example.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(csvFile)
	data, _ := reader.ReadAll()

	for _, value := range data {
		transaction = append(transaction, Transaction{
			TransactionId:      value[0],
			RequestId:          value[1],
			TerminalId:         value[2],
			ParterObjectId:     value[3],
			AmountTotal:        value[4],
			CommissionPS:       value[5],
			CommissionClient:   value[6],
			CommissionProvider: value[7],
			DateInput:          value[8],
			DatePost:           value[9],
			Status:             value[10],
			PaymentType:        value[11],
			PaymentNumber:      value[12],
			ServiceId:          value[13],
			Service:            value[14],
			PayeeId:            value[15],
			PayeeName:          value[16],
			PayeeBankMfo:       value[17],
			PayeeBankAccount:   value[18],
			PaymentNarrative:   value[19],
		})
	}
	for i:=1; i<len(transaction); i++{
    TransactionId,_ := strconv.Atoi(transaction[i].TransactionId)
	RequestId,_ := strconv.Atoi(transaction[i].RequestId)
	TerminalId,_ := strconv.Atoi(transaction[i].TerminalId)
	ParterObjectId,_ := strconv.Atoi(transaction[i].ParterObjectId)
	AmountTotal,_ := strconv.Atoi(transaction[i].AmountTotal)
	CommissionPS,_ := strconv.ParseFloat(transaction[i].CommissionPS)
	CommissionClient,_ := strconv.Atoi(transaction[i].CommissionClient)
	CommissionProvider,_ := strconv.ParseFloat(transaction[i].CommissionProvider)
	ServiceId,_ := strconv.Atoi(transaction[i].ServiceId)
	PayeeId,_ := strconv.Atoi(transaction[i].PayeeId)
	PayeeBankMfo,_ := strconv.Atoi(transaction[i].PayeeBankMfo)
	_, err := db.Exec("insert into 
	test(TransactionId,
		 RequestId,
		 TerminalId,
		 ParterObjectId,
		 AmountTotal,
		 CommissionPS
		 CommissionClient,
		 CommissionProvider,
		 DateInput,
		 DatePost,
		 Status,
		 PaymentType,
		 PaymentNumber,
		 ServiceId,
		 Service,
		 PayeeId,
		 PayeeName,
		 PayeeBankMfo,
		 PayeeBankAccount,
		 PaymentNarrative) 
	values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
	TransactionId, RequestId, TerminalId, ParterObjectId,
    AmountTotal, CommissionPS, CommissionClient, CommissionProvider,
    transaction[i].DateInput, transaction[i].DatePost, transaction[i].Status, 
	transaction[i].PaymentType, transaction[i].PaymentNumber, 
    ServiceId, transaction[i].Service, PayeeId, transaction[i].PayeeName, 
	PayeeBankMfo, transaction[i].PayeeBankAccount, transaction[i].PaymentNarrative )
	if err != nil {
		fmt.Printf("" + err.Error())
	}
	}

	fmt.Println("All inserted")

	transactionJSON, _ := json.Marshal(transaction)
	fmt.Println(string(transactionJSON)) 


}
