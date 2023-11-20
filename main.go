package main

import (
	"TaskExecutorGO/src/executor"
	"TaskExecutorGO/src/models"
	"TaskExecutorGO/src/utils"
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	var N, T, E int
	utils.LoadVariables(&N, &T, &E)
	var wg sync.WaitGroup
	var fileMutex = sync.Mutex{}
	bufferResultsChannel := int64(math.Pow(10, float64(N)))
	resultsChannel := make(chan models.Result, bufferResultsChannel)
	tasksChannel := make(chan models.Task, T)
	tasks := utils.Load(N, E)
	fmt.Println("Expected value:", utils.GetExpectedValue(tasks))
	startTime := time.Now()
	executor.Executor(T, tasks, &wg, tasksChannel, tasksChannel, resultsChannel, &fileMutex)
	duration := time.Since(startTime)

	for result := range resultsChannel {
		fmt.Println("Result:", result.Result, "ID:", result.ID, "Time:", result.Time)
	}

	fmt.Println("Expected value:", utils.GetExpectedValue(tasks))
	fmt.Println("Duration:", duration)
}
