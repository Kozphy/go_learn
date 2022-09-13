package learn_bbgo

import (
	"bytes"
)

type Notifier interface {
	NotifyTo(channel string, obj interface{}, args ...interface{})
	Notify(obj interface{}, args ...interface{})
	SendPhotoTo(channel string, buffer *bytes.Buffer)
	SendPhoto(buffer *bytes.Buffer)
}

type Notifiability struct {
	notifiers            []Notifier
	SessionChannelRouter *PatternChannelRouter `json:"-"`
	SymbolChannelRouter  *PatternChannelRouter `json:"-"`
	ObjectChannelRouter  *ObjectChannelRouter  `json:"-"`
}
