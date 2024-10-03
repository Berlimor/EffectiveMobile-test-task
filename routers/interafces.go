package routers

type GenericResponse struct {
	Status string `json:"status"`
}

type ErrorResponse struct {
	GenericResponse
	Error string `json:"error"`
}

type GenericMusic struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Group string `json:"group"`
}

type Music struct {
	Title string `json:"title"`
	Group string `json:"group"`
}

type SongDetail struct {
	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type SongDetailRequest struct {
	Music

	ReleaseDate string `json:"release_date"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

type MusicWithText struct {
	GenericMusic
	Text string `json:"text"`
}

type MusicText struct {
	Title string         `json:"title"`
	Text  map[int]string `json:"text"`
}

type MusicWithTextResponse struct {
	GenericMusic
	Status int    `json:"status"`
	Text   string `json:"text"`
}

type MusicListReaponse struct {
	Status int            `json:"status"`
	Musics []GenericMusic `json:"musics"`
}