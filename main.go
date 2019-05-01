package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"mod.go/models"
)

// Env - Used for database dependency injection
type Env struct {
	db models.Datastore
}

func (env *Env) workerGetTasks(tasks chan models.Task, shutdown chan os.Signal, wg *sync.WaitGroup) {
	for {
		select {
		case sig := <-shutdown:
			shutdown <- sig
			//We're done!
			fmt.Println("Shutdwon received!")
			wg.Done()
			return
		default:
			fmt.Println("Scraper started")
			fmt.Println(len(tasks))

			rndSleep := rand.Int63n(120)
			fmt.Println("No task in que, waiting", rndSleep, "seconds until rescraping tasks")
			time.Sleep(time.Duration(rndSleep) * time.Second)

			fmt.Println("Doing my tasks")
		}
	}
}

func initDB() {
	db, err := models.NewDB("postgres://user:pass@localhost/bookstore")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	fmt.Println(wg)

	// go worker(messages,shutdown,wg)

	// Initzialize channels
	tasks := make(chan models.Task, 100)
	shutdown := make(chan os.Signal, 1)

	go getTasks(tasks, shutdown, wg)

	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	// fileURL := "http://api.filer.net/dl/h4ugxfbwsrmm1s2s.json"

	// if err := DownloadFile(fileURL, "south1.part01.rar"); err != nil {
	// 	panic(err)
	// }

	for i := 0; i < 3; i++ {
		// Register the new goroutine in the WaitGroup
		wg.Add(1)
		// Create new goroutine "thread"
		go downloadWorker(tasks, shutdown, wg)
	}

	//close(shutdown) //Signal to shutdown
	wg.Wait() //Wait for the goroutine to shutdown
	close(shutdown)
	close(tasks)
}

// DownloadFile worker will download a url and store it in local filepath.
// It writes to the destination file as it downloads it, without
// loading the entire file into memory.
func (env *Env) workerDownload(tasks <-chan models.Task, shutdown chan os.Signal, wg *sync.WaitGroup) {
	for {
		select {
		case sig := <-shutdown:
			shutdown <- sig
			//We're done!
			fmt.Println("Shutdwon received!")
			wg.Done()
			return
		case models.Task := <-tasks:
			// Create the HTTP Download request and add the basic auth to it
			req, err := http.NewRequest("GET", models.Task.Link, nil)
			req.SetBasicAuth("cyrill.naef@gmail.com", "KmxxspHE")

			client := &http.Client{}

			// Create the file
			out, err := os.Create(models.Task.FileName)
			if err != nil {
				fmt.Println(err.Error())
			}
			defer out.Close()

			// Get the data
			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err.Error())
			}
			defer resp.Body.Close()

			// Write the body to file
			_, err = io.Copy(out, resp.Body)
			if err != nil {
				fmt.Println(err.Error())
			}

			// return nil
		default:
			rndSleep := rand.Int63n(120)
			fmt.Println("No task in que, waiting", rndSleep, "seconds until rescraping tasks")
			time.Sleep(time.Duration(rndSleep) * time.Second)
		}
	}
}
