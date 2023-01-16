package atmail

import (
	"context"
	"encoding/base64"
	"log"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

func Base64Message(msg EmailMessage) string {
	boundary := randStr(32, "alphanum")

	messageBody := []byte("Content-Type: multipart/mixed; boundary=" + boundary + " \n" +
		"MIME-Version: 1.0\n" +
		"from: " + msg.From + "\n" +
		"to: " + msg.To + "\n" +
		"subject: " + msg.Subject + "\n\n" +

		"--" + boundary + "\n" +
		"Content-Type: text/html; charset=" + string('"') + "UTF-8" + string('"') + "\n" +
		"MIME-Version: 1.0\n" +
		"Content-Transfer-Encoding: 7bit\n\n" +
		msg.Body + "\n\n" +
		generateAttachmentsMessage(msg.Attachments, boundary) +
		"--" + boundary + "--")
	return base64.URLEncoding.EncodeToString(messageBody)
}

func GetGmailService(secret string, accesstokenfile string, scope ...string) (srv *gmail.Service, err error) {
	ctx := context.Background()
	b, err := os.ReadFile(secret)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, scope...)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config, accesstokenfile)

	srv, err = gmail.NewService(ctx, option.WithHTTPClient(client))
	return srv, err
}
