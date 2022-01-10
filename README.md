# Lab : State design pattern
Lab to learn State design pattern

A web API enables to upload and download a file.
When a user creates a video, before the upload finish, a download link is given to the user.

## TODO
Code this feature on the web API :

When the video’s author calls the download link & the video is still uploading :

=> it shows “Your video is not ready”

When a user calls the download link & the video is still uploading :

=> it shows “The video is not uploading”

When anyone calls the download link & the video is ready :

=> it shows “the video is downloaded”

