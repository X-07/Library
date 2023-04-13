package tsFunction

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type fileStruct struct {
	CpuN      string
	ChargeIo  int64
	ChargeUsr int64
	ChargeIrq int64
	ChargeSys int64
	Total     int64
}

// nm-tool
//
// NetworkManager Tool
//
// State: connected (global)
//
// - Device: eth1  [Connexion filaire 1] ------------------------------------------
//   Type:              Wired
//   Driver:            r8169
//   State:             connected
//   Default:           yes
//   HW Address:        00:25:22:A7:FF:C8
//
//   Capabilities:
//     Carrier Detect:  yes
//     Speed:           1000 Mb/s
//
//   Wired Properties
//     Carrier:         on
//
//   IPv4 Settings:
//     Address:         192.168.1.20
//     Prefix:          24 (255.255.255.0)
//     Gateway:         192.168.1.1
//
//     DNS:             192.168.1.1
//
//
// - Device: eth0 -----------------------------------------------------------------
//   Type:              Wired
//   Driver:            r8169
//   State:             unavailable
//   Default:           no
//   HW Address:        00:25:22:A7:FF:C7
//
//   Capabilities:
//     Carrier Detect:  yes
//
//   Wired Properties
//     Carrier

// GetConnexion : recherche du device de connexion Ethernet (ex: eth0)
func GetConnexion() string {
	var connect string
	// execution de la commande shell "nm-tool"
	cmd := exec.Command("nmcli")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		CallNotifySend([]string{">>> FATAL ERROR <<<", "tsSys.GetConnexion()", "-t", "20000", "-i", "/usr/share/icons/gnome/32x32/status/dialog-error.png"})
		panic(fmt.Sprint("tsSys - getConnexion: ", err))
	}
	// recherche du device sur lequel on est connecté et extraction de son nom
	var re = regexp.MustCompile(`(.*): connect. .`)
	matches := re.FindStringSubmatch(out.String())
	if len(matches) == 2 {
		connect = matches[1]
		PrintConsole("device : " + matches[1])
	} else {
		connect = ""
	}

	return connect
}

// ReadStatsDown (device) : lecture du nombre d'octets reçus
func ReadStatsDown(connect string) int64 {
	var rxBytes64 int64
	// lecture du fichier "/sys/class/net/$interface/statistics/rx_bytes" par le cat du bash !
	// 2114086
	cmd := exec.Command("cat", "/sys/class/net/"+connect+"/statistics/rx_bytes")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		CallNotifySend([]string{">>> FATAL ERROR <<<", "tsSys.ReadStatsDown()", "-t", "20000", "-i", "/usr/share/icons/gnome/32x32/status/dialog-error.png"})
		panic(fmt.Sprint("tsSys - readStatsDown > 'cat': ", err))
	}
	m := strings.Split(out.String(), "\n")
	rxBytes := m[0]
	rxBytes64, err = strconv.ParseInt(rxBytes, 10, 64)
	if err != nil {
		CallNotifySend([]string{">>> FATAL ERROR <<<", "tsSys.ReadStatsDown()", "-t", "20000", "-i", "/usr/share/icons/gnome/32x32/status/dialog-error.png"})
		panic(fmt.Sprint("tsSys - readStatsDown > parseInt: ", err))
	}

	return rxBytes64
}

// ReadStatsUp (device) : lecture du nombre d'octets émis
func ReadStatsUp(connect string) int64 {
	var txBytes64 int64
	// lecture du fichier "/sys/class'/net/$interface/statistics/tx_bytes" par le cat du bash !
	// 770240
	cmd := exec.Command("cat", "/sys/class/net/"+connect+"/statistics/tx_bytes")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		CallNotifySend([]string{">>> FATAL ERROR <<<", "tsSys.ReadStatsUp()", "-t", "20000", "-i", "/usr/share/icons/gnome/32x32/status/dialog-error.png"})
		panic(fmt.Sprint("tsSys - readStatsUp > 'cat': ", err))
	}
	m := strings.Split(out.String(), "\n")
	txBytes := m[0]
	txBytes64, err = strconv.ParseInt(txBytes, 10, 64)
	if err != nil {
		CallNotifySend([]string{">>> FATAL ERROR <<<", "tsSys.ReadStatsUp()", "-t", "20000", "-i", "/usr/share/icons/gnome/32x32/status/dialog-error.png"})
		panic(fmt.Sprint("tsSys - readStatsUp > parseInt: ", err))
	}

	return txBytes64
}

// GetDataCPU : recherche les infos du CPU
func GetDataCPU(core string) fileStruct {
	var file fileStruct
	enr, err := ReadFileForValue("/proc/stat", "cpu"+core)
	if err != nil {
		CallNotifySend([]string{">>> FATAL ERROR <<<", "tsSys.GetDataCPU()", "-t", "20000", "-i", "/usr/share/icons/gnome/32x32/status/dialog-error.png"})
		panic(fmt.Sprint("tsSys - getDataCPU > ReadFileForValue: ", err))
	}
	mots := strings.Fields(enr)
	file.CpuN = mots[0]
	var user, nice, system, idle, irq, softirq int64
	user, err = strconv.ParseInt(mots[1], 10, 64)
	if err != nil {
		CallNotifySend([]string{">>> FATAL ERROR <<<", "tsSys.GetDataCPU()", "-t", "20000", "-i", "/usr/share/icons/gnome/32x32/status/dialog-error.png"})
		panic(fmt.Sprint("tsSys - getDataCPU > parseInt: ", err))
	}
	nice, err = strconv.ParseInt(mots[2], 10, 64)
	if err != nil {
		CallNotifySend([]string{">>> FATAL ERROR <<<", "tsSys.GetDataCPU()", "-t", "20000", "-i", "/usr/share/icons/gnome/32x32/status/dialog-error.png"})
		panic(fmt.Sprint("tsSys - getDataCPU > parseInt: ", err))
	}
	system, err = strconv.ParseInt(mots[3], 10, 64)
	if err != nil {
		CallNotifySend([]string{">>> FATAL ERROR <<<", "tsSys.GetDataCPU()", "-t", "20000", "-i", "/usr/share/icons/gnome/32x32/status/dialog-error.png"})
		panic(fmt.Sprint("tsSys - getDataCPU > parseInt: ", err))
	}
	idle, err = strconv.ParseInt(mots[4], 10, 64)
	if err != nil {
		CallNotifySend([]string{">>> FATAL ERROR <<<", "tsSys.GetDataCPU()", "-t", "20000", "-i", "/usr/share/icons/gnome/32x32/status/dialog-error.png"})
		panic(fmt.Sprint("tsSys - getDataCPU > parseInt: ", err))
	}
	file.ChargeIo, err = strconv.ParseInt(mots[5], 10, 64)
	if err != nil {
		CallNotifySend([]string{">>> FATAL ERROR <<<", "tsSys.GetDataCPU()", "-t", "20000", "-i", "/usr/share/icons/gnome/32x32/status/dialog-error.png"})
		panic(fmt.Sprint("tsSys - getDataCPU > parseInt: ", err))
	}
	irq, err = strconv.ParseInt(mots[6], 10, 64)
	if err != nil {
		CallNotifySend([]string{">>> FATAL ERROR <<<", "tsSys.GetDataCPU()", "-t", "20000", "-i", "/usr/share/icons/gnome/32x32/status/dialog-error.png"})
		panic(fmt.Sprint("tsSys - getDataCPU > parseInt: ", err))
	}
	softirq, err = strconv.ParseInt(mots[7], 10, 64)
	if err != nil {
		CallNotifySend([]string{">>> FATAL ERROR <<<", "tsSys.GetDataCPU()", "-t", "20000", "-i", "/usr/share/icons/gnome/32x32/status/dialog-error.png"})
		panic(fmt.Sprint("tsSys - getDataCPU > parseInt: ", err))
	}
	file.ChargeUsr = user + nice
	file.ChargeIrq = irq + softirq
	file.ChargeSys = system + file.ChargeIrq

	file.Total = file.ChargeUsr + file.ChargeSys + idle + file.ChargeIo

	return file
}

// GetDataAllDisk : recherche des données de taux de transfert de tous les disques
func GetDataAllDisk() (float64, float64) {
	var totalLec, totalEcr float64
	//Calcul du nb de HDD/SDD connectés
	hddAll := []string{}
	listDisk, _ := filepath.Glob("/sys/block/sd*")
	for _, device := range listDisk {
		input, err := os.ReadFile(device + "/size")
		if err != nil {
			CallNotifySend([]string{">>> FATAL ERROR <<<", "tsSys.GetDataAllDisk()", "-t", "20000", "-i", "/usr/share/icons/gnome/32x32/status/dialog-error.png"})
			panic(fmt.Sprint("tsSys - getDataAllDisk > ReadFile: ", err))
		}
		taille := strings.Split(string(input), "\n")[0]
		if taille != "0" {
			disk := strings.Split(device, "/")[3]
			hddAll = append(hddAll, disk)
		}
	}
	PrintConsole("hddAll = ", hddAll)

	//Récupérer la taille du secteur
	sectorSize := []int64{}
	for _, hdd := range hddAll {
		sector, err := os.ReadFile("/sys/block/" + hdd + "/queue/hw_sector_size")
		if err == nil {
			var val int64
			val, err = strconv.ParseInt(strings.Fields(string(sector))[0], 10, 64)
			if err != nil {
				CallNotifySend([]string{">>> FATAL ERROR <<<", "tsSys.GetDataAllDisk()", "-t", "20000", "-i", "/usr/share/icons/gnome/32x32/status/dialog-error.png"})
				panic(fmt.Sprint("tsSys - getDataAllDisk > ReadFile: ", err))
			}
			sectorSize = append(sectorSize, val)
		} else {
			sectorSize = append(sectorSize, 512)
		}
	}
	PrintConsole("sectorSize = ", sectorSize)

	for idx, hdd := range hddAll {
		var lec, ecr int64
		input, err := os.ReadFile("/sys/block/" + hdd + "/stat")
		if err != nil {
			CallNotifySend([]string{">>> FATAL ERROR <<<", "tsSys.GetDataAllDisk()", "-t", "20000", "-i", "/usr/share/icons/gnome/32x32/status/dialog-error.png"})
			panic(fmt.Sprint("tsSys - getDataAllDisk > ReadFile: ", err))
		}
		mots := strings.Fields(string(input))
		lec, err = strconv.ParseInt(mots[3], 10, 64)
		if err != nil {
			CallNotifySend([]string{">>> FATAL ERROR <<<", "tsSys.GetDataAllDisk()", "-t", "20000", "-i", "/usr/share/icons/gnome/32x32/status/dialog-error.png"})
			panic(fmt.Sprint("tsSys - getDataAllDisk > ParseInt: ", err))
		}
		ecr, err = strconv.ParseInt(mots[7], 10, 64)
		if err != nil {
			CallNotifySend([]string{">>> FATAL ERROR <<<", "tsSys.GetDataAllDisk()", "-t", "20000", "-i", "/usr/share/icons/gnome/32x32/status/dialog-error.png"})
			panic(fmt.Sprint("tsSys - getDataAllDisk > ParseInt: ", err))
		}

		totalLec += float64(lec*sectorSize[idx]/(1024*1024)) / 60
		totalEcr += float64(ecr*sectorSize[idx]/(1024*1024)) / 60
	}

	return totalLec, totalEcr
}

// GetNbProcess retourne la liste des PID correspondant au nom d'un processus
func GetNbProcess(processName string) int64 {
	out, err := exec.Command("pgrep", "-c", processName).Output()
	if err != nil {
		if fmt.Sprint(err) != "exit status 1" {
			CallNotifySend([]string{">>> FATAL ERROR <<<", "tsSys.GetNbProcess()", "-t", "20000", "-i", "/usr/share/icons/gnome/32x32/status/dialog-error.png"})
			panic(fmt.Sprint("tsSys - getProcessPIDs > Command: pgrep ", err))
		}
	}
	var nb int64
	nbx := strings.Split(string(out), "\n")[0]
	nb, err = strconv.ParseInt(nbx, 10, 64)
	if err != nil {
		CallNotifySend([]string{">>> FATAL ERROR <<<", "tsSys.GetNbProcess()", "-t", "20000", "-i", "/usr/share/icons/gnome/32x32/status/dialog-error.png"})
		panic(fmt.Sprint("tsSys - getProcessPIDs > parseInt: ", err))
	}

	return nb
}

// CallNotifySend affiche un notification popup à l'écran
func CallNotifySend(cmdArgs []string) {
	cmdName := "notify-send"
	cmd := exec.Command(cmdName, cmdArgs...)
	err := cmd.Run()
	if err != nil {
		panic(fmt.Sprint("notify-send ", err))
	}
}

// GetCurrentUser ; retourne l'utilisateur connecté
func GetCurrentUser() (string, error) {
	var currentUser string
	user, err := user.Current()
	if err == nil {
		currentUser = user.Username // jluc20mx
	}
	return currentUser, err
}
