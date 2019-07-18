package soundcloud

import "encoding/json"

func (sc SoundCloudAuth) Followings(sc_id string) (*SoundCloudFollowings, error) {

	url := Concat(SOUNDCLOUD_API_URL, "/users/", sc_id, "/followings?client_id=", sc.ClientId)

	resp, err := Request(url)

	if err != nil {
		return nil, err
	}

	/** Load response into SounCloud following struct **/
	var followings SoundCloudFollowings

	err = json.Unmarshal(resp, &followings)
	if err != nil {
		return nil, err
	}

	return &followings, nil
}
