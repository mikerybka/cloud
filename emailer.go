package cloud

type Emailer interface {
	SendEmail(to, subject, body string) error
}
