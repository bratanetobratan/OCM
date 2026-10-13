package telegram

import (
	"log"
	"strings"
)

const (
	SaveCmd  = "/save"
	RmvCmd   = "/remove"
	JoinCmd  = "/join"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, username)

	switch text {
	case SaveCmd:
	case RmvCmd:
	case JoinCmd:
	case HelpCmd:
	case StartCmd:
	default:

	}
}

func (p *Processor) savePlayer(chatID int) error {
	p.tg.SendMessage(chatID, msgEnterFullName)

}
