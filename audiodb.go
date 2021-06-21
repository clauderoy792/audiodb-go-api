package audiodb

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

//API doc at https://www.theaudiodb.com/api_guide.php
type AudioDB struct {
	client Client
}

func NewAudioDB() *AudioDB {
	return &AudioDB{
		client: NewClient(),
	}
}

func (audioDB *AudioDB) SearchTitle(artist, title string) (TrackCollection, error) {
	var tracks TrackCollection
	v := url.Values{}
	v.Set("s",artist)
	v.Set("t",title)
	resp, err := audioDB.client.Get("searchtrack",v.Encode())
	if err != nil {
		return tracks, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, &tracks)
	}

	return tracks, err
}

func (audioDB *AudioDB) SearchArtist(artist string) (ArtistCollection, error) {
	var artistCollection ArtistCollection
	v := url.Values{}
	v.Set("s",artist)
	resp, err := audioDB.client.Get("search",v.Encode())
	if err != nil {
		return artistCollection, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	
	if err == nil {
		err = json.Unmarshal(body, &artistCollection)
	}

	return artistCollection, err
}

type ArtistCollection struct {
	Artists []struct {
		ID        		int 	`json:"idArtist,string"`
		Name          	string 	`json:"strArtist"`
		Stripped  		string 	`json:"strArtistStripped"`
		Alternate 		string 	`json:"strArtistAlternate"`
		Label           string 	`json:"strLabel"`
		LabelID         string 	`json:"idLabel"`
		FormedYear      string 	`json:"intFormedYear"`
		BornYear        int 	`json:"intBornYear,string"`
		DiedYear        int    	`json:"intDiedYear,string"`
		Disbanded       string 	`json:"strDisbanded"`
		Style           string 	`json:"strStyle"`
		Genre           string 	`json:"strGenre"`
		Mood            string 	`json:"strMood"`
		Website         string 	`json:"strWebsite"`
		Facebook        string 	`json:"strFacebook"`
		Twitter         string 	`json:"strTwitter"`
		BiographyEN     string 	`json:"strBiographyEN"`
		BiographyDE     string 	`json:"strBiographyDE"`
		BiographyFR     string 	`json:"strBiographyFR"`
		BiographyCN     string 	`json:"strBiographyCN"`
		BiographyIT     string 	`json:"strBiographyIT"`
		BiographyJP     string 	`json:"strBiographyJP"`
		BiographyRU     string 	`json:"strBiographyRU"`
		BiographyES     string 	`json:"strBiographyES"`
		BiographyPT     string 	`json:"strBiographyPT"`
		BiographySE     string 	`json:"strBiographySE"`
		BiographyNL     string 	`json:"strBiographyNL"`
		BiographyHU     string 	`json:"strBiographyHU"`
		BiographyNO     string 	`json:"strBiographyNO"`
		BiographyIL     string 	`json:"strBiographyIL"`
		BiographyPL     string 	`json:"strBiographyPL"`
		Gender          string 	`json:"strGender"`
		Members         int 	`json:"intMembers,string"`
		Country         string `json:"strCountry"`
		CountryCode     string `json:"strCountryCode"`
		ArtistThumb     string `json:"strArtistThumb"`
		ArtistLogo      string `json:"strArtistLogo"`
		ArtistClearart  string `json:"strArtistClearart"`
		ArtistWideThumb string `json:"strArtistWideThumb"`
		ArtistFanart    string `json:"strArtistFanart"`
		ArtistFanart2   string `json:"strArtistFanart2"`
		ArtistFanart3   string `json:"strArtistFanart3"`
		ArtistBanner    string `json:"strArtistBanner"`
		MusicBrainzID   string `json:"strMusicBrainzID"`
		LastFMChart     string `json:"strLastFMChart"`
		Charted         string `json:"intCharted"`
		Locked          string `json:"strLocked,omitempty"`
	} `json:"artists"`
}

type AlbumCollection struct {
	Albums []struct {
		ID             		int 	`json:"idAlbum,string"`
		ArtistID            int 	`json:"idArtist,string"`
		LabelID             int 	`json:"idLabel,string"`
		Name               	string 	`json:"strAlbum"`
		AlbumStripped       string 	`json:"strAlbumStripped"`
		Artist              string 	`json:"strArtist"`
		Artistipped         string 	`json:"strArtistStripped"`
		YearReleased        int    	`json:"intYearReleased"`
		Style               string 	`json:"strStyle"`
		Genre               string 	`json:"strGenre"`
		Label               string 	`json:"strLabel"`
		ReleaseFormat       string 	`json:"strReleaseFormat"`
		Sales               int    	`json:"intSales"`
		Thumb          		string 	`json:"strAlbumThumb"`
		ThumbHQ        		string 	`json:"strAlbumThumbHQ"`
		ThumbBack      		string 	`json:"strAlbumThumbBack"`
		CDart          		string 	`json:"strAlbumCDart"`
		Spine          		string 	`json:"strAlbumSpine"`
		A3DCase         	string 	`json:"strAlbum3DCase"`
		A3DFlat         	string 	`json:"strAlbum3DFlat"`
		A3DFace         	string 	`json:"strAlbum3DFace"`
		A3DThumb        	string 	`json:"strAlbum3DThumb"`
		DescriptionEN       string 	`json:"strDescriptionEN,omitempty"`
		DescriptionDE       string 	`json:"strDescriptionDE"`
		DescriptionFR       string 	`json:"strDescriptionFR"`
		DescriptionCN       string 	`json:"strDescriptionCN"`
		DescriptionIT       string 	`json:"strDescriptionIT"`
		DescriptionJP       string 	`json:"strDescriptionJP"`
		DescriptionRU       string 	`json:"strDescriptionRU"`
		DescriptionES       string 	`json:"strDescriptionES"`
		DescriptionPT       string 	`json:"strDescriptionPT"`
		DescriptionSE       string 	`json:"strDescriptionSE"`
		DescriptionNL       string 	`json:"strDescriptionNL"`
		DescriptionHU       string 	`json:"strDescriptionHU"`
		DescriptionNO       string 	`json:"strDescriptionNO"`
		DescriptionIL       string 	`json:"strDescriptionIL"`
		DescriptionPL       string 	`json:"strDescriptionPL"`
		Loved               int    	`json:"intLoved,string"`
		Score               float64 `json:"intScore,string"`
		ScoreVotes          string 	`json:"intScoreVotes"`
		Review              string 	`json:"strReview"`
		Mood                string 	`json:"strMood"`
		Theme               string 	`json:"strTheme"`
		Speed               string 	`json:"strSpeed"`
		Location            string 	`json:"strLocation"`
		MusicBrainzID       string 	`json:"strMusicBrainzID"`
		MusicBrainzArtistID string 	`json:"strMusicBrainzArtistID"`
		AllMusicID          string 	`json:"strAllMusicID"`
		BBCReviewID         string 	`json:"strBBCReviewID"`
		RateYourMusicID     string 	`json:"strRateYourMusicID"`
		DiscogsID           string 	`json:"strDiscogsID"`
		WikidataID          string 	`json:"strWikidataID"`
		WikipediaID         string 	`json:"strWikipediaID"`
		GeniusID            string 	`json:"strGeniusID"`
		LyricWikiID         string 	`json:"strLyricWikiID"`
		MusicMozID          string 	`json:"strMusicMozID"`
		ItunesID            string 	`json:"strItunesID"`
		AmazonID            string 	`json:"strAmazonID"`
		Locked              string 	`json:"strLocked"`
		Description         string 	`json:"strDescription,omitempty"`
	} `json:"album"`
}

type TrackCollection struct {
	Tracks []struct {
		ID             		int 	`json:"idTrack,string,omitempty"`
		AlbumID             int 	`json:"idAlbum,string,omitempty"`
		ArtistID            int 	`json:"idArtist,string,omitempty"`
		LyricID             int 	`json:"idLyric,string,omitempty"`
		IMVDBID             int 	`json:"idIMVDB,string,omitempty"`
		Title               string  `json:"strTrack"`
		Album               string  `json:"strAlbum"`
		Artist              string  `json:"strArtist"`
		ArtistAlternate     string  `json:"strArtistAlternate"`
		CD                  int     `json:"intCD,string,omitempty"`
		Duration            int     `json:"intDuration,string"`
		Genre               string  `json:"strGenre"`
		Mood                string  `json:"strMood"`
		Style               string  `json:"strStyle"`
		Theme               string  `json:"strTheme"`
		DescriptionEN       string  `json:"strDescriptionEN"`
		TrackThumb          string  `json:"strTrackThumb"`
		Track3DCase         string  `json:"strTrack3DCase"`
		TrackLyrics         string  `json:"strTrackLyrics"`
		MusicVid            string  `json:"strMusicVid"`
		MusicVidDirector    string  `json:"strMusicVidDirector"`
		MusicVidCompany     string  `json:"strMusicVidCompany"`
		MusicVidScreen1     string  `json:"strMusicVidScreen1"`
		MusicVidScreen2     string  `json:"strMusicVidScreen2"`
		MusicVidScreen3     string  `json:"strMusicVidScreen3"`
		MusicVidViews       int     `json:"intMusicVidViews,string,omitempty"`
		MusicVidLikes       int     `json:"intMusicVidLikes,string,omitempty"`
		MusicVidDislikes    int     `json:"intMusicVidDislikes,string,omitempty"`
		MusicVidFavorites   int     `json:"intMusicVidFavorites,string,omitempty"`
		MusicVidComments    int     `json:"intMusicVidComments,string,omitempty"`
		TrackNumber         int     `json:"intTrackNumber,string,omitempty"`
		Loved               int     `json:"intLoved,string,omitempty"`
		Score               float64 `json:"intScore,string,omitempty"`
		ScoreVotes          int     `json:"intScoreVotes,string,omitempty"`
		TotalListeners      int     `json:"intTotalListeners,string,omitempty"`
		TotalPlays          int     `json:"intTotalPlays,string,omitempty"`
		MusicBrainzID       string  `json:"strMusicBrainzID"`
		MusicBrainzAlbumID  string  `json:"strMusicBrainzAlbumID"`
		MusicBrainzArtistID string  `json:"strMusicBrainzArtistID"`
		Locked              string  `json:"strLocked,omitempty"`
	} `json:"track"`
}

type Discography struct {
	Album []struct {
		Album        string `json:"strAlbum"`
		YearReleased int    `json:"intYearReleased,string"`
	} `json:"album"`
}

type MusicVideoCollection struct {
	MusicVideos []struct {
		ArtistID      int		`json:"idArtist,string"`
		AlbumID       int 		`json:"idAlbum,string"`
		TrackID       int 		`json:"idTrack,string"`
		Track         string 	`json:"strTrack"`
		TrackThumb    string 	`json:"strTrackThumb"`
		MusicVid      string 	`json:"strMusicVid"`
		DescriptionEN string 	`json:"strDescriptionEN"`
	} `json:"mvids"`
}

const baseAddress = "https://www.theaudiodb.com/api/v1/json/1/"

// Client used to work with TheAudioDB API.
type Client struct {
	http    *http.Client
	baseURL string

	AutoRetry      bool
	AcceptLanguage string
}

func (client *Client) Get(action, query string) (*http.Response,error) {
	url := client.baseURL + action + ".php?" + query
	return http.Get(url)
	
}

// NewClient returns a client for working with the Spotify Web API.
// The provided HTTP client must include the user's access token in each request;
// if you do not have such a client, use the `Authenticator.NewClient` method instead.
func NewClient() Client {
	return Client{
		http:    &http.Client{},
		baseURL: baseAddress,
	}
}
