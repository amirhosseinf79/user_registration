package repository

type SMSRepository interface {
	SendOne(to string, message string) error
	SendMany(receptors []string, message string) error
}
