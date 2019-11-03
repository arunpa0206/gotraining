package main

import (
	"fmt"
	"log"
	"runtime"
)

// bigBytes allocates 100 megabytes
func bigBytes() *[]byte {
	s := make([]byte, 100000000)
	return &s
}
func main() {
	var mem runtime.MemStats

	fmt.Println("memory baseline...")

	runtime.ReadMemStats(&mem)
	log.Println("mem.Alloc:", mem.Alloc)
	log.Println("total allocated meomory:", mem.TotalAlloc)
	log.Println("mem.HeapAlloc:", mem.HeapAlloc)
	log.Println("mem.Heapsys:", mem.HeapSys)

	for i := 0; i < 10; i++ {
		s := bigBytes()
		if s == nil {
			log.Println("oh noes")
		}
	}

	fmt.Println("memory comparison...")

	runtime.ReadMemStats(&mem)
	log.Println("mem.Alloc", mem.Alloc)
	log.Println("mem.TotalAlloc:", mem.TotalAlloc)
	log.Println("mem.HeapAlloc:", mem.HeapAlloc)
	log.Println("mem.HeapSys:", mem.HeapSys)
}
