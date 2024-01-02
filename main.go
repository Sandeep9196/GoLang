package main

import (
	"log"
	"strconv"
	"strings"
	"time"

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
	appendingString := ""
	appendingPaymentString := ""
	// afterDecuctionAmount := ""
	// afterDeductionUsdt := ""
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
			beforeDeductionUsdt := float64(transactionAmount) / 8.5
			// Calculate due amount based on the provided information
			afterDecuctionAmount := float64(transactionAmount) * 0.97
			afterDeductionUsdt := float64(afterDecuctionAmount) / 8.5

			robot.DueAmount += afterDecuctionAmount
			robot.DueAmountUsdt += afterDeductionUsdt

			currentTime := time.Now().Format("15:04")

			if appendingString == "" {
				appendingString = currentTime + "   " + strconv.Itoa(transactionAmount) + " Yuan/8.5=" + strconv.FormatFloat(beforeDeductionUsdt, 'f', 2, 64) + "USDT\n"
			} else {
				appendingString += " " + currentTime + "   " + strconv.Itoa(transactionAmount) + " Yuan/8.5=" + strconv.FormatFloat(beforeDeductionUsdt, 'f', 2, 64) + "USDT\n"

			}
			lineOfDashes := strings.Repeat("-", 50)

			replyText := "<b>Today new transaction(" + strconv.Itoa(robot.TotalTransactions) + " slip)</b>\n" + appendingString + lineOfDashes + "\n<b>Today payment(" + strconv.Itoa(robot.TotalPayments) + " slip)</b>\n" + lineOfDashes +
				"\n<b>Total Chinese Yuan:</b>" + strconv.Itoa(robot.TotalChineseAmount) + "\n" +
				"<b>Exchange rate:</b>8.5000\n<b>Per-transaction fee rate:</b>3%\n" +
				"<b>Total Payment:</b> " + strconv.FormatFloat(afterDecuctionAmount, 'f', 2, 64) + " Yuan | " +
				strconv.FormatFloat((afterDeductionUsdt), 'f', 2, 64) + " USDT\n" + lineOfDashes + "\n" +
				"<b>Paid amount:</b> 0 Yuan | 0 USDT\n" +
				"<b>Due amount:</b> " + strconv.FormatFloat(robot.DueAmount, 'f', 2, 64) + " Yuan | " +
				strconv.FormatFloat(robot.DueAmountUsdt, 'f', 2, 64) + " USDT"

			reply := tgbotapi.NewMessage(update.Message.Chat.ID, replyText)
			reply.ParseMode = tgbotapi.ModeHTML
			bot.Send(reply)
		} else if strings.HasPrefix(update.Message.Text, "-") {
			amountStr := strings.TrimPrefix(update.Message.Text, "-")
			transactionAmount, err := strconv.Atoi(amountStr)
			if err != nil {
				log.Println("Error parsing transaction amount:", err)
				continue
			}

			robot.TotalPayments++
			// Multiply the transaction amount by the exchange rate (8.5)
			deductedAmount := float64(transactionAmount) * 8.5
			robot.TotalChineseAmount = int(float64(robot.TotalChineseAmount) - deductedAmount)

			currentTime := time.Now().Format("15:04")

			if appendingPaymentString == "" {
				appendingPaymentString = currentTime + "   " + strconv.FormatFloat(deductedAmount, 'f', 2, 64) + " Yuan/8.5=" + strconv.Itoa(transactionAmount) + "USDT\n"
			} else {
				appendingPaymentString += currentTime + "   " + strconv.FormatFloat(deductedAmount, 'f', 2, 64) + " Yuan/8.5=" + strconv.Itoa(transactionAmount) + "USDT\n"

			}
			lineOfDashes := strings.Repeat("-", 50)

			replyText := "<b>Today new transaction(" + strconv.Itoa(robot.TotalTransactions) + " slip)</b>\n" + appendingString + lineOfDashes + "\n<b>Today payment(" + strconv.Itoa(robot.TotalPayments) + " slip)</b>\n" + lineOfDashes + "\n" + appendingPaymentString +
				"\n<b>Total Chinese Yuan:</b>" + strconv.Itoa(robot.TotalChineseAmount) + "\n" +
				"<b>Exchange rate:</b>8.5000\n<b>Per-transaction fee rate:</b>3%\n" +
				"<b>Total Payment:</b>  Yuan |  0 USDT\n" + lineOfDashes + "\n" +
				"<b>Paid amount:</b> 0 Yuan | 0 USDT\n" +
				"<b>Due amount:</b> " + strconv.FormatFloat(robot.DueAmount, 'f', 2, 64) + " Yuan | " +
				strconv.FormatFloat(robot.DueAmountUsdt, 'f', 2, 64) + " USDT"

			reply := tgbotapi.NewMessage(update.Message.Chat.ID, replyText)
			reply.ParseMode = tgbotapi.ModeHTML
			bot.Send(reply)
		}
	}
}
