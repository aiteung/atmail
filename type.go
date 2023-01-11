package atmail

type FileAttachment struct {
	MIMEType string `json:"mimetype,omitempty" bson:"filemimetype,omitempty"`
	Name     string `json:"name,omitempty" bson:"filename,omitempty"`
	Base64   string `json:"base64,omitempty" bson:"base64string,omitempty"`
}

type EmailMessage struct {
	From        string           `json:"from,omitempty" bson:"from,omitempty"`
	To          string           `json:"to,omitempty" bson:"to,omitempty"`
	Subject     string           `json:"subject,omitempty" bson:"subject,omitempty"`
	Body        string           `json:"body,omitempty" bson:"body,omitempty"`
	Attachments []FileAttachment `json:"attachments,omitempty" bson:"attachments,omitempty"`
}
