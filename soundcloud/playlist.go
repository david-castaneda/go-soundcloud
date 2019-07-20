package soundcloud

import (
	"encoding/json"
)

func (sc SoundCloudClient) Playlists(username string) (*[]SoundCloudPlaylist, error) {

	url := Concat(SOUNDCLOUD_API_URL, "/users/", username, "/playlists?client_id=", sc.ClientId)

	resp, err := Request(url)

	if err != nil {
		return nil, err
	}

	/** Load response into SounCloud playlist struct **/
	var playlists []SoundCloudPlaylist

	err = json.Unmarshal(resp, &playlists)
	if err != nil {
		return nil, err
	}

	return &playlists, nil
}
