package mediaInfo

import (
	"encoding/xml"
	"fmt"
	"math"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	tsIO "tsFunction"
)

const version = 0.5

var start time.Time

var traceConsole = false //trace sur la console
var appRep string

// liste des extentions contenant des médias
var containers = map[string]bool{
	".mkv":  true,
	".mka":  true,
	".mks":  true,
	".ogg":  true,
	".ogm":  true,
	".avi":  true,
	".wav":  true,
	".mpeg": true,
	".mpg":  true,
	".vob":  true,
	".mp4":  true,
	".mpgv": true,
	".mpv":  true,
	".m1v":  true,
	".m2v":  true,
	".mp2":  true,
	".mp3":  true,
	".asf":  true,
	".wma":  true,
	".wmv":  true,
	".qt":   true,
	".mov":  true,
	".rm":   true,
	".rmvb": true,
	".ra":   true,
	".ifo":  true,
	".ac3":  true,
	".dts":  true,
	".aac":  true,
	".ape":  true,
	".mac":  true,
	".flac": true,
	".dat":  true,
	".aiff": true,
	".aifc": true,
	".au":   true,
	".iff":  true,
	".paf":  true,
	".sd2":  true,
	".irca": true,
	".w64":  true,
	".mat":  true,
	".pvf":  true,
	".xi":   true,
	".sds":  true,
	".avr":  true,
	".m4v":  true,
}

// structure en retour de l'appel à mediainfo
type mediainfoXml struct {
	XMLName xml.Name `xml:"Mediainfo"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	File    struct {
		Text  string `xml:",chardata"`
		Track []struct {
			Text                     string `xml:",chardata"`
			Type                     string `xml:"type,attr"`
			Streamid                 string `xml:"streamid,attr"`
			UniqueID                 string `xml:"Unique_ID"`
			CompleteName             string `xml:"Complete_name"`
			Format                   string `xml:"Format"`
			FormatVersion            string `xml:"Format_version"`
			FileSize                 string `xml:"File_size"`
			Duration                 string `xml:"Duration"`
			OverallBitRate           string `xml:"Overall_bit_rate"`
			EncodedDate              string `xml:"Encoded_date"`
			WritingApplication       string `xml:"Writing_application"`
			WritingLibrary           string `xml:"Writing_library"`
			DURATION                 string `xml:"DURATION"`
			NUMBEROFFRAMES           string `xml:"NUMBER_OF_FRAMES"`
			NUMBEROFBYTES            string `xml:"NUMBER_OF_BYTES"`
			STATISTICSWRITINGAPP     string `xml:"_STATISTICS_WRITING_APP"`
			STATISTICSWRITINGDATEUTC string `xml:"_STATISTICS_WRITING_DATE_UTC"`
			STATISTICSTAGS           string `xml:"_STATISTICS_TAGS"`
			ID                       string `xml:"ID"`
			FormatInfo               string `xml:"Format_Info"`
			FormatProfile            string `xml:"Format_profile"`
			FormatSettingsCABAC      string `xml:"Format_settings__CABAC"`
			FormatSettingsReFrames   string `xml:"Format_settings__ReFrames"`
			CodecID                  string `xml:"Codec_ID"`
			CodecIDInfo              string `xml:"Codec_ID_Info"`
			BitRate                  string `xml:"Bit_rate"`
			Width                    string `xml:"Width"`
			Height                   string `xml:"Height"`
			DisplayAspectRatio       string `xml:"Display_aspect_ratio"`
			FrameRateMode            string `xml:"Frame_rate_mode"`
			FrameRate                string `xml:"Frame_rate"`
			OriginalFrameRate        string `xml:"Original_frame_rate"`
			MinimumFrameRate         string `xml:"Minimum_frame_rate"`
			MaximumFrameRate         string `xml:"Maximum_frame_rate"`
			ColorSpace               string `xml:"Color_space"`
			ChromaSubsampling        string `xml:"Chroma_subsampling"`
			BitDepth                 string `xml:"Bit_depth"`
			ScanType                 string `xml:"Scan_type"`
			BitsPixelFrame           string `xml:"Bits__Pixel_Frame_"`
			StreamSize               string `xml:"Stream_size"`
			Language                 string `xml:"Language"`
			Default                  string `xml:"Default"`
			Forced                   string `xml:"Forced"`
			ColorPrimaries           string `xml:"Color_primaries"`
			TransferCharacteristics  string `xml:"Transfer_characteristics"`
			MatrixCoefficients       string `xml:"Matrix_coefficients"`
			ModeExtension            string `xml:"Mode_extension"`
			FormatSettingsEndianness string `xml:"Format_settings__Endianness"`
			EncodingSettings         string `xml:"Encoding_settings"`
			BitRateMode              string `xml:"Bit_rate_mode"`
			ChannelS                 string `xml:"Channel_s_"`
			ChannelPositions         string `xml:"Channel_positions"`
			SamplingRate             string `xml:"Sampling_rate"`
			CompressionMode          string `xml:"Compression_mode"`
		} `xml:"track"`
	} `xml:"File"`
}

// structure MediaInfo
type MediaInfo_struct struct {
	General General_struct
	Video   []Video_struct
	Audio   []Audio_struct
	Text    []Text_struct
}

// structure Générale
type General_struct struct {
	Format          string  // MPEG-4
	FormatVersion   string  // Version 2
	FileSize        float64 // 1.43 ( < 1.43 GiB)
	Duration        int64   // 2413 (en sec < 40mn 13s)
	OverallBitRate  int64   // 5098 ( < 5 098 Kbps)
	AudioMultiPiste MultiPiste_struct
	TextMultiPiste  MultiPiste_struct
}

// structure Vidéo
type Video_struct struct {
	Format        string  // AVC
	FormatInfo    string  // Advanced Video Codec
	FormatProfile string  // High@L4.0
	CodecID       string  // V_MPEG4/ISO/AVC
	Duration      int64   // 2413 (en sec < 40mn 13s)
	BitRate       int64   // 4613 ( < 4 613 Kbps)
	Width         int64   // 1920 ( < 1 920 pixels)
	Height        int64   // 1080 ( < 1 080 pixels)
	FrameRateMode string  // Constant/Variable
	FrameRate     float64 // 23.976 ( < 23.976 fps)
	BitDepth      int64   // 8 ( < 8 bits)
	Language      string  // English
}

// structure Audio
type Audio_struct struct {
	Format           string // AC-3
	FormatInfo       string // Audio Coding 3
	CodecID          string // A_AC3
	Duration         int64  // 2413 (en sec < 40mn 13s)
	BitRateMode      string // Constant/Variable
	BitRate          int64  // 384 ( < 384 Kbps)
	ChannelS         int64  // 6 ( < 6 channels)
	ChannelPositions string // Front: L C R, Side: L R, LFE
	ChannelDetail    ChannelDetail_struct
	SamplingRate     float64 // 48.0 ( < 48.0 KHz)
	BitDepth         int64   // 16 ( < 16 bits)
	CompressionMode  string  // Lossy
	Language         string  // English
}

// structure Channel
type ChannelDetail_struct struct {
	FrontL bool
	FrontC bool
	FrontR bool
	RearL  bool
	RearR  bool
	Sub    bool
}

// structure sous-titre
type Text_struct struct {
	Format      string // UTF-8
	CodecID     string // S_TEXT/UTF8
	CodecIDInfo string // UTF-8 Plain Text
	Language    string // English
}

// structure sous-titre
type MultiPiste_struct struct {
	Format   string // UTF-8 / UTF-8
	Language string // English / French
	NoFrench bool
}

// init() : initialisation du composant
func init() {
	start = time.Now()

	tsIO.TraceConsole = &traceConsole //trace sur la console

	var err error
	// récupère le répertoire de l'application
	appRep, err = tsIO.GetAppPath()
	if err != nil {
		panic(fmt.Sprint("  init > ", err))
	}
	tsIO.PrintConsole("App path : " + appRep)
}

// IsMediaFile() - détemine si le suffixe du fichier correspond à un media (audio ou vidéo)
func IsMediaFile(ext string) bool {
	result := false
	if _, ok := containers[strings.ToLower(ext)]; ok {
		result = true
	}

	return result
}

// GetMediaInfo() : récupère les infos du média dans MediaInfo_struct
func GetMediaInfo(fileName string) MediaInfo_struct {
	var mediainfo_cmd string
	mediainfo_cmd, err := exec.LookPath("mediainfo")
	if err != nil {
		panic(fmt.Sprint("  could not find path to 'mediainfo': ", err))
	}
	tsIO.PrintConsole("-- found 'mediainfo' command: ", mediainfo_cmd)

	out, err := exec.Command(mediainfo_cmd, "--Output=XML", fileName).Output()
	if err != nil {
		panic(fmt.Sprint("Command: mediainfo ", err))
	}
	var result mediainfoXml
	err = xml.Unmarshal(out, &result) //DECODAGE

	//	fmt.Println(result)
	var mediaInfo MediaInfo_struct
	for _, track := range result.File.Track {
		switch track.Type {
		case "General":
			var general General_struct
			general.Format = track.Format
			general.FormatVersion = track.FormatVersion
			general.FileSize = extractFileSize(track.FileSize)
			general.Duration = extractDuration(track.Duration)
			general.OverallBitRate = extractBitRate(track.OverallBitRate)
			mediaInfo.General = general
		case "Video":
			var video Video_struct
			video.Format = track.Format
			video.FormatInfo = track.FormatInfo
			video.FormatProfile = track.FormatProfile
			video.CodecID = track.CodecID
			video.Duration = extractDuration(track.Duration)
			video.BitRate = extractBitRate(track.BitRate)
			video.Width = extractSize(track.Width)
			video.Height = extractSize(track.Width)
			video.FrameRateMode = track.FrameRateMode
			if track.FrameRate != "" {
				video.FrameRate = extractFrameRate(track.FrameRate)
			} else if track.OverallBitRate != "" {
				video.FrameRate = extractFrameRate(track.OverallBitRate)
			}
			video.BitDepth = extractBitDepth(track.BitDepth)
			video.Language = track.Language
			mediaInfo.Video = append(mediaInfo.Video, video)
		case "Audio":
			var audio Audio_struct
			audio.Format = track.Format
			audio.FormatInfo = track.FormatInfo
			audio.CodecID = track.CodecID
			audio.Duration = extractDuration(track.Duration)
			audio.BitRateMode = track.BitRateMode
			audio.BitRate = extractBitRate(track.BitRate)
			audio.ChannelS = extractChannel(track.ChannelS)
			audio.ChannelPositions = track.ChannelPositions
			audio.ChannelDetail = getChannelDetail(track.ChannelPositions)
			audio.SamplingRate = extractSamplingRate(track.SamplingRate)
			audio.BitDepth = extractBitDepth(track.BitDepth)
			audio.CompressionMode = track.CompressionMode
			audio.Language = track.Language
			mediaInfo.Audio = append(mediaInfo.Audio, audio)
		case "Text":
			var text Text_struct
			text.Format = track.Format
			text.CodecID = track.CodecID
			text.CodecIDInfo = track.CodecIDInfo
			text.Language = track.Language
			mediaInfo.Text = append(mediaInfo.Text, text)
		}
	}
	if len(mediaInfo.Audio) > 1 {
		var lang []string
		var format []string
		for _, audio := range mediaInfo.Audio {
			format = append(format, audio.Format)
			lang = append(lang, audio.Language)
			if audio.Language != "French" {
				mediaInfo.General.AudioMultiPiste.NoFrench = true
			}
		}
		mediaInfo.General.AudioMultiPiste.Format = strings.Join(format, " / ")
		mediaInfo.General.AudioMultiPiste.Language = strings.Join(lang, " / ")
	}

	if len(mediaInfo.Text) > 1 {
		var lang []string
		var format []string
		for _, text := range mediaInfo.Text {
			format = append(format, text.Format)
			lang = append(lang, text.Language)
			if text.Language != "French" {
				mediaInfo.General.TextMultiPiste.NoFrench = true
			}
		}
		mediaInfo.General.TextMultiPiste.Format = strings.Join(format, " / ")
		mediaInfo.General.TextMultiPiste.Language = strings.Join(lang, " / ")
	}

	return mediaInfo
}

// extractFileSize() return size in GiB (1.43 GiB --> 1.43  ou  785 MiB --> 0.766)
func extractFileSize(size string) float64 {
	if size == "" {
		return 0.0
	} else {
		mots := strings.Fields(size)
		val, err := strconv.ParseFloat(mots[0], 64)
		if err != nil {
			panic(fmt.Sprint("  extractFileSize > ParseFloat ", err))
		}
		if mots[1] == "MiB" {
			val /= 1024
		}
		val = math.RoundToEven(val*10) / 10

		return val
	}
}

// extractSize() return size in pixel (1 920 pixels --> 1920)
func extractSize(size string) int64 {
	if size == "" {
		return 0
	} else {
		mots := strings.Fields(size)
		var tmp string
		for _, val := range mots[:len(mots)-1] {
			tmp += val
		}
		result, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			panic(fmt.Sprint("  extractSize > ParseInt ", err))
		}
		return result
	}
}

// extractDuration() return durée in sec (40mn 13s --> 2413)
func extractDuration(duration string) int64 {
	if duration == "" {
		return 0
	} else {
		mots := strings.Fields(duration)
		var result int64
		result = 0
		for _, val := range mots {
			var re = regexp.MustCompile(`([0-9]*)([a-zA-Z]*)`)
			matches := re.FindStringSubmatch(val)
			if len(matches) == 3 {
				tmp, err := strconv.ParseInt(matches[1], 10, 64)
				if err != nil {
					panic(fmt.Sprint("  extractDuration > ParseInt ", err))
				}
				switch matches[2] {
				case "h":
					result += tmp * 3600
				case "mn":
					result += tmp * 60
				case "s":
					result += tmp
				}
			} else {
				panic(fmt.Sprint("  extractDuration > regexp ", val))
			}
		}
		return result
	}
}

// extractBitRate() return bitRate en Kbps (5 098 Kbps  --> 5098)
func extractBitRate(bitRate string) int64 {
	if bitRate == "" {
		return 0
	} else {
		mots := strings.Fields(bitRate)
		var tmp string
		for _, val := range mots[:len(mots)-1] {
			tmp += val
		}
		result, err := strconv.ParseInt(tmp, 10, 64)
		if err != nil {
			panic(fmt.Sprint("  extractBitRate > ParseInt ", err))
		}
		return result
	}
}

// extractFrameRate() return frameRate in fps (23.976 fps --> 23.976)
func extractFrameRate(frame string) float64 {
	if frame == "" {
		return 0.0
	} else {
		mots := strings.Fields(frame)
		val, err := strconv.ParseFloat(mots[0], 64)
		if err != nil {
			panic(fmt.Sprint("  extractFrameRate > ParseFloat ", err))
		}
		return val
	}
}

// extractBitDepth() return bitDepth in bits (8 bits --> 8)
func extractBitDepth(bitDepth string) int64 {
	if bitDepth == "" {
		return 0
	} else {
		mots := strings.Fields(bitDepth)
		val, err := strconv.ParseInt(mots[0], 10, 64)
		if err != nil {
			panic(fmt.Sprint("  extractBitDepth > ParseInt ", err))
		}
		return val
	}
}

// extractChannel() return nb audio channel (6 channels --> 6)
func extractChannel(channel string) int64 {
	if channel == "" {
		return 0
	} else {
		mots := strings.Fields(channel)
		val, err := strconv.ParseInt(mots[0], 10, 64)
		if err != nil {
			panic(fmt.Sprint("  extractChannel > ParseInt ", err))
		}
		return val
	}
}

// getChannelDetail() // Front: L C R, Side: L R, LFE
func getChannelDetail(channelPositions string) ChannelDetail_struct {
	var channelDetail ChannelDetail_struct
	if channelPositions != "" {
		lines := strings.Split(channelPositions, ",")
		for _, val := range lines {
			mots := strings.Fields(strings.TrimSpace(val))
			switch mots[0] {
			case "Front:":
				for _, val := range mots[1:] {
					switch val {
					case "L":
						channelDetail.FrontL = true
					case "C":
						channelDetail.FrontC = true
					case "R":
						channelDetail.FrontR = true
					}
				}
			case "Side:":
				for _, val := range mots[1:] {
					switch val {
					case "L":
						channelDetail.RearL = true
					case "R":
						channelDetail.RearR = true
					}
				}
			case "LFE":
				channelDetail.Sub = true
			}

		}
	}
	return channelDetail
}

// extractSamplingRate() return sampling rate in Khz  (48.0 KHz --> 48.0)
func extractSamplingRate(rate string) float64 {
	if rate == "" {
		return 0.0
	} else {
		mots := strings.Fields(rate)
		val, err := strconv.ParseFloat(mots[0], 64)
		if err != nil {
			panic(fmt.Sprint("  extractSamplingRate > ParseFloat ", err))
		}
		return val
	}
}
