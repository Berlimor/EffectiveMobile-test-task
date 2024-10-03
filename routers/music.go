package routers

import (
	orm "EffectiveMobile/m/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Get Music List
// @Summary Get all music posted
// @Description Get all music
// @Tags Music
// @Accept  json
// @Produce  json
// @Success 200 {object} routers.MusicListReaponse
// @Failure 400 {object} routers.ErrorResponse
// @Router /music/all [get]
func GetMusicList(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	music, err := orm.GetAllMusic(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Error", 
			"error": err.Error(),
		})
		return
	}

	var musicResp []GenericMusic
	for _, val := range music {
		musicResp = append(musicResp, GenericMusic{
			ID:     val.ID,
			Title:  val.Title,
			Group:  val.Group,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok", 
		"music": musicResp,
	})
}


// Get Music List
// @Summary Get all music posted
// @Description Get all music
// @Tags Music
// @Accept  json
// @Produce  json
// @Param title query string false "title"
// @Param group query string false "group"
// @Success 200 {object} routers.MusicListReaponse
// @Failure 400 {object} routers.ErrorResponse
// @Router /music/filter [get]
func GetMusicFiltered(c *gin.Context) {
	title := c.Query("title")
	group := c.Query("group")

	db := c.MustGet("db").(*gorm.DB)

	// Get full query with all neccessary parameters
	var query strings.Builder
	if title != "" {
		// query.WriteString("title LIKE '%" + title + "%'")
		query.WriteString("title = '" + title + "' ")
	}
	if group != "" {
		if query.Len() > 0 {
			query.WriteString("AND ")
		}
		query.WriteString("\"group\" = '" + group + "' ")
	}
	music, err := orm.GetMusicFiltered(db, query.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Error", 
			"error": err.Error(),
		})
		return
	}

	var musicResp []GenericMusic
	for _, val := range music {
		musicResp = append(musicResp, GenericMusic{
			ID:     val.ID,
			Title:  val.Title,
			Group:  val.Group,
		})
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "music": musicResp})
}

// Create new music post
// @Summary Post music
// @Description Create new music post with text
// @Tags Music
// @Accept  json
// @Produce  json
// @Param music body routers.Music true "music"
// @Success 200 {object} routers.GenericResponse
// @Failure 400 {object} routers.ErrorResponse
// @Router /music [post]
func PostMusic(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var music Music
	if err := c.BindJSON(&music); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Error", 
			"error": err.Error(),
		})
		return
	}
	if err := orm.PostMusic(db, &orm.Music{
		Title:  music.Title,
		Group:  music.Group,
	}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Error", 
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// Get Music Text
// @Summary Get music text
// @Description Get the text of a specified music
// @Tags Music
// @Accept  json
// @Produce  json
// @Param title query string true "music"
// @Param group query string true "group"
// @Success 200 {object} routers.MusicText
// @Failure 400 {object} routers.ErrorResponse
// @Failure 500 {object} routers.ErrorResponse
// @Router /music/text [get]
func GetMusicText(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	title := c.Query("title")
	group := c.Query("group")
	if title == "" || group == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Error", 
			"error": "title and group are required",
		})
		return
	}
	musicId, err := orm.GetMusicIdByTitleGroup(db, title, group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error", 
			"error": err.Error(),
		})
		return
	}
	text, err := orm.GetSongText(db, musicId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Error", 
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "text": text})
}

// Get Music Details
// @Summary Get music details
// @Description Returns the details of a specified music
// @Tags Music
// @Accept  json
// @Produce  json
// @Param title query string true "music"
// @Param group query string true "group"
// @Success 200 {object} routers.SongDetail
// @Failure 400 {object} routers.ErrorResponse
// @Failure 500 {object} routers.ErrorResponse
// @Router /music/details [get]
func GetMusicDetails(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	title := c.Query("title")
	group := c.Query("group")
	if title == "" || group == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Error", 
			"error": "title and group are required",
		})
		return
	}
	musicId, err := orm.GetMusicIdByTitleGroup(db, title, group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error", 
			"error": err.Error(),
		})
		return
	}

	details, err := orm.GetSongDetails(db, musicId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Error", 
			"error": err.Error(),
		})
		return
	}
	detailsResp := SongDetail{
		ReleaseDate: details.ReleaseDate,
		Text:       details.Text,
		Link:       details.Link,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok", 
		"details": detailsResp,
	})
}

// Update song details
// @Summary Update song details
// @Description Update or create song details
// @Tags Music
// @Accept  json
// @Produce  json
// @Param details body routers.SongDetailRequest true "details"
// @Success 200 {object} routers.GenericResponse
// @Failure 400 {object} routers.ErrorResponse
// @Router /music/details [post]
func UpdateSongDetails(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var details SongDetailRequest
	if err := c.BindJSON(&details); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Error", 
			"error": err.Error(),
		})
		return
	}
	musicId, err := orm.GetMusicIdByTitleGroup(db, details.Title, details.Group)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Error", 
			"error": err.Error(),
		})
		return
	}

	if err := orm.UpdateSongDetails(db, musicId, details.ReleaseDate, details.Text, details.Link); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error", 
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// Delete music
// @Summary Delete music
// @Description Delete music and it's details
// @Tags Music
// @Accept  json
// @Produce  json
// @Param title query string true "music"
// @Param group query string true "group"
// @Success 200 {object} routers.GenericResponse
// @Failure 400 {object} routers.ErrorResponse
// @Router /music [delete]
func DeleteMusic(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	title := c.Query("title")
	group := c.Query("group")
	if title == "" || group == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "Error", 
			"error": "title and group are required",
		})
		return
	}
	musicId, err := orm.GetMusicIdByTitleGroup(db, title, group)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error", 
			"error": err.Error(),
		})
		return
	}
	if err := orm.DeleteMusic(db, musicId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error", 
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}