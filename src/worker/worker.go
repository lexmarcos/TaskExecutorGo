package worker

import (
	"TaskExecutorGO/src/models"
	"TaskExecutorGO/src/utils"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

func DoWork(tasksChan <-chan models.Task, results chan<- models.Result, wg *sync.WaitGroup, fileMutex *sync.Mutex) {
	defer wg.Done()

	for task := range tasksChan {
		startTime := time.Now()

		time.Sleep(time.Duration(task.Cost * 1000))
		var result models.Result

		if task.Type == "WRITE" {
			fileMutex.Lock()
			err := error(nil)
			result, err = writeTask(task, startTime)
			if err != nil {
				fileMutex.Unlock()
				continue
			}
			fileMutex.Unlock()
		} else {
			var err = error(nil)
			result, err = readTask(task, startTime)
			if err != nil {
				continue
			}
		}

		results <- result
	}
}

func writeTask(task models.Task, startTime time.Time) (models.Result, error) {
	fmt.Println("WRITING from file. Task ID: ", task.ID)
	currentValue, err := utils.ReadFromFile("output.txt")
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return models.Result{}, err
	}

	currentValueInt, err := strconv.Atoi(strings.TrimSpace(currentValue))
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return models.Result{}, err
	}

	valueToWrite := currentValueInt + task.Value
	utils.WriteToFile("output.txt", strconv.Itoa(valueToWrite))

	duration := time.Since(startTime)
	result := *models.NewResult(task.ID, valueToWrite, duration)

	return result, nil
}

func readTask(task models.Task, startTime time.Time) (models.Result, error) {
	fmt.Println("Reading from file. Task ID: ", task.ID)
	currentFileValue, _ := utils.ReadFromFile("output.txt")
	currentFileValueInt, _ := strconv.Atoi(strings.TrimSpace(currentFileValue))
	duration := time.Since(startTime)
	result := *models.NewResult(task.ID, currentFileValueInt, duration)
	return result, nil
}
