package main

import (
	"fmt"
	"github.com/pkg/profile"
	"io/ioutil"
	"log"
	"net/http"
	//"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	defer profile.Start(profile.CPUProfile).Stop()

	uri := "http://localhost:9090"
	count := 100
	start := time.Now()

	// for i := 0; i < count; i++ {
	//      response, err := http.Get(uri)
	//      if err != nil {
	//              log.Fatal(err)
	//      } else {
	//              defer response.Body.Close()
	//              _, err := ioutil.ReadAll(response.Body)
	//              if err != nil {
	//                      log.Fatal(err)
	//              }
	//      }
	// }

	var wg sync.WaitGroup

	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(uri string) {
			response, err := http.Get(uri)
			if err != nil {
				log.Fatal(err)
			} else {
				defer response.Body.Close()
				_, err := ioutil.ReadAll(response.Body)
				if err != nil {
					log.Fatal(err)
				}
				wg.Done()
			}
		}(uri)
	}
	wg.Wait()

	elapsed := time.Since(start)

	fmt.Printf("Execution time: %s\n", elapsed)
}
