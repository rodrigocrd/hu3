package cpf

import (
	"fmt"
	"github.com/rodrigocrd/hu3/pkg/domain"
	"github.com/rodrigocrd/hu3/pkg/util"
	"math/rand"
	"time"
)

type CPFGenerator struct {
}

func (cg CPFGenerator) Generate() domain.Document {
	source := rand.NewSource(time.Now().UnixMicro())
	r := rand.New(source)

	inscription := r.Intn(999_999_999 + 1)
	validationDigits := calculateValidationDigits(inscription)
	cpf, err := NewCpf(inscription, validationDigits)

	if err != nil {
		panic(err)
	}

	return cpf
}

func calculateValidationDigits(inscription int) int {

	baseCpf := fmt.Sprintf("%09d", inscription)
	weights := []int{10, 9, 8, 7, 6, 5, 4, 3, 2}

	return util.Checksum([][]int{weights, append([]int{11}, weights...)}, baseCpf)
}
