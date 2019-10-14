package mediaInfo

import (
	"encoding/xml"
	"fmt"
	"math"
	"os/exec"
	"path/filepath"
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

const CHANNEl1 = "1ch: Mono"
const CHANNEl2 = "2ch: Stereo"
const CHANNEl3 = "3ch: Stereo 2.1"
const CHANNEl5 = "5ch: Surround"
const CHANNEl6 = "6ch: Surround"
const CHANNEl8 = "8ch: Surround +"

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
	Conteneur       string  // mkv
	Format          string  // MPEG-4
	FormatVersion   string  // Version 2
	FileSize        float64 // 1.43 ( < 1.43 GiB)
	Duration        int64   // 2413 (en sec < 40mn 13s)
	DurationAff     int64   // 40
	OverallBitRate  int64   // 5098 ( < 5 098 Kbps)
	XFileSize       string  // 1.43 ( < 1.43 GiB)
	XDuration       string  // 2413 (en sec < 40mn 13s)
	XDurationAff    string  // 40
	XOverallBitRate string  // 5098 ( < 5 098 Kbps)
	AudioMultiPiste MultiPiste_struct
	TextMultiPiste  MultiPiste_struct
}

// structure Vidéo
type Video_struct struct {
	Format        string // AVC
	FormatInfo    string // Advanced Video Codec
	FormatProfile string // High@L4.0
	CodecID       string // V_MPEG4/ISO/AVC
	CodecIDInfo   string
	CodecV        string
	Duration      int64  // 2413 (en sec < 40mn 13s)
	DurationAff   int64  // 40
	BitRate       int64  // 4613 ( < 4 613 Kbps)
	Width         int64  // 1920 ( < 1 920 pixels)
	Height        int64  // 1080 ( < 1 080 pixels)
	FrameRateMode string // Constant/Variable
	FrameRate     string // 23.976 ( < 23.976 fps)
	BitDepth      int64  // 8 ( < 8 bits)
	Language      string // English
	XDuration     string // 2413 (en sec < 40mn 13s)
	XDurationAff  string // 40
	XBitRate      string // 4613 ( < 4 613 Kbps)
	XWidth        string // 1920 ( < 1 920 pixels)
	XHeight       string // 1080 ( < 1 080 pixels)
	XBitDepth     string // 8 ( < 8 bits)
}

// structure Audio
type Audio_struct struct {
	Format           string // AC-3
	FormatInfo       string // Audio Coding 3
	CodecID          string // A_AC3
	CodecIDInfo      string
	CodecA           string
	Duration         int64  // 2413 (en sec < 40mn 13s)
	DurationAff      int64  // 40
	BitRateMode      string // Constant/Variable
	BitRate          int64  // 384 ( < 384 Kbps)
	Channel          int64  // 6 ( < 6 channels)
	ChannelPositions string // Front: L C R, Side: L R, LFE
	ChannelDetail    ChannelDetail_struct
	ChannelAff       string  // 6ch: Surround
	SamplingRate     float64 // 48.0 ( < 48.0 KHz)
	BitDepth         int64   // 16 ( < 16 bits)
	CompressionMode  string  // Lossy
	Language         string  // English
	XDuration        string  // 2413 (en sec < 40mn 13s)
	XDurationAff     string  // 40
	XBitRate         string  // 384 ( < 384 Kbps)
	XChannel         string  // 6 ( < 6 channels)
	XSamplingRate    string  // 48.0 ( < 48.0 KHz)
	XBitDepth        string  // 16 ( < 16 bits)
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
			general.FileSize, general.XFileSize = extractFileSize(track.FileSize)
			general.Duration, general.XDuration = extractDuration(track.Duration)
			general.DurationAff, general.XDurationAff = extractDurationMN(general.Duration
			general.OverallBitRate, general.XOverallBitRate = extractBitRate(track.OverallBitRate)
			mediaInfo.General = general
		case "Video":
			var video Video_struct
			video.Format = track.Format
			video.FormatInfo = track.FormatInfo
			video.FormatProfile = track.FormatProfile
			video.CodecID = track.CodecID
			video.CodecIDInfo = track.CodecIDInfo
			video.CodecV = getCodecVideo(video.Format, video.FormatProfile, video.CodecID)
			video.Duration, video.XDuration = extractDuration(track.Duration)
			video.DurationAff, video.XDurationAff = extractDurationMN(video.Duration)
			video.BitRate, video.XBitRate = extractBitRate(track.BitRate)
			video.Width, video.XWidth = extractSize(track.Width)
			video.Height, video.XHeight = extractSize(track.Height)
			video.FrameRateMode = track.FrameRateMode
			if track.FrameRate != "" {
				video.FrameRate = transcodeVideoFrameRate(extractFrameRate(track.FrameRate))
			} else if track.OverallBitRate != "" {
				video.FrameRate = transcodeVideoFrameRate(extractFrameRate(track.OverallBitRate))
			}
			video.BitDepth, video.XBitDepth = extractBitDepth(track.BitDepth)
			video.Language = track.Language
			mediaInfo.Video = append(mediaInfo.Video, video)
		case "Audio":
			var audio Audio_struct
			audio.Format = track.Format
			audio.FormatInfo = track.FormatInfo
			audio.CodecID = track.CodecID
			audio.CodecIDInfo = track.CodecIDInfo
			audio.CodecA = getCodeCodecAudio(audio.FormatInfo, audio.CodecID, audio.CodecIDInfo)
			audio.Duration, audio.XDuration = extractDuration(track.Duration)
			audio.DurationAff, audio.XDurationAff = extractDurationMN(audio.Duration)
			audio.BitRateMode = track.BitRateMode
			audio.BitRate, audio.XBitRate = extractBitRate(track.BitRate)
			audio.Channel, audio.XChannel = extractChannel(track.ChannelS)
			audio.ChannelPositions = track.ChannelPositions
			audio.ChannelDetail = getChannelDetail(track.ChannelPositions)
			audio.ChannelAff = getChannelAff(audio.Channel)
			audio.SamplingRate, audio.XSamplingRate = extractSamplingRate(track.SamplingRate)
			audio.BitDepth, audio.XBitDepth = extractBitDepth(track.BitDepth)
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
	if len(mediaInfo.Audio) > 0 {
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

	if len(mediaInfo.Text) > 0 {
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

	mediaInfo.General.Conteneur = strings.ToLower(filepath.Ext(fileName))

	return mediaInfo
}

// extractFileSize() return size in GiB (1.43 GiB --> 1.43  ou  785 MiB --> 0.766)
func extractFileSize(size string) (float64, string) {
	if size == "" {
		return 0.0, "?"
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

		return val, strconv.FormatFloat(val, 'f', 2, 64)
	}
}

// extractSize() return size in pixel (1 920 pixels --> 1920)
func extractSize(size string) (int64, string) {
	if size == "" {
		return 0, "?"
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
		return result, strconv.FormatInt(result, 10)
	}
}

// extractDuration() return durée in sec (40mn 13s --> 2413)
func extractDuration(duration string) (int64, string) {
	if duration == "" {
		return 0, "?"
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
		return result, strconv.FormatInt(result, 10)
	}
}

// extractDurationMN() convertir la durée sec -> mn
func extractDurationMN(duration int64) (int64, string) {
	if duration == 0 {
		return 0, "?"
	} else {
		result := duration / 60
		return result, strconv.FormatInt(result, 10)
	}
}

// extractBitRate() return bitRate en Kbps (5 098 Kbps  --> 5098)
func extractBitRate(bitRate string) (int64, string) {
	if bitRate == "" {
		return 0, "?"
	} else {
		mots := strings.Fields(bitRate)
		var tmp string
		for _, val := range mots[:len(mots)-1] {
			tmp += val
		}
		var result int64
		var err error
		if strings.Contains(tmp, ".") {
			var tmp float64
			tmp, err = strconv.ParseFloat(tmp, 64)
			result = int64(tmp)
		} else {
			result, err = strconv.ParseInt(tmp, 10, 64)
		}
		if err != nil {
			panic(fmt.Sprint("  extractBitRate > ParseInt/PaeseFloat ", err))
		}
		return result, strconv.FormatInt(result, 10)
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
func extractBitDepth(bitDepth string) (int64, string) {
	if bitDepth == "" {
		return 0, "?"
	} else {
		mots := strings.Fields(bitDepth)
		val, err := strconv.ParseInt(mots[0], 10, 64)
		if err != nil {
			panic(fmt.Sprint("  extractBitDepth > ParseInt ", err))
		}
		return val, strconv.FormatInt(val, 10)
	}
}

// extractChannel() return nb audio channel (6 channels --> 6)
func extractChannel(channel string) (int64, string) {
	if channel == "" {
		return 0, "?"
	} else {
		mots := strings.Fields(channel)
		val, err := strconv.ParseInt(mots[0], 10, 64)
		if err != nil {
			panic(fmt.Sprint("  extractChannel > ParseInt ", err))
		}
		return val, strconv.FormatInt(val, 10)
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

//### getChannelAff() - transcode les canaux audio pour faciliter la lecture
func getChannelAff(channel int64) string {
	var retour string
	switch channel {
	case 1:
		retour = CHANNEl1 // "1ch: Mono"
	case 2:
		retour = CHANNEl2 // "2ch: Stéréo"
	case 3:
		retour = CHANNEl3 // "3ch: Stéréo 2.1"
	case 5:
		retour = CHANNEl5 // "5ch: Surround"
	case 6:
		retour = CHANNEl6 // "6ch: Surround"
	case 8:
		retour = CHANNEl8 // "8ch: Surround +"
	default:
		retour = strconv.FormatInt(channel, 10)
	}
	return retour
}

// extractSamplingRate() return sampling rate in Khz  (48.0 KHz --> 48.0)
func extractSamplingRate(rate string) (float64, string) {
	if rate == "" {
		return 0.0, "?"
	} else {
		mots := strings.Fields(rate)
		val, err := strconv.ParseFloat(mots[0], 64)
		if err != nil {
			panic(fmt.Sprint("  extractSamplingRate > ParseFloat ", err))
		}
		return val, strconv.FormatFloat(val, 'f', 1, 64)
	}
}

//### getCodecVideo() - transcode le codec vidéo pour faciliter la lecture
//		Format: 'AVC'
//		FormatInfo: 'Advanced Video Codec'
//		FormatProfile: 'High@L4.1'
//		CodecID: 'V_MPEG4/ISO/AVC'
//		CodecIDInfo

func getCodecVideo(format string, formatProfile string, codecID string) string {
	var codecV string
	//	if videoCodecHint == "divx 3 low" {
	//		codecV = "DivX 3 Low"
	if codecID == "dx50" {
		codecV = "DivX 5"
	} else {
		switch format {
		case "XVID", "xvid":
			codecV = "XviD"
		case "DIV3":
			codecV = "DivX 3"
		case "DIV4":
			codecV = "DivX 4"
		case "MPEGVIDEO": //&& codec == "mpeg-1v" {
			codecV = "MPEG-1"
		case "MPEG-4VISUAL":
			switch codecID {
			case "mp42":
				codecV = "MPEG-4"
			case "divx":
				codecV = "DivX 4"
			case "xvid":
				codecV = "XviD"
			default:
				codecV = "MPEG-4"
			}
		case "MPEG-4":
			codecV = format
		case "AVC":
			codecV = "X264"
			mots := strings.Split(formatProfile, "@")
			val := mots[1][1:]
			if !strings.Contains(val, ".") {
				val += ".0"
			}
			codecV += " - " + val
		case "hevc", "HEVC":
			codecV = "X265"
		default:
			codecV = "????"
		}
	}
	return codecV
}

//### getCodeCodecAudio() - transcode le codec audio pour faciliter la lecture
func getCodeCodecAudio(format string, codec string, codecHint string) string {
	var codecA string
	if codecHint == "mp3" || codec == "mpa1l3" {
		codecA = "MP3"
	} else if codecHint == "mp2" || codec == "mpa1l2" {
		codecA = "MP2"
	} else if strings.ToUpper(format) == "VORBIS" {
		codecA = "Vorbis"
	} else {
		codecA = strings.ToUpper(format)
	}
	return codecA
}

//### transcodeVideoFrameRate - transcode le framerate vidéo pour faciliter la lecture
func transcodeVideoFrameRate(frameRate float64) string {
	result := "?"
	if frameRate != 0 {
		result = ""
		frameRates := []float64{23.000, 23.976, 24.000, 25.000, 26.000, 29.970, 30.000, 48.000, 50.000, 60.000}
		for _, val := range frameRates {
			if val == frameRate {
				result = strconv.FormatFloat(val, 'f', 3, 64) // 3 decimales
				break
			}
		}
		if result == "" {
			result = strconv.FormatFloat(frameRate, 'f', 0, 64) + ".xxx"
		}
	}
	return result
}
