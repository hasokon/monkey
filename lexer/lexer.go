package lexer

type Lexer struct {
	input string
	position int // The current position of input
	readPosition int // The next position of input
	ch byte // The charactor under checking
}

func New(input string) *Lexer {
	l := &Lexer{
		input : input,
	}

	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition = l.readPosition + 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
		case '=':
			tok = newToken(token.ASSIGN, l.ch)
