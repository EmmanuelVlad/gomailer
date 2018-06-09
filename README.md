# gomailer

Quick mailer in go

Example:
```go
mail := gomailer.New()
mail.Remote = "localhost:25"
mail.Sender = "sender@gomailer.com"
mail.Recipient = "recipient@gomailer.com"
mail.Content = "example"
if err := mail.Send(); err != nil {
    panic(err)
}
```