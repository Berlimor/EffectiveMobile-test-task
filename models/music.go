package models

import (
	"gorm.io/gorm"
)

type Music struct {
	gorm.Model
	Title  string `gorm:"not null"`
	Group  string `gorm:"not null"`
}

// GetAllMusic returns all music in slice
func GetAllMusic(db *gorm.DB) ([]Music, error) {
	var music []Music
	err := db.Find(&music).Error
	return music, err
}

//
func PostMusic(db *gorm.DB, music *Music) error {
	err := db.Create(music).Error
	return err
}

// GetMusicFiltered returns only filtered music in slice
func GetMusicFiltered(db *gorm.DB, filter string) ([]Music, error) {
	if filter == "" {
		return GetAllMusic(db)
	}
	var musics []Music
	err := db.Where(filter).Find(&musics).Error
	return musics, err
}

// Returns music id by title and group
func GetMusicIdByTitleGroup(db *gorm.DB, title string, group string) (uint, error) {
	var music Music
	err := db.Where("title = ? AND \"group\" = ?", title, group).First(&music).Error
	return music.ID, err
}

func DeleteMusic(db *gorm.DB, musicId uint) error {
	var music Music
	return db.Where("id = ?", musicId).Delete(&music).Error
}