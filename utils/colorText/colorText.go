package colorText

import "fmt"

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"

func BlueText(text string) {
	fmt.Println(Blue + text + Reset)
}

func RedText(text string) {
	fmt.Errorf(Red + text + Reset)
}

func GreenText(text string) {
	fmt.Println(Green + text + Reset)
}

func CyanText(text string) {
	fmt.Println(Cyan + text + Reset)
}

func MagentaText(text string) {
	fmt.Println(Magenta + text + Reset)
}

func YellowText(text string) {
	fmt.Println(Yellow + text + Reset)
}
