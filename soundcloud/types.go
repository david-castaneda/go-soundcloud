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
	Id                   int      `json:"id"`
	FirstName            string   `json:"first_name,omitempty"`
	LastName             string   `json:"last_name,omitempty"`
	FullName             string   `json:"full_name,omitempty"`
	Username             string   `json:"username,omitempty"`
	AvatarURL            string   `json:"avatar_url,omitempty"`
	URI                  string   `json:"uri,omitempty"`
	FollowingsCount      int      `json:"followings_count,omitempty"`
	FollowersCount       int      `json:"followers_count,omitempty"`
	Online               bool     `json:"online,omitempty"`
	CommentsCount        int      `json:"comments_count,omitempty"`
	LikesCount           int      `json:"likes_count,omitempty"`
	Plan                 string   `json:"plan,omitempty"`
	PlaylistCount        int      `json:"playlist_count,omitempty"`
	Website              string   `json:"website,omitempty"`
	Kind                 string   `json:"kind,omitempty"`
	Country              string   `json:"country,omitempty"`
	Subscriptions        []string `json:"subscriptions,omitempty"`
	City                 string   `json:"city,omitempty"`
	Description          string   `json:"description,omitempty"`
	TrackCount           int      `json:"track_count,omitempty"`
	DiscogsName          string   `json:"discogs_name,omitempty"`
	PublicFavoritesCount int      `json:"public_favorites_count,omitempty"`
	RepostsCount         int      `json:"reposts_count,omitempty"`
	MySpaceName          string   `json:"myspace_name,omitempty"`
	WebsiteTitle         string   `json:"website_title,omitempty"`
	Permalink            string   `json:"permalink,omitempty"`
	PermalinkURL         string   `json:"permalink_url,omitempty"`
	LastModified         string   `json:"last_modified,omitempty"`
}

type SoundCloudPlaylist struct {
	Id            int             `json:"id"`
	Title         string          `json:"title,omitempty"`
	Tracks        []PlaylistTrack `json:"tracks,omitempty"`
	User          SoundCloudUser  `json:"user,omitempty"`
	PlaylistType  string          `json:"playlist_type,omitempty"`
	Uri           string          `json:"uri,omitempty"`
	ArtworkURL    string          `json:"artwork_url,omitempty"`
	Streamable    bool            `json:"streamable,omitempty"`
	Downloadable  bool            `json:"downloadable,omitempty"`
	UserId        int             `json:"user_id,omitempty"`
	Duration      int             `json:"duration,omitempty"`
	Kind          string          `json:"kind,omitempty"`
	Type          string          `json:"type,omitempty"`
	Permalink     string          `json:"permalink,omitempty"`
	PermalinkURL  string          `json:"permalink_url,omitempty"`
	RepostsCount  int             `json:"reposts_count,omitempty"`
	LikesCount    int             `json:"likes_count,omitempty"`
	Genre         string          `json:"genre,omitempty"`
	PurchaseURL   string          `json:"purchase_url,omitempty"`
	ReleaseDay    string          `json:"release_day,omitempty"`
	ReleaseMonth  string          `json:"release_month,omitempty"`
	ReleaseYear   string          `json:"release_year,omitempty"`
	Description   string          `json:"description,omitempty"`
	LabelName     string          `json:"label_name,omitempty"`
	TagList       []string        `json:"tag_list,omitempty"`
	TrackCount    int             `json:"track_count,omitempty"`
	LastModified  string          `json:"last_modified,omitempty"`
	License       string          `json:"license,omitempty"`
	Sharing       string          `json:"sharing,omitempty"`
	PurchaseTitle string          `json:"purchase_title,omitempty"`
	EmbeddableBy  string          `json:"embeddable_by,omitempty"`
	LabelId       int             `json:"label_id,omitempty"`
	Release       string          `json:"release,omitempty"`
	Ean           string          `json:"ean,omitempty"`
	CreatedAt     string          `json:"created_at,omitempty"`
}

type PlaylistTrack struct {
	Id                  int            `json:"id"`
	User                SoundCloudUser `json:"user,omitempty"`
	State               string         `json:"state,omitempty"`
	CommentCount        int            `json:"comment_count,omitempty"`
	Release             string         `json:"release,omitempty"`
	OriginalContentSize int            `json:"original_content_size,omitempty"`
	TrackType           string         `json:"track_type,omitempty"`
	OriginalFormat      string         `json:"original_format,omitempty"`
	Streamable          bool           `json:"streamable,omitempty"`
	DownloadURL         string         `json:"download_url,omitempty"`
	LastModified        string         `json:"last_modified,omitempty"`
	FavoritingCount     int            `json:"favoriting_count,omitempty"`
	Kind                string         `json:"kind,omitempty"`
	PurchaseURL         string         `json:"purchase_url,omitempty"`
	ReleaseDay          string         `json:"release_day,omitempty"`
	ReleaseMonth        string         `json:"release_month,omitempty"`
	ReleaseYear         string         `json:"release_year,omitempty"`
	Sharing             string         `json:"sharing,omitempty"`
	AttachmentURI       string         `json:"attachments_uri,omitempty"`
	License             string         `json:"license,omitempty"`
	UserId              int            `json:"user_id,omitempty"`
	WaveformURL         string         `json:"waveform_url,omitempty"`
	Permalink           string         `json:"permalink,omitempty"`
	PermalinkURL        string         `json:"permalink_url,omitempty"`
	PlaybackCount       int            `json:"playback_count,omitempty"`
	Downloadable        bool           `json:"downloadable,omitempty"`
	CreatedAt           string         `json:"created_at,omitempty"`
	Description         string         `json:"description,omitempty"`
	Duration            int            `json:"duration,omitempty"`
	VideoURL            string         `json:"video_url,omitempty"`
	ArtworkURL          string         `json:"artwork_url,omitempty"`
	TagList             string         `json:"tag_list,omitempty"`
	Genre               string         `json:"genre,omitempty"`
	RepostsCount        int            `json:"reposts_count,omitempty"`
	LabelName           string         `json:"label_name,omitempty"`
	Commentable         bool           `json:"commentable,omitempty"`
	Bpm                 string         `json:"bpm,omitempty"`
	Policy              string         `json:"policy,omitempty"`
	KeySignature        string         `json:"key_signature,omitempty"`
	Isrc                string         `json:"isrc,omitempty"`
	Uri                 string         `json:"uri,omitempty"`
	DownloadCount       int            `json:"download_count,omitempty"`
	LikesCount          int            `json:"likes_count,omitempty"`
	PurchaseTitle       string         `json:"purchase_title,omitempty"`
	EmbeddableBy        string         `json:"embeddable_by,omitempty"`
	MonetizationModel   string         `json:"monetization_model,omitempty"`
	StreamURL           string         `json:"stream_url,omitempty"`
	LabelId             string         `json:"label_id,omitempty"`
}
