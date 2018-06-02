package util

func CheckVedio(extension string) (flag bool) {
	if extension == ".ogg" || extension == ".mp4" || extension == ".webm" {
		flag = true
	}
	return
}
