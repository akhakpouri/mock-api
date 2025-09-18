package services

type WebChecker func(string) bool
type result struct {
	// url
	string
	// response code
	bool
}

func CheckWebsites(wc WebChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	channel := make(chan result)

	for _, url := range urls {
		/*
			following goroutine function is sending the result (type) to the channel.
		*/
		go func() {
			channel <- result{url, wc(url)}
		}()
	}

	for range urls {
		//receives the value from channel and adds it to the results collection/dictionary.
		r := <-channel
		results[r.string] = r.bool
	}

	return results
}
