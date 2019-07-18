package soundcloud

const SOUNDCLOUD_API_URL string = "https://api.soundcloud.com"
const SOUNDCLOUD_BASE_URL string = "https://soundcloud.com"

type SoundCloud interface {
	Resolve(username string) (*SoundCloudUser, error)
	GetPlaylists(username string) []SoundCloudPlaylist
}

type SoundCloudAuth struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SoundCloudUser struct {
	Id                   uint32   `json:"id"`
	FirstName            string   `json:"first_name,omitempty"`
	LastName             string   `json:"last_name,omitempty"`
	FullName             string   `json:"full_name,omitempty"`
	Username             string   `json:"username,omitempty"`
	AvatarURL            string   `json:"avatar_url,omitempty"`
	URI                  string   `json:"uri,omitempty"`
	FollowingsCount      uint32   `json:"followings_count,omitempty"`
	FollowersCount       uint32   `json:"followers_count,omitempty"`
	Online               bool     `json:"online,omitempty"`
	CommentsCount        uint32   `json:"comments_count,omitempty"`
	LikesCount           uint32   `json:"likes_count,omitempty"`
	Plan                 string   `json:"plan,omitempty"`
	PlaylistCount        uint32   `json:"playlist_count,omitempty"`
	Website              string   `json:"website,omitempty"`
	Kind                 string   `json:"kind,omitempty"`
	Country              string   `json:"country,omitempty"`
	Subscriptions        []string `json:"subscriptions,omitempty"`
	City                 string   `json:"city,omitempty"`
	Description          string   `json:"description,omitempty"`
	TrackCount           uint32   `json:"track_count,omitempty"`
	DiscogsName          string   `json:"discogs_name,omitempty"`
	PublicFavoritesCount uint32   `json:"public_favorites_count,omitempty"`
	RepostsCount         uint32   `json:"reposts_count,omitempty"`
	MySpaceName          string   `json:"myspace_name,omitempty"`
	WebsiteTitle         string   `json:"website_title,omitempty"`
	Permalink            string   `json:"permalink,omitempty"`
	PermalinkURL         string   `json:"permalink_url,omitempty"`
	LastModified         string   `json:"last_modified,omitempty"`
}

type SoundCloudPlaylist struct {
	Id            uint32          `json:"id"`
	Title         string          `json:"title,omitempty"`
	Tracks        []PlaylistTrack `json:"tracks,omitempty"`
	User          SoundCloudUser  `json:"user,omitempty"`
	PlaylistType  string          `json:"playlist_type,omitempty"`
	Uri           string          `json:"uri,omitempty"`
	ArtworkURL    string          `json:"artwork_url,omitempty"`
	Streamable    bool            `json:"streamable,omitempty"`
	Downloadable  bool            `json:"downloadable,omitempty"`
	UserId        uint32          `json:"user_id,omitempty"`
	Duration      uint32          `json:"duration,omitempty"`
	Kind          string          `json:"kind,omitempty"`
	Type          string          `json:"type,omitempty"`
	Permalink     string          `json:"permalink,omitempty"`
	PermalinkURL  string          `json:"permalink_url,omitempty"`
	RepostsCount  uint32          `json:"reposts_count,omitempty"`
	LikesCount    uint32          `json:"likes_count,omitempty"`
	Genre         string          `json:"genre,omitempty"`
	PurchaseURL   string          `json:"purchase_url,omitempty"`
	ReleaseDay    uint            `json:"release_day,omitempty"`
	ReleaseMonth  uint            `json:"release_month,omitempty"`
	ReleaseYear   uint            `json:"release_year,omitempty"`
	Description   string          `json:"description,omitempty"`
	LabelName     string          `json:"label_name,omitempty"`
	TagList       string          `json:"tag_list,omitempty"`
	TrackCount    uint32          `json:"track_count,omitempty"`
	LastModified  string          `json:"last_modified,omitempty"`
	License       string          `json:"license,omitempty"`
	Sharing       string          `json:"sharing,omitempty"`
	PurchaseTitle string          `json:"purchase_title,omitempty"`
	EmbeddableBy  string          `json:"embeddable_by,omitempty"`
	LabelId       uint32          `json:"label_id,omitempty"`
	Release       string          `json:"release,omitempty"`
	Ean           string          `json:"ean,omitempty"`
	CreatedAt     string          `json:"created_at,omitempty"`
}

type PlaylistTrack struct {
	Id                  uint32         `json:"id"`
	User                SoundCloudUser `json:"user,omitempty"`
	State               string         `json:"state,omitempty"`
	CommentCount        uint32         `json:"comment_count,omitempty"`
	Release             string         `json:"release,omitempty"`
	OriginalContentSize uint32         `json:"original_content_size,omitempty"`
	TrackType           string         `json:"track_type,omitempty"`
	OriginalFormat      string         `json:"original_format,omitempty"`
	Streamable          bool           `json:"streamable,omitempty"`
	DownloadURL         string         `json:"download_url,omitempty"`
	LastModified        string         `json:"last_modified,omitempty"`
	FavoritingCount     uint32         `json:"favoriting_count,omitempty"`
	Kind                string         `json:"kind,omitempty"`
	PurchaseURL         string         `json:"purchase_url,omitempty"`
	ReleaseDay          uint           `json:"release_day,omitempty"`
	ReleaseMonth        uint           `json:"release_month,omitempty"`
	ReleaseYear         uint           `json:"release_year,omitempty"`
	Sharing             string         `json:"sharing,omitempty"`
	AttachmentURI       string         `json:"attachments_uri,omitempty"`
	License             string         `json:"license,omitempty"`
	UserId              uint32         `json:"user_id,omitempty"`
	WaveformURL         string         `json:"waveform_url,omitempty"`
	Permalink           string         `json:"permalink,omitempty"`
	PermalinkURL        string         `json:"permalink_url,omitempty"`
	PlaybackCount       uint32         `json:"playback_count,omitempty"`
	Downloadable        bool           `json:"downloadable,omitempty"`
	CreatedAt           string         `json:"created_at,omitempty"`
	Description         string         `json:"description,omitempty"`
	Duration            uint32         `json:"duration,omitempty"`
	VideoURL            string         `json:"video_url,omitempty"`
	ArtworkURL          string         `json:"artwork_url,omitempty"`
	TagList             string         `json:"tag_list,omitempty"`
	Genre               string         `json:"genre,omitempty"`
	RepostsCount        uint32         `json:"reposts_count,omitempty"`
	LabelName           string         `json:"label_name,omitempty"`
	Commentable         bool           `json:"commentable,omitempty"`
	Bpm                 string         `json:"bpm,omitempty"`
	Policy              string         `json:"policy,omitempty"`
	KeySignature        string         `json:"key_signature,omitempty"`
	Isrc                string         `json:"isrc,omitempty"`
	Uri                 string         `json:"uri,omitempty"`
	DownloadCount       uint32         `json:"download_count,omitempty"`
	LikesCount          uint32         `json:"likes_count,omitempty"`
	PurchaseTitle       string         `json:"purchase_title,omitempty"`
	EmbeddableBy        string         `json:"embeddable_by,omitempty"`
	MonetizationModel   string         `json:"monetization_model,omitempty"`
	StreamURL           string         `json:"stream_url,omitempty"`
	LabelId             string         `json:"label_id,omitempty"`
}

type SoundCloudFollowers struct {
	Collection []SoundCloudUser `json:"collection"`
}

type SoundCloudFollowings struct {
	Collection []SoundCloudUser `json:"collection"`
}
