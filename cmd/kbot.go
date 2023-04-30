/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	TeleToken = os.Getenv("TELE_TOKEN")
)

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("kbot %s started", appVersion)
		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Please check TELE_TOKEN env variable. %s", err)
			return
		}

		kbot.Handle(telebot.OnText, handleMessage)
		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)
}

func handleMessage(m telebot.Context) error {
	loc, _ := time.LoadLocation("Europe/Kiev")
	currentTime := time.Now().In(loc)
	switch m.Message().Payload {
	case "hello":
		err := m.Send(fmt.Sprintf("Hello, I'm KBot version %s", appVersion))
		return err
	case "date":
		err := m.Send(currentTime.Format("02.01.2006 15:04:05"))
		return err
	case "day":
		err := m.Send(currentTime.Format("Monday"))
		return err
	default:
		err := m.Send("Sorry, I don't understand you..")
		return err
	}
}
