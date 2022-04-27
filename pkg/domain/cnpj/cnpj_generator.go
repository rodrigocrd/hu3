package cnpj

import (
	"fmt"
	"github.com/rodrigocrd/hu3/pkg/domain"
	"github.com/rodrigocrd/hu3/pkg/util"
	"math/rand"
	"time"
)

type CNPJGenerator struct {
}

func (cg CNPJGenerator) Generate() domain.Document {
	source := rand.NewSource(time.Now().UnixMicro())
	r := rand.New(source)

	inscription := r.Intn(99_999_999 + 1)
	index := r.Intn(9999 + 1)
	validationDigits := calculateValidationDigits(inscription, index)
	cnpj, err := NewCNPJ(inscription, index, validationDigits)

	if err != nil {
		panic(err)
	}

	return cnpj
}

func calculateValidationDigits(inscription int, index int) int {

	baseCnpj := fmt.Sprintf("%08d%04d", inscription, index)
	weights := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	return util.Checksum([][]int{weights, append([]int{6}, weights...)}, baseCnpj)
}
