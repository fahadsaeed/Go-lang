package main

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"net/http"
	"strconv"
	"math"

)

func dealwithErr(err error) {
	if err != nil {
		fmt.Println(err)
		//os.Exit(-1)
	}
}

func GetHardwareData(w http.ResponseWriter, r *http.Request) {


	// host or machine kernel, uptime, platform Info
	hostStat, err := host.Info()
	dealwithErr(err)
	//fmt.Println("hostStat", hostStat)


	html :=  "<html><br>===================== HOST INFORMATION =============================<br><br>"
	html = html + "OS: " + hostStat.OS + "<br>"
	html = html + "Hostname: " + hostStat.Hostname + "<br>"
	html = html + "Platform: " + hostStat.Platform + "<br>"
	html = html + "Platform Family: " + hostStat.PlatformFamily + "<br>"
	html = html + "Platform PlatformVersion: " + hostStat.PlatformVersion + "<br>"
	html = html + "Number of processes running: " + strconv.FormatUint(hostStat.Procs, 10) + "<br>"
	html = html + "Uptime: " + strconv.FormatUint(hostStat.Uptime, 10) + "<br>"
	html = html + "Host ID(uuid): " + hostStat.HostID + "<br>"



	// memory
	vmStat, err := mem.VirtualMemory()
	dealwithErr(err)
	//fmt.Println("VirtualMemory", vmStat)

	html  = html + "<br><br>===================== Memory RAM INFORMATION =============================<br><br>"

	html = html + "Total memory in bytes: "  + fmt.Sprint(vmStat.Total) + " and " + bytesToSize(vmStat.Total) + "<br>"
	html = html + "Used memory in bytes: "  + fmt.Sprint(vmStat.Used) + " and " + bytesToSize(vmStat.Used) + "<br>"
	html = html + "Unused memory in bytes: "  + fmt.Sprint(vmStat.Available) + " and " + bytesToSize(vmStat.Available) + "<br>"
	html = html + "Free memory  in bytes: " + fmt.Sprint(vmStat.Free) + " and " +  bytesToSize(vmStat.Free)+" <br>"
	html = html + "Percentage used memory: " + fmt.Sprint(vmStat.UsedPercent) + " and " + strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64) + "%<br>"

	// Disk
	diskPartitions, err := disk.Partitions(true)
	dealwithErr(err)

	html = html + "<br><br>===================== DISK SPACE INFORMATION =============================<br><br>"

	for partitionIndex, partition := range diskPartitions {
		partitionStat, err := disk.Usage(partition.Mountpoint)
		dealwithErr(err)
		fmt.Println("partition index", partitionIndex)

		html = html + "<br>--------------- Partition "+   partition.Mountpoint +"  --------------- <br>"
		html = html + "Total space  in bytes: " + fmt.Sprint(partitionStat.Total) + " and "+  bytesToSize(partitionStat.Total)  + " <br>"
		html = html + "Used space in bytes: " +  fmt.Sprint(partitionStat.Used) + " and " + bytesToSize(partitionStat.Used)  +  " <br>"
		html = html + "Free space in bytes`: " +  fmt.Sprint(partitionStat.Free) + " and " + bytesToSize(partitionStat.Free)  + " <br>"
		html = html + "Percentage space usage: " + fmt.Sprint(partitionStat.UsedPercent) + "and " + strconv.FormatFloat(partitionStat.UsedPercent, 'f', 2, 64) + "%<br>"
	}



	// cpu - get CPU number of cores and speed
	cpuStat, err := cpu.Info()
	dealwithErr(err)
	percentage, err := cpu.Percent(0, true)
	dealwithErr(err)

	//fmt.Println("cpuStat", cpuStat)
	//fmt.Println("cpuStat percentage ", percentage)

	html = html + "<br><br>===================== CPU INFORMATION =============================<br><br>"
	html = html + "CPU index number: " + strconv.FormatInt(int64(cpuStat[0].CPU), 10) + "<br>"
	html = html + "CPU index number: " + fmt.Sprint(cpuStat[0].Model) + "<br>"
	html = html + "VendorID: " + cpuStat[0].VendorID + "<br>"
	html = html + "Family: " + cpuStat[0].Family + "<br>"
	html = html + "Number of cores: " + strconv.FormatInt(int64(cpuStat[0].Cores), 10) + "<br>"
	html = html + "Model Name: " + cpuStat[0].ModelName + "<br>"
	html = html + "Speed: " + strconv.FormatFloat(cpuStat[0].Mhz, 'f', 2, 64) + " MHz <br>"



	//
	//CPU        int32    `json:"cpu"`
	//VendorID   string   `json:"vendorId"`
	//Family     string   `json:"family"`
	//Model      string   `json:"model"`
	//Stepping   int32    `json:"stepping"`
	//PhysicalID string   `json:"physicalId"`
	//CoreID     string   `json:"coreId"`
	//Cores      int32    `json:"cores"`
	//ModelName  string   `json:"modelName"`
	//Mhz        float64  `json:"mhz"`
	//CacheSize  int32    `json:"cacheSize"`
	//Flags      []string `json:"flags"`
	//Microcode  string   `json:"microcode"`


	for idx, cpupercent := range percentage {
		html = html + "Current CPU utilization: [" + strconv.Itoa(idx) + "] " + strconv.FormatFloat(cpupercent, 'f', 2, 64) + "%<br>"
	}


	// Load
	//loadStat, err := load.Avg();
	//dealwithErr(err)
	loadMisc, err := load.Misc();
	dealwithErr(err)
	//fmt.Println("loadStat", loadStat)
	//fmt.Println("loadMisc.................", loadMisc)
	html = html + "<br><br>===================== LOAD INFORMATION =============================<br><br>"
	html = html + "Procs Running: " + fmt.Sprint(loadMisc.ProcsRunning) + "<br>"
	html = html + "Procs ProcsBlocked: " + fmt.Sprint(loadMisc.ProcsBlocked) + "<br>"
	html = html + "Procs Ctxt: " + fmt.Sprint(loadMisc.Ctxt) + "<br>"


	html = html + "</html>"

	w.Write([]byte(html))

}



func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", GetHardwareData)

	http.ListenAndServe(":8002", mux)

}

func bytesToSize(bytes uint64) string {
	sizes := []string {"Bytes", "KB", "MB", "GB", "TB"};
	if bytes == 0 {
		return fmt.Sprint(float64(0), "bytes")
	} else {
		var bytes1 = float64(bytes)
		var i = math.Floor(math.Log(bytes1) / math.Log(1024));
		var count = bytes1 / math.Pow(1024, i)
		var j = int(i)
		var val = fmt.Sprintf("%.1f", count)
		//fmt.Println("bytes, count,  val........", bytes, count,  val);
		return fmt.Sprint(val , sizes[j])
		//return fmt.Sprint(float64(count) , sizes[j])
		//return fmt.Sprintf("%.2f", float64(count) )
	}
};

//reference
//https://www.socketloop.com/tutorials/golang-get-hardware-information-such-as-disk-memory-and-cpu-usage
//https://stackoverflow.com/questions/15900485/correct-way-to-conver