package soundcloud

func Auth(client_id string) *SoundCloudClient {
	return &SoundCloudClient{ClientId: client_id}
}
