package main

type Uploading struct {
	video *Video
}

func (u *Uploading) GetDownloadMessage(user string) string {
	if u.video.author == user {
		return "YOUR video not ready : \n" + u.video.name + "\nby\n" + u.video.author
	}
	return "video not ready : \n" + u.video.name + "\nby\n" + u.video.author
}