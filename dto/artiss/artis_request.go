package artis

type CreateArtisRequest struct {
	Name         string `json:"name" form:"name" validate:"required"`
	Old          int    `json:"old" form:"old"  validate:"required" gorm:"type: int"`
	Type_Artis   string `json:"type_artis" form:"type_artis"  validate:"required"`
	Start_Career int    `json:"start_career" form:"start_career" validate:"required" gorm:"type: int"`
}

type UpdateArtisRequest struct {
	Name         string `json:"name"`
	Old          int    `json:"old"`
	Type_Artis   string `json:"type_artis"`
	Start_Career int    `json:"start_career"`
}
