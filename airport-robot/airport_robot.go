package airportrobot

// Greeter is an interface that defines the greeting method signatures regardless the language.
type Greeter interface {
	LanguageName() string
	Greet(string) string
}

// Italian is a struct that contains the concrete implementation of the greeter in italian language
type Italian struct {
	vLanguageName string
}

func (italian Italian) LanguageName() string {
	return "Italian"
}
func (italian Italian) Greet(name string) string {
	return "Ciao " + name + "!"
}

// Portuguese is a struct that contains the concrete implementation of the greeter in italian language
type Portuguese struct {
	vLanguageName string
}

func (portuguese Portuguese) LanguageName() string {
	return "Portuguese"
}

func (portuguese Portuguese) Greet(name string) string {
	return "Olá " + name + "!"
}

// SayHello is a function that greets the user with the language associated to them
// by the implementation of the greeter.
func SayHello(name string, specificGreeter Greeter) string {
	return "I can speak " + specificGreeter.LanguageName() + ": " + specificGreeter.Greet(name)
}

// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
