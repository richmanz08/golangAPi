package movies

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"os/exec"
	"path"
)


func createFolder(movie Movie) bool {
	var bucketFolder = "D:\\streamingfile\\"

	// Replace "YourDesiredFolderName" with the name of the folder you want to create
	folderName := movie.DirectoryName
	// Replace "D:\\" with the desired path on the D drive where you want to create the folder
	folderPath := bucketFolder + folderName

	// Create the command to execute
	cmd := exec.Command("cmd", "/c", "mkdir", folderPath)

	errCommand := cmd.Run()
	if errCommand != nil {
		fmt.Println("Error creating folder:", errCommand)
		// os.Exit(1)
		return false
		
	}

	// thumbnail sub folder
	thumbnailPathFolder := bucketFolder + folderName + "\\thumbnail"
	// Create the command to execute
	cmdthumbnail := exec.Command("cmd", "/c", "mkdir", thumbnailPathFolder)
	cmdthumbnail.Run()

	// subtitle sub folder
	subtitlePathFolder := bucketFolder + folderName + "\\subtitle"
	// Create the command to execute
	cmdsubtitle := exec.Command("cmd", "/c", "mkdir", subtitlePathFolder)
	cmdsubtitle.Run()

	// audio sub folder
	audioPathFolder := bucketFolder + folderName + "\\audio"
	// Create the command to execute
	cmdaudio := exec.Command("cmd", "/c", "mkdir", audioPathFolder)
	cmdaudio.Run()

	// video sub folder
	videoPathFolder := bucketFolder + folderName + "\\video"
	// Create the command to execute
	cmdvideo := exec.Command("cmd", "/c", "mkdir", videoPathFolder)
	cmdvideo.Run()
	// os.Exit(1)
	return true
}
func uploadImage(file multipart.File, header *multipart.FileHeader) (string, error) {
	filename := header.Filename
	filepath := path.Join("public", filename)
	out, err := os.Create(filepath)
	if err != nil {
		log.Printf("Error creating file: %s", err)
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Printf("Error copying file: %s", err)
		return "", err
	}

	// Construct the full URL dynamically
	baseURL := "http://localhost:8080" // Change this to your actual base URL
	fullURL := baseURL + "/" + filepath
	return fullURL, nil
}