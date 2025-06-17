package main

import (
	"fmt"
	"os"
	"os/exec"
)

func downloader() {
	var url string
	fmt.Print("Enter YouTube URL: ")
	fmt.Scanln(&url)

	cmd := exec.Command("yt-dlp", "-x", "--audio-format", "mp3", "--audio-quality", "0", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Error downloading:", err)
	} else {
		fmt.Println("Download complete.")
	}
}

/* Uncomment the following lines to run the downloader function directly
func main() {
	downloader()
}
*/
