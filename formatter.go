package ylog

type Formatter interface {
	Format(entry *Entry) error
}
