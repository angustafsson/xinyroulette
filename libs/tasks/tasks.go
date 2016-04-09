package tasks

import (
	"math/rand"
	"time"
)

type (
	// Task : Struct for each task
	Task struct {
		TaskTitle string
		TaskDesc  string
	}
)

func GetRandomTask() (task Task) {
	slice := tasks()
	rand.Seed(time.Now().Unix())
	return slice[rand.Intn(len(slice))]
}

func tasks() (tasks []Task) {
	tasks = []Task{
		Task{
			TaskTitle: "Blog",
			TaskDesc:  "You are to write a blog with the following features:\n\n- Frontpage\n\tShould list the 5 latest blog posts\n- Single Post Page\n\tWhen clicking on a title send visitor to a single post page viewing that single blog post\n- Admin interface\n\tA place for the blog author to login to access the following tools:\n\t-- New Post\n\t\tWrite new blog posts\n\t-- Delete Post\n\t\tA way to delete written blog posts",
		},
		Task{
			TaskTitle: "Simple \"Crypt\" tool",
			TaskDesc:  "You are to write a simple \"Crypt\" tool that someone can then use to \"encrypt\" and \"decrypt\" a text file.\n\nThe user should run the application with a password and a text-file as its two arguments. The application should the produce another text-file containing the \"encrypted\" content. The content of this new text-file should be un-readable to the human eye. The application should also be able to \"decrypt\" this new file (with the file and password as arguments) and produce a text-file with the correct original content in it.",
		},
		Task{
			TaskTitle: "Compiler",
			TaskDesc: "You are to write a simple compiler. The user should be able to write Hello World on an appropiate interface (depending on language). Maybe use this as inspiration: https://github.com/thejameskyle/the-super-tiny-compiler",
		},
		Task{
			TaskTitle: "Morse code converter",
			TaskDesc: "Write a program that automatically converts English text to Morse code and vice versa.",
		},
		Task{
			TaskTitle: "Pong!",
			TaskDesc: "Write a single-play pong game",
		},
		Task{
			TaskTitle: "Crude AI",
			TaskDesc: "Write a program that plays Rock, Paper, Scissors better than random against a human. Try to exploit that humans are very bad at generating random numbers.",
		},
		Task{
			TaskTitle: "Battle ship",
			TaskDesc: "Write a program that plays Battle Ship against human opponents. It takes coordinates as input and outputs whether that was a hit or not and its own shot’s coordinates.",
		},
	}
	return
}
