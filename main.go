package main

import (
	"fmt"

	"github.com/dvwallin/xinyroulette/libs/languages"
	"github.com/dvwallin/xinyroulette/libs/tasks"
)

func main() {
	lang := languages.GetRandomLanguage()
	task := tasks.GetRandomTask()
	fmt.Printf("\n\n---- TASK ----\nWrite %s in %s (%s)\n---- DESC ----\n%s\n\n\n", task.TaskTitle, lang.LanguageName, lang.LanguageDocURL, task.TaskDesc)
}
