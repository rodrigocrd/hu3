package cpf

import "fmt"

type CPF struct {
	inscription int
	digit       int
}

func NewCpf(inscription int, digits int) (*CPF, error) {
	return &CPF{
		inscription: inscription,
		digit:       digits,
	}, nil
}

func (c *CPF) Formatted() string {

	baseInscription := fmt.Sprintf("%9d", c.inscription)

	return fmt.Sprintf("%03s.%03s.%03s-%02d",
		baseInscription[0:3],
		baseInscription[3:6],
		baseInscription[6:9],
		c.digit,
	)
}

func (c *CPF) String() string {
	baseInscription := fmt.Sprintf("%9d", c.inscription)

	return fmt.Sprintf("%09s%02d",
		baseInscription,
		c.digit,
	)
}
