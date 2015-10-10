package checkdb

import (
	"io/ioutil"
	"net/http"
)

var (
	run_url string
	body    string
)

func RunPHP(run_url string) (str []byte) {
	resp, _ := http.Get(run_url)

	body, _ := ioutil.ReadAll(resp.Body)

	return body

}
