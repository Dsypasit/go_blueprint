package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
)

var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar URL")

type TryAvatars []Avatar

func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatarURL
}

type Avatar interface {
	GetAvatarURL(c ChatUser) (string, error)
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

func (AuthAvatar) GetAvatarURL(c ChatUser) (string, error) {
	url := c.AvatarURL()
	if len(url) == 0 {
		return "", ErrNoAvatarURL
	}
	return url, nil
}

type GravatarAvatar struct{}

var UseGravatar GravatarAvatar

func (GravatarAvatar) GetAvatarURL(c ChatUser) (string, error) {
	return fmt.Sprintf("http://www.gravatar.com/avatar/%s", c.UniqueID()), nil
}

type FileSystemAvatar struct{}

var UseFileSystemAvatar FileSystemAvatar

func (FileSystemAvatar) GetAvatarURL(c ChatUser) (string, error) {
	if files, err := ioutil.ReadDir("avatars"); err == nil {
		if err != nil {
			return "", ErrNoAvatarURL
		}
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if match, _ := path.Match(c.UniqueID()+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}
	return "", ErrNoAvatarURL
}
