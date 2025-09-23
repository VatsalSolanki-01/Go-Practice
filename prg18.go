package main

import "fmt"

var TaskList []Task

type Task struct {
	task_id    int
	task       string
	date       string
	isCoreTask bool
	taskStatus bool
}

// method returns a Task object
func (t Task) createTask(task_id int, date string, isCoreTask bool, taskStatus bool, task string) Task {
	return Task{task_id: task_id, date: date, isCoreTask: isCoreTask, taskStatus: taskStatus, task: task}
}

func NextDay() {
	for {
		fmt.Println("************************************************************")
		fmt.Println("Plan your next day by bifurcating core tasks and other tasks")
		fmt.Println("enter 1 to plan the Upcoming day :- ")
		fmt.Println("enter 2 to display plan for the day :- ")
		fmt.Println("enter 3 to update the task status :- ")
		fmt.Println("enter 4 to exit :- ")
		fmt.Println("************************************************************")
		var input string
		fmt.Scan(&input)

		if input == "4" {
			break
		} else {
			switch input {
			case "1":
				createTasks()
			case "2":
				displayTasks()
			case "3":
				updateTaskStatus()
			}
		}
	}
}

func createTasks() {
	var task_id int
	var task, date, isThisTaskCore string
	var isCoreTask, taskStatus bool

	fmt.Println("enter id of task :- ")
	fmt.Scan(&task_id)
	fmt.Println("enter the task name only :- ")
	fmt.Scan(&task)
	fmt.Println("enter date in dd/mm/yy format :-  ")
	fmt.Scan(&date)
	fmt.Println("say yes if the task is core task :- ")
	fmt.Scan(&isThisTaskCore)
	if isThisTaskCore == "yes" {
		isCoreTask = true
	}
	var Task Task

	TaskList = append(TaskList, Task.createTask(task_id, date, isCoreTask, taskStatus, task))
}

func displayTasks() {
	if TaskList == nil {
		fmt.Println("No tasks found 404")
		return
	}
	for i, v := range TaskList {
		fmt.Println("taskno. :- ", i+1, "||", "date :- ", v.date, "||", "taskid :-", v.task_id, "||", "is this a core task :- ", v.isCoreTask, "||", "task status :-", v.taskStatus, "||")

	}

}

func updateTaskStatus() {
	var flag bool
	var task_id int
	fmt.Println("enter taskid:-")
	fmt.Scan(&task_id)

	for i, v := range TaskList {
		if v.task_id == task_id {
			TaskList[i].taskStatus = true
			flag = true
			break
		}
	}
	if flag != true {
		fmt.Println("task not found")
	}
}
