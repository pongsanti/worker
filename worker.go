package worker

import (
	"log"
)

type Worker struct {
	workChannel    chan int
	numberOfWorker int
	workFunction   func(int)
	workerStarted  bool
}

func NewWorker(workFunction func(payload int), numberOfWorker int) *Worker {
	workChannel := make(chan int, 1)

	worker := &Worker{
		workChannel:    workChannel,
		numberOfWorker: numberOfWorker,
		workFunction:   workFunction,
	}
	return worker
}

func (w *Worker) workingFunction(number int) {
	for {
		val, ok := <-w.workChannel
		if !ok {
			log.Print("Channel closed")
			break
		} else {
			log.Printf("Running from worker No.%d\n", number+1)
			w.workFunction(val)
		}
	}
}

func (w *Worker) Start() {
	if !w.workerStarted {
		for i := 0; i < w.numberOfWorker; i++ {
			go w.workingFunction(i)
			log.Printf("Worker number %d started!\n", i+1)
		}
		w.workerStarted = true
	} else {
		log.Print("Worker already started")
	}
}

func (w *Worker) Send(payload int) {
	w.workChannel <- payload
	log.Printf("Sent payload %d to the channel\n", payload)
}
