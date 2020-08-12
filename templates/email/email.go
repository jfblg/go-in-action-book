package main

type EmailMessage stuct {
	From, Subject, Body string
	To []string
}

type EmailCredentials struct {
	Username, Password, Server string
	Port int
}

const emailTemplate = `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}

Body: {{.Body}}
`

var t *template.Template

func init() {
	t = template.New("email")
	t.Parse(emailTemplate)
}

func main() {
	message := &EmailMessage {
		From: "me@example.com",
		To: []string{"you@exampl.ecom"},
		Subject: "Test",
		Body: "Just a message",
	}

	var body bytes.Buffer
	t.Execute(&body, message)
	
	authCreds := &EmailCredentials{
		Username: "myUsername",
		Password: "myPass",
		Server: "smtp.example.com",
		Port: 25,
		}
		auth := smtp.PlainAuth("",
		authCreds.Username,
		authCreds.Password,
		authCreds.Server,
		)
		smtp.SendMail(authCreds.Server+":"+strconv.Itoa(authCreds.Port),
		auth,
		message.From,
		message.To,
		body.Bytes())
		}

}