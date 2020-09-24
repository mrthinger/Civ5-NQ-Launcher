package main

//CurrentLinks response for current getting current links from server
type CurrentLinks struct {
	Map string `json:"map"`
	Mod string `json:"mod"`
}

//BuildInfo contains info about specific version of binary
type BuildInfo struct {
	Version     int    `json:"version"`
	DownloadURL string `json:"downloadUrl"`
}
