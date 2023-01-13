package atmail

type FileAttachment struct {
	MIMEType string `json:"mimetype,omitempty" bson:"mimetype,omitempty"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Base64   string `json:"base64,omitempty" bson:"base64,omitempty"`
}

type EmailMessage struct {
	From        string           `json:"from,omitempty" bson:"from,omitempty"`
	To          string           `json:"to,omitempty" bson:"to,omitempty"`
	Subject     string           `json:"subject,omitempty" bson:"subject,omitempty"`
	Body        string           `json:"body,omitempty" bson:"body,omitempty"`
	Attachments []FileAttachment `json:"attachments,omitempty" bson:"attachments,omitempty"`
}

type Success struct {
	LabelIds []string `json:"labels,omitempty" bson:"labels,omitempty"`
	Queue    Queue    `json:"queue,omitempty" bson:"queue,omitempty"`
}

type Queue struct {
	Session  string   `json:"session,omitempty" bson:"session,omitempty"`
	Function string   `json:"function,omitempty" bson:"function,omitempty"`
	Scope    []string `json:"scope,omitempty" bson:"scope,omitempty"`
	Base64   string   `json:"base64,omitempty" bson:"base64,omitempty"`
}
