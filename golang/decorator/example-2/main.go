package main

import (
	"fmt"
	"time"
)

type Args map[string]string
type Data map[string]string

type Fetcher interface {
	Fetch(args Args) (Data, error)
}

type Repository struct {}

func (r *Repository) Fetch(args Args) (Data, error) {
	if (len(args) == 0) {
		return Data{}, fmt.Errorf("No arguments were provided")
	}

	data := Data{
		"user": "root",
		"pass": "swordfish",
	}
	fmt.Printf("Repo fetched data successfully: %v\n", data)
	return data, nil
}

type Retrier struct {
	RetryCount int
	WaitInterval time.Duration
	Fetcher Fetcher
}

func (r *Retrier) Fetch(args Args) (Data, error)  {
	var err error
	var data Data
	for i:= 0; i < r.RetryCount; i++ {
		fmt.Printf("Retrier number: %d\n", i+1)
		if data, err = r.Fetcher.Fetch(args); err == nil {
			fmt.Printf("Success after %d times\n", i+1)
			return data, err
		}
		fmt.Printf("waiting...\n")
		time.Sleep(r.WaitInterval)
	}
	fmt.Printf("Retrier failed to fetch for %d times\n", r.RetryCount)
	return Data{}, err
}

func main() {
	repo := &Repository{}
	repo.Fetch(Args{"id": "1"})
	
	retrier := &Retrier{
		RetryCount:   3,
		WaitInterval: time.Second,
		Fetcher:      repo,
	}

	retrier.Fetch(Args{})
	retrier.Fetch(Args{"id": "1"})
}