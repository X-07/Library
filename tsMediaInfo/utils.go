package tsMediaInfo

func GetFormat(width, height int) string {
	switch {
	case width == 0 && height == 0:
		return ""
	case height <= 288:
		return "VCD" // 352 x 288 or 352 x 240
	case width < 640 || height < 380:
		return "VHS" // 310 x 576 or 320 x 480
	case width < 900 && height <= 576:
		return "DVD" // 720 x 576 or 720 x 480
	case width < 1900 && height <= 768:
		return "HD" // 1280 x 720 or 1366 x 768
	case width < 2500 && height <= 1080:
		return "FHD" // 1920 x 1080
	case width < 3800 && height <= 1440:
		return "2K" // 2560 x 1440
	case width < 7600 && height <= 2160:
		return "4K" // 3840 x 2160
	case width < 15200 && height <= 4320:
		return "8K" // 7680 x 4320
	default:
		return "NA"
	}
}

func GetEncQuality(format string, debitV int, otherPb *string, toReEncode *bool, replace *bool) string {
	encQuality := ""
	switch format {
	case "NA":
		encQuality = "NA"
	case "HD":
		switch {
		case debitV == 0:
			encQuality = "NA"
		case debitV < 1300:
			encQuality = "Bad"
		case debitV < 2000:
			encQuality = "Light"
		case debitV < 3000:
			encQuality = "Good"
		default: // >= 3000
			encQuality = "Hight"
			*otherPb = "Débit V."
			*toReEncode = true
		}
	case "FHD":
		switch {
		case debitV == 0:
			encQuality = "NA"
		case debitV < 2500:
			encQuality = "Bad"
		case debitV < 4000:
			encQuality = "Light"
		case debitV < 5000:
			encQuality = "Good"
		default: // >= 5000 {
			encQuality = "Hight"
			*otherPb = "Débit V."
			*toReEncode = true
		}
	case "2K", "4K", "8K":
		*otherPb = "Format"
		*toReEncode = true
	default:
		*replace = true
	}
	return encQuality
}
