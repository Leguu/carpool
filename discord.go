package main

import (
	"flag"
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

var token = flag.String("token", "", "The token of the Discord bot")
var legu = "437154938527678465"
var discord *discordgo.Session

func init() {
	flag.Parse()
	if token == nil || *token == "" {
		panic("token not provided")
	}
	var err error
	discord, err = discordgo.New("Bot " + *token)
	if err != nil {
		panic(err)
	}

	discord.Identify.Intents = discordgo.IntentsGuildMessages
}

func sendMessageToLegu(message string) {
	channel, err := discord.UserChannelCreate(legu)
	if err != nil {
		slog.Error(fmt.Sprintf("could not create discord channel: %s", err))
	} else {
		_, err := discord.ChannelMessageSend(channel.ID, message)
		if err != nil {
			slog.Error(fmt.Sprintf("could not send message: %s", err))
		}
	}
}
