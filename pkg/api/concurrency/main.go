package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/domain/repositories"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/services"
	"github.com/FreeCodeUserJack/GoRESTMicroservicePart1/pkg/api/utils/errors"
)

var (
	success map[string]string
	failed map[string]errors.ApiError
)


func getRequests() []repositories.CreateRepoRequest {
	res := make([]repositories.CreateRepoRequest, 0)

	file, err := os.Open("requests.txt")
	if err != nil {
		log.Fatalf("expected no error opening file but got an error: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		line := scanner.Text()
		res = append(res, repositories.CreateRepoRequest{
			Name: line,
		})
	}

	return res
}

func main() {
	requests := getRequests()
	fmt.Printf(fmt.Sprintf("about to process %d requests", len(requests)) + "\n")

	outChan := make(chan createRepoResult)
	buff := make(chan bool, 10)

	var wg sync.WaitGroup

	go handleResults(&wg, outChan)

	for _, req := range requests {
		buff <- true
		wg.Add(1)
		go createRepo(buff, req, outChan)
	}

	wg.Wait()
	close(outChan)

	fmt.Printf("%d %d\n", len(success), len(failed))
}

func handleResults(wg *sync.WaitGroup, outChan chan createRepoResult) {
	for res := range outChan {
		if res.Error != nil {
			failed[res.Request.Name] = res.Error
			continue
		} else {
			success[res.Request.Name] = res.Result.Name
		}
		wg.Done()
	}
}

type createRepoResult struct {
	Request repositories.CreateRepoRequest
	Result *repositories.CreateRepoResponse
	Error errors.ApiError
}

func createRepo(buff chan bool, request repositories.CreateRepoRequest, output chan createRepoResult) {
	result, apiErr := services.RepositoryService.CreateRepo(request)

	output <- createRepoResult{
		Request: request,
		Result: result,
		Error: apiErr,
	}

	<-buff
}