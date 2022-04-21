package ytdl

import (
	"os"

	"github.com/Stridsvagn69420/ganyu/utils"
	"github.com/Stridsvagn69420/pringo"
)

func Download(url string, quality string, output string) {
	if utils.CommandExists("yt-dlp") {
		utils.RunShell(false, "yt-dlp", url, "-f", quality, "-o", output)
	} else if utils.CommandExists("youtube-dl") {
		utils.RunShell(false, "youtube-dl", url, "-f", quality, "-o", output)
	} else {
		utils.Printer.Errorln("Neither youtube-dl nor yt-dlp is installed!", pringo.Red)
		os.Exit(1)
	}
}

func PrintAvailable(url string) {
	if utils.CommandExists("yt-dlp") {
		utils.RunShell(false, "yt-dlp", url, "-F")
	} else if utils.CommandExists("youtube-dl") {
		utils.RunShell(false, "youtube-dl", url, "-F")
	} else {
		utils.Printer.Errorln("Neither youtube-dl nor yt-dlp is installed!", pringo.Red)
		os.Exit(1)
	}
}
