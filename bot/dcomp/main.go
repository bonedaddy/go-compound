package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"flag"

	"github.com/Necroforger/dgwidgets"
	"github.com/bwmarrin/discordgo"
	"github.com/postables/go-compound/client"
	"github.com/tbruyelle/imgur"
)

var (
	configFile  = flag.String("config.file", "config.json", "config file to use")
	imgurClient *imgur.Client
	url         = "https://api.compound.finance/api/v2"
	help        string
)

func init() {
	flag.Parse()
	var hlp string
	hlp = fmt.Sprintf("%scommands:\neth-price: get ethereum price from cmc", hlp)
	hlp = fmt.Sprintf("%s\ndai-price: get dai price from cmc", hlp)
	hlp = fmt.Sprintf("%s\nliqqable: get liquidatable accounts", hlp)
	hlp = fmt.Sprintf("%s\nhealth-check <acct>: get account health", hlp)
	hlp = fmt.Sprintf("%s\ncollateral-value <acct>: get account collateral value", hlp)
	hlp = fmt.Sprintf("%s\nborrow-value <acct>: get account borrow value", hlp)
	help = hlp
}

func main() {
	var discordToken = os.Getenv("DISCORD_TOKEN")
	if discordToken == "" {
		log.Fatal("DISCORD_TOKEN env var is empty")
	}
	// we need to prepend Bot to allow discord
	// to assign permissions properly
	dg, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		fmt.Println("failed to authenticate with discord")
		log.Fatal(err)
	}
	dg.AddHandler(messageCreate)
	if err := dg.Open(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("bot is now running")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func sendHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	if _, err := s.ChannelMessageSend(m.ChannelID, help); err != nil {
		fmt.Println("error sending message: ", err.Error())
		return
	}
}

func liqqable(s *discordgo.Session, m *discordgo.MessageCreate) {
	cl := client.NewClient(url)
	accts, err := cl.GetLiquidatableAccounts()
	if err != nil {
		fmt.Println("failed to get liquidatable accounts ", err.Error())
		return
	}
	p := dgwidgets.NewPaginator(s, m.ChannelID)
	for k, v := range accts {
		// get the account collateral value
		collatValue, err := cl.GetTotalCollateralValueInEth(k)
		if err != nil {
			fmt.Println("failed to get data for account ", k)
			continue
		}
		borrowValue, err := cl.GetTotalBorrowValueInEth(k)
		if err != nil {
			fmt.Println("failed to get data for account ", k)
			continue
		}
		var fields []*discordgo.MessageEmbedField
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:   "Account Address",
			Value:  k,
			Inline: true,
		}, &discordgo.MessageEmbedField{
			Name:   "Account Value (ETH)",
			Value:  fmt.Sprintf("%v", borrowValue+collatValue),
			Inline: true,
		}, &discordgo.MessageEmbedField{
			Name:   "Account Health",
			Value:  fmt.Sprintf("%v", v),
			Inline: true,
		})
		embed := &discordgo.MessageEmbed{
			Author:      &discordgo.MessageEmbedAuthor{},
			Color:       0x00ff00,
			Description: "Liquidatable Account",
		}
		embed.Fields = fields
		p.Add(embed)
	}
	p.SetPageFooters()

	// enable looping
	p.Loop = true

	// When the paginator is done listening set the colour to yellow
	p.ColourWhenDone = 0xffff

	// Stop listening for reaction events after five minutes
	p.Widget.Timeout = time.Minute * 1

	// start
	p.Spawn()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// parse the message contents based off string fields
	args := strings.Fields(m.Content)
	if len(args) == 0 {
		return
	}
	// ensure the first field is a valid invocation of dpinner
	if args[0] != "!moneybags" {
		fmt.Println("invalid invocation")
		return
	}
	if len(args) == 1 && args[0] == "!moneybags" {
		sendHelp(s, m)
		return
	}
	// commands which only requrie two arguments
	if len(args) == 2 {
		if args[1] == "eth-price" {
			val, err := RetrieveUsdPrice("ethereum")
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "[ERROR]: failed to get eth price: "+err.Error())
			}
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("eth price: $%.3fUSD\n", val))
		}
		if args[1] == "dai-price" {
			val, err := RetrieveUsdPrice("dai")
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "[ERROR]: failed to get dai price: "+err.Error())
			}
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("eth price: $%.3fUSD\n", val))
		}
		if args[1] == "liqqable" {
			liqqable(s, m)
			return
		}
	}
	if len(args) > 2 {
		// If the message is "ping" reply with "Pong!"
		if args[1] == "ping" {
			if _, err := s.ChannelMessageSend(m.ChannelID, "Pong!"); err != nil {
				fmt.Println(err)
			}
			return
		}
		// If the message is "pong" reply with "Ping!"
		if args[1] == "pong" {
			if _, err := s.ChannelMessageSend(m.ChannelID, "Ping!"); err != nil {
				fmt.Println(err)
			}
			return
		}
		if args[1] == "health-check" {
			cl := client.NewClient(url)
			resp, err := cl.GetAccount(args[2])
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "[ERROR]: failed to get account health: "+err.Error())
				return
			}
			if len(resp.Accounts) == 0 {
				s.ChannelMessageSend(m.ChannelID, "[ERROR]: an unexpected error occurred")
			}
			health, err := strconv.ParseFloat(resp.Accounts[0].Health.Value, 64)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "[ERROR]: failed to get account health: "+err.Error())
				return
			}
			var status string
			if health < 1.0 {
				status = "danger"
			} else if health < 1.15 {
				status = "risky"
			} else if health < 1.25 {
				status = "okay"
			} else {
				status = "good"
			}
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("health: %v\nstatus: %s\n", health, status))
		}
		if args[1] == "collateral-value" {
			cl := client.NewClient(url)
			value, err := cl.GetTotalCollateralValueInEth(args[2])
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "[ERROR]: failed to get collateral value: "+err.Error())
				return
			}
			if _, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("collateral value: %vETH", value)); err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		if args[1] == "borrow-value" {
			cl := client.NewClient(url)
			value, err := cl.GetTotalBorrowValueInEth(args[2])
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "[ERROR]: failed to get collateral value: "+err.Error())
				return
			}
			if _, err := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("borrow value: %vETH", value)); err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}
}
