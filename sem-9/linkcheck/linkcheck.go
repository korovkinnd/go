//go:build !solution

package linkcheck

import (
	"context"
	"net/http"
	"sync"
)

type Result struct {
	Status int
	Err    error
}

func Check(ctx context.Context, client *http.Client, urls []string, limit int) []Result {
	// TODO: schedule the requests and collect results.
	results := make([]Result, len(urls))
	indexes := make(chan int)
	var wg sync.WaitGroup
	for range min(limit, len(urls)) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range indexes {
				results[i] = check(ctx, client, urls[i])
			}
		}()
	}
	for i := range urls {
		indexes <- i
	}
	close(indexes)
	wg.Wait()
	return results
}

func check(ctx context.Context, client *http.Client, u string) Result {
	if err := ctx.Err(); err != nil {
		return Result{Err: err}
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, u, nil)
	if err != nil {
		return Result{Err: err}
	}
	resp, err := client.Do(req)
	if err != nil {
		return Result{Err: err}
	}
	defer resp.Body.Close()
	return Result{Status: resp.StatusCode}
}
