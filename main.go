package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadDriver() {
	//sizeOfEachEntry := []int{3, 8, 2, 1, 2, 1, 2, 2, 1, 2, 2, 2, 2, 4, 4, 2, 1, 4, 11, 8}
	f, err := os.Open("mydriver")
	check(err)
	defer f.Close()

	reader := bufio.NewReader(f)
	buf := make([]byte, 512)

	for i := 0; i < 1; i++ {
		_, err := reader.Read(buf)

		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}

		//fmt.Printf("%s", hex.Dump(buf))
	}

	// infos := read.BootSector{}
	// values := reflect.TypeOf(infos)
	// for i := 0; i < values.NumField(); i++ {
	// 	for _, value := range sizeOfEachEntry {
	// 		switch i {
	// 		case 0:
	// 			copy(infos.JumpBoot[:], buf[:value])
	// 		}
	// 		buf = buf[value:]
	// 	}
	// }

	// fmt.Println(infos)

	// for _, value := range sizeOfEachEntry {
	// 	//fmt.Printf("%x\n", buf[:value])
	// 	fmt.Printf("%d", value)
	// 	fmt.Print(buf[:value])
	// 	// for i, v := range buf[:value] {
	// 	// 	fmt.Printf("%d: ", i)
	// 	// 	fmt.Printf("%c", v)
	// 	// }

	// 	buf = buf[value:]
	// }
	fmt.Print(buf[:0x03])

	// for i, v := range allValues {
	// 	fmt.Printf("%d: ", i)
	// 	fmt.Printf("[%d];", v)
	// }
}

func main() {
	ReadDriver()
}
