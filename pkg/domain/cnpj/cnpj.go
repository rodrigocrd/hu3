package cnpj

import "fmt"

type CNPJ struct {
	inscription int
	index       int
	digit       int
}

func NewCNPJ(inscription int, index int, digits int) (*CNPJ, error) {
	return &CNPJ{
		inscription: inscription,
		index:       index,
		digit:       digits,
	}, nil
}

func (c *CNPJ) Formatted() string {

	baseInscription := fmt.Sprintf("%08d", c.inscription)

	return fmt.Sprintf("%02s.%03s.%03s/%04d-%02d",
		baseInscription[0:2],
		baseInscription[2:5],
		baseInscription[5:8],
		c.index,
		c.digit,
	)
}

func (c *CNPJ) String() string {
	return fmt.Sprintf("%08d%04d%02d",
		c.inscription,
		c.index,
		c.digit,
	)
}
