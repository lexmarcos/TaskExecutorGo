package utils

import (
	"TaskExecutorGO/src/models"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func Load(N int, E int) []*models.Task {
	var queue []*models.Task
	rand.NewSource(time.Now().UnixNano())
	totalTasks := math.Pow(10, float64(N))
	writeTasksCount := (float64(E) * totalTasks) / 100
	writeTasksAdded := 0
	fmt.Printf("Generating %.0f tasks, with %.0f write tasks and %.0f read tasks\n", totalTasks, writeTasksCount, totalTasks-writeTasksCount)

	for i := 0; i < int(totalTasks); i++ {
		id := i + 1
		cost := rand.Float64() * 0.01
		value := 0
		taskType := "READ"

		isPossibleToAddWriteTask := writeTasksAdded < int(writeTasksCount)
		numberOfWriteTasksLeft := writeTasksCount - float64(writeTasksAdded)
		tasksLeft := totalTasks - float64(i)

		shouldAddWriteTask := isPossibleToAddWriteTask && rand.Float64() < numberOfWriteTasksLeft/tasksLeft

		if shouldAddWriteTask {
			taskType = "WRITE"
			value = rand.Intn(10)
			writeTasksAdded++
		}

		task := models.NewTask(id, cost, taskType, value)
		queue = append(queue, task)
	}
	return queue
}

func GetExpectedValue(queue []*models.Task) int {
	expectedValue := 0
	for _, task := range queue {
		if task.Type == "WRITE" {
			expectedValue += task.Value
		}
	}
	return expectedValue
}
