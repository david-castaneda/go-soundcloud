package soundcloud

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (sc SoundCloudAuth) Users(username string) (*SoundCloudUser, error) {

	url := Concat(SOUNDCLOUD_API_URL, "/users/", username, "?client_id=", sc.ClientId)

	resp, err := http.Get(url)

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

	/** Load response into SounCloud playlist struct **/
	var usr SoundCloudUser

	err = json.Unmarshal(body, &usr)
	if err != nil {
		return nil, err
	}

	return &usr, nil
}
