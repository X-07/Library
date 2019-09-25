package tsFunction

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	tsIO "tsFunction"
)

var TraceConsole *bool //trace sur la console

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

type MediaInfo_struct struct {
	General General_struct
	Video 	[]Video_struct
	Audio 	[]Audio_struct
	Text 	[]Text_struct
}

type General_struct {
	Format                   string		// MPEG-4
	FormatVersion            string		// Version 2
	FileSize                 int64		// 1.43 ( < 1.43 GiB)
	Duration                 string		// 2413 (en sec < 40mn 13s)
	OverallBitRate           int64		// 5098 ( < 5 098 Kbps)
}

type Video_struct {
	Format                   string		// AVC
	FormatInfo               string		// Advanced Video Codec
	FormatProfile            string		// High@L4.0
	CodecID                  string		// V_MPEG4/ISO/AVC
	Duration                 int64		// 2413 (en sec < 40mn 13s)
	BitRate                  int64		// 4613 ( < 4 613 Kbps)
	Width                    int64 		// 1920 ( < 1 920 pixels)
	Height                   int64	 	// 1080 ( < 1 080 pixels)
	FrameRateMode            string 	// Constant/Variable
	FrameRate                float64	// 23.976 ( < 23.976 fps)
	BitDepth                 int 64 	// 8 ( < 8bits)
	Language                 string 	// English
}

type Audio_struct {
	Format                   string		// AC-3
	FormatInfo               string		// Audio Coding 3
	CodecID                  string		// A_AC3
	Duration                 int64		// 2413 (en sec < 40mn 13s)
	BitRateMode              string 	// Constant/Variable
	BitRate                  int64		// 384 ( < 384 Kbps)
	ChannelS                 int64	 	// 6 ( < 6 channels)
	ChannelPositions         string 	// Front: L C R, Side: L R, LFE
	SamplingRate             float64 	// 48.0 ( < 48.0 KHz)
	BitDepth                 int64 		// 16 ( < 16 bits)
	CompressionMode          string 	// Lossy
	Language                 string 	// English
}

type Text_struct {
	Format                   string		// UTF-8
	CodecID                  string		// S_TEXT/UTF8
	CodecIDInfo              string		// UTF-8 Plain Text
	Language                 string 	// English
}

//
func getMediaInfo(fileName string) bool {
	var mediainfo_cmd string
	mediainfo_cmd, err = exec.LookPath("mediainfo")
	if err != nil {
		panic(fmt.Sprint("  could not find path to 'mediainfo': ", err))
	}
	tsIO.PrintConsole("-- found 'mediainfo' command: ", mediainfo_cmd)

	out, err := exec.Command(mediainfo_cmd, "--Output=XML", fileName).Output()
	if err != nil {
		panic(fmt.Sprint("Command: mediainfo", err))
	}
	var result MediainfoXmlXml
	err := xml.Unmarshal(out, &result) //DECODAGE

}
