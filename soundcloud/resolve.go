package soundcloud

func (sc SoundCloudAuth) Resolve(url string) (*[]byte, error) {

	req_url := Concat(SOUNDCLOUD_API_URL, "/resolve?url=", url, "&client_id=", sc.ClientId)

	resp, err := Request(req_url)

	if err != nil {
		return nil, err
	}

	return &resp, nil
}
