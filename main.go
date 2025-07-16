package main

import (
	"fmt"
	"os"
	"strconv"
	"encoding/json"
	"log"
)

func main(){

	//слайс тасок
	tasks := []Task{}

	//считали таски с файла и пишем их в слайс, если файла нет или пуст выведется что тасок нету
	tasksFromFile, _ := os.ReadFile("tasks.json")
	
	if err := json.Unmarshal(tasksFromFile, &tasks); err != nil && len(tasks) != 0{
		log.Fatal(err)
	}

	//приветствие если прога запущена без аргументов
	if len(os.Args) == 1{
		greet()
		return 
	}

	if len(os.Args) >= 2 {
		
		if len(os.Args) == 2 && os.Args[1] == "list"{
			if len(tasks) == 0{
				fmt.Println("There is no tasks")
				return
			}
			fmt.Println("\tListing of tasks")
			//выводим листинг тасок
			printTasks(tasks)
			return
		} else if os.Args[1] == "add" && len(os.Args) == 3{
			//добавление таски
			number := len(tasks) + 1
			taska := string(os.Args[2])
			newTask := Task{Id: number, Task: taska}
			tasks = append(tasks, newTask)

			updateId(tasks) //обновляем id тасок
			saveTasks(tasks)//записываем таски в json

		} else if os.Args[1] == "update" && len(os.Args) == 4{
			//изменение таски
			num, err := strconv.Atoi(os.Args[2]) //номер таски
			if err != nil{ log.Fatal(err)}
			taska := os.Args[3]
			for i := 0; i < len(tasks); i++{
				if tasks[i].Id == num{
					tasks[i].Task = taska
				}
			}
			saveTasks(tasks)

		} else if os.Args[1] == "delete" && len(os.Args) == 3{

			if os.Args[2] == "all"{
				//удалить все таски
				if err := os.Truncate("tasks.json", 0); err != nil{
					log.Fatal(err)
				}
				return
			}
			number, err := strconv.Atoi(os.Args[2]) //номер таски для удаления
			number--
			tasks, err = removeById(tasks, number)
			if err != nil{
				log.Fatal(err)
			}
			updateId(tasks) //обновляем id тасок
			saveTasks(tasks) //сохраняем таски в файл
		} else {
		log.Fatal("Wrong input")
		}
	}
}

type Task struct{
	Id 		int
	Task 	string
}

func greet(){
	fmt.Println("\t\t---Task Manager---")
	fmt.Println("Run programm like: Tas [command] [numberTask] [\"taskString\"]")
	fmt.Println("For example: tas add \"Buy Milk\"")
	fmt.Println("\t     tas delete 1")
	fmt.Println("\t     tas delete all")
	fmt.Println("\t     tas list")
	fmt.Println("\t     tas update 1 \"buy milk\"")
	fmt.Println("\t\t---  Good luck ---")
}
//удаление таски по номеру
func removeById(slice []Task, idx int) ([]Task, error){
	if idx < 0 || idx >= len(slice) {
		return nil, fmt.Errorf("нет такого id")
	}
	return append(slice[:idx], slice[idx+1:]...), nil
}
//обновление тасок, после удаления или добавления таски
func updateId(tasks []Task){
	for i := 0; i < len(tasks); i++{
		tasks[i].Id = i + 1
	}
}
//запись тасок в json
func saveTasks(sliceTasks []Task){
	file, _ := os.Create("tasks.json")
	defer file.Close()
	tasksJson, _ := json.Marshal(sliceTasks)
	file.Write(tasksJson)
}
//вывод тасок
func printTasks(tasks []Task){
	for _, t := range tasks{
		fmt.Fprintf(os.Stdout, "%d - %s\n", t.Id, t.Task)
	}
}

