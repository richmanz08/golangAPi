package video

import (
	// "bufio"
	// "bytes"

	"fmt"
	"net/http"

	// "net/http"

	// "log"

	"strings"

	// "time"

	// "io"
	// "log"
	// "net/http"
	// "os"

	"github.com/gin-gonic/gin"
)

// type WV struct {
// 	AudioChannels          uint
// 	AudioFormat            uint
// 	AudioProfileIDC        uint
// 	AudioSampleSize        uint
// 	AudioSamplingFrequency uint
// 	CypherVersion          string
// 	ECM                    string
// 	VideoFormat            uint
// 	VideoFrameRate         uint
// 	VideoLevelIDC          uint
// 	VideoProfileIDC        uint
// 	VideoResolution        string
// 	VideoSAR               string
// }
// type Key struct {
// 	Method            string
// 	URI               string
// 	IV                string
// 	Keyformat         string
// 	Keyformatversions string
// }
// type Alternative struct {
// 	GroupId         string
// 	URI             string
// 	Type            string
// 	Language        string
// 	Name            string
// 	Default         bool
// 	Autoselect      string
// 	Forced          string
// 	Characteristics string
// 	Subtitles       string
// }
// type MediaSegment struct {
// 	SeqId           uint64
// 	Title           string // optional second parameter for EXTINF tag
// 	URI             string
// 	Duration        float64   // first parameter for EXTINF tag; duration must be integers if protocol version is less than 3 but we are always keep them float
// 	Limit           int64     // EXT-X-BYTERANGE <n> is length in bytes for the file under URI
// 	Offset          int64     // EXT-X-BYTERANGE [@o] is offset from the start of the file under URI
// 	Key             *Key      // displayed before the segment and means changing of encryption key (in theory each segment may have own key)
// 	Discontinuity   bool      // EXT-X-DISCONTINUITY indicates an encoding discontinuity between the media segment that follows it and the one that preceded it (i.e. file format, number and type of tracks, encoding parameters, encoding sequence, timestamp sequence)
// 	ProgramDateTime time.Time // EXT-X-PROGRAM-DATE-TIME tag associates the first sample of a media segment with an absolute date and/or time
// }
// type MediaType uint
// type VariantParams struct {
// 	ProgramId    uint32
// 	Bandwidth    uint32
// 	Codecs       string
// 	Resolution   string
// 	Audio        string
// 	Video        string
// 	Subtitles    string
// 	Iframe       bool // EXT-X-I-FRAME-STREAM-INF
// 	Alternatives []*Alternative
// }
// type MediaPlaylist struct {
// 	TargetDuration float64
// 	SeqNo          uint64 // EXT-X-MEDIA-SEQUENCE
// 	Segments       []*MediaSegment
// 	Args           string // optional arguments placed after URIs (URI?Args)
// 	Iframe         bool   // EXT-X-I-FRAMES-ONLY
// 	Closed         bool   // is this VOD (closed) or Live (sliding) playlist?
// 	MediaType      MediaType

// 	Key *Key // encryption key displayed before any segments
// 	WV  *WV  // Widevine related tags
// 	// contains filtered or unexported fields
// }
// type Variant struct {
// 	URI       string
// 	Chunklist *MediaPlaylist
// 	VariantParams
// }
// type MasterPlaylist struct {
// 	Variants      []*Variant
// 	Args          string // optional arguments placed after URI (URI?Args)
// 	CypherVersion string // non-standard tag for Widevine (see also WV struct)
// 	// contains filtered or unexported fields
// }
// type Playlist interface {
// 	Encode() *bytes.Buffer
// 	Decode(bytes.Buffer, bool) error
// 	DecodeFrom(reader io.Reader, strict bool) error
// }

// const (
// 	DATETIME = time.RFC3339Nano // Format for EXT-X-PROGRAM-DATE-TIME defined in section 3.4.5
// )

// func NewMasterPlaylist() *MasterPlaylist
// func (p *MasterPlaylist) DecodeFrom(reader io.Reader, strict bool) error
// func (p *MasterPlaylist) Encode() *bytes.Buffer
// func NewMediaPlaylist(winsize uint, capacity uint) (*MediaPlaylist, error)
// func (p *MediaPlaylist) Append(uri string, duration float64, title string) error
// func (p *MediaPlaylist) Close()
// func (p *MediaPlaylist) DecodeFrom(reader io.Reader, strict bool) error
// func (p *MediaPlaylist) DurationAsInt(yes bool)
// func (p *MediaPlaylist) Encode() *bytes.Buffer
// func (p *MediaPlaylist) Remove() (err error)
// func (p *MediaPlaylist) ResetCache()
// func (p *MediaPlaylist) SetDefaultKey(method, uri, iv, keyformat, keyformatversions string)
// func (p *MediaPlaylist) SetDiscontinuity() error
// func (p *MediaPlaylist) SetKey(method, uri, iv, keyformat, keyformatversions string) error
// func (p *MediaPlaylist) SetProgramDateTime(value time.Time) error
// func (p *MediaPlaylist) SetRange(limit, offset int64) error
// func (p *MediaPlaylist) Slide(uri string, duration float64, title string)
// func (p *MasterPlaylist) Decode(data bytes.Buffer, strict bool) error
// func Decode(data bytes.Buffer, strict bool) (Playlist, ListType, error)
// func DecodeFrom(reader io.Reader, strict bool) (Playlist, ListType, error)

// type ListType uint
// const (
// 	// use 0 for not defined type
// 	MASTER ListType = iota + 1
// 	MEDIA
// )
// type Reader struct {
// 	// contains filtered or unexported fields
// }
// func NewReader(rd io.Reader) *Reader
func ServerFileMedia(c *gin.Context) {

	file_name := c.Param("mID")
	// fmt.Println("Filename was connected : ",file_name)
	fileRoot := "assets/"
	typeFile := strings.Split(file_name, ".")
	typeFileName := typeFile[1]
	// fmt.Println("result file type TS :::",typeFileName == "ts")

	if typeFileName == "ts" {
		c.Writer.Header().Set("Content-Type", "application/octet-stream")
	} else {
		c.Writer.Header().Set("Content-Type", "application/x-mpegURL")
	}

	//  c.File(fileRoot+file_name)

	c.File(fileRoot + file_name)
	// c.Status(http.StatusOK)
	// f, err := os.Open(fileRoot+file_name)
	// if err != nil {
	// 	fmt.Println(err)
	//   }
	//   p := NewMasterPlaylist()
	// //   err = p.Decode(bufio.NewReader(f), false)
	//   err = p.Decode(bytes.Buffer(*bufio.NewReader(f)),false)
	//   if err != nil {
	// 	fmt.Println(err)
	//   }

	//   fmt.Printf("Playlist object: %+v\n", p)
	// err = p.Decode(bufio.NewReader(f), false)
	// defer file.Close()
	// stFile, err := file.Stat()
	// data := make([]byte, 1024)
	// count, err := file.Read(data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(count)
	// newFile := http.FileServer(http.Dir(fileRoot+file_name))
	// http.ServeFile(file)
	// fmt.Println(newFile)
	// c.Status(http.StatusOK,{file:newFile})
	// wrFile := os.WriteFile(stFile)

	// c.File(file)
	// c.Status(http.StatusOK,file)

}

type SubtitleURLStruct struct {
	MovieID  string `json:"mID"`
	Language string `json:"lang"`
}
type MediaURLStruct struct {
	MovieID string `json:"mID"`
}

func ServerURLFileSubtitle(c *gin.Context) {
	var subtitleOptions SubtitleURLStruct

	movieID := c.Request.URL.Query().Get("mID")
	subtitle_lang := c.Request.URL.Query().Get("lang")

	subtitleOptions.MovieID = movieID
	subtitleOptions.Language = subtitle_lang

	// fmt.Println("movieID :::",movieID)
	// fmt.Println("subtitle_lang :::",subtitle_lang)
	// c.JSON(http.StatusOK,subtitleOptions)
	fileRoot := "assets/"
	fileName := "example_subtitle" // waiting... db for know name file
	fileType := ".vtt"
	fileLang := strings.ToUpper(subtitleOptions.Language)
	result_file_name := fmt.Sprintf("http://localhost:8080/%s%s%s%s", fileRoot, fileName, fileLang, fileType)
	// fmt.Println("fileName :::",result_file_name)
	// fmt.Println("results path :::",fileRoot+result_file_name)

	// c.Writer.Header().Set("Content-Type","WEBVTT")
	c.JSON(http.StatusOK, result_file_name)
	// c.File(fileRoot+result_file_name)

}

func ServerURLFileMedia(c *gin.Context) {
	var mediaOptions MediaURLStruct
	movieID := c.Request.URL.Query().Get("mID")
	mediaOptions.MovieID = movieID

	fileRoot := "movie/"
	fileName := "hotd_fhd" // waiting... db for know name file
	fileType := ".m3u8"

	result_file_name := fmt.Sprintf("http://localhost:8080/%s%s%s", fileRoot, fileName, fileType)
	c.JSON(http.StatusOK, result_file_name)
}

func ServerFileThumbnail(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "image/jpg")
	Filename := c.Param("file")
	fileRoot := "assets/"
	c.File(fileRoot + Filename)
}

//https://github.com/aofiee/Music-Streaming-HLS-Go-fiber/blob/main/main.go
