package main

type Video struct {
	name string
	author string
	currentState State
}

func (v *Video) GetDownloadMessage(user string) string {
	return v.currentState.GetDownloadMessage(user)
}

func (v *Video) setState(s State) {
	v.currentState = s
}

func (v *Video) getState() State {
	return v.currentState
}