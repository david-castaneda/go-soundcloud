package soundcloud

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (sc SoundCloudAuth) Resolve(username string) (*SoundCloudUser, error) {

	/** Construct url **/
	var res_url string = SOUNDCLOUD_API_URL + "/resolve?url="
	var query_url = SOUNDCLOUD_BASE_URL + "/" + username
	var req_url = res_url + query_url + "&client_id=" + sc.ClientId

	/** Make http request **/
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

	/** Load response into SounCloud user struct **/
	var usr SoundCloudUser

	err = json.Unmarshal(body, &usr)
	if err != nil {
		return nil, err
	}

	return &usr, nil
}
