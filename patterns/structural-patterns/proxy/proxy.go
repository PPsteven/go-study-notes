package main

import (
	"fmt"
	"time"
)

var (
	videoInfo   map[string]string
	video       map[string][]byte
	videoExpire map[string]time.Time
)

type ThirdPartyYouTubeLib interface {
	ListVideos()
	GetVideoInfo(string)
	DownloadVideo(string)
}

type thirdPartyYouTubeLib struct {

}

func (y *thirdPartyYouTubeLib) ListVideos() {
	fmt.Println("list videos")
}

func (y *thirdPartyYouTubeLib) GetVideoInfo(name string) {
	if videoInfo == nil {
		videoInfo = make(map[string]string)
	}
	videoInfo[name] = fmt.Sprintf("movie: %s", name)
	fmt.Println("get video info: ", videoInfo[name])
}

func (y *thirdPartyYouTubeLib) DownloadVideo(name string) {
	if video == nil || videoExpire == nil {
		video = make(map[string][]byte)
		videoExpire = make(map[string]time.Time)
	}
	video[name] = []byte(fmt.Sprintf("movie %s: huge video...", name))
	videoExpire[name] = time.Now().Add(time.Second * 3)
	fmt.Println("download video info: ", string(video[name]))
}

type CachedThirdPartyYouTubeLib struct {
	tube ThirdPartyYouTubeLib
}

func (y *CachedThirdPartyYouTubeLib) ListVideos() {
	y.tube.ListVideos()
}

func (y *CachedThirdPartyYouTubeLib) GetVideoInfo(name string) {
	if y.cacheExists(name) {
		fmt.Println("video cache exists use cache, ", videoInfo[name])
	}
	y.tube.GetVideoInfo(name)
}

func (y *CachedThirdPartyYouTubeLib) DownloadVideo(name string) {
	if y.downloadExists(name) {
		fmt.Println("download exists use cache")
	}
	y.tube.DownloadVideo(name)
}

func (y *CachedThirdPartyYouTubeLib) downloadExists(name string) bool {
	if len(video[name]) > 0 && videoExpire[name].After(time.Now()) {
		return true
	}
	return false
}

func (y *CachedThirdPartyYouTubeLib) cacheExists(name string) bool {
	if len(videoInfo[name]) > 0 {
		return true
	}
	return false
}

func main() {
	cachedYoutub := &CachedThirdPartyYouTubeLib{tube: &thirdPartyYouTubeLib{}}
	cachedYoutub.GetVideoInfo("forrest gump")
	cachedYoutub.DownloadVideo("forrest gump")
	cachedYoutub.GetVideoInfo("forrest gump")
	cachedYoutub.DownloadVideo("forrest gump")
	time.Sleep(4*time.Second)
	fmt.Println("after expire...")
	cachedYoutub.DownloadVideo("forrest gump")
}