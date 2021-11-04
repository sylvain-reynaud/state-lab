package main

type State interface {
	GetDownloadMessage(user string) string
}