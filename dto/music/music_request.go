package musicdto

type MusicRequest struct {
	Title          string `json:"title" form:"title" gorm:"type: varchar(255)" validate:"required"`
	ThumbnailMusic string `json:"thumbnailMusic" form:"thumbnailMusic" gorm:"type: varchar(255)" `
	Year           int    `json:"year" form:"year" gorm:"type: int" validate:"required"`
	ArtisID        int    `json:"artis_id" form:"artis_id" gorm:"type: int"`
	Attache        string `json:"attache" form:"attache" gorm:"type: varchar(255)"`
}

type UpdateMusicRequest struct {
	Title          string `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailMusic string `json:"thumbnailMusic" form:"thumbnailMusic" gorm:"type: varchar(255)" `
	Year           int    `json:"year" form:"year" gorm:"type: int"`
	ArtisID        int    `json:"artis_id" form:"artis_id" gorm:"type: int"`
	Attache        string `json:"attache" form:"attache" gorm:"type: varchar(255)"`
}
