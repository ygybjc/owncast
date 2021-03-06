package config

import (
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

func getDefaults() config {
	defaults := config{}
	defaults.WebServerPort = 8080
	defaults.FFMpegPath = getDefaultFFMpegPath()
	defaults.VideoSettings.ChunkLengthInSeconds = 4
	defaults.Files.MaxNumberInPlaylist = 5
	defaults.YP.Enabled = false
	defaults.YP.YPServiceURL = "https://yp.owncast.online"

	defaultQuality := StreamQuality{
		IsAudioPassthrough: true,
		VideoBitrate:       1200,
		EncoderPreset:      "veryfast",
		Framerate:          24,
	}
	defaults.VideoSettings.StreamQualities = []StreamQuality{defaultQuality}

	return defaults
}

func getDefaultFFMpegPath() string {
	cmd := exec.Command("which", "ffmpeg")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Panicln("Unable to determine path to ffmpeg. Please specify it in the config file.")
	}

	path := strings.TrimSpace(string(out))

	return path
}
