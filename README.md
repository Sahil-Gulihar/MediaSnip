
# Go Project Setup Guide

This guide will walk you through setting up and initializing the Go project to utilize the functionalities provided in the `main.go` file.

## Prerequisites

Before you begin, ensure you have the following installed on your system:

- Go Programming Language: [Install Go](https://golang.org/doc/install)
- YouTube-DL: [Install YouTube-DL](https://github.com/ytdl-org/youtube-dl#installation)
- FFmpeg: [Install FFmpeg](https://ffmpeg.org/download.html)

## Project Initialization

1. Clone the repository or create a new directory for your project.

   ```bash
   git clone <repository-url>
   ```

2. Navigate to the project directory.

   ```bash
   cd <project-directory>
   ```

3. Create a new Go module (if not already initialized).

   ```bash
   go mod init <module-name>
   ```

## Running the Project

1. Open the `main.go` file in your preferred text editor.

2. Make sure to update the values of `inputFilePath`, `outputFilePath`, and `outputAudio` variables if needed.
   It takes input in both HH:MM:SS format or simple SS only 

3. Run the project using the following command:

   ```bash
   go run main.go
   ```

4. Follow the on-screen prompts to interact with the program:

   - Select the desired operation (Video Download Full, Download Audio, Trimmed Video, or Audio Trim).
   - Provide the necessary input such as URL, start time, and end time as prompted.

## Demo:
![Demo Video](output.mp4)



## Additional Notes

- This project uses external commands such as `yt-dlp` and `ffmpeg`, make sure they are properly installed and accessible in your system's PATH.
- Ensure that you have appropriate permissions to read, write, and execute files in the specified directories.
- Modify the code as per your requirements and extend its functionalities as needed.

---
