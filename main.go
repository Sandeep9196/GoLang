package main

import (
	"log"
	"math"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Robot struct {
	TotalTransactions   int
	TotalChineseAmount  float64
	TotalPayments       int
	afterDeduction      float64
	afterDeductionUsdt  float64
	TotalPaidAmount     float64
	TotalPaidAmountUsdt float64
	PaidAmount          float64
	PaidAmountUsdt      float64
	DueAmount           float64
	DueAmountUsdt       float64
	ExchangeRate        float64
}

func main() {
	bot, err := tgbotapi.NewBotAPI("6654622790:AAGgV-mwffMfHw81MRwzCC9060Xb1NwKWGQ")
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

		if strings.Contains(update.Message.Text, "开始") {
			reply := tgbotapi.NewMessage(update.Message.Chat.ID, "机器人已开启，请开始记账。")
			bot.Send(reply)
		} else if strings.Contains(update.Message.Text, "清除今日账单") {
			reply := tgbotapi.NewMessage(update.Message.Chat.ID, "今日账单已清除，可重新开始记录")
			bot.Send(reply)
		} else if strings.Contains(update.Message.Text, "设置汇率8.5") {
			reply := tgbotapi.NewMessage(update.Message.Chat.ID, "固定汇率设置成功， 当前固定汇率为8.5")
			bot.Send(reply)
		} else if strings.HasPrefix(update.Message.Text, "+") {
			amountStr := strings.TrimPrefix(update.Message.Text, "+")
			transactionAmount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				log.Println("Error parsing transaction amount:", err)
				continue
			}

			if transactionAmount != 0 {
				robot.TotalTransactions++
			} else {
				robot.TotalTransactions = 0
				robot.TotalChineseAmount = 0
				robot.TotalPayments = 0
				robot.afterDeduction = 0
				robot.afterDeductionUsdt = 0
				robot.TotalPaidAmount = 0
				robot.TotalPaidAmountUsdt = 0
				robot.PaidAmount = 0
				robot.PaidAmountUsdt = 0
				robot.DueAmount = 0
				robot.DueAmountUsdt = 0

			}
			robot.TotalChineseAmount += float64(transactionAmount)
			beforeDeductionUsdt := float64(transactionAmount) / robot.ExchangeRate
			// Calculate due amount based on the provided information
			afterDeductionAmount := float64(transactionAmount) * 0.97
			afterDeductionUsdt := float64(afterDeductionAmount) / robot.ExchangeRate
			robot.TotalPaidAmount += afterDeductionAmount
			robot.TotalPaidAmountUsdt += afterDeductionUsdt
			robot.DueAmount += afterDeductionAmount
			robot.DueAmountUsdt += afterDeductionUsdt
			robot.afterDeduction = afterDeductionAmount
			robot.afterDeductionUsdt = afterDeductionUsdt

			currentTime := time.Now().Format("15:04")

			if transactionAmount != 0 {
				if appendingString == "" {
					appendingString = currentTime + "   " + strconv.FormatFloat(transactionAmount, 'f', 2, 64) + "/" + strconv.FormatFloat(robot.ExchangeRate, 'f', 2, 64) + "=" + strconv.FormatFloat(beforeDeductionUsdt, 'f', 2, 64) + "U\n"
				} else {

					appendingString += " " + currentTime + "   " + strconv.FormatFloat(transactionAmount, 'f', 2, 64) + "/" + strconv.FormatFloat(robot.ExchangeRate, 'f', 2, 64) + "=" + strconv.FormatFloat(beforeDeductionUsdt, 'f', 2, 64) + "U\n"

				}
			} else {
				appendingString = ""
			}
			lineOfDashes := strings.Repeat("-", 50)

			replyText := "<b>今日入款(" + strconv.Itoa(robot.TotalTransactions) + " 笔)</b>\n" + lineOfDashes + "\n" + appendingString + lineOfDashes + "\n<b>今日下发(" + strconv.Itoa(robot.TotalPayments) + " 笔)</b>\n" + lineOfDashes +
				"\n<b>总入款:</b>" + strconv.FormatFloat(robot.TotalChineseAmount, 'f', 2, 64) + "\n" +
				"<b>汇率:</b>" + strconv.FormatFloat(robot.ExchangeRate, 'f', 2, 64) + "\n<b>交易费率:</b>3%\n" + lineOfDashes + "\n" +
				"<b>应下发:</b> " + strconv.FormatFloat(robot.TotalPaidAmount, 'f', 2, 64) + " | " +
				strconv.FormatFloat(robot.TotalPaidAmountUsdt, 'f', 2, 64) + " U\n" +
				"<b>已下发:</b> " + strconv.FormatFloat(robot.PaidAmount, 'f', 2, 64) + " | " + strconv.FormatFloat(robot.PaidAmountUsdt, 'f', 2, 64) + " U\n" +
				"<b>未下发:</b> " + strconv.FormatFloat(robot.DueAmount, 'f', 2, 64) + " | " +
				strconv.FormatFloat(math.Abs(robot.DueAmountUsdt), 'f', 2, 64) + " U"

			reply := tgbotapi.NewMessage(update.Message.Chat.ID, replyText)
			reply.ParseMode = tgbotapi.ModeHTML
			bot.Send(reply)
		} else if strings.HasPrefix(update.Message.Text, "-") {
			amountStr := strings.TrimPrefix(update.Message.Text, "-")
			transactionAmount, err := strconv.ParseFloat(amountStr, 64)
			if err != nil {
				log.Println("Error parsing transaction amount:", err)
				continue
			}

			robot.TotalPayments++
			// Multiply the transaction amount by the exchange rate (8.5)

			currentTime := time.Now().Format("15:04")

			if appendingPaymentString == "" {
				appendingPaymentString = currentTime + " " + strconv.FormatFloat(transactionAmount, 'f', 2, 64) + "U\n"
			} else {
				appendingPaymentString += currentTime + " " + strconv.FormatFloat(transactionAmount, 'f', 2, 64) + "U\n"

			}
			lineOfDashes := strings.Repeat("-", 50)

			PaidUsdt := float64(transactionAmount)
			PaidAmount := float64(PaidUsdt) * robot.ExchangeRate

			robot.PaidAmount += PaidAmount
			robot.PaidAmountUsdt += PaidUsdt
			robot.DueAmount -= PaidAmount
			robot.DueAmountUsdt -= PaidUsdt

			// replyText := "<b>Today new transaction(" + strconv.Itoa(robot.TotalTransactions) + " slip)</b>\n" + lineOfDashes + "\n" + appendingString + lineOfDashes + "\n<b>Today payment(" + strconv.Itoa(robot.TotalPayments) + " slip)</b>\n" + lineOfDashes + "\n" + appendingPaymentString +
			// 	"\n<b>Total Chinese Yuan:</b>" + strconv.FormatFloat(robot.TotalChineseAmount, 'f', 2, 64) + "\n" +
			// 	"<b>Exchange rate:</b>8.5000\n<b>Per-transaction fee rate:</b>3%\n" + lineOfDashes + "\n" +
			// 	"<b>Total Payment:</b>" + strconv.FormatFloat(robot.TotalPaidAmount, 'f', 2, 64) + "  Yuan |  " + strconv.FormatFloat(robot.TotalPaidAmountUsdt, 'f', 2, 64) + " USDT\n" +
			// 	"<b>Paid amount:</b> " + strconv.FormatFloat(robot.PaidAmount, 'f', 2, 64) + " Yuan | " + strconv.FormatFloat(robot.PaidAmountUsdt, 'f', 2, 64) + " USDT\n" +
			// 	"<b>Due amount:</b> " + strconv.FormatFloat(robot.DueAmount, 'f', 2, 64) + " Yuan | " +
			// 	strconv.FormatFloat(math.Abs(robot.DueAmountUsdt), 'f', 2, 64) + " USDT"

			replyText := "<b>今日入款(" + strconv.Itoa(robot.TotalTransactions) + " 笔)</b>\n" + lineOfDashes + "\n" + appendingString + lineOfDashes + "\n<b>今日下发(" + strconv.Itoa(robot.TotalPayments) + " 笔)</b>\n" + lineOfDashes + "\n" + appendingPaymentString +
				"\n<b>总入款:</b>" + strconv.FormatFloat(robot.TotalChineseAmount, 'f', 2, 64) + "\n" +
				"<b>汇率:</b>" + strconv.FormatFloat(robot.ExchangeRate, 'f', 2, 64) + "\n<b>交易费率:</b>3%\n" + lineOfDashes + "\n" +
				"<b>应下发:</b> " + strconv.FormatFloat(robot.TotalPaidAmount, 'f', 2, 64) + " | " +
				strconv.FormatFloat(robot.TotalPaidAmountUsdt, 'f', 2, 64) + " U\n" +
				"<b>已下发:</b> " + strconv.FormatFloat(robot.PaidAmount, 'f', 2, 64) + " | " + strconv.FormatFloat(robot.PaidAmountUsdt, 'f', 2, 64) + " U\n" +
				"<b>未下发:</b> " + strconv.FormatFloat(robot.DueAmount, 'f', 2, 64) + " | " +
				strconv.FormatFloat(math.Abs(robot.DueAmountUsdt), 'f', 2, 64) + " U"

			reply := tgbotapi.NewMessage(update.Message.Chat.ID, replyText)
			reply.ParseMode = tgbotapi.ModeHTML
			bot.Send(reply)
		} else if strings.HasPrefix(update.Message.Text, "设置动态汇率") {

			trimRate := strings.TrimPrefix(update.Message.Text, "设置动态汇率")
			dynamicExchangeRate, err := strconv.ParseFloat(trimRate, 64)
			if err != nil {
				log.Println("Error parsing dynamic exchange rate:", err)
				continue
			}

			robot.ExchangeRate = dynamicExchangeRate
			replyText := "汇率更新为 " + trimRate

			reply := tgbotapi.NewMessage(update.Message.Chat.ID, replyText)
			reply.ParseMode = tgbotapi.ModeHTML
			bot.Send(reply)
		}
	}
}
