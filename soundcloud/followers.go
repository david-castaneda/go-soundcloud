package soundcloud

import (
	"encoding/json"
)

func (sc SoundCloudAuth) Followers(sc_id string) (*SoundCloudFollowers, error) {

	url := Concat(SOUNDCLOUD_API_URL, "/users/", sc_id, "/followers?client_id=", sc.ClientId)

	resp, err := Request(url)

	if err != nil {
		return nil, err
	}

	/** Load response into SounCloud followers struct **/
	var followers SoundCloudFollowers

	err = json.Unmarshal(resp, &followers)
	if err != nil {
		return nil, err
	}

	return &followers, nil
}
