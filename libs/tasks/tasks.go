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
	}
	return
}
