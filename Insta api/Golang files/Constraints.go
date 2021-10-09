func ParseLocationId(lid LocationId) string

type Api struct {
	ClientId             string
	ClientSecret         string
	AccessToken          string
	EnforceSignedRequest bool
	Header               http.Header
}


type Attribution struct {
	Website   string
	ItunesUrl string
	Name      string
}

type Caption Comment
type Comment struct {
	CreatedTime StringUnixTime `json:"created_time"`
	Text        string
	From        *User
	Id          string
}

type Comments struct {
	Count int64
	Data  []Comment
}

type CommentsResponse struct {
	MetaResponse
	Comments []Comment `json:"data"`
}

type Image struct {
	Url    string
	Width  int64
	Height int64
}

type Images struct {
	LowResolution      *Image `json:"low_resolution"`
	Thumbnail          *Image
	StandardResolution *Image `json:"standard_resolution"`
}

type Likes struct {
	Count int64
	Data  []User
}

type Location struct {
	Id        LocationId
	Name      string
	Latitude  float64
	Longitude float64
}

type LocationId interface{}
type LocationResponse struct {
	MetaResponse
	Location *Location `json:"data"`
}

type Media struct {
	Type         string
	Id           string
	UsersInPhoto []UserPosition `json:"users_in_photo"`
	Filter       string
	Tags         []string
	Comments     *Comments
	Caption      *Caption
	Likes        *Likes
	Link         string
	User         *User
	CreatedTime  StringUnixTime `json:"created_time"`
	Images       *Images
	Videos       *Videos
	Location     *Location
	UserHasLiked bool `json:"user_has_liked"`
	Attribution  *Attribution
}

type MediaPagination struct {
	*Pagination
}
type MediaResponse struct {
	MetaResponse
	Media *Media `json:"data"`
}
type Meta struct {
	Code         int
	ErrorType    string `json:"error_type"`
	ErrorMessage string `json:"error_message"`
}
func (m *MetaError) Error() string
type MetaResponse struct {
	Meta *Meta
}
type PaginatedMediasResponse struct {
	MediasResponse
	Pagination *MediaPagination
}
type PaginatedUsersResponse struct {
	UsersResponse
	Pagination *UserPagination
}
type Pagination struct {
	NextUrl   string `json:"next_url"`
	NextMaxId string `json:"next_max_id"`

	// Used only on GetTagRecentMedia()
	NextMaxTagId string `json:"next_max_tag_id"`
	// Used only on GetTagRecentMedia()
	MinTagId string `json:"min_tag_id"`
}

type User struct {
	Id             string `json:"id"`
	Username       string `json:"username"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	FullName       string `json:"full_name"`
	ProfilePicture string `json:"profile_picture"`
	Bio            string
	Website        string
	Counts         *UserCounts
}
