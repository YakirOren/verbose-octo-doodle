package files

import (
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// DownloadFileFromUrl returns the string value of a URL body
func DownloadFileFromUrl(url string) (result string, err error) {
	/* #nosec we have disabled the calling function*/
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	if resp == nil {
		return "", fmt.Errorf("failed to get data from url %v", url)
	}

	defer closeReader(resp.Body)

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("failed to get data from url %v, status code: %v", url, resp.StatusCode)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func closeReader(body io.ReadCloser) {
	if body == nil {
		return
	}

	if err := body.Close(); err != nil {
		log.Errorf("failed to closeReader error: %v", err)
	}
}
