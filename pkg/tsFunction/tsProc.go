package tsIO

import "strconv"

// MiseEnFormeByte(nb octets)
func MiseEnFormeByte(bps int) string {
	var unit string
	var speed float64
	var xSpeed string
	// strconv.FormatFloat(f, 'f', 5, 64) - 'f' format - 5 is the number of decimals - 64 is for float64 type

	if bps >= 1073741824 {
		unit = "G"
		speed = float64(bps) / 1073741824
		if bps >= 107374182400 {
			xSpeed = strconv.FormatFloat(speed, 'f', 0, 64) // 0 decimale
		} else if bps >= 10737418240 {
			xSpeed = strconv.FormatFloat(speed, 'f', 1, 64) // 1 decimale
		} else {
			xSpeed = strconv.FormatFloat(speed, 'f', 2, 64) // 2 decimale
		}
	} else if bps >= 1048576 {
		unit = "M"
		speed = float64(bps) / 1048576
		if bps >= 104857600 {
			xSpeed = strconv.FormatFloat(speed, 'f', 0, 64) // 0 decimale
		} else if bps >= 10485760 {
			xSpeed = strconv.FormatFloat(speed, 'f', 1, 64) // 1 decimale
		} else {
			xSpeed = strconv.FormatFloat(speed, 'f', 2, 64) // 2 decimale
		}
	} else if bps >= 1024 {
		unit = "K"
		speed = float64(bps) / 1024
		if bps >= 102400 {
			xSpeed = strconv.FormatFloat(speed, 'f', 1, 64) // 1 decimale
		} else if bps >= 10240 {
			xSpeed = strconv.FormatFloat(speed, 'f', 1, 64) // 1 decimale
		} else {
			xSpeed = strconv.FormatFloat(speed, 'f', 2, 64) // 2 decimale
		}
	} else {
		unit = "o"
		xSpeed = strconv.Itoa(bps)
	}

	return xSpeed + unit
}

// MiseEnFormeByte(nb bits)
func MiseEnFormeBit(bps int) string {
	var unit string
	var speed float64
	var xSpeed string
	// strconv.FormatFloat(f, 'f', 5, 64) - 'f' format - 5 is the number of decimals - 64 is for float64 type

	if bps >= 1000000000 {
		unit = "Gbit/s"
		speed = float64(bps) / 1000000000
		if bps >= 100000000000 {
			xSpeed = strconv.FormatFloat(speed, 'f', 0, 64) // 0 decimale
		} else if bps >= 10000000000 {
			xSpeed = strconv.FormatFloat(speed, 'f', 1, 64) // 1 decimale
		} else {
			xSpeed = strconv.FormatFloat(speed, 'f', 2, 64) // 2 decimale
		}
	} else if bps >= 1000000 {
		unit = "Mbit/s"
		speed = float64(bps) / 1000000
		if bps >= 100000000 {
			xSpeed = strconv.FormatFloat(speed, 'f', 0, 64) // 0 decimale
		} else if bps >= 10000000 {
			xSpeed = strconv.FormatFloat(speed, 'f', 1, 64) // 1 decimale
		} else {
			xSpeed = strconv.FormatFloat(speed, 'f', 2, 64) // 2 decimale
		}
	} else if bps >= 1000 {
		unit = "Kbit/s"
		speed = float64(bps) / 1000
		if bps >= 100000 {
			xSpeed = strconv.FormatFloat(speed, 'f', 0, 64) // 0 decimale
		} else if bps >= 10000 {
			xSpeed = strconv.FormatFloat(speed, 'f', 1, 64) // 1 decimale
		} else {
			xSpeed = strconv.FormatFloat(speed, 'f', 2, 64) // 2 decimale
		}
	} else {
		unit = "bit/s"
		xSpeed = strconv.Itoa(bps)
	}

	return xSpeed + unit
}
