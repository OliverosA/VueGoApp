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

// create the data for json
func CreateEmailJsonData(emailFile io.Reader) (EmailJsonData, error) {
	// struct EmailJsonData
	email := EmailJsonData{}
	emailContent := ""
	scanner := bufio.NewScanner(emailFile)

	// scanner emails
	for scanner.Scan() {
		textLine := scanner.Text()

		if strings.Contains(textLine, "Subject:") {
			// getting subject email
			subjectLine := strings.Split(textLine, "Subject:")

			// if subject isnt empty
			if len(subjectLine) > 0 {
				email.Subject = subjectLine[1]
			}
		}

		// get email "from"
		if strings.Contains(textLine, "From:") {
			senderLine := strings.Split(textLine, "From:")

			if len(senderLine) > 0 {
				email.From = senderLine[1]
			}
		}

		// get email "To"
		if strings.Contains(textLine, "To:") {
			recipientLine := strings.Split(textLine, "To:")

			if len(recipientLine) > 0 {
				email.To = recipientLine[1]
			}
		}

		// get email content
		if strings.Contains(textLine, "X-FileName:") {
			for scanner.Scan() {
				emailContent += scanner.Text() + " "
			}
		}

	}

	// set all the email content
	email.Content = emailContent

	return email, nil
}

func ParseEmailFileToJson(emailFilePath string) EmailJsonData {

	// open the file
	emailFile, emailFileErr := os.Open(emailFilePath)

	if emailFileErr != nil {
		log.Printf("Error reading file %s: %s", emailFilePath, emailFileErr)
	}

	/*
	*	return the email parsed with:
	*	- Subject
	*	- To
	*	- From
	*	- Content
	 */
	emailJsonData, _ := CreateEmailJsonData(emailFile)

	return emailJsonData
}
