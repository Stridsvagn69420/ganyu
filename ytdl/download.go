package ytdl

import (
	"os"

	"github.com/Stridsvagn69420/ganyu/utils"
	"github.com/Stridsvagn69420/pringo"
)

type Ytdl struct {
	Audio      string `json:"audio"`
	Video      string `json:"video"`
	AudioVideo string `json:"audio+video"`
	Website    string `json:"website"`
}

func DownloadHandle(location int, frmts []Ytdl) {
	if location == -1 {
		utils.Printer.Errorln("Website not found in config!", pringo.Yellow)
		// Print available formats
		PrintAvailable(os.Args[3])
		// Ask user what format to use
		format := utils.Printer.Promtln("Please enter the format you'd like to use: ", pringo.None)
		// Download
		Download(os.Args[3], format, os.Args[4])
	} else {
		switch os.Args[2] {
		case "audio":
			Download(os.Args[3], frmts[location].Audio, os.Args[4])
		case "video":
			Download(os.Args[3], frmts[location].Video, os.Args[4])
		case "combined":
			Download(os.Args[3], frmts[location].AudioVideo, os.Args[4])

		default:
			utils.Printer.Errorln("Invalid media type!", pringo.RedBright)
			os.Exit(1)
		}
	}
}
