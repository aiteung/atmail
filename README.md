# atmail
aiteung email helper for gmail. Authentication using access token token.json


## Example sending email


```go

package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/aiteung/atmq/helper"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	b, err := os.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailSendScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := helper.GetClient(config)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	var message gmail.Message

	// Compose the message
	messageStr := []byte(
		"From: youremail@gmail.com\r\n" +
			"To: awangga@gmail.com\r\n" +
			"Subject: My first Gmail API message\r\n\r\n" +
			"Message body goes here!")

	// Place messageStr into message.Raw in base64 encoded format
	message.Raw = base64.URLEncoding.EncodeToString(messageStr)

	// Send the message
	_, err = srv.Users.Messages.Send("me", &message).Do()
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Println("Message sent!")
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