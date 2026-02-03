package main

func main() {
	var taskManager TaskLister
	taskManager.decode()

	taskManager.Add("Learn Go")
	taskManager.Show()
}
