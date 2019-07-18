package soundcloud

import (
	"io/ioutil"
	"log"
	"net/http"
)

func (sc SoundCloudAuth) Resolve(url string) (*[]byte, error) {

	req_url := Concat(SOUNDCLOUD_API_URL, "/resolve?url=", url, "&client_id=", sc.ClientId)

	resp, err := http.Get(req_url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return &body, nil
}
