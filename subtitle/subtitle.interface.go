package subtitle

type ParamsSubtitle struct {
	Directory string `json:"directory"`
}

type SubtitleFile struct {
	Language string
	FileName string
}