package app

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (a *App) ReplyToMsg(srcmsg *tgbotapi.Message, text string) {
	msg := tgbotapi.NewMessage(srcmsg.Chat.ID, tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, text))
	msg.ParseMode = tgbotapi.ModeMarkdownV2
	msg.ReplyToMessageID = srcmsg.MessageID

	_, _ = a.bot.Send(msg)
}

func (a *App) SendToChat(chatId int64, text string) {
	msg := tgbotapi.NewMessage(chatId, tgbotapi.EscapeText(tgbotapi.ModeMarkdownV2, text))
	msg.ParseMode = tgbotapi.ModeMarkdownV2

	_, _ = a.bot.Send(msg)
}
