package movies

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetOptionsSeries(c *gin.Context) {
	movieGroupID := c.Param("MovieGroupID")
	seriesMovies, err := GetAndCountSeriesMovies(DB,movieGroupID)
	if err != nil {
	
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get and count series movies"})
	}

	// Print retrieved series movies
	for _, entry := range seriesMovies {
		println("Season:", entry.Season, "- Episode Total:", entry.EpisodeTotal)
	}
	c.JSON(http.StatusOK,seriesMovies)
}

func GetAndCountSeriesMovies(db *gorm.DB, movieGroupID string) ([]ResponseEpisodesOption, error) {
	var seriesMovies []Movie
	var options []ResponseEpisodesOption

	// Find series movies for the specified MovieGroupID
	query := db.Where("type = ? AND movie_group_id = ?", "SERIES", movieGroupID).Find(&seriesMovies)
	if query.Error != nil {
		return nil, query.Error
	}

	// Count episodes by season
	seasonMap := make(map[int32]int)
	for _, movie := range seriesMovies {
		seasonMap[movie.Season] += 1
	}

	// Convert to desired response format
	for season, episodeCount := range seasonMap {
		entry := ResponseEpisodesOption{
			Season:        season,
			EpisodeTotal: episodeCount,
		}
		options = append(options, entry)
	}

	// Sort options by season number
	sort.Slice(options, func(i, j int) bool {
		return options[i].Season < options[j].Season
	})

	return options, nil
}