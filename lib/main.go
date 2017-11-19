package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var Token string

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New(Token)

	// Catch the error
	if err != nil {
		fmt.Println("Error while creating discord session, ", err)
		return
	}

	err = dg.Open()
	// Catch the opening error
	if err != nil {
		fmt.Println("Error while opening connection, ", err)
		return
	}

	fmt.Println("Selfrobot in running. CTRL + C for exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}