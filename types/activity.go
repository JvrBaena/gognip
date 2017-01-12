package types

/*
Activity ...
*/
type Activity struct {
	ID                      string                          `json:"id,omitempty"`
	Actor                   activityActor                   `json:"actor,omitempty"`
	Verb                    string                          `json:"verb,omitempty"`
	Generator               activityGenerator               `json:"generator,omitempty"`
	Provider                activityProvider                `json:"provider,omitempty"`
	InReplyTo               activityInReplyTo               `json:"inReplyTo,omitempty"`
	Location                activityLocation                `json:"location,omitempty"`
	TwitterEntities         activityTwitterEntities         `json:"twitter_entities,omitempty"`
	TwitterExtendedEntities activityTwitterExtendedEntities `json:"twitter_extended_entities,omitempty"`
	Link                    string                          `json:"link,omitempty"`
	Body                    string                          `json:"body,omitempty"`
	ObjectType              string                          `json:"objectType,omitempty"`
	Object                  activityObject                  `json:"object,omitempty"`
	PostedTime              string                          `json:"postedTime,omitempty"`
	TimestampMs             string                          `json:"timestampMs,omitempty"`
	TwitterLang             string                          `json:"twitter_lang,omitempty"`
	Gnip                    activityGnip                    `json:"gnip,omitempty"`
}

type activityActor struct {
	ObjectType        string        `json:"objectType,omitempty"`
	ID                string        `json:"id,omitempty"`
	Link              string        `json:"link,omitempty"`
	DisplayName       string        `json:"displayName,omitempty"`
	PostedTime        string        `json:"postedTime,omitempty"`
	Image             string        `json:"image,omitempty"`
	Summary           string        `json:"summary,omitempty"`
	Links             []actorLinks  `json:"links,omitempty"`
	FriendsCount      int           `json:"friendsCount,omitempty"`
	FollowersCount    int           `json:"followersCount,omitempty"`
	ListedCount       int           `json:"listedCount,omitempty"`
	StatusesCount     int           `json:"statusesCount,omitempty"`
	FavoritesCount    int           `json:"favoritesCount,omitempty"`
	TwitterTimeZone   string        `json:"twitterTimeZone,omitempty"`
	Verified          bool          `json:"verified,omitempty"`
	UTCOffset         string        `json:"utcOffset,omitempty"`
	PreferredUsername string        `json:"preferredUsername,omitempty"`
	Languages         []string      `json:"languages,omitempty"`
	Location          actorLocation `json:"location,omitempty"`
}

type actorLinks struct {
	Href string `json:"href,omitempty"`
	Rel  string `json:"rel,omitempty"`
}

type actorLocation struct {
	ObjectType  string `json:"objectType,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

type activityGenerator struct {
	DisplayName string `json:"displayName,omitempty"`
	Link        string `json:"link,omitempty"`
}

type activityProvider struct {
	ObjectType  string `json:"objectType,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Link        string `json:"link,omitempty"`
}

type activityInReplyTo struct {
	Link string `json:"link,omitempty"`
}

type activityLocation struct {
	ObjectType         string      `json:"objectType,omitempty"`
	DisplayName        string      `json:"displayName,omitempty"`
	Name               string      `json:"name,omitempty"`
	CountryCode        string      `json:"country_code,omitempty"`
	TwitterCountryCode string      `json:"twitter_country_code,omitempty"`
	Link               string      `json:"link,omitempty"`
	Geo                locationGeo `json:"geo,omitempty"`
	TwitterPlaceType   string      `json:"twitter_place_type,omitempty"`
	StreetAddress      string      `json:"streetAddress,omitempty"`
}

type locationGeo struct {
	Type        string `json:"type,omitempty"`
	Coordinates [][][]geoCoordinate
}

type geoCoordinate float32

type simpleGeo struct {
	Type        string `json:"type,omitempty"`
	Coordinates []geoCoordinate
}

type activityTwitterEntities struct {
	Hashtags     []entityText    `json:"hashtags,omitempty"`
	Symbols      []entityText    `json:"symbols,omitempty"`
	Urls         []entityURL     `json:"urls,omitempty"`
	UserMentions []entityMention `json:"user_mentions,omitempty"`
	Media        []entityMedia   `json:"media,omitempty"`
}

type entityText struct {
	Text    string `json:"text,omitempty"`
	Indices []int  `json:"indices,omitempty"`
}

type entityURL struct {
	URL         string `json:"url,omitempty"`
	ExpandedURL string `json:"expanded_url,omitempty"`
	DisplayURL  string `json:"display_url,omitempty"`
	Indices     []int  `json:"indices,omitempty"`
}

type entityMention struct {
	ScreenName string `json:"screen_name,omitempty"`
	Name       string `json:"name,omitempty"`
	ID         int64  `json:"id,omitempty"`
	IDStr      string `json:"id_str,omitempty"`
	Indices    []int  `json:"indices,omitempty"`
}

type entityMedia struct {
	ID            int64     `json:"id,omitempty"`
	IDStr         string    `json:"id_str,omitempty"`
	Indices       []int     `json:"indices,omitempty"`
	MediaURL      string    `json:"media_url,omitempty"`
	MediaURLHttps string    `json:"media_url_https,omitempty"`
	URL           string    `json:"url,omitempty"`
	DisplayURL    string    `json:"display_url,omitempty"`
	ExpandedURL   string    `json:"expanded_url,omitempty"`
	Type          string    `json:"photo,omitempty"`
	Sizes         mediaSize `json:"sizes,omitempty"`
}

type mediaSize struct {
	Small  size `json:"small,omitempty"`
	Medium size `json:"medium,omitempty"`
	Large  size `json:"large,omitempty"`
	Thumb  size `json:"thumb,omitempty"`
}

type size struct {
	W      int    `json:"w,omitempty"`
	H      int    `json:"h,omitempty"`
	Resize string `json:"resize,omitempty"`
}

type activityTwitterExtendedEntities struct {
	Media []entityMedia `json:"media,omitempty"`
}

type activityObject struct {
	ObjectType string `json:"objectType,omitempty"`
	ID         string `json:"id,omitempty"`
	Summary    string `json:"summary,omitempty"`
	Link       string `json:"link,omitempty"`
	PostedTime string `json:"postedTime,omitempty"`
}

type activityGnip struct {
	MatchingRules    []gnipRule            `json:"matching_rules,omitempty"`
	Urls             []gnipURL             `json:"urls,omitempty"`
	ProfileLocations []gnipProfileLocation `json:"profileLocations,omitempty"`
}

type gnipRule struct {
	Tag string `json:"tag,omitempty"`
	ID  int64  `json:"id,omitempty"`
}

type gnipURL struct {
	URL                    string `json:"url,omitempty"`
	ExpandedURL            string `json:"expanded_url,omitempty"`
	ExpandedStatus         string `json:"expanded_status,omitempty"`
	ExpandedURLTitle       string `json:"expanded_url_title,omitempty"`
	ExpandedURLDescription string `json:"expanded_url_description,omitempty"`
}

type gnipProfileLocation struct {
	Address     profileLocationAddress `json:"address,omitempty"`
	ObjectType  string                 `json:"objectType,omitempty"`
	DisplayName string                 `json:"displayName,omitempty"`
	Geo         simpleGeo              `json:"geo,omitempty"`
}

type profileLocationAddress struct {
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	Locality    string `json:"locality,omitempty"`
	Region      string `json:"region,omitempty"`
	SubRegion   string `json:"subRegion,omitempty"`
}
