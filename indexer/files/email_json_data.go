package files

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type EmailJsonData struct {
	Subject string `json:"subject"`
	To      string `json:"to"`
	From    string `json:"from"`
	Content string `json:"content"`
}

func CreateEmailJsonData(emailFile io.Reader) (EmailJsonData, error) {
	email := EmailJsonData{}
	emailContent := ""
	finalLineSpace := ""
	scanner := bufio.NewScanner(emailFile)

	for scanner.Scan() {
		textLine := scanner.Text()

		if strings.Contains(textLine, "Subject:") {
			subjectLine := strings.Split(textLine, "Subject:")

			if len(subjectLine) > 0 {
				email.Subject = subjectLine[1]
			}
		}

		if strings.Contains(textLine, "From:") {
			senderLine := strings.Split(textLine, "From:")

			if len(senderLine) > 0 {
				email.From = senderLine[1]
			}
		}

		if strings.Contains(textLine, "To:") {
			recipientLine := strings.Split(textLine, "To:")

			if len(recipientLine) > 0 {
				email.To = recipientLine[1]
			}
		}

		if strings.Contains(textLine, "X-FileName:") {
			for scanner.Scan() {
				emailContent += scanner.Text() + finalLineSpace
			}
		}

	}

	email.Content = emailContent

	return email, nil
}

func ParseEmailFileToJson(emailFilePath string) EmailJsonData {

	emailFile, emailFileErr := os.Open(emailFilePath)

	if emailFileErr != nil {
		log.Printf("Error reading file %s: %s", emailFilePath, emailFileErr)
	}

	emailJsonData, _ := CreateEmailJsonData(emailFile)

	return emailJsonData
}
