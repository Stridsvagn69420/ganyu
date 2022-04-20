package ytdl

import (
	"github.com/Stridsvagn69420/ganyu/utils"
)

func Download(url string, quality string, output string) {
	if utils.CommandExists("yt-dlp") {
		utils.RunShell(false, "yt-dlp", url, "-f", quality, "-o", output)
	} else {
		utils.RunShell(false, "youtube-dl", url, "-f", quality, "-o", output)
	}
}

func PrintAvailable(url string) {
	if utils.CommandExists("yt-dlp") {
		utils.RunShell(false, "yt-dlp", url, "-F")
	} else {
		utils.RunShell(false, "youtube-dl", url, "-F")
	}
}
