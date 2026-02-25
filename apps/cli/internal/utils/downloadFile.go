package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		HandleError(fmt.Errorf("Could not download from url: %s", url), "downloadFile")
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		ExitIfError(fmt.Errorf("Server returned status: %s", resp.Status), "downloadFile")
	}

	out, err := os.Create(filepath)
	if err != nil {
		HandleError(fmt.Errorf("File path can't be created on %s", filepath), "downloadFile")
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
