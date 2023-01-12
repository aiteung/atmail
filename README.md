# atmail
aiteung email helper for gmail. Authentication using access token token.json


## Example sending email


```go
package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aiteung/atmail"
	"google.golang.org/api/gmail/v1"
)

func filereadmime(fname string) (fileData, fileMIMEType string) {
	fileBytes, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fileMIMEType = http.DetectContentType(fileBytes)
	fileData = base64.StdEncoding.EncodeToString(fileBytes)
	return fileData, fileMIMEType
}

func main() {
	msg := &atmail.EmailMessage{
		From:    "Penerbit Buku Pedia<penerbit@bukupedia.co.id>",
		To:      "awangga@gmail.com",
		Subject: "subjek email",
		Body:    "ini bodi isinya ya <b>luar biasa</b>",
	}

	filename := "README.md"
	fileData, fileMIMEType := filereadmime(filename) //"text/plain; charset=utf-8"
	attachment := &atmail.FileAttachment{
		Name:     filename,
		MIMEType: fileMIMEType,
		Base64:   fileData,
	}
	msg.Attachments = append(msg.Attachments, *attachment)

	srv := atmail.GetGmailService("client_secret.json", "token.json", gmail.GmailSendScope)
	var message gmail.Message
	message.Raw = atmail.Base64Message(*msg)
	// Send the message
	resp, err := srv.Users.Messages.Send("me", &message).Do()
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Println(resp.LabelIds)
	}
}

```


## Dev
set env 

GOPROXY=proxy.golang.org

```sh
git tag v0.0.1
git push origin --tags
go list -m github.com/aiteung/atmail@v0.0.1
```
