package cmd

import (
	"log"
	"strings"

	"github.com/go-programming-tour-book/tour/internal/word"
	"github.com/spf13/cobra"
)

var str string
var mode int8

func init() {
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "please input string")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "please input mode of transferation")
}

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderscoreToLowerCamelCase
	ModeCamelCaseToUnderscore
)

var desc = strings.Join([]string{
	"the mode of transfering word format:",
	"1: To Upper",
	"2: To Lower",
	"3: Underscore To UpperCamelCase",
	"4: Underscore To LowerCamelCase",
	"5, CamelCase To Underscore",
}, "\n")

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "transfer word format",
	Long:  "support many word formats trasfer",
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUppperCamelCase(str)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderScore(str)
		default:
			log.Fatalf("no support for this mode, please check help doc")
		}
		log.Printf("Output: %s", content)
	},
}
