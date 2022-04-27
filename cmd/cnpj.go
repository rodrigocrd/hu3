package cmd

import (
	"fmt"
	"github.com/rodrigocrd/hu3/pkg/domain/cnpj"

	"github.com/spf13/cobra"
)

// cnpjCmd represents the cnpj command
var cnpjCmd = &cobra.Command{
	Use:   "cnpj",
	Short: "Gera um número de CNPJ aleatório",
	Long:  `Gera um número de CNPJ aleatório`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cnpj.CNPJGenerator{}.Generate().Formatted())
	},
}

func init() {
	rootCmd.AddCommand(cnpjCmd)
}
