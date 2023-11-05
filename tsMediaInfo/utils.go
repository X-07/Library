package tsMediaInfo

func GetFormat(width, height int) string {
	switch {
	case width == 0 && height == 0:
		return "NA"
	case width <= 352 && height <= 288:
		return "VCD" // 352 x 288 or 352 x 240
	case width <= 320 && height <= 576:
		return "VHS" // 310 x 576 or 320 x 480
	case width <= 720 && height <= 576:
		return "DVD" // 720 x 576 or 720 x 480
	case width <= 1366 && height <= 768:
		return "HD" // 1280 x 720 or 1366 x 768
	case width <= 1920 && height <= 1080:
		return "FHD" // 1920 x 1080
	case width <= 2560 && height <= 1440:
		return "2K" // 2560 x 1440
	case width <= 3840 && height <= 2160:
		return "4K" // 3840 x 2160
	case width <= 7680 && height <= 4320:
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
