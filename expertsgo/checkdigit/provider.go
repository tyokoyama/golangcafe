package checkdigit

import (
	"errors"
	"unicode"
)

var ErrInvalidArgument = errors.New("Invalid Argument")

type Generator interface {
	Generate(seed string) (int, error)
}

type Verifier interface {
	Verify(code string) bool
}

type Provider interface {
	Verifier
	Generator
}

type luhn struct {}

func (l luhn) Verify(code string) bool {
	if len(code) < 2 {
		return false
	}

	i, err := l.Generate(code[:len(code)-1])

	return err == nil && i == int(code[len(code)-1]-'0')
}

func (l luhn) Generate(seed string) (int, error) {
	if seed == "" {
		return 0, ErrInvalidArgument
	}

	sum, parity := 0, (len(seed) + 1) % 2
	for i, n := range seed {
		if isNotNumber(rune(n)) {
			return 0, ErrInvalidArgument
		}
		d := int(n - '0')
		if i % 2 == parity {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
	}

	return sum * 9 % 10, nil
}

func isNotNumber(n rune) bool {
	return !unicode.IsNumber(n)
}

func NewLuhn() Provider {
	return &luhn{}
}
