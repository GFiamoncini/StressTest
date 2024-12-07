package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	// Parâmetros de entrada
	url := flag.String("url", "", "www.URLDEEXEMPLO")
	requests := flag.Int("requests", 0, "100")
	concurrency := flag.Int("concurrency", 1, "30")
	flag.Parse()

	if *url == "" || *requests <= 0 || *concurrency <= 0 {
		fmt.Println("Parâmetros inválidos. Use --url, --requests e --concurrency corretamente.")
		return
	}

	// Canal para resultados
	results := make(chan int, *requests)
	var wg sync.WaitGroup

	start := time.Now()

	// Execução concorrente
	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case results <- makeRequest(*url):
				default:
					return
				}
			}
		}()
	}

	// Aguarda conclusão
	wg.Wait()
	close(results)

	// Processa resultados
	total := 0
	status200 := 0
	statusDist := make(map[int]int)

	for res := range results {
		total++
		if res == 200 {
			status200++
		}
		statusDist[res]++
	}

	elapsed := time.Since(start)

	// Relatório
	fmt.Printf("Tempo total: %v\n", elapsed)
	fmt.Printf("Total de requests: %d\n", total)
	fmt.Printf("Requests com status 200: %d\n", status200)
	fmt.Println("Distribuição de status HTTP:")
	for code, count := range statusDist {
		fmt.Printf("Status %d: %d\n", code, count)
	}
}

func makeRequest(url string) int {
	resp, err := http.Get(url)
	if err != nil {
		return 0
	}
	defer resp.Body.Close()
	return resp.StatusCode
}
