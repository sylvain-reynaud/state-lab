package main

type Ready struct {
	video *Video
}

func (u *Ready) GetDownloadMessage(user string) string {
	result := "video READY : \n" + u.video.name + "\nby\n" + u.video.author
	return result
}