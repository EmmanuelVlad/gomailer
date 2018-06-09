# gomailer

Quick mailer in go

Example:
```go
mail := gomailer.New()
mail.Remote = "localhost:25"
mail.Sender = "sender@gomailer.com"
mail.Recipient = "recipient@gomailer.com"
mail.Headers.Subject = "Example"
mail.Headers.ContentType = "text/html"
mail.Content = "<h1>Example</h1>"
if err := mail.Send(); err != nil {
    panic(err)
}
```