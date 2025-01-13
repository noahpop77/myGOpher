package main

// Implement Format method for PlainText
func (p PlainText) Format() string {
	return p.message
}

// Implement Format method for Bold
func (b Bold) Format() string {
	return "**" + b.message + "**"
}

// Implement Format method for Code
func (c Code) Format() string {
	return "`" + c.message + "`"
}

// Formatter interface with a Format method
type Formatter interface {
	Format() string
}

// PlainText struct
type PlainText struct {
	message string
}

// Bold struct
type Bold struct {
	message string
}

// Code struct
type Code struct {
	message string
}

// Don't Touch below this line

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}
