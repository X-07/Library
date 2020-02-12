package mediaInfo

import (
	"encoding/xml"
	"fmt"
	"math"
	"os/exec"
	"path/filepath"
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

// structure en retour de l'appel à mediainfo le programme
type MediainfoXml struct {
	XMLName xml.Name `xml:"Mediainfo"`
	Version string   `xml:"version,attr"`
	File    struct {
		Track []struct {
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
			CodecIDHint              string `xml:"Codec_ID_Hint"`
			BitRate                  string `xml:"Bit_rate"`
			NominalBitRate           string `xml:"Nominal_bit_rate"`
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

// type MediainfoFullXml struct {
// 	XMLName xml.Name `xml:"Mediainfo"`
// 	Version string   `xml:"version,attr"`
// 	File    struct {
// 		Track []struct {
// 			Type                          string   `xml:"type,attr"`
// 			Count                         string   `xml:"Count"`
// 			CountOfStreamOfThisKind       string   `xml:"Count_of_stream_of_this_kind"`
// 			KindOfStream                  []string `xml:"Kind_of_stream"`
// 			StreamIdentifier              string   `xml:"Stream_identifier"`
// 			CountOfVideoStreams           string   `xml:"Count_of_video_streams"`
// 			CountOfAudioStreams           string   `xml:"Count_of_audio_streams"`
// 			VideoFormatList               string   `xml:"Video_Format_List"`
// 			VideoFormatWithHintList       string   `xml:"Video_Format_WithHint_List"`
// 			CodecsVideo                   string   `xml:"Codecs_Video"`
// 			VideoLanguageList             string   `xml:"Video_Language_List"`
// 			AudioFormatList               string   `xml:"Audio_Format_List"`
// 			AudioFormatWithHintList       string   `xml:"Audio_Format_WithHint_List"`
// 			AudioCodecs                   string   `xml:"Audio_codecs"`
// 			AudioLanguageList             string   `xml:"Audio_Language_List"`
// 			CompleteName                  string   `xml:"Complete_name"`
// 			FileName                      string   `xml:"File_name"`
// 			FileExtension                 string   `xml:"File_extension"`
// 			Format                        []string `xml:"Format"`
// 			FormatURL                     string   `xml:"Format_Url"`
// 			FormatExtensionsUsuallyUsed   string   `xml:"Format_Extensions_usually_used"`
// 			CommercialName                string   `xml:"Commercial_name"`
// 			FormatVersion                 string   `xml:"Format_version"`
// 			Codec                         []string `xml:"Codec"`
// 			CodecURL                      string   `xml:"Codec_Url"`
// 			CodecExtensionsUsuallyUsed    string   `xml:"Codec_Extensions_usually_used"`
// 			FileSize                      []string `xml:"File_size"`
// 			Duration                      []string `xml:"Duration"`
// 			OverallBitRate                []string `xml:"Overall_bit_rate"`
// 			FrameRate                     []string `xml:"Frame_rate"`
// 			FrameCount                    string   `xml:"Frame_count"`
// 			StreamSize                    []string `xml:"Stream_size"`
// 			ProportionOfThisStream        string   `xml:"Proportion_of_this_stream"`
// 			FileLastModificationDate      string   `xml:"File_last_modification_date"`
// 			FileLastModificationDateLocal string   `xml:"File_last_modification_date__local_"`
// 			WritingApplication            []string `xml:"Writing_application"`
// 			WritingLibrary                []string `xml:"Writing_library"`
// 			StreamOrder                   string   `xml:"StreamOrder"`
// 			ID                            []string `xml:"ID"`
// 			UniqueID                      string   `xml:"Unique_ID"`
// 			FormatInfo                    string   `xml:"Format_Info"`
// 			FormatProfile                 string   `xml:"Format_profile"`
// 			FormatSettings                string   `xml:"Format_settings"`
// 			FormatSettingsCABAC           []string `xml:"Format_settings__CABAC"`
// 			FormatSettingsReFrames        []string `xml:"Format_settings__ReFrames"`
// 			InternetMediaType             string   `xml:"Internet_media_type"`
// 			CodecID                       string   `xml:"Codec_ID"`
// 			CodecIDURL                    string   `xml:"Codec_ID_Url"`
// 			CodecFamily                   string   `xml:"Codec_Family"`
// 			CodecInfo                     string   `xml:"Codec_Info"`
// 			CodecProfile                  string   `xml:"Codec_profile"`
// 			CodecSettings                 string   `xml:"Codec_settings"`
// 			CodecSettingsCABAC            string   `xml:"Codec_settings__CABAC"`
// 			CodecSettingsRefFrames        string   `xml:"Codec_Settings_RefFrames"`
// 			BitRate                       []string `xml:"Bit_rate"`
// 			Width                         []string `xml:"Width"`
// 			Height                        []string `xml:"Height"`
// 			PixelAspectRatio              string   `xml:"Pixel_aspect_ratio"`
// 			DisplayAspectRatio            []string `xml:"Display_aspect_ratio"`
// 			FrameRateMode                 []string `xml:"Frame_rate_mode"`
// 			Resolution                    []string `xml:"Resolution"`
// 			Colorimetry                   string   `xml:"Colorimetry"`
// 			ColorSpace                    string   `xml:"Color_space"`
// 			ChromaSubsampling             string   `xml:"Chroma_subsampling"`
// 			BitDepth                      []string `xml:"Bit_depth"`
// 			ScanType                      []string `xml:"Scan_type"`
// 			Interlacement                 []string `xml:"Interlacement"`
// 			BitsPixelFrame                string   `xml:"Bits__Pixel_Frame_"`
// 			Delay                         []string `xml:"Delay"`
// 			DelayOrigin                   []string `xml:"Delay__origin"`
// 			EncodedLibraryName            string   `xml:"Encoded_Library_Name"`
// 			EncodedLibraryVersion         string   `xml:"Encoded_Library_Version"`
// 			EncodingSettings              string   `xml:"Encoding_settings"`
// 			Language                      []string `xml:"Language"`
// 			Default                       []string `xml:"Default"`
// 			Forced                        []string `xml:"Forced"`
// 			ColorRange                    string   `xml:"Color_range"`
// 			ColourDescriptionPresent      string   `xml:"colour_description_present"`
// 			MatrixCoefficients            string   `xml:"Matrix_coefficients"`
// 			ModeExtension                 string   `xml:"Mode_extension"`
// 			FormatSettingsEndianness      string   `xml:"Format_settings__Endianness"`
// 			BitRateMode                   []string `xml:"Bit_rate_mode"`
// 			ChannelS                      []string `xml:"Channel_s_"`
// 			ChannelPositions              []string `xml:"Channel_positions"`
// 			ChannelLayout                 string   `xml:"ChannelLayout"`
// 			SamplingRate                  []string `xml:"Sampling_rate"`
// 			SamplesCount                  string   `xml:"Samples_count"`
// 			CompressionMode               []string `xml:"Compression_mode"`
// 			DelayRelativeToVideo          []string `xml:"Delay_relative_to_video"`
// 			Video0Delay                   []string `xml:"Video0_delay"`
// 			Bsid                          string   `xml:"bsid"`
// 			Dialnorm                      string   `xml:"dialnorm"`
// 			DialnormString                string   `xml:"dialnorm_String"`
// 			Compr                         string   `xml:"compr"`
// 			ComprString                   string   `xml:"compr_String"`
// 			Acmod                         string   `xml:"acmod"`
// 			Lfeon                         string   `xml:"lfeon"`
// 			DialnormAverage               string   `xml:"dialnorm_Average"`
// 			DialnormAverageString         string   `xml:"dialnorm_Average_String"`
// 			DialnormMinimum               string   `xml:"dialnorm_Minimum"`
// 			DialnormMinimumString         string   `xml:"dialnorm_Minimum_String"`
// 			DialnormMaximum               string   `xml:"dialnorm_Maximum"`
// 			DialnormMaximumString         string   `xml:"dialnorm_Maximum_String"`
// 			DialnormCount                 string   `xml:"dialnorm_Count"`
// 			ComprAverage                  string   `xml:"compr_Average"`
// 			ComprAverageString            string   `xml:"compr_Average_String"`
// 			ComprMinimum                  string   `xml:"compr_Minimum"`
// 			ComprMinimumString            string   `xml:"compr_Minimum_String"`
// 			ComprMaximum                  string   `xml:"compr_Maximum"`
// 			ComprMaximumString            string   `xml:"compr_Maximum_String"`
// 			ComprCount                    string   `xml:"compr_Count"`
// 		} `xml:"track"`
// 	} `xml:"File"`
// }

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

// // GetMediaInfo() : récupère les infos du média dans MediainfoXml (données brutes)
// func GetMediaInfoFullData(fileName string) MediainfoFullXml {
// 	var mediainfo_cmd string
// 	mediainfo_cmd, err := exec.LookPath("mediainfo")
// 	if err != nil {
// 		panic(fmt.Sprint("  could not find path to 'mediainfo': ", err))
// 	}
// 	tsIO.PrintConsole("-- found 'mediainfo' command: ", mediainfo_cmd)

// 	out, err := exec.Command(mediainfo_cmd, "-f --Output=XML", fileName).Output()
// 	if err != nil {
// 		panic(fmt.Sprint("Command: mediainfo ", err))
// 	}
// 	var result MediainfoFullXml
// 	err = xml.Unmarshal(out, &result) //DECODAGE
// 	if err != nil {
// 		panic(fmt.Sprint("GetMediaInfoData: Unmarshal ", err))
// 	}

// 	return result
// }

// GetMediaInfo() : récupère les infos du média dans MediainfoXml (données brutes)
func GetMediaInfoData(fileName string) MediainfoXml {
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
	var result MediainfoXml
	err = xml.Unmarshal(out, &result) //DECODAGE
	if err != nil {
		panic(fmt.Sprint("GetMediaInfoData: Unmarshal ", err))
	}

	return result
}

// GetMediaInfo() : récupère les infos du média dans MediaInfo_struct
func GetMediaInfo(fileName string) MediaInfo_struct {
	result := GetMediaInfoData(fileName)

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
			general.DurationAff, general.XDurationAff = extractDurationMN(general.Duration)
			general.OverallBitRate, general.XOverallBitRate = extractBitRate(track.OverallBitRate, track.NominalBitRate)
			mediaInfo.General = general
		case "Video":
			var video Video_struct
			video.Format = track.Format
			video.FormatInfo = track.FormatInfo
			video.FormatProfile = track.FormatProfile
			video.CodecID = track.CodecID
			video.CodecIDInfo = track.CodecIDInfo
			video.CodecV = getCodecVideo(video.Format, video.FormatProfile, video.CodecID, track.CodecIDHint)
			video.Duration, video.XDuration = extractDuration(track.Duration)
			video.DurationAff, video.XDurationAff = extractDurationMN(video.Duration)
			video.BitRate, video.XBitRate = extractBitRate(track.BitRate, track.NominalBitRate)
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
			audio.CodecA = getCodeCodecAudio(audio.Format, audio.CodecID, track.CodecIDHint, track.FormatVersion, track.FormatProfile)
			audio.Duration, audio.XDuration = extractDuration(track.Duration)
			audio.DurationAff, audio.XDurationAff = extractDurationMN(audio.Duration)
			audio.BitRateMode = track.BitRateMode
			audio.BitRate, audio.XBitRate = extractBitRate(track.BitRate, track.NominalBitRate)
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

	if len(mediaInfo.Video) == 0 {
		var video Video_struct
		video.Format = "?"
		video.FormatInfo = "?"
		video.FormatProfile = "?"
		video.CodecID = "?"
		video.CodecIDInfo = "?"
		video.CodecV = "?"
		video.Duration, video.XDuration = 0, "?"
		video.DurationAff, video.XDurationAff = 0, "?"
		video.BitRate, video.XBitRate = 0, "?"
		video.Width, video.XWidth = 0, "?"
		video.Height, video.XHeight = 0, "?"
		video.FrameRateMode = "?"
		video.BitDepth, video.XBitDepth = 0, "?"
		video.Language = "?"
		mediaInfo.Video = append(mediaInfo.Video, video)
	}
	if len(mediaInfo.Audio) == 0 {
		var audio Audio_struct
		audio.Format = "?"
		audio.FormatInfo = "?"
		audio.CodecID = "?"
		audio.CodecIDInfo = "?"
		audio.CodecA = "?"
		audio.Duration, audio.XDuration = 0, "?"
		audio.DurationAff, audio.XDurationAff = 0, "?"
		audio.BitRateMode = "?"
		audio.BitRate, audio.XBitRate = 0, "?"
		audio.Channel, audio.XChannel = 0, "?"
		audio.ChannelPositions = "?"
		audio.ChannelDetail = getChannelDetail("")
		audio.ChannelAff = "?"
		audio.SamplingRate, audio.XSamplingRate = 0, "?"
		audio.BitDepth, audio.XBitDepth = 0, "?"
		audio.CompressionMode = "?"
		audio.Language = "?"
		mediaInfo.Audio = append(mediaInfo.Audio, audio)
	}

	return mediaInfo
}

// extractFileSize() return size in GiB (1.43 GiB --> 1.43  ou  785 MiB --> 0.766)
func extractFileSize(size string) (float64, string) {
	if size == "" {
		return 0.00, "?"
	} else {
		mots := strings.Fields(size)
		val, err := strconv.ParseFloat(mots[0], 64)
		if err != nil {
			return 0.00, "-X-"
		}
		if mots[1] == "MiB" {
			val /= 1024
		}
		if val < 0.1 {
			val = math.RoundToEven(val*100) / 100
			return val, strconv.FormatFloat(val, 'f', 2, 64)
		} else {
			val = math.RoundToEven(val*10) / 10
			return val, strconv.FormatFloat(val, 'f', 1, 64)
		}
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
		result, err := strconv.ParseFloat(tmp, 64)
		if err != nil {
			return 0, "-X-"
		}
		return int64(result), strconv.FormatInt(int64(result), 10)
	}
}

// extractDuration() return durée in sec (40mn 13s --> 2413 (int & string))
func extractDuration(duration string) (int64, string) {
	if duration == "" {
		return 0, "?"
	} else {
		duration := strings.ReplaceAll(strings.Join(strings.Fields(duration), ""), "mn", "m") // 40mn 13s  ==>  40m13s
		duree, err := time.ParseDuration(duration)
		if err != nil {
			return 0, "-X-"
		}
		result := int64(duree.Seconds())

		return result, strconv.FormatInt(result, 10)
	}
}

// extractDurationMN() convertir la durée sec -> mn
func extractDurationMN(duration int64) (int64, string) {
	if duration == 0 {
		return 0, "?"
	} else {
		result := (duration + 30) / 60
		return result, strconv.FormatInt(result, 10)
	}
}

// extractBitRate() return bitRate en Kbps (5 098 Kbps  --> 5098)
func extractBitRate(bitRate string, nominalBitRate string) (int64, string) {
	if bitRate == "" && nominalBitRate == "" {
		return 0, "?"
	} else {
		if bitRate == "" {
			bitRate = nominalBitRate
		}
		mots := strings.Fields(bitRate)
		var tmp string
		for _, val := range mots[:len(mots)-1] {
			if val == "Kbps" || val == "Mbps" {
				break
			}
			if _, err := strconv.ParseFloat(val, 64); err == nil {
				tmp += val
			}
		}
		result, err := strconv.ParseFloat(tmp, 64)
		if err != nil {
			return 0, "-X-"
		}
		if mots[len(mots)-1] == "Mbps" {
			result *= 1024
		}
		return int64(result), strconv.FormatInt(int64(result), 10)
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
			return 0.0
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
		val, err := strconv.ParseFloat(mots[0], 64)
		if err != nil {
			return 0, "-X-"
		}
		return int64(val), strconv.FormatInt(int64(val), 10)
	}
}

// extractChannel() return nb audio channel (6 channels --> 6)
func extractChannel(channel string) (int64, string) {
	if channel == "" {
		return 0, "?"
	} else {
		mots := strings.Fields(channel)
		val, err := strconv.ParseFloat(mots[0], 64)
		if err != nil {
			return 0, "-X-"
		}
		return int64(val), strconv.FormatInt(int64(val), 10)
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
		var valX string
		mots := strings.Fields(rate)
		val, err := strconv.ParseFloat(mots[0], 64)
		if err != nil {
			return 0.0, "-X-"
		}
		if strings.HasSuffix(mots[0], ".0") {
			valX = strconv.FormatInt(int64(val), 10)
		} else {
			valX = strconv.FormatFloat(val, 'f', 1, 64)
		}
		return val, valX
	}
}

//### getCodecVideo() - transcode le codec vidéo pour faciliter la lecture
//		Format: 'AVC'
//		FormatInfo: 'Advanced Video Codec'
//		FormatProfile: 'High@L4.1'
//		CodecID: 'V_MPEG4/ISO/AVC'
//		CodecIDInfo

func getCodecVideo(format string, formatProfile string, codecID string, codecIDHint string) string {
	if format == "" && formatProfile == "" && codecID == "" && codecIDHint == "" {
		return "????"
	}

	var codecV string
	if strings.ToUpper(codecIDHint) == "DIVX 3 LOW" {
		return "DivX 3 Low"
	}
	//-------------------------------
	switch strings.ToUpper(codecID) {
	case "DX50":
		codecV = "DivX 5"
	case "XVID":
		codecV = "XviD"
	case "DIV3":
		codecV = "DivX 3"
		if strings.ToUpper(codecIDHint) == "DIVX 3 LOW" {
			codecV = "DivX 3 Low"
		}
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
		case "MPEG-4 VISUAL":
			switch strings.ToUpper(codecID) {
			case "MP42":
				codecV = "MPEG-4"
			case "DIVX":
				codecV = "DivX 4"
			case "XVID", "V_MS/VFW/FOURCC / XVID":
				codecV = "XviD"
			case "V_MS/VFW/FOURCC / DX50":
				codecV = "DivX 5"
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
		case "HEVC":
			codecV = "X265"
		default:
			codecV = "????"
		}
	}
	return codecV
}

//### getCodeCodecAudio() - transcode le codec audio pour faciliter la lecture
func getCodeCodecAudio(format string, codec string, codecHint string, formatVersion string, formatProfile string) string {
	var codecA string
	if strings.ToUpper(format) == "MPEG AUDIO" && strings.ToUpper(formatVersion) == "VERSION 1" && strings.ToUpper(formatProfile) == "LAYER 3" {
		codecA = "MP3"
	} else if strings.ToUpper(format) == "MPEG AUDIO" && strings.ToUpper(formatVersion) == "VERSION 1" && strings.ToUpper(formatProfile) == "LAYER 2" {
		codecA = "MP2"
	} else if strings.ToUpper(codecHint) == "MP3" || strings.ToUpper(codec) == "MPA1L3" {
		codecA = "MP3"
	} else if strings.ToUpper(codecHint) == "MP2" || strings.ToUpper(codec) == "MPA1L2" {
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
