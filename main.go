package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var userInput int; 
	var Url, startTime, endTime string

	inputFilePath := "Input1.webm"
	outputFilePath := "output.webm"
	outputAudio := "output.opus"

	err := RemoveFileIfExists(outputFilePath)
	if err != nil {
		fmt.Println("Error removing output file:", err)
		return
	}

	err = RemoveFileIfExists(outputAudio)
	if err != nil {
		fmt.Println("Error removing input file:", err)
		return
	}

	fmt.Println("Enter a number \n 1. Video Download Full \n 2. Download Audio \n 3. Trimmed Video \n 4. Audio Trim.")
	fmt.Scanln(&userInput)

	switch userInput {
	case 1:
		fmt.Println("Enter the URL to download:")
		fmt.Scanln(&Url)
		cmd := exec.Command("yt-dlp", "-o", "output", Url)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Video downloaded and saved as output.mkv")
	case 2:
		fmt.Println("Enter the URL to download:")
		fmt.Scanln(&Url)
		cmd := exec.Command("yt-dlp", "-x", "-o", "output", Url)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Audio saved as output.opus")
	case 3:
		fmt.Println("Enter the URL to download:")
		fmt.Scanln(&Url)
		fmt.Println("Enter Start:")
		fmt.Scanln(&startTime)

		fmt.Println("Enter End:")
		fmt.Scanln(&endTime)
		cmd := exec.Command("yt-dlp", "-o", "Input1", Url)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		cmd = exec.Command("ffmpeg", "-i", "Input1.webm", "-ss", startTime, "-to", endTime, "-c:v", "copy", "-c:a", "copy", "output.webm")
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Video downloaded and saved as output.mkv")
	case 4:
		fmt.Println("Enter the URL to download:")
		fmt.Scanln(&Url)
		fmt.Println("Enter Start:")
		fmt.Scanln(&startTime)
		fmt.Println("Enter End:")
		fmt.Scanln(&endTime)
		cmd := exec.Command("yt-dlp", "-x", "-o", "FullAudio", Url)
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		cmd = exec.Command("ffmpeg", "-ss", startTime, "-to", endTime, "-i", "FullAudio.opus", "output.opus")
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Audio trimmed succesully")
	default:
		fmt.Println("Invalid input. Please enter a number between 1 and 4.")
	}
	err = RemoveFileIfExists(inputFilePath)
	if err != nil {
		fmt.Println("Error removing input file:", err)
		return
	}
	err = RemoveFileIfExists("FullAudio.opus")
	if err != nil {
		fmt.Println("Error removing input file:", err)
		return
	}
}

func RemoveFileIfExists(filePath string) error {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}
	err = os.Remove(filePath)
	if err != nil {
		return err
	}
	fmt.Println("File", filePath, "removed successfully!")
	return nil

}
