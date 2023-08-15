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
	case width >= 1900 && height <= 1080:
		return "FHD" // 1920 x 1080
	case width == 1828 && height == 1332:
		return "2K" // 1828 × 1332
	case width == 3840 && height == 2160:
		return "4K" // 3840 x 2160
	case width == 7680 && height == 4320:
		return "8K" // 7680 x 4320
	default:
		return "NA"
	}
}

func GetEncQuality(format string, debitV int, otherPb *string, toReEncode *bool, replace *bool) string {
	encQuality := ""
	switch format {
	case "HD":
		switch {
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
		case debitV < 4000:
			encQuality = "Light"
		case debitV < 5000:
			encQuality = "Good"
		default: // >= 5000 {
			encQuality = "Hight"
			*otherPb = "Débit V."
			*toReEncode = true
		}
	case "4K", "8K":
		*otherPb = "Format"
		*toReEncode = true
	default:
		*replace = true
	}
	return encQuality
}
