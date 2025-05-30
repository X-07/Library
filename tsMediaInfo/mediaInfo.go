package tsMediaInfo

import (
	"encoding/xml"
	"fmt"
	"math"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	tsIO "github.com/X-07/Library/tsUtils"
)

// const version = 0.8

var start time.Time

var traceConsole = false //trace sur la console
var appRep string

const channel1 = "1ch: Mono"
const channel2 = "2ch: Stereo"
const channel3 = "3ch: Stereo 2.1"
const channel5 = "5ch: Surround"
const channel6 = "6ch: Surround"
const channel8 = "8ch: Surround +"

// liste des extensions contenant des médias
var containers = map[string]bool{
	".mkv":  true,
	".mka":  true,
	".mks":  true,
	".ogg":  true,
	".ogv":  true,
	".ogm":  true,
	".avi":  true,
	".wav":  true,
	".mpeg": true,
	".mpg":  true,
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

// // MediaInfoXML1404 structure en retour de l'appel à mediainfo le programme
// type MediaInfoXML1404 struct {
// 	XMLName xml.Name `xml:"MediaInfo"`
// 	Version string   `xml:"version,attr"`
// 	File    struct {
// 		Track []struct {
// 			Type                     string `xml:"type,attr"`
// 			StreamID                 string `xml:"streamid,attr"`
// 			UniqueID                 string `xml:"Unique_ID"`
// 			CompleteName             string `xml:"Complete_name"`
// 			Format                   string `xml:"Format"`
// 			FormatVersion            string `xml:"Format_version"`
// 			FileSize                 string `xml:"File_size"`
// 			Duration                 string `xml:"Duration"`
// 			OverallBitRate           string `xml:"Overall_bit_rate"`
// 			EncodedDate              string `xml:"Encoded_date"`
// 			WritingApplication       string `xml:"Writing_application"`
// 			WritingLibrary           string `xml:"Writing_library"`
// 			DURATION                 string `xml:"DURATION"`
// 			NUMBEROFFRAMES           string `xml:"NUMBER_OF_FRAMES"`
// 			NUMBEROFBYTES            string `xml:"NUMBER_OF_BYTES"`
// 			STATISTICSWRITINGAPP     string `xml:"_STATISTICS_WRITING_APP"`
// 			STATISTICSWRITINGDATEUTC string `xml:"_STATISTICS_WRITING_DATE_UTC"`
// 			STATISTICSTAGS           string `xml:"_STATISTICS_TAGS"`
// 			ID                       string `xml:"ID"`
// 			FormatInfo               string `xml:"Format_Info"`
// 			FormatProfile            string `xml:"Format_profile"`
// 			FormatSettingsCABAC      string `xml:"Format_settings__CABAC"`
// 			FormatSettingsReFrames   string `xml:"Format_settings__ReFrames"`
// 			CodecID                  string `xml:"Codec_ID"`
// 			CodecIDInfo              string `xml:"Codec_ID_Info"`
// 			CodecIDHint              string `xml:"Codec_ID_Hint"`
// 			BitRate                  string `xml:"Bit_rate"`
// 			NominalBitRate           string `xml:"Nominal_bit_rate"`
// 			Width                    string `xml:"Width"`
// 			Height                   string `xml:"Height"`
// 			DisplayAspectRatio       string `xml:"Display_aspect_ratio"`
// 			FrameRateMode            string `xml:"Frame_rate_mode"`
// 			FrameRate                string `xml:"Frame_rate"`
// 			OriginalFrameRate        string `xml:"Original_frame_rate"`
// 			MinimumFrameRate         string `xml:"Minimum_frame_rate"`
// 			MaximumFrameRate         string `xml:"Maximum_frame_rate"`
// 			ColorSpace               string `xml:"Color_space"`
// 			ChromaSubsampling        string `xml:"Chroma_subsampling"`
// 			BitDepth                 string `xml:"Bit_depth"`
// 			ScanType                 string `xml:"Scan_type"`
// 			BitsPixelFrame           string `xml:"Bits__Pixel_Frame_"`
// 			StreamSize               string `xml:"Stream_size"`
// 			Language                 string `xml:"Language"`
// 			Default                  string `xml:"Default"`
// 			Forced                   string `xml:"Forced"`
// 			ColorPrimaries           string `xml:"Color_primaries"`
// 			TransferCharacteristics  string `xml:"Transfer_characteristics"`
// 			MatrixCoefficients       string `xml:"Matrix_coefficients"`
// 			ModeExtension            string `xml:"Mode_extension"`
// 			FormatSettingsEndianness string `xml:"Format_settings__Endianness"`
// 			EncodingSettings         string `xml:"Encoding_settings"`
// 			BitRateMode              string `xml:"Bit_rate_mode"`
// 			ChannelS                 string `xml:"Channel_s_"`
// 			ChannelPositions         string `xml:"Channel_positions"`
// 			SamplingRate             string `xml:"Sampling_rate"`
// 			CompressionMode          string `xml:"Compression_mode"`
// 		} `xml:"track"`
// 	} `xml:"File"`
// }

// MediaInfoXML structure en retour de l'appel au programme 'mediainfo'
type MediaInfoXML struct {
	XMLName         xml.Name `xml:"MediaInfo"`
	Version         string   `xml:"version,attr"`
	CreatingLibrary struct {
		Version string `xml:"version,attr"`
	} `xml:"creatingLibrary"`
	Media struct {
		CompleteName string `xml:"ref,attr"`
		Track        []struct {
			Type                     string `xml:"type,attr"`
			Typeorder                string `xml:"typeorder,attr"`
			UniqueID                 string `xml:"UniqueID"`
			VideoCount               string `xml:"VideoCount"`
			AudioCount               string `xml:"AudioCount"`
			TextCount                string `xml:"TextCount"`
			MenuCount                string `xml:"MenuCount"`
			FileExtension            string `xml:"FileExtension"`
			Format                   string `xml:"Format"`
			FormatVersion            string `xml:"Format_Version"`
			FileSize                 string `xml:"FileSize"`
			Duration                 string `xml:"Duration"`
			OverallBitRate           string `xml:"OverallBitRate"`
			NominalBitRate           string `xml:"NominalBitRate"`
			BitRateNominal           string `xml:"BitRate_Nominal"`
			FrameRate                string `xml:"FrameRate"`
			FrameCount               string `xml:"FrameCount"`
			StreamSize               string `xml:"StreamSize"`
			IsStreamable             string `xml:"IsStreamable"`
			Title                    string `xml:"Title"`
			Movie                    string `xml:"Movie"`
			EncodedBy                string `xml:"EncodedBy"`
			ContentType              string `xml:"ContentType"`
			EncodedDate              string `xml:"Encoded_Date"`
			FileModifiedDate         string `xml:"File_Modified_Date"`
			FileModifiedDateLocal    string `xml:"File_Modified_Date_Local"`
			EncodedApplication       string `xml:"Encoded_Application"`
			EncodedLibrary           string `xml:"Encoded_Library"`
			Cover                    string `xml:"Cover"`
			StreamOrder              string `xml:"StreamOrder"`
			ID                       string `xml:"ID"`
			FormatProfile            string `xml:"Format_Profile"`
			FormatLevel              string `xml:"Format_Level"`
			FormatSettingsCABAC      string `xml:"Format_Settings_CABAC"`
			FormatSettingsRefFrames  string `xml:"Format_Settings_RefFrames"`
			CodecID                  string `xml:"CodecID"`
			BitRate                  string `xml:"BitRate"`
			Width                    string `xml:"Width"`
			Height                   string `xml:"Height"`
			SampledWidth             string `xml:"Sampled_Width"`
			SampledHeight            string `xml:"Sampled_Height"`
			PixelAspectRatio         string `xml:"PixelAspectRatio"`
			DisplayAspectRatio       string `xml:"DisplayAspectRatio"`
			FrameRateMode            string `xml:"FrameRate_Mode"`
			ColorSpace               string `xml:"ColorSpace"`
			ChromaSubsampling        string `xml:"ChromaSubsampling"`
			BitDepth                 string `xml:"BitDepth"`
			ScanType                 string `xml:"ScanType"`
			Delay                    string `xml:"Delay"`
			EncodedLibraryName       string `xml:"Encoded_Library_Name"`
			EncodedLibraryVersion    string `xml:"Encoded_Library_Version"`
			EncodedLibrarySettings   string `xml:"Encoded_Library_Settings"`
			Language                 string `xml:"Language"`
			Default                  string `xml:"Default"`
			Forced                   string `xml:"Forced"`
			ColourRange              string `xml:"colour_range"`
			ColourDescriptionPresent string `xml:"colour_description_present"`
			ColourPrimaries          string `xml:"colour_primaries"`
			TransferCharacteristics  string `xml:"transfer_characteristics"`
			MatrixCoefficients       string `xml:"matrix_coefficients"`
			FormatSettingsEndianness string `xml:"Format_Settings_Endianness"`
			BitRateMode              string `xml:"BitRate_Mode"`
			Channels                 string `xml:"Channels"`
			ChannelPositions         string `xml:"ChannelPositions"`
			ChannelLayout            string `xml:"ChannelLayout"`
			SamplesPerFrame          string `xml:"SamplesPerFrame"`
			SamplingRate             string `xml:"SamplingRate"`
			SamplingCount            string `xml:"SamplingCount"`
			CompressionMode          string `xml:"Compression_Mode"`
			DelaySource              string `xml:"Delay_Source"`
			StreamSizeProportion     string `xml:"StreamSize_Proportion"`
			ServiceKind              string `xml:"ServiceKind"`
			ElementCount             string `xml:"ElementCount"`
		} `xml:"track"`
	} `xml:"media"`
}

// MediaInfo : structure MediaInfo (MediaInfo)
type MediaInfo struct {
	General mediaInfoGeneral
	Video   []mediaInfoVideo
	Audio   []mediaInfoAudio
	Text    []mediaInfoText
}

// mediaInfoGeneral : structure Générale (General_struct)
type mediaInfoGeneral struct {
	Conteneur       string  // mkv
	Format          string  // Matroska
	FormatVersion   string  // 4
	FileSize        float64 // 1.89 ( < 2025275547 B)
	Duration        int64   // 5315 ( < 5314.940 s)
	DurationAff     int64   // 89
	OverallBitRate  int64   // 3048 ( < 3048426 bps)
	XFileSize       string  // 1.89 ( < 2025275547 B)
	XDuration       string  // 5315 ( < 5314.940 s)
	XDurationAff    string  // 89
	XOverallBitRate string  // 3048 ( < 3048426 bps)
	AudioMultiPiste mediaInfoMultiPiste
	TextMultiPiste  mediaInfoMultiPiste
	Cover           bool
}

// mediaInfoVideo : structure Vidéo (Video_struct)
type mediaInfoVideo struct {
	Format        string // AVC
	FormatProfile string // High
	FormatLevel   string // 4.1
	CodecID       string // V_MPEG4/ISO/AVC
	CodecV        string
	Duration      int64  // 5315 ( < 5314.935000000 s)
	DurationAff   int64  // 89
	BitRate       int64  // 2600 ( < 2600000 bps)
	Width         int64  // 1920
	Height        int64  // 1080
	FrameRateMode string // Constant/Variable ( < CFR)
	FrameRate     string // 23.976 ( < 23.976 fps)
	BitDepth      int64  // 8 ( < 8 bits)
	ScanType      string // Progressive
	Language      string // en
	XDuration     string // 5315
	XDurationAff  string // 89
	XBitRate      string // 2600 ( < 2600000 bps)
	XWidth        string // 1920
	XHeight       string // 1080
	XBitDepth     string // 8 ( < 8 bits)
}

// mediaInfoAudio : structure Audio (Audio_struct)
type mediaInfoAudio struct {
	Format           string // AC-3
	CodecID          string // A_AC3
	CodecA           string
	Duration         int64  // 5315 (5314.656000000 s)
	DurationAff      int64  // 89
	BitRateMode      string // Constant/Variable (CBR)
	BitRate          int64  // 448 ( < 448 Kbps)
	Channel          int64  // 6 ( < 6 channels)
	ChannelPositions string // Front: L C R, Side: L R, LFE
	ChannelDetail    mediaInfoChannelDetail
	ChannelAff       string  // 6ch: Surround
	SamplingRate     float64 // 48.0 ( < 48.0 KHz)
	BitDepth         int64   // 16 ( < 16 bits)
	CompressionMode  string  // Lossy
	Language         string  // fr
	XDuration        string  // 5315 (5314.656000000 s)
	XDurationAff     string  // 89
	XBitRate         string  // 448 ( < 448 Kbps)
	XChannel         string  // 6 ( < 6 channels)
	XSamplingRate    string  // 48.0 ( < 48.0 KHz)
	XBitDepth        string  // 16 ( < 16 bits)
}

// mediaInfoChannelDetail : structure Channel (ChannelDetail_struct)
type mediaInfoChannelDetail struct {
	FrontL bool
	FrontC bool
	FrontR bool
	RearL  bool
	RearR  bool
	Sub    bool
}

// mediaInfoText : structure sous-titre (Text_struct)
type mediaInfoText struct {
	Format   string // UTF-8
	CodecID  string // S_TEXT/UTF8
	Language string // fr
}

// mediaInfoMultiPiste : structure sous-titre (MultiPiste_struct)
type mediaInfoMultiPiste struct {
	Format   string // UTF-8 / UTF-8
	Language string // en / fr
	NoFrench bool
}

// init : initialisation du composant
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
	duree := time.Since(start).Seconds()
	tsIO.PrintConsole(fmt.Sprintf("\n\n\nExécutée en %.3fs", duree))
}

// IsMediaFile - determine si le suffixe du fichier correspond à un media (audio ou vidéo)
func IsMediaFile(ext string) bool {
	result := false
	if _, ok := containers[strings.ToLower(ext)]; ok {
		result = true
	}

	return result
}

// GetMediaInfoAudio : récupère les infos du média Audio (Durée, BitRate, SamplingRate)
func GetMediaInfoAudio(fileName string) (string, string, string) {
	var mediainfoCmd string
	mediainfoCmd, err := exec.LookPath("mediainfo")
	if err != nil {
		panic(fmt.Sprint("  could not find path to 'mediainfo': ", err))
	}
	tsIO.PrintConsole("-- found 'mediainfo' command: ", mediainfoCmd)

	out, err := exec.Command(mediainfoCmd, "--Inform=Audio;%Duration/String%,%BitRate/String%,%SamplingRate/String%", fileName).Output()
	if err != nil {
		panic(fmt.Sprint("Command: mediainfo ", err))
	}

	result := strings.Split(strings.Split(string(out), "\n")[0], ",")
	if len(result) == 3 {
		return strings.Replace(result[0], "min", "mn", 1), result[1], result[2]
	} else {
		return "", "", ""
	}
}

// GetMediaInfoData : récupère les infos du média dans MediaInfoXML1404 (données brutes)
func GetMediaInfoData(fileName string) MediaInfoXML {
	var mediainfoCmd string
	mediainfoCmd, err := exec.LookPath("mediainfo")
	if err != nil {
		panic(fmt.Sprint("  could not find path to 'mediainfo': ", err))
	}
	tsIO.PrintConsole("-- found 'mediainfo' command: ", mediainfoCmd)

	out, err := exec.Command(mediainfoCmd, "--Output=XML", fileName).Output()
	if err != nil {
		panic(fmt.Sprint("Command: mediainfo ", err))
	}
	var result MediaInfoXML
	err = xml.Unmarshal(out, &result) //DÉCODAGE
	if err != nil {
		panic(fmt.Sprint("GetMediaInfoData: Unmarshal ", err))
	}

	return result
}

// GetMediaInfo : récupère les infos du média dans MediaInfo
func GetMediaInfo(fileName string) MediaInfo {
	result := GetMediaInfoData(fileName)

	//	fmt.Println(result)
	var mediaInfo MediaInfo
	for _, track := range result.Media.Track {
		switch track.Type {
		case "General":
			var general mediaInfoGeneral
			general.Format = track.Format
			general.FormatVersion = track.FormatVersion
			general.FileSize, general.XFileSize = extractFileSize(track.FileSize)
			general.Duration, general.XDuration = extractDuration(track.Duration)
			general.DurationAff, general.XDurationAff = extractDurationMN(track.Duration)
			general.OverallBitRate, general.XOverallBitRate = extractBitRate(track.OverallBitRate, track.NominalBitRate, track.BitRateNominal)
			general.Cover = false
			if strings.ToLower(track.Cover) == "yes" {
				general.Cover = true
			}
			mediaInfo.General = general
		case "Video":
			var video mediaInfoVideo
			video.Format = track.Format
			video.FormatProfile = track.FormatProfile
			video.FormatLevel = track.FormatLevel
			video.CodecID = track.CodecID
			video.CodecV = getCodecVideo(video.Format, video.FormatProfile, video.FormatLevel, video.CodecID)
			video.Duration, video.XDuration = extractDuration(track.Duration)
			video.DurationAff, video.XDurationAff = extractDurationMN(track.Duration)
			video.BitRate, video.XBitRate = extractBitRate(track.BitRate, track.NominalBitRate, track.BitRateNominal)
			video.Width, video.XWidth = extractSize(track.Width)
			video.Height, video.XHeight = extractSize(track.Height)
			if track.FrameRateMode == "CFR" {
				video.FrameRateMode = "Constant"
			} else {
				video.FrameRateMode = "Variable"
			}
			if track.FrameRate != "" {
				video.FrameRate = transcodeVideoFrameRate(extractFrameRate(track.FrameRate))
			} else if track.OverallBitRate != "" {
				video.FrameRate = transcodeVideoFrameRate(extractFrameRate(track.OverallBitRate))
			}
			video.BitDepth, video.XBitDepth = extractBitDepth(track.BitDepth)
			video.ScanType = track.ScanType
			video.Language = track.Language
			mediaInfo.Video = append(mediaInfo.Video, video)
		case "Audio":
			var audio mediaInfoAudio
			audio.Format = track.Format
			audio.CodecID = track.CodecID
			audio.CodecA = getCodeCodecAudio(audio.Format)
			audio.Duration, audio.XDuration = extractDuration(track.Duration)
			audio.DurationAff, audio.XDurationAff = extractDurationMN(track.Duration)
			audio.BitRateMode = track.BitRateMode
			audio.BitRate, audio.XBitRate = extractBitRate(track.BitRate, track.NominalBitRate, track.BitRateNominal)
			audio.Channel, audio.XChannel = extractChannel(track.Channels)
			audio.ChannelPositions = track.ChannelPositions
			audio.ChannelDetail = getChannelDetail(track.ChannelPositions)
			audio.ChannelAff = getChannelAff(audio.Channel)
			audio.SamplingRate, audio.XSamplingRate = extractSamplingRate(track.SamplingRate)
			audio.BitDepth, audio.XBitDepth = extractBitDepth(track.BitDepth)
			audio.CompressionMode = track.CompressionMode
			audio.Language = track.Language
			mediaInfo.Audio = append(mediaInfo.Audio, audio)
		case "Text":
			var text mediaInfoText
			text.Format = track.Format
			text.CodecID = track.CodecID
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
			if audio.Language != "fr" {
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
			if text.Language != "fr" {
				mediaInfo.General.TextMultiPiste.NoFrench = true
			}
		}
		mediaInfo.General.TextMultiPiste.Format = strings.Join(format, " / ")
		mediaInfo.General.TextMultiPiste.Language = strings.Join(lang, " / ")
	}

	mediaInfo.General.Conteneur = strings.ToLower(filepath.Ext(fileName))

	// if len(mediaInfo.Video) == 0 {
	// 	var video mediaInfoVideo
	// 	video.Format = "?"
	// 	video.FormatProfile = "?"
	// 	video.CodecID = "?"
	// 	video.CodecV = "?"
	// 	video.Duration, video.XDuration = 0, "?"
	// 	video.DurationAff, video.XDurationAff = 0, "?"
	// 	video.BitRate, video.XBitRate = 0, "?"
	// 	video.Width, video.XWidth = 0, "?"
	// 	video.Height, video.XHeight = 0, "?"
	// 	video.FrameRateMode = "?"
	// 	video.BitDepth, video.XBitDepth = 0, "?"
	// 	video.Language = "?"
	// 	mediaInfo.Video = append(mediaInfo.Video, video)
	// }
	// if len(mediaInfo.Audio) == 0 {
	// 	var audio mediaInfoAudio
	// 	audio.Format = "?"
	// 	audio.CodecID = "?"
	// 	audio.CodecA = "?"
	// 	audio.Duration, audio.XDuration = 0, "?"
	// 	audio.DurationAff, audio.XDurationAff = 0, "?"
	// 	audio.BitRateMode = "?"
	// 	audio.BitRate, audio.XBitRate = 0, "?"
	// 	audio.Channel, audio.XChannel = 0, "?"
	// 	audio.ChannelPositions = "?"
	// 	audio.ChannelDetail = getChannelDetail("")
	// 	audio.ChannelAff = "?"
	// 	audio.SamplingRate, audio.XSamplingRate = 0, "?"
	// 	audio.BitDepth, audio.XBitDepth = 0, "?"
	// 	audio.CompressionMode = "?"
	// 	audio.Language = "?"
	// 	mediaInfo.Audio = append(mediaInfo.Audio, audio)
	// }

	return mediaInfo
}

// extractFileSize return size in GiB (2025275547 (B) --> 1.9 (GiB))
func extractFileSize(size string) (float64, string) {
	resultStr := ""
	if size == "" {
		return 0.00, "?"
	}
	result, err := strconv.ParseFloat(size, 64)
	if err != nil {
		return 0, "-X-"
	}

	result /= 1024 * 1024 * 1024
	switch {
	case result < 1:
		resultStr = fmt.Sprintf("%.3f", result+0.0005)
	case result < 10:
		resultStr = fmt.Sprintf("%.2f", result+0.005)
	default:
		resultStr = fmt.Sprintf("%.1f", result+0.05)
	}

	result = math.Round(result*100) / 100
	return result, resultStr
}

// extractSize return size in pixel (1920 pixels --> 1920)
func extractSize(size string) (int64, string) {
	if size == "" {
		return 0, "?"
	}
	result, err := strconv.ParseInt(size, 10, 64)
	if err != nil {
		return 0, "-X-"
	}
	return result, strconv.FormatInt(result, 10)
}

// extractDuration return durée in sec (5314.940 (s) --> 5315 (s))
func extractDuration(duration string) (int64, string) {
	if duration == "" {
		return 0, "?"
	}
	result, err := strconv.ParseFloat(duration, 64)
	if err != nil {
		return 0, "-X-"
	}
	result = math.Round(result)
	return int64(result), strconv.FormatInt(int64(result), 10)
}

// extractDurationMN convertir la durée sec -> h,mn,sec
func extractDurationMN(duration string) (int64, string) {
	if duration == "" {
		return 0, "?"
	}
	secF, err := strconv.ParseFloat(duration, 64)
	if err != nil {
		return 0, "-X-"
	}
	mn := (secF + 30) / 60

	res := ""
	sec := int64(secF + 0.5)
	hrs, sec := sec/3600, sec%3600
	mins, sec := sec/60, sec%60
	if hrs != 0 {
		res += fmt.Sprintf("%dh", hrs)
	}
	if mins != 0 {
		if res != "" {
			res += " "
		}
		res += fmt.Sprintf("%dmn", mins)
	}
	if sec != 0 {
		if res != "" {
			res += " "
		}
		res += fmt.Sprintf("%ds", sec)
	}
	return int64(mn), res
}

// extractBitRate return bitRate en Kbps (3048426 (bps) --> 3048 (Kbps))
func extractBitRate(bitRate string, nominalBitRate, bitRateNominal string) (int64, string) {
	if bitRate == "" && nominalBitRate == "" && bitRateNominal == "" {
		return 0, "?"
	}
	if bitRate == "" {
		bitRate = nominalBitRate
		if bitRate == "" {
			bitRate = bitRateNominal
		}
	}
	result, err := strconv.ParseInt(bitRate, 10, 64)
	if err != nil {
		return 0, "-X-"
	}
	result /= 1000
	return int64(result), strconv.FormatInt(int64(result), 10)
}

// extractFrameRate return frameRate in fps (23.976 (fps) --> 23.976 (fps))
func extractFrameRate(frame string) float64 {
	if frame == "" {
		return 0.0
	}
	val, err := strconv.ParseFloat(frame, 64)
	if err != nil {
		return 0.0
	}
	return val
}

// extractBitDepth return bitDepth in bits (8 (bits) --> 8 (bits))
func extractBitDepth(bitDepth string) (int64, string) {
	if bitDepth == "" {
		return 0, "?"
	}
	result, err := strconv.ParseInt(bitDepth, 10, 64)
	if err != nil {
		return 0, "-X-"
	}
	return result, strconv.FormatInt(result, 10)
}

// extractChannel return nb audio channel (6 --> 6)
func extractChannel(channel string) (int64, string) {
	if channel == "" {
		return 0, "?"
	}
	result, err := strconv.ParseInt(channel, 10, 64)
	if err != nil {
		return 0, "-X-"
	}
	return result, strconv.FormatInt(result, 10)
}

// getChannelDetail // Front: L C R, Side: L R, LFE
func getChannelDetail(channelPositions string) mediaInfoChannelDetail {
	var channelDetail mediaInfoChannelDetail
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

// ### getChannelAff - transcode les canaux audio pour faciliter la lecture
func getChannelAff(channel int64) string {
	var retour string
	switch channel {
	case 1:
		retour = channel1 // "1ch: Mono"
	case 2:
		retour = channel2 // "2ch: Stéréo"
	case 3:
		retour = channel3 // "3ch: Stéréo 2.1"
	case 5:
		retour = channel5 // "5ch: Surround"
	case 6:
		retour = channel6 // "6ch: Surround"
	case 8:
		retour = channel8 // "8ch: Surround +"
	default:
		retour = strconv.FormatInt(channel, 10)
	}
	return retour
}

// extractSamplingRate return sampling rate in Khz  (48000 (Hz) --> 48 (KHz)) et (44100 (Hz) --> 44.1 (KHz))
func extractSamplingRate(rate string) (float64, string) {
	if rate == "" {
		return 0.0, "?"
	}
	result, err := strconv.ParseFloat(rate, 64)
	if err != nil {
		return 0.0, "-X-"
	}
	result /= 1000
	var resultX string
	if result == float64(int64(result)) {
		resultX = strconv.FormatInt(int64(result), 10)
	} else {
		resultX = strconv.FormatFloat(result, 'f', 1, 64)
	}
	return result, resultX
}

// ### getCodecVideo - transcode le codec vidéo pour faciliter la lecture
//
//	format        : 'AVC'
//	formatProfile : 'High'
//	formatLevel   : '4.1'
//	codecID       : 'V_MPEG4/ISO/AVC'
func getCodecVideo(format string, formatProfile string, formatLevel string, codecID string) string {
	if format == "" && formatProfile == "" && formatLevel == "" && codecID == "" {
		return "????"
	}
	if formatLevel != "" && len(formatLevel) == 1 && strings.ToUpper(format) == "AVC" {
		formatLevel += ".0"
	}

	var codecV string

	switch strings.ToUpper(codecID) {
	case "DX50":
		codecV = "DivX 5"
	case "XVID":
		codecV = "XviD"
	default:
		switch strings.ToUpper(format) {
		case "XVID":
			codecV = "XviD"
		case "DIV3":
			codecV = "DivX 3"
		case "DIV4":
			codecV = "DivX 4"
		case "MPEGVIDEO", "MPEG VIDEO": //&& codec == "mpeg-1v" {
			codecV = "MPEG-1"
		case "MPEG-4VISUAL":
			switch strings.ToUpper(codecID) {
			case "MP42":
				codecV = "MPEG-4"
			case "DIVX":
				codecV = "DivX 4"
			case "XVID":
				codecV = "XviD"
			default:
				codecV = "MPEG-4"
			}
		case "MPEG-4":
			codecV = format
		case "AVC":
			codecV = "X264" + " - " + formatLevel
		case "HEVC":
			codecV = "X265"
		case "THEORA":
			codecV = "Theora"
		default:
			codecV = "????"
		}
	}
	return codecV
}

// //### getCodecVideo - transcode le codec vidéo pour faciliter la lecture
// //		Format: 'AVC'
// //		FormatInfo: 'Advanced Video Codec'
// //		FormatProfile: 'High@L4.1'
// //		CodecID: 'V_MPEG4/ISO/AVC'
// //		CodecIDInfo

// func getCodecVideoOLD(format string, formatProfile string, codecID string, codecIDHint string) string {
// 	if format == "" && formatProfile == "" && codecID == "" && codecIDHint == "" {
// 		return "????"
// 	}

// 	var codecV string
// 	if codecIDHint == "divx 3 low" {
// 		return "DivX 3 Low"
// 	}
// 	//-------------------------------
// 	switch strings.ToUpper(codecID) {
// 	case "DX50":
// 		codecV = "DivX 5"
// 	case "XVID":
// 		codecV = "XviD"
// 	default:
// 		switch strings.ToUpper(format) {
// 		case "XVID":
// 			codecV = "XviD"
// 		case "DIV3":
// 			codecV = "DivX 3"
// 		case "DIV4":
// 			codecV = "DivX 4"
// 		case "MPEGVIDEO", "MPEG VIDEO": //&& codec == "mpeg-1v" {
// 			codecV = "MPEG-1"
// 		case "MPEG-4VISUAL":
// 			switch strings.ToUpper(codecID) {
// 			case "MP42":
// 				codecV = "MPEG-4"
// 			case "DIVX":
// 				codecV = "DivX 4"
// 			case "XVID":
// 				codecV = "XviD"
// 			default:
// 				codecV = "MPEG-4"
// 			}
// 		case "MPEG-4":
// 			codecV = format
// 		case "AVC":
// 			codecV = "X264"
// 			mots := strings.Split(formatProfile, "@")
// 			val := mots[1][1:]
// 			if !strings.Contains(val, ".") {
// 				val += ".0"
// 			}
// 			codecV += " - " + val
// 		case "HEVC":
// 			codecV = "X265"
// 		default:
// 			codecV = "????"
// 		}
// 	}
// 	return codecV
// }

//### getCodeCodecAudio - transcode le codec audio pour faciliter la lecture
//		format     : 'AC-3'
//		codecID    : 'A_AC3'

func getCodeCodecAudio(format string) string {
	var codecA string
	if strings.ToUpper(format) == "MPEG AUDIO" {
		codecA = "MP3"
	} else if strings.ToUpper(format) == "VORBIS" {
		codecA = "Vorbis"
	} else {
		codecA = strings.ToUpper(format)
	}
	return codecA
}

// //### getCodeCodecAudioOLD - transcode le codec audio pour faciliter la lecture
// func getCodeCodecAudioOLD(format string, codec string, codecHint string, formatVersion string, formatProfile string) string {
// 	var codecA string
// 	if strings.ToUpper(format) == "MPEG AUDIO" && strings.ToUpper(formatVersion) == "VERSION 1" && strings.ToUpper(formatProfile) == "LAYER 3" {
// 		codecA = "MP3"
// 	} else if strings.ToUpper(format) == "MPEG AUDIO" && strings.ToUpper(formatVersion) == "VERSION 1" && strings.ToUpper(formatProfile) == "LAYER 2" {
// 		codecA = "MP2"
// 	} else if strings.ToUpper(codecHint) == "MP3" || strings.ToUpper(codec) == "MPA1L3" {
// 		codecA = "MP3"
// 	} else if strings.ToUpper(codecHint) == "MP2" || strings.ToUpper(codec) == "MPA1L2" {
// 		codecA = "MP2"
// 	} else if strings.ToUpper(format) == "VORBIS" {
// 		codecA = "Vorbis"
// 	} else {
// 		codecA = strings.ToUpper(format)
// 	}
// 	return codecA
// }

// ### transcodeVideoFrameRate - transcode le framerate vidéo pour faciliter la lecture
func transcodeVideoFrameRate(frameRate float64) string {
	result := "?"
	if frameRate != 0 {
		result = ""
		frameRates := []float64{23.000, 23.976, 24.000, 25.000, 26.000, 29.970, 30.000, 48.000, 50.000, 60.000}
		for _, val := range frameRates {
			if val == frameRate {
				result = strconv.FormatFloat(val, 'f', 3, 64) // 3 décimales
				break
			}
		}
		if result == "" {
			result = strconv.FormatFloat(frameRate, 'f', 0, 64) + ".xxx"
		}
	}
	return result
}
