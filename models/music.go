package models

import "time"

type Music struct {
	ID             int           `json:"id" gorm:"primary_key:auto_increment"`
	Title          string        `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailMusic string        `json:"thumbnailMusic" form:"thumbnailMusic" gorm:"type: varchar(255)"`
	Year           int           `json:"year" form:"year" gorm:"type: int"`
	ArtisID        int           `json:"artis_id" form:"artis_id"`
	Artis          ArtisResponse `json:"artis"`
	Attache        string        `json:"attache" form:"attache" gorm:"type: varchar(255)"`
	CreatedAt      time.Time     `json:"-"`
	UpdatedAt      time.Time     `json:"-"`
}

type MusicResponse struct {
	ID             int           `json:"id"`
	Title          string        `json:"title"`
	ThumbnailMusic string        `json:"thumbnailMusic"`
	Year           int           `json:"year"`
	ArtisID        int           `json:"-"`
	Artis          ArtisResponse `json:"artis"`
	Attache        string        `json:"attache"`
}

type MusicResponseEps struct {
	ID             int           `json:"id"`
	Title          string        `json:"title"`
	ThumbnailMusic string        `json:"thumbnailMusic"`
	Year           int           `json:"year"`
	ArtisID        int           `json:"-"`
	Artis          ArtisResponse `json:"artis"`
	Attache        string        `json:"attache"`
}

func (MusicResponse) TableName() string {
	return "musics"
}
func (MusicResponseEps) TableName() string {
	return "musics"
}
