package learn_bbgo

import (
	"regexp"
)

type PatternChannelRouter struct {
	routes map[*regexp.Regexp]string
}

type ObjectChannelHandler func(obj interface{}) (channel string, ok bool)

type ObjectChannelRouter struct {
	routes []ObjectChannelHandler
}
