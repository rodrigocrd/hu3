package cmd

import (
	"fmt"
	"github.com/rodrigocrd/hu3/pkg/domain/cpf"

	"github.com/spf13/cobra"
)

var cpfCmd = &cobra.Command{
	Use:   "cpf",
	Short: "Gera um número de CPF aleatório",
	Long:  `Gera um número de CPF aleatório`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cpf.CPFGenerator{}.Generate().Formatted())
	},
}

func init() {
	rootCmd.AddCommand(cpfCmd)
}
