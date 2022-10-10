package artis

type ArtisResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Old          int    `json:"old"`
	Type_Artis   string `json:"type_artis"`
	Start_Career int    `json:"start_career"`
}
