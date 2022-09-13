package episodedto

import "dumbflix_be/models"

type EpisodeResponse struct {
	ID						int						`json:"id" gorm:"primary_key:auto_increment"`
	Title					string					`json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm			string					`json:"thumbnailEpisode" form:"thumbnailEpisode" gorm:"type: varchar(255)"`
	LinkFilm				string					`json:"linkFilm" form:"linkFilm" gorm:"type: varchar(255)"`
	FilmID 					int                		`json:"film_id" form:"film_id"`
	Film   					models.FilmResponse    	`json:"film" gorm:"foreignKey:FilmID"`
}

type EpisodeUpdateResponse struct {
	ID						int						`json:"id" gorm:"primary_key:auto_increment"`
	Title					string					`json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm			string					`json:"thumbnailEpisode" form:"thumbnailEpisode" gorm:"type: varchar(255)"`
	LinkFilm				string					`json:"linkFilm" form:"linkFilm" gorm:"type: varchar(255)"`
	FilmID 					int                		`json:"film_id" form:"film_id"`
	Film   					models.FilmResponse    	`json:"film" `
}

type EpisodeDeleteResponse struct {
	ID					int							`json:"id"`
}