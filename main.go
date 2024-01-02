// package main

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
// )

// var (
// 	botToken  = "6603083264:AAEmOV1kIuNWYIxgIrj5kye7NfsgP-Ud2m8"
// 	expenses  = make(map[int]float64)
// 	startTime = time.Now()
// )

// func main() {
// 	bot, err := tgbotapi.NewBotAPI(botToken)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	bot.Debug = true

// 	log.Printf("Authorized on account %s", bot.Self.UserName)

// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 60

// 	updates, err := bot.GetUpdatesChan(u)

// 	for update := range updates {
// 		if update.Message == nil {
// 			continue
// 		}

// 		processMessage(update.Message, bot)
// 	}
// }

// func processMessage(msg *tgbotapi.Message, bot *tgbotapi.BotAPI) {
// 	if msg.IsCommand() {
// 		command := msg.Command()
// 		switch command {
// 		case "start":
// 			sendMessage(bot, msg.Chat.ID, "Welcome to the bookkeeping bot! Type /master for available commands.")
// 		case "master":
// 			sendMessage(bot, msg.Chat.ID, "Available Records:\n"+
// 				"Today new transaction(1 slip)\n"+
// 				"Today payment(0 slip)\n"+
// 				"Total Chinese Yuan:45652\n"+
// 				"Exchange rate:8.5000\n"+
// 				"Per-transaction fee rate:3%\n"+
// 				"Total Payment: 44282.44 Yuan | 5209.70 USDT\n"+
// 				"Paid amount: 0 Yuan | 0 USDT\n"+
// 				"Due amount: 44282.44 Yuan | 5209.70 USDT\n")

// 		default:
// 			sendMessage(bot, msg.Chat.ID, "Unknown command.")
// 		}
// 	} else {
// 		sendMessage(bot, msg.Chat.ID, "Unknown command.")
// 	}
// }

// func handleTimeCommand(msg *tgbotapi.Message, bot *tgbotapi.BotAPI) {
// 	elapsedTime := time.Since(startTime)
// 	sendMessage(bot, msg.Chat.ID, fmt.Sprintf("Bot has been running for %s", elapsedTime))
// }

// func sendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
// 	msg := tgbotapi.NewMessage(chatID, text)
// 	bot.Send(msg)
// }

// // package main

// // import (
// // 	"fmt"
// // 	"log"
// // 	"strings"

// // 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
// // 	"github.com/jinzhu/gorm"
// // 	_ "github.com/jinzhu/gorm/dialects/sqlite"
// // )

// // // Transaction represents a transaction record
// // type Transaction struct {
// // 	gorm.Model
// // 	UserID    int
// // 	UserName  string
// // 	Amount    float64
// // 	Currency  string
// // 	Operation string
// // }

// // var db *gorm.DB

// // func init() {
// // 	var err error
// // 	db, err = gorm.Open("sqlite3", "transactions.db")
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // 	db.AutoMigrate(&Transaction{})
// // }

// // func main() {
// // 	token := "6603083264:AAEmOV1kIuNWYIxgIrj5kye7NfsgP-Ud2m8"
// // 	bot, err := tgbotapi.NewBotAPI(token)
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}

// // 	bot.Debug = true

// // 	log.Printf("Authorized on account %s", bot.Self.UserName)

// // 	u := tgbotapi.NewUpdate(0)
// // 	u.Timeout = 60

// // 	updates, err := bot.GetUpdatesChan(u)

// // 	for update := range updates {
// // 		if update.Message == nil {
// // 			continue
// // 		}

// // 		handleMessage(update.Message, bot)
// // 	}
// // }

// // func handleMessage(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
// // 	if message.IsCommand() {
// // 		switch message.Command() {
// // 		case "start":
// // 			msg := tgbotapi.NewMessage(message.Chat.ID, "Welcome to the transaction bot!")
// // 			bot.Send(msg)
// // 		case "record":
// // 			handleRecordCommand(message, bot)
// // 		default:
// // 			msg := tgbotapi.NewMessage(message.Chat.ID, "Unknown command. Use /start to begin.")
// // 			bot.Send(msg)
// // 		}
// // 	}
// // }

// // func handleRecordCommand(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
// // 	// Parse the message text to get transaction details
// // 	args := strings.Fields(message.CommandArguments())
// // 	if len(args) != 4 {
// // 		msg := tgbotapi.NewMessage(message.Chat.ID, "Invalid command format. Use /record UserID UserName Amount Currency")
// // 		bot.Send(msg)
// // 		return
// // 	}

// // 	// Convert the amount to a float64
// // 	amount := 0.0
// // 	fmt.Sscanf(args[2], "%f", &amount)

// // 	// Create a new transaction record
// // 	transaction := Transaction{
// // 		UserID:    message.From.ID,
// // 		UserName:  args[1],
// // 		Amount:    amount,
// // 		Currency:  args[3],
// // 		Operation: "buy_usdt",
// // 	}

// // 	// Save the transaction record to the database
// // 	if err := db.Create(&transaction).Error; err != nil {
// // 		msg := tgbotapi.NewMessage(message.Chat.ID, "Error recording transaction.")
// // 		bot.Send(msg)
// // 		log.Println(err)
// // 		return
// // 	}

// // 	// Respond with a confirmation message
// // 	response := fmt.Sprintf("Transaction recorded:\nUser: %s\nAmount: %.2f %s", args[1], amount, args[3])
// // 	msg := tgbotapi.NewMessage(message.Chat.ID, response)
// // 	bot.Send(msg)
// // }

// go get -u github.com/go-telegram-bot-api/telegram-bot-api

// package main

// import (
// 	"log"
// 	"strings"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
// )

// func main() {
// 	// Set up the Telegram bot using your Bot Token
// 	bot, err := tgbotapi.NewBotAPI("6603083264:AAEmOV1kIuNWYIxgIrj5kye7NfsgP-Ud2m8")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Set up an update configuration
// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 60

// 	// Get updates from the bot
// 	updates, err := bot.GetUpdatesChan(u)

// 	for update := range updates {
// 		// Check if the update contains a message
// 		if update.Message == nil {
// 			continue
// 		}

// 		// Process the master's commands
// 		if strings.Contains(update.Message.Text, "Ready to keep records now!") {
// 			reply := tgbotapi.NewMessage(update.Message.Chat.ID, "Robot is here now for you.")
// 			bot.Send(reply)
// 		} else if strings.Contains(update.Message.Text, "Please prepare new finance sheet!") {
// 			reply := tgbotapi.NewMessage(update.Message.Chat.ID, "New finance sheet is done.")
// 			bot.Send(reply)
// 		} else if strings.Contains(update.Message.Text, "Set exchange rate Chinese Yuan/ USDT at 8.5") {
// 			reply := tgbotapi.NewMessage(update.Message.Chat.ID, "Already set exchange rate at 8.5")
// 			bot.Send(reply)
// 		} else if strings.Contains(update.Message.Text, "+0") {
// 			// Process the +0 command
// 			// You can customize the response based on your requirements
// 			replyText := "Today new transaction(0 slip)\nToday payment(0 slip)\n" +
// 				"Total Chinese Yuan:0\nExchange rate:8.5000\nPer-transaction fee rate:3%\n" +
// 				"Total Payment: 0 Yuan | 0 USDT\nPaid amount: 0 Yuan | 0 USDT\nDue amount: 0 Yuan | 0 USDT"
// 			reply := tgbotapi.NewMessage(update.Message.Chat.ID, replyText)
// 			bot.Send(reply)
// 		}

// 	}
// }

// package main

// import (
// 	"log"
// 	"strconv"
// 	"strings"

// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
// )

// type Robot struct {
// 	TotalTransactions int
// 	TotalPayments     int
// 	ExchangeRate      float64
// }

// func main() {
// 	bot, err := tgbotapi.NewBotAPI("6603083264:AAEmOV1kIuNWYIxgIrj5kye7NfsgP-Ud2m8")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 60

// 	updates, err := bot.GetUpdatesChan(u)

// 	robot := &Robot{
// 		ExchangeRate: 8.5,
// 	}

// 	for update := range updates {
// 		if update.Message == nil {
// 			continue
// 		}

// 		if strings.Contains(update.Message.Text, "Ready to keep records now!") {
// 			reply := tgbotapi.NewMessage(update.Message.Chat.ID, "Robot is here now for you.")
// 			bot.Send(reply)
// 		} else if strings.Contains(update.Message.Text, "Please prepare new finance sheet!") {
// 			reply := tgbotapi.NewMessage(update.Message.Chat.ID, "New finance sheet is done.")
// 			bot.Send(reply)
// 		} else if strings.Contains(update.Message.Text, "Set exchange rate Chinese Yuan/ USDT at 8.5") {
// 			reply := tgbotapi.NewMessage(update.Message.Chat.ID, "Already set exchange rate at 8.5")
// 			bot.Send(reply)
// 		} else if strings.Contains(update.Message.Text, "+0") {
// 			robot.TotalTransactions++
// 			robot.TotalPayments += 0 // Assuming the payment is always 0 in this example

// 			// Calculate due amount based on the provided information
// 			dueAmount := float64(robot.TotalTransactions) * robot.ExchangeRate * 0.03

// 			replyText := "Today new transaction(0 slip)\nToday payment(0 slip)\n" +
// 				"Total Chinese Yuan:0\nExchange rate:8.5000\nPer-transaction fee rate:3%\n" +
// 				"Total Payment: " + strconv.Itoa(robot.TotalPayments) + " Yuan | 0 USDT\n" +
// 				"Paid amount: " + strconv.Itoa(robot.TotalPayments) + " Yuan | 0 USDT\n" +
// 				"Due amount: " + strconv.FormatFloat(dueAmount, 'f', 2, 64) + " Yuan | 0 USDT"

// 			reply := tgbotapi.NewMessage(update.Message.Chat.ID, replyText)
// 			bot.Send(reply)
// 		} else if strings.Contains(update.Message.Text, "+45652") {
// 			amountStr := strings.TrimPrefix(update.Message.Text, "+")
// 			transactionAmount, err := strconv.Atoi(amountStr)
// 			if err != nil {
// 				log.Println("Error parsing transaction amount:", err)
// 				continue
// 			}

// 			robot.TotalTransactions++
// 			robot.TotalPayments += transactionAmount

// 			// Calculate due amount based on the provided information
// 			dueAmount := float64(robot.TotalPayments) * robot.ExchangeRate * 0.03

// 			replyText := "Today new transaction(1 slip)\nToday payment(0 slip)\n" +
// 				"Total Chinese Yuan:" + strconv.Itoa(robot.TotalPayments) + "\n" +
// 				"Exchange rate:8.5000\nPer-transaction fee rate:3%\n" +
// 				"Total Payment: " + strconv.FormatFloat(dueAmount, 'f', 2, 64) + " Yuan | " +
// 				strconv.FormatFloat((dueAmount/robot.ExchangeRate), 'f', 2, 64) + " USDT\n" +
// 				"Paid amount: 0 Yuan | 0 USDT\n" +
// 				"Due amount: " + strconv.FormatFloat(dueAmount, 'f', 2, 64) + " Yuan | " +
// 				strconv.FormatFloat((dueAmount/robot.ExchangeRate), 'f', 2, 64) + " USDT"

// 			reply := tgbotapi.NewMessage(update.Message.Chat.ID, replyText)
// 			bot.Send(reply)

// 			// Additional calculation message
// 			calculationMsg := strconv.Itoa(transactionAmount) + " Yuan/" +
// 				strconv.FormatFloat(robot.ExchangeRate, 'f', 2, 64) + "=" +
// 				strconv.FormatFloat(float64(transactionAmount)/robot.ExchangeRate, 'f', 2, 64) + " USDT"
// 			replyCalculation := tgbotapi.NewMessage(update.Message.Chat.ID, calculationMsg)
// 			bot.Send(replyCalculation)
// 		}

// 	}
// }

package main

import (
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Robot struct {
	TotalTransactions  int
	TotalChineseAmount int
	TotalPayments      int
	PaidAmount         float64
	DueAmount          float64
	DueAmountUsdt      float64
	ExchangeRate       float64
}

func main() {
	bot, err := tgbotapi.NewBotAPI("6603083264:AAEmOV1kIuNWYIxgIrj5kye7NfsgP-Ud2m8")
	if err != nil {
		log.Fatal(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	robot := &Robot{
		ExchangeRate: 8.5,
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if strings.Contains(update.Message.Text, "Ready to keep records now!") {
			reply := tgbotapi.NewMessage(update.Message.Chat.ID, "Robot is here now for you.")
			bot.Send(reply)
		} else if strings.Contains(update.Message.Text, "Please prepare new finance sheet!") {
			reply := tgbotapi.NewMessage(update.Message.Chat.ID, "New finance sheet is done.")
			bot.Send(reply)
		} else if strings.Contains(update.Message.Text, "Set exchange rate Chinese Yuan/ USDT at 8.5") {
			reply := tgbotapi.NewMessage(update.Message.Chat.ID, "Already set exchange rate at 8.5")
			bot.Send(reply)
		} else if strings.HasPrefix(update.Message.Text, "+") {
			amountStr := strings.TrimPrefix(update.Message.Text, "+")
			transactionAmount, err := strconv.Atoi(amountStr)
			if err != nil {
				log.Println("Error parsing transaction amount:", err)
				continue
			}

			robot.TotalTransactions++
			robot.TotalChineseAmount += transactionAmount

			// Calculate due amount based on the provided information
			afterDecuctionAmount := float64(transactionAmount) * 0.97
			afterDeductionUsdt := float64(afterDecuctionAmount) / 8.5

			robot.DueAmount += afterDecuctionAmount
			robot.DueAmountUsdt += afterDeductionUsdt

			replyText := "Today new transaction(" + strconv.Itoa(robot.TotalTransactions) + " slip)\nToday payment(" + strconv.Itoa(robot.TotalPayments) + " slip)\n" +
				"Total Chinese Yuan:" + strconv.Itoa(robot.TotalChineseAmount) + "\n" +
				"Exchange rate:8.5000\nPer-transaction fee rate:3%\n" +
				"Total Payment: " + strconv.FormatFloat(afterDecuctionAmount, 'f', 2, 64) + " Yuan | " +
				strconv.FormatFloat((afterDeductionUsdt), 'f', 2, 64) + " USDT\n" +
				"Paid amount: 0 Yuan | 0 USDT\n" +
				"Due amount: " + strconv.FormatFloat(robot.DueAmount, 'f', 2, 64) + " Yuan | " +
				strconv.FormatFloat(robot.DueAmountUsdt, 'f', 2, 64) + " USDT"

			reply := tgbotapi.NewMessage(update.Message.Chat.ID, replyText)
			bot.Send(reply)

			// Additional calculation message
			calculationMsg := strconv.Itoa(transactionAmount) + " Yuan/" +
				strconv.FormatFloat(robot.ExchangeRate, 'f', 2, 64) + "=" +
				strconv.FormatFloat(float64(transactionAmount)/robot.ExchangeRate, 'f', 2, 64) + " USDT"
			replyCalculation := tgbotapi.NewMessage(update.Message.Chat.ID, calculationMsg)
			bot.Send(replyCalculation)
		}
	}
}
