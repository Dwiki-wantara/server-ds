package models

type Artis struct {
	ID           int    `json:"id" gorm:"primary_key:auto_increment"`
	Name         string `json:"name" form:"name" gorm:"type: varchar(255)"`
	Old          int    `json:"old" form:"old" gorm:"type: int"`
	Type_Artis   string `json:"type_artis" form:"type_artis" gorm:"type: varchar(255)"`
	Start_Career int    `json:"start_career" form:"start_career" gorm:"type: int"`
}

type ArtisResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Old          int    `json:"old"`
	Type_Artis   string `json:"type_artis"`
	Start_Career int    `json:"start_career"`
}

func (ArtisResponse) TableName() string {
	return "artis"
}
