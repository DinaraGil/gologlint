package core

type Checker interface {
	Check(string) error
	Name() string
}
