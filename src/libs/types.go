package libs

import (
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

type IClient struct {
	WA *whatsmeow.Client
}

type ICommand struct {
	Name        string
	As          []string
	Description string
	Tags        string
	IsPrefix    bool
	IsOwner     bool
	IsMedia     bool
	IsQuery     bool
	IsGroup     bool
	IsWait      bool
	IsPrivate   bool
	Before      func(conn *IClient, m *IMessage)
	Execute     func(conn *IClient, m *IMessage) bool
}

type IMessage struct {
	Info       types.MessageInfo
	IsOwner    bool
	Body       string
	Text       string
	Args       []string
	Command    string
	Message    *waProto.Message
	Media      whatsmeow.DownloadableMessage
	IsMedia    string
	Expiration uint32
	Quoted     *waProto.ContextInfo
	Reply      func(text string, opts ...whatsmeow.SendRequestExtra) (whatsmeow.SendResponse, error)
	React      func(emoji string, opts ...whatsmeow.SendRequestExtra) (whatsmeow.SendResponse, error)
}
