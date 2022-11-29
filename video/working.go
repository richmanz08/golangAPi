package video

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func indexPage(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "index.html")
// }
// func StreamHandle(c  *gin.Context){
// 	mediaBase := "1"
// 	m3u8Name := "testm3u8file.m3u8"
// 	serveHlsM3u8(c, mediaBase, m3u8Name)
// }
// func getMediaBase(mId int) string {
// 	mediaRoot := "public/"
// 	return fmt.Sprintf("%s/%d", mediaRoot, mId)
// }
// func serveHlsM3u8(ces *gin.Context, mediaBase, m3u8Name string) {

// 	mediaFile := fmt.Sprintf("%s/hls/%s", mediaBase, m3u8Name)
// 	ces.File(mediaFile)
// 	ces.Writer.Header().Set("Content-Type", "application/x-mpegURL")
// 	// http.ServeFile( ces, mediaFile)
// 	// ces.Header().Set("Content-Type", "application/x-mpegURL")
// }
// func serveHlsTs(w http.ResponseWriter, r *http.Request, mediaBase, segName string) {
// 	mediaFile := fmt.Sprintf("%s/hls/%s", mediaBase, segName)
// 	http.ServeFile(w, r, mediaFile)
// 	w.Header().Set("Content-Type", "video/MP2T")
// }
// // https://www.rohitmundra.com/video-streaming-server