package movies

import (
	"api-webapp/common"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"os/exec"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func createFolder(movie Movie) bool {
	var bucketFolder = "D:\\streamingfile\\"

	// Replace "YourDesiredFolderName" with the name of the folder you want to create
	folderName := movie.DirectoryName
	// Replace "D:\\" with the desired path on the D drive where you want to create the folder
	folderPath := bucketFolder + folderName

	// Create the command to execute
	cmd := exec.Command("cmd", "/c", "mkdir", folderPath)

	errCommand := cmd.Run()
	if errCommand != nil {
		fmt.Println("Error creating folder:", errCommand)
		// os.Exit(1)
		return false
		
	}

	// thumbnail sub folder
	thumbnailPathFolder := bucketFolder + folderName + "\\thumbnail"
	// Create the command to execute
	cmdthumbnail := exec.Command("cmd", "/c", "mkdir", thumbnailPathFolder)
	cmdthumbnail.Run()

	// subtitle sub folder
	subtitlePathFolder := bucketFolder + folderName + "\\subtitle"
	// Create the command to execute
	cmdsubtitle := exec.Command("cmd", "/c", "mkdir", subtitlePathFolder)
	cmdsubtitle.Run()

	// audio sub folder
	audioPathFolder := bucketFolder + folderName + "\\audio"
	// Create the command to execute
	cmdaudio := exec.Command("cmd", "/c", "mkdir", audioPathFolder)
	cmdaudio.Run()

	// video sub folder
	videoPathFolder := bucketFolder + folderName + "\\video"
	// Create the command to execute
	cmdvideo := exec.Command("cmd", "/c", "mkdir", videoPathFolder)
	cmdvideo.Run()
	// os.Exit(1)
	return true
}
func uploadImage(file multipart.File, header *multipart.FileHeader) (string, error) {
	filename := header.Filename
	bucket := os.Getenv("BUCKET_FILE_URL")
	result := fmt.Sprintf("%spublic", bucket)
	filepath := path.Join(result, filename)

	out, err := os.Create(filepath)
	if err != nil {
		log.Printf("Error creating file: %s", err)
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Printf("Error copying file: %s", err)
		return "", err
	}
	return filename, nil
}

// this help funtion for refactor code cleanup
func parseQueryParams(c *gin.Context) (ParamsMovies, error) {
	var params ParamsMovies

	movieGroupIdStr := c.DefaultQuery("MovieGroupID", "0")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	currentStr := c.DefaultQuery("current", "1")
	seasonStr := c.DefaultQuery("season", "0")

	movieGroupID, err := strconv.Atoi(movieGroupIdStr)
	if err != nil {
		return params, err
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		return params, err
	}

	current, err := strconv.Atoi(currentStr)
	if err != nil {
		return params, err
	}

	season, err := strconv.Atoi(seasonStr)
	if err != nil {
		return params, err
	}

	params.MovieGroupID = movieGroupID
	params.PageSize = pageSize
	params.Current = current
	params.Season = season

	return params, nil
}
func createDynamicQuery(params ParamsMovies) *gorm.DB {
	dynamicQuery := DB.Model(&Movie{})

	if params.MovieGroupID != 0 {
		dynamicQuery = dynamicQuery.Where("movie_group_id = ?", params.MovieGroupID)
	}

	if params.Season != 0 {
		dynamicQuery = dynamicQuery.Where("season = ?", params.Season)
	}

	return dynamicQuery
}
func countTotalRecords(dynamicQuery *gorm.DB) (int64, error) {
	var totalCount int64
	countTotal := dynamicQuery.Count(&totalCount)
	if countTotal.Error != nil {
		return 0, countTotal.Error
	}
	return totalCount, nil
}
func handlePaginationAndQuery(dynamicQuery *gorm.DB, params ParamsMovies, movies *[]Movie) error {
	paginatedDB, err := common.Paginate(dynamicQuery, params.Current, params.PageSize)
	if err != nil {
		return err
	}

	queryFindAll := paginatedDB.Find(movies)
	if queryFindAll.Error != nil {
		return queryFindAll.Error
	}

	if len(*movies) == 0 && params.Current > 1 {
		params.Current = 1
		dynamicQuery = createDynamicQuery(params)
		paginatedDB, err = common.Paginate(dynamicQuery, params.Current, params.PageSize)
		if err != nil {
			return err
		}

		queryFindAll = paginatedDB.Find(movies)
		if queryFindAll.Error != nil {
			return queryFindAll.Error
		}
	}

	return nil
}

// this help GetAllMovieGroup
func readQueryStringMovieGroup(c *gin.Context) (ParamsMovieGroup, error ){
	var params ParamsMovieGroup
	pageSizeStr := c.Query("pageSize")
    currentStr := c.Query("current")
	if pageSizeStr != "" {
        ps, err := strconv.Atoi(pageSizeStr)
        if err == nil {
            params.PageSize = ps
        } else {
           
            return params,err
        }
    }

	if currentStr != "" {
        ps, err := strconv.Atoi(currentStr)
        if err == nil {
            params.Current = &ps
        } else {
           
			return params,err
        }
    }

	if name := c.Query("Name"); name != "" {
		params.Name = name
	}
	if status := c.Query("Status"); status != "" {
		params.Status = status
	}
	if includesIds := c.Query("IncludesMovieGroupIds"); includesIds != "" {
		params.InCludesMovieGroupIds = includesIds
	}
	if excludesIds := c.Query("ExcludesMovieGroupIds"); excludesIds != "" {
		params.ExCludesMovieGroupIds = excludesIds
	}

	return params, nil

}
func createDynamicQueryMovieGroup(params ParamsMovieGroup) (*gorm.DB, error) {
	dynamicQuery := DB.Model(&MovieGroup{})

	if params.Name != "" {
		dynamicQuery = dynamicQuery.Where("name_eng LIKE ? OR name_local LIKE ?", "%"+params.Name+"%", "%"+params.Name+"%")
	}

	if params.Status != "" {
		dynamicQuery = dynamicQuery.Where("status IN (?)", params.Status)
	}

	if params.InCludesMovieGroupIds != "" {
        includesIds, err := common.ConvertToIDSlice(params.InCludesMovieGroupIds)
        if err != nil {
			return nil, err
        }
        dynamicQuery = dynamicQuery.Where("id IN ?", includesIds)
    }

    if params.ExCludesMovieGroupIds != "" {
        excludesIds, err := common.ConvertToIDSlice(params.ExCludesMovieGroupIds)
        if err != nil {
			return nil, err
        }
        dynamicQuery = dynamicQuery.Where("id NOT IN ?", excludesIds)
    }
	return dynamicQuery,nil
}
 func handlePaginationAndQueryMovieGroup(dynamicQuery *gorm.DB, params *ParamsMovieGroup, movies *[]MovieGroup) error {

	queryRows, err := common.Paginate(dynamicQuery, *params.Current, params.PageSize)
	if err != nil {
		return err
	} 

	queryFindAll := queryRows.Find(&movies)
	if queryFindAll.Error != nil {
		return queryFindAll.Error
	}

	
	if len(*movies) == 0 && *params.Current > 1 {
		*params.Current = 1

		// Create a fresh query for the fallback
		fallbackQuery, err := createDynamicQueryMovieGroup(*params)
		if err != nil {
			return err
		}
		fallbackDB, err := common.Paginate(fallbackQuery, *params.Current, params.PageSize)

		if err != nil {
			return err
		}

		// Empty the movies slice before performing the fallback query
		*movies = nil

		queryFindAllFallback := fallbackDB.Find(&movies)

		if queryFindAllFallback.Error != nil {
			return queryFindAllFallback.Error
		}
	}

	return nil
}