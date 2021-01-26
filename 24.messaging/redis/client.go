package main

func main() {
	TaskKey := "task"
	broker := New()
	broker.sendTask(TaskKey, "7,2")
}
