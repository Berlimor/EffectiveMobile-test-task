package models

import (
	"bufio"
	"log"
	"strings"

	"gorm.io/gorm"
)

type SongDetail struct {
	gorm.Model
	MusicID 	uint	`gorm:"unique; not null"`
	Music 		Music 	`gorm:"foreignKey:MusicID;constraint:OnDelete:CASCADE"`

	ReleaseDate string	
	Text 		string
	Link 		string
}

// GetSongDetails returns the details of the specified song
func GetSongDetails(db *gorm.DB, musicId uint) (*SongDetail, error) {
	var songDetail SongDetail
	if err := db.Where("music_id = ?", musicId).First(&songDetail).Error; err != nil {
		return nil, err
	}
	return &songDetail, nil
}

// UpdateSongDetails updates the specified song details
func UpdateSongDetails(db *gorm.DB, musicId uint, date string, text string, link string) error {
	var origDetail SongDetail
	if err := db.Where("music_id = ?", musicId).First(&origDetail).Error; err != nil {
		log.Println("SongDetail not found, creating new one")
		return db.Create(&SongDetail{
			MusicID: musicId, 
			ReleaseDate: date, 
			Text: text, 
			Link: link,
		}).Error
	}
	origDetail.ReleaseDate = date
	origDetail.Text = text
	origDetail.Link = link
	return db.Save(&origDetail).Error
}

// GetSongText returns the text of the specified song
func GetSongText(db *gorm.DB, musicId uint) (map[int]string, error) {
	var songDetail SongDetail
	if err := db.Where("music_id = ?", musicId).First(&songDetail).Error; err != nil {
		return nil, err
	}

	// Iterate over text to find new lines and return the map with readable text
	result := make(map[int]string)
	i := 1
	scanner := bufio.NewScanner(strings.NewReader(songDetail.Text))
	for scanner.Scan() {
		result[i] = scanner.Text()
		i++
	}
	return result, nil
}