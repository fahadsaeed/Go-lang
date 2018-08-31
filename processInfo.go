package main

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
	"math"
	"net/http"
)

func dealwithErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func GetHardwareData(w http.ResponseWriter, r *http.Request) {

	html := "<html><br>===================== Proccess INFORMATION =============================<br><br>"

	processes, err := process.Processes()
	dealwithErr(err)
	//fmt.Println("processes = ", processes)

	for processesIndex, proc := range processes {
		fmt.Println("partition index", processesIndex)

		name, err := proc.Name()
		dealwithErr(err)
		fmt.Println("Name = ", name)

		cpuPercent, err := proc.CPUPercent()
		dealwithErr(err)
		fmt.Println("cPUPercent = ", cpuPercent)

		isRunning, err := proc.IsRunning()
		dealwithErr(err)
		fmt.Println("IsRunning = ", isRunning)

		status, err := proc.Status()
		dealwithErr(err)
		fmt.Println("Status = ", status)

		threads, err := proc.Threads()
		dealwithErr(err)
		fmt.Println("Threads = ", threads)

		numThreads, err := proc.NumThreads()
		dealwithErr(err)
		fmt.Println("NumThreads = ", numThreads)

		memoryInfo, err := proc.MemoryInfo()
		dealwithErr(err)
		fmt.Println("MemoryInfo = ", memoryInfo)

		//MemoryInfoEx, err := proc.MemoryInfoEx()
		//dealwithErr(err)
		//fmt.Println("MemoryInfoEx = ", MemoryInfoEx)

		memoryPercent, err := proc.MemoryPercent()
		dealwithErr(err)
		fmt.Println("MemoryPercent = ", memoryPercent)

		username, err := proc.Username()
		dealwithErr(err)
		fmt.Println("username = ", username)

		connections, err := proc.Connections()
		dealwithErr(err)
		fmt.Println("Connections = ", connections)

		createTime, err := proc.CreateTime()
		dealwithErr(err)
		fmt.Println("CreateTime = ", createTime)

		times, err := proc.Times()
		dealwithErr(err)
		fmt.Println("Times = ", times)

		cwd, err := proc.Cwd()
		dealwithErr(err)
		fmt.Println("Cwd = ", cwd)

		exe, err := proc.Exe()
		dealwithErr(err)
		fmt.Println("Exe = ", exe)

		IOCounters, err := proc.IOCounters()
		dealwithErr(err)
		fmt.Println("IOCounters = ", IOCounters)

		NetIOCounters, err := proc.NetIOCounters(true)
		dealwithErr(err)
		fmt.Println("IOCounters = ", NetIOCounters)

		html = html + "<br>--------------- Process  " + name + "  --------------- <br>"
		html = html + "Cpu Percent : " + fmt.Sprint(cpuPercent) + " <br>"
		html = html + "Process is Running : " + fmt.Sprint(isRunning) + " <br>"
		html = html + "Status : " + status + " <br>"
		html = html + "Process Threads : " + fmt.Sprint(threads) + " <br>"
		html = html + "Num of Threads : " + fmt.Sprint(numThreads) + " <br>"
		html = html + "Process memory : " + fmt.Sprint(memoryInfo) + " <br>"
		html = html + "username : " + fmt.Sprint(username) + " <br>"
		html = html + "connections: " + fmt.Sprint(connections) + " <br>"
		html = html + "Create Time : " + fmt.Sprint(createTime) + " <br>"
		html = html + "Percentage of memory : " + fmt.Sprint(times) + " <br>"
		html = html + "Create working directory : " + fmt.Sprint(cwd) + " <br>"
		html = html + "Execute : " + fmt.Sprint(exe) + " <br>"
		html = html + "IO Counters : " + fmt.Sprint(IOCounters) + " <br>"
	}

	html = html + "</html>"

	w.Write([]byte(html))

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", GetHardwareData)
	http.ListenAndServe(":7002", mux)

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
