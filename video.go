package main

type Video struct {
	name string
	author string
}

func (v *Video) GetDownloadMessage() string {
	result := "Downloading video : \n" + v.name + "\nby\n" + v.author
	return result
}