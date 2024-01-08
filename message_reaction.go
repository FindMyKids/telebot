package telebot

import "strconv"

type MessageReaction struct {
	// The Chat containing the message the user reacted to
	Chat *Chat `json:"chat"`

	// Optional. The ActorChat on behalf of which the reaction was changed, if the user is anonymous
	ActorChat *Chat `json:"actor_chat"`

	// Optional. The User that changed the reaction, if the user isn't anonymous
	User *User `json:"user"`

	// Unique identifier of the message inside the chat
	MessageId int `json:"message_id"`

	// Date of the change in Unix time
	Date int64 `json:"date"`

	// Previous list of reaction types that were set by the user
	OldReactions []Reaction `json:"old_reaction"`

	// New list of reaction types that have been set by the user
	NewReactions []Reaction `json:"new_reaction"`
}

type Reaction struct {
	// Type of the reaction, always “emoji”
	Type string `json:"type"`

	// Reaction emoji
	Emoji string `json:"emoji"`
}

func (mr *MessageReaction) MessageSig() (messageID string, chatID int64) {
	return strconv.Itoa(mr.MessageId), mr.Chat.ID
}
