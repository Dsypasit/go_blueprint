package main

import "errors"

var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar URL")

type Avatar interface {
	GetAvatarURL(c *client) (string, error)
}
