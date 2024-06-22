package lexer

type Lexer struct {
	input        string
	position     int // current position in input
	readPosition int // next position to be read (position + 1)
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}

	return l
}

func (l *Lexer) readChar() {
	totalNoOfCharactersInInput := len(l.input)

	if l.readPosition > totalNoOfCharactersInInput {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition

	l.readPosition += 1

}
