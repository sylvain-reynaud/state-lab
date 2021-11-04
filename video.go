package main

type Video struct {
	name string
	author string
	currentState State
}

func (v *Video) GetDownloadMessage(user string) string {
	/* */
}

func (v *Video) setState(s State) {
	/* */
}