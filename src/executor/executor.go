package executor

import (
	"TaskExecutorGO/src/models"
	"TaskExecutorGO/src/utils"
	"TaskExecutorGO/src/worker"
	"fmt"
	"sync"
)

func Executor(T int, tasks []*models.Task, wg *sync.WaitGroup, tasksChan chan<- models.Task, tasksChanReceiver <-chan models.Task, resultsChan chan<- models.Result, fileMutex *sync.Mutex) {
	utils.WriteToFile("output.txt", "0")
	for i := 1; i <= T; i++ {
		wg.Add(1)
		go worker.DoWork(tasksChanReceiver, resultsChan, wg, fileMutex)
	}

	for _, task := range tasks {
		tasksChan <- *task
	}

	go func() {
		fmt.Println("Closing tasks channel")
		close(tasksChan)
		wg.Wait()

		close(resultsChan)
	}()
}
