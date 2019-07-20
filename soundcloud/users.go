package soundcloud

import (
	"encoding/json"
)

func (sc SoundCloudAuth) Users(username string) (*SoundCloudUser, error) {

	url := Concat(SOUNDCLOUD_API_URL, "/users/", username, "?client_id=", sc.ClientId)

	resp, err := Request(url)

	/** Load response into SounCloud playlist struct **/
	var usr SoundCloudUser

	err = json.Unmarshal(resp, &usr)
	if err != nil {
		return nil, err
	}

	return &usr, nil
}
