package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/bwmarrin/discordgo"
)

var token = os.Getenv("DISCORD_TOKEN")
var legu = os.Getenv("DISCORD_LEGU_ID")
var discord *discordgo.Session

func init() {
	var err error
	discord, err = discordgo.New("Bot " + token)
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
