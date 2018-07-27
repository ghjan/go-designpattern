package barrier

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var timeoutMilliseconds = 5000

type barrierResp struct {
	Resp string
	Err  error
}

func barrier(endpoints ...string) []string {
	requestNumber := len(endpoints)
	results := make([]string, requestNumber)

	var in []chan barrierResp
	for i := 0; i < requestNumber; i++ {
		in = append(in, make(chan barrierResp))
	}

	responses := make([]barrierResp, requestNumber)

	for i, endpoint := range endpoints {
		go makeRequest(in[i], endpoint)

	}

	var hasError bool
	for i := 0; i < requestNumber; i++ {
		resp := <-in[i]
		if resp.Err != nil {
			results[i] = fmt.Sprintln("ERROR: ", resp.Err)
			hasError = true
			break
		}
		responses[i] = resp
	}

	if !hasError {
		for i, resp := range responses {
			results[i] = fmt.Sprintln(resp.Resp)
		}
	}
	for i := 0; i < requestNumber; i++ {
		close(in[i])
	}
	return results
}

func makeRequest(out chan<- barrierResp, url string) {
	res := barrierResp{}
	client := http.Client{
		Timeout: time.Duration(time.Duration(timeoutMilliseconds) * time.Millisecond),
	}

	resp, err := client.Get(url)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	res.Resp = string(byt)
	out <- res
}

func captureBarrierOutput(endpoints ...string) string {

	results := barrier(endpoints...)
	temp := ""
	for _, result := range results {
		if result != "" {
			temp += result
		}
	}
	return temp
}
