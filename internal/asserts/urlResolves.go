package asserts

import "net/http"

func UrlResolves(url string) bool {
	resolves := false

	response, err := http.Get(url)

	if err == nil {
		if response.StatusCode >= 200 && response.StatusCode <= 399 {
			resolves = true
		}
	}

	return resolves
}
