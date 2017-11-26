package emojis

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func Emojis(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID != s.State.User.ID {
		return
	}

	emojis := map[string]string{
		"innocent": "ʘ‿ʘ",
		"cutebear": "ʕ•ᴥ•ʔ",
		"disagree": "٩◔̯◔۶",
		"seal": "(ᵔᴥᵔ)",
		"injured": "(҂◡_◡)",
		"careless": "◔_◔",
	}

	thinking := "thinking360"
	if strings.Contains(m.Content, thinking) {
		message := strings.Replace(m.Content, thinking, "http://s.gc.gy/thinking360.webp", -1)
		s.ChannelMessageEdit(m.ChannelID, m.ID, message)
	}

	for command, emoji := range emojis {
		if strings.Contains(m.Content, command) {
			message := strings.Replace(m.Content, command, "`" + emoji + "`", -1)
			s.ChannelMessageEdit(m.ChannelID, m.ID, message)
		}
	}
}
