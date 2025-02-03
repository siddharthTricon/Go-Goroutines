package main

import (
	"context"
	"fmr"
	"fmt"
	"net/http"
	"time"
)

func fetchURL(ctx context.Context, url string, results chan<- string, errors chan<- error){
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		errors <- fmt.Errorf("Failed to create request for %s: %v", url, err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil{
		errors <- fmt.Errorf("failed to fetch %s: %v", url, err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK{
		errors <- fmt.Errorf("non-200 respones from %s: %d", url, resp.StatusCode)
		return
	}

	results <- fmt.Sprintf("Successfully fetched %s with status: %d", url, resp.StatusCode)
	
}

func main(){

	urls := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	results := make(chan string)
	errors := make(chan error)

	for _, url := range urls{
		go fetchURL(ctx, url, results, errors)
	}

	done := make(chan struct{})

	go func(){
		for range urls{
			select{
				
			}
		}
	}
}