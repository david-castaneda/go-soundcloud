package soundcloud

func Auth(client_id string) *SoundCloudAuth {
	return &SoundCloudAuth{ClientId: client_id}
}
