package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

//UpdateTemplate downloads the latest master templates from Theme.API
func UpdateTemplate(web *http.Client, url string) error {
	lst, err := FindTemplates(web, url)

	if err != nil {
		return err
	}

	for _, v := range lst {
		err = DownloadTemplate(web, v, url)

		if err != nil {
			return err
		}
	}

	return nil
}

func FindTemplates(web *http.Client, themeUrl string) ([]string, error) {
	fullURL := fmt.Sprintf("%s/asset/html", themeUrl)

	resp, err := web.Get(fullURL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	var result []string
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func DownloadTemplate(web *http.Client, template, themeURL string) error {
	fullURL := fmt.Sprintf("%s/asset/html/%s", themeURL, template)

	resp, err := web.Get(fullURL)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("%v: %s", resp.StatusCode, string(bdy))
	}

	out, err := os.Create("/views/_shared/" + template)

	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	return err
}
