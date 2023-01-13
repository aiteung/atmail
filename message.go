package atmail

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
)

func generateAttachmentsMessage(attachments []FileAttachment, boundary string) (content string) {
	content = ""
	for _, attachment := range attachments {
		a := "--" + boundary + "\n" +
			"Content-Type: " + attachment.MIMEType + "; name=" + string('"') + attachment.Name + string('"') + " \n" +
			"MIME-Version: 1.0\n" +
			"Content-Transfer-Encoding: base64\n" +
			"Content-Disposition: attachment; filename=" + string('"') + attachment.Name + string('"') + " \n\n" +
			chunkSplit(attachment.Base64, 76, "\n")
		content = content + a
	}
	return content
}

func randStr(strSize int, randType string) string {

	var dictionary string

	if randType == "alphanum" {
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	}

	var strBytes = make([]byte, strSize)
	_, _ = rand.Read(strBytes)
	for k, v := range strBytes {
		strBytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(strBytes)
}

func chunkSplit(body string, limit int, end string) string {
	var charSlice []rune

	// push characters to slice
	for _, char := range body {
		charSlice = append(charSlice, char)
	}

	var result = ""

	for len(charSlice) >= 1 {
		if len(charSlice) < limit {
			limit = len(charSlice)
		} else {
			result = result + string(charSlice[:limit]) + end
			charSlice = charSlice[limit:]
		}

	}
	return result
}

func SetMIMEandNameifEmpty(msg *EmailMessage) {
	for i, attachment := range msg.Attachments {
		if attachment.MIMEType == "" {
			msg.Attachments[i].MIMEType = getMIMETypefromBase64(attachment.Base64)
		}
		if attachment.Name == "" {
			msg.Attachments[i].Name = "Attachment"
		}
	}
}

func getMIMETypefromBase64(base64str string) string {
	fileBytes, err := base64.StdEncoding.DecodeString(base64str)
	if err != nil {
		fmt.Println(err)
	}
	return http.DetectContentType(fileBytes)
}
