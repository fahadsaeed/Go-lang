package main

import (
	"fmt"
	"math"
	"net/http"
	"github.com/shirou/gopsutil/process"
	//"context"
)

func dealwithErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func GetHardwareData(w http.ResponseWriter, r *http.Request) {

	//ctx := context.Background();
	// host or machine kernel, uptime, platform Info
	//processStat, err := process.PidExists()
	processStat, err := process.Processes()
	dealwithErr(err)
	//processStatExist, err := process.PidExists(96)
	//dealwithErr(err)
	//PidsWithContext, err := process.
	//dealwithErr(err)

	//fmt.Println("processStat = ", processStat)
	cpuffinity, err := processStat[5].Username()
	dealwithErr(err)

	//fmt.Println("processStatExist = ", processStatExist)
	fmt.Println("Connections = ", cpuffinity)
	//fmt.Println("PidsWithContext = ", PidsWithContext)

	html := "<html><br>===================== Proccess INFORMATION =============================<br><br>"
	//html = html + "OS: " + hostStat.OS + "<br>"
	//html = html + "Hostname: " + hostStat.Hostname + "<br>"
	//html = html + "Platform: " + hostStat.Platform + "<br>"
	//html = html + "Platform Family: " + hostStat.PlatformFamily + "<br>"
	//html = html + "Platform PlatformVersion: " + hostStat.PlatformVersion + "<br>"
	//html = html + "Number of processes running: " + strconv.FormatUint(hostStat.Procs, 10) + "<br>"
	//html = html + "Uptime: " + strconv.FormatUint(hostStat.Uptime, 10) + "<br>"
	//html = html + "Host ID(uuid): " + hostStat.HostID + "<br>"



	html = html + "</html>"

	w.Write([]byte(html))

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", GetHardwareData)
	http.ListenAndServe(":7001", mux)

}

func bytesToSize(bytes uint64) string {
	sizes := []string{"Bytes", "KB", "MB", "GB", "TB"}
	if bytes == 0 {
		return fmt.Sprint(float64(0), "bytes")
	} else {
		var bytes1 = float64(bytes)
		var i = math.Floor(math.Log(bytes1) / math.Log(1024))
		var count = bytes1 / math.Pow(1024, i)
		var j = int(i)
		var val = fmt.Sprintf("%.1f", count)
		return fmt.Sprint(val, sizes[j])
	}
}

//reference
//https://www.socketloop.com/tutorials/golang-get-hardware-information-such-as-disk-memory-and-cpu-usage
//https://stackoverflow.com/questions/15900485/correct-way-to-conver
