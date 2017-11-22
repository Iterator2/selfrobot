package poll

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func Poll(s *discordgo.Session, m *discordgo.MessageCreate) {
	if (m.Author.ID != s.State.User.ID) || !strings.HasPrefix(m.Content, "?sondage") {
		return
	}

	question := strings.Replace(m.Content, "?sondage ", "", 1)
	embed := discordgo.MessageEmbed{
		Title: question + " 📣",
		Description: "Oui ✅\n Non ❌",
		Color: 16771337,
	}

	s.ChannelMessageDelete(m.ChannelID, m.ID)
	message, error := s.ChannelMessageSendEmbed(m.ChannelID, &embed)

	if error != nil {
		s.ChannelMessageSend(m.ChannelID, "An error happened")
		return
	}

	s.MessageReactionAdd(m.ChannelID, message.ID, "✅")
	s.MessageReactionAdd(m.ChannelID, message.ID, "❌")
}