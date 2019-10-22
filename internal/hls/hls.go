package hls

import (
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
)

var HomeDir = ".gohls"
var FFProbePath = "ffprobe"
var FFMPEGPath = "ffmpeg"
type Video struct {
	title string
	size int64
	codec string
	path string
	status string
}

//
func Convert(mode string, v Video, distPath string) string {
	args := make([]string, 1)
	if mode == "convert_h264" {
		args = []string{
			"-i", v.path,
			"-codec", "copy",
			"-map", "0",
			"-f", "segment",
			"-vbsf", "h264_mp4toannexb",
			"-segment_format", "mpegts",
			"-segment_time","10",
			"-segment_list", distPath + "/"+ v.title +"/" + v.title +".m3u8",
			distPath + "/" + v.title + "/%03d.ts",
		}
		log.Infof("Executing: ffmpeg %v", args)
		cmd := exec.Command(FFMPEGPath, args...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Infof(string(output))
			//log.Infof("Error ffmpeg", err)
			return "error"
		}

	}
	return "ok"
}
func getCodec(videoPath string) string{
	args := []string{
		"-i", videoPath,
	}
	//log.Infof("Executing: ffmpeg %v", args)
	cmd := exec.Command(FFProbePath, args...)
	r := regexp.MustCompile(`(Video:.+)`)
	output, err := cmd.CombinedOutput()
	result := r.FindAllString(string(output), -1)
	log.Println(string(output))
	if err != nil {
		log.Infof("Error ffmpeg", err)
		//log.Infof(string(output))
	}
	return result[0]
}

func convertHls(v Video, distPath string) Video{
	if _, err := os.Stat(distPath+v.title); err == nil {
		log.Infof("skip %v", v.title)
		return v
	}
	_ = os.Mkdir(distPath+v.title, 0777)
	//存在したらスキップ

	if strings.Contains(v.codec, "h264"){
		log.Infof("convert h264", v.title)
		v.status = Convert("convert_h264", v, distPath)
		if v.status == "error"{
			_ = syscall.Rmdir(distPath + v.title)
		}
	}
	return v
}
const (
	logfile = "./logs/output.log"
)


func ConvertAllFIle(basePath string, distPath string) {
	logfile, err := os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Info(err)
	}
	log.SetOutput(logfile)
	// Find ffmpeg
	ffmpeg, err := exec.LookPath("ffmpeg")
	if err != nil {
		log.Fatal("ffmpeg could not be found in your path", err)
	}

	// Find ffprobe
	ffprobe, err := exec.LookPath("ffprobe")
	if err != nil {
		log.Fatal("ffprobe could not be found in your path", err)
	}

	homeDir, err := homedir.Dir()
	if err != nil {
		log.Fatal("Could not determine home directory", err)
	}
	FFMPEGPath = ffmpeg
	FFProbePath = ffprobe
	HomeDir = path.Join(homeDir, ".gohls")

	files, err :=ioutil.ReadDir(basePath)
	videos := make([]Video,1)
	for _, f := range files{
		if f.Name() == ".gitkeep" {
			continue
		}
		videoPath := filepath.Base(f.Name())
		video := Video{
			title: f.Name(),
			size:  f.Size(),
			codec: getCodec(basePath + videoPath),
			path:  basePath + videoPath,
		}
		videos = append(videos, video)
		video = convertHls(video, distPath)
		log.Info(video)
	}

}
