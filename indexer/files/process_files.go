package files

import (
	"log"
)

// Process the file
func ProcessFiles(datasetFilePath string, inboxFlag bool) []EmailJsonData {

	log.Println("Phase 1 Started")

	//Stage 1 file decompression
	decompressionErr := Untar(datasetFilePath, inboxFlag)

	if decompressionErr != nil {
		log.Fatalf("Error decompressing file: %s", decompressionErr)
	}

	log.Println("Phase 2 Started")

	//Stage 2 List Files
	filePaths, filePathsErr := GetFilePaths("./enron_mail_20110402")

	if filePathsErr != nil {
		log.Fatalf("Error getting file paths for indexation: %s", filePathsErr)
	}

	log.Println("Phase 3 Started")

	//Stage 3 Parse Files
	parsedJsonFiles := ParseFilesByPath(filePaths)

	return parsedJsonFiles
}
