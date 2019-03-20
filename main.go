package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// does spawner.xdp exist
	if _, err := os.Stat("./spawner.xdp"); os.IsNotExist(err) {
		// We might want to automatically fetch it if not the case or something
		fmt.Println("No spawner.xdp! You need one of those!")
		os.Exit(-1)
	} else {
		if _, err := os.Stat("C:\\Program Files (x86)\\Origin Games\\Command and Conquer Red Alert II\\spawner.xdp"); os.IsNotExist(err) {
			// In this case it means I have much more work to do
			fmt.Println("AAAAAAAAAAAAAAAAAA")
			os.Exit(-1)
		}
		// I'm well aware that the path to the game might not be this. But I think it should
		os.Rename("C:\\Program Files (x86)\\Origin Games\\Command and Conquer Red Alert II\\spawner.xdp", "C:\\Program Files (x86)\\Origin Games\\Command and Conquer Red Alert II\\spawner.xdp.tmp")

		// move the local spawner into prog files
		os.Rename("./spawner.xdp", "C:\\Program Files (x86)\\Origin Games\\Command and Conquer Red Alert II\\spawner.xdp")
		// Execute the CnC net Launcher
		cmd := exec.Command(`C:\Program Files (x86)\Origin Games\Command and Conquer Red Alert II\Resources\clientdx.exe`) // this works???
		// Probably not necessary to keep the stdout and stderr but yknow why not
		var outbuf, errbuf bytes.Buffer
		cmd.Stdout = &outbuf
		cmd.Stderr = &errbuf
		// actually run the command
		err = cmd.Run()
		// print stdout and stderr if they exist
		fmt.Println(outbuf.String())
		fmt.Println(errbuf.String())
		// is there an error with cmd.Run?
		if err != nil {
			fmt.Println("wrryyyy")
			fmt.Println(err)
		}
		// rename our file back
		os.Rename("C:\\Program Files (x86)\\Origin Games\\Command and Conquer Red Alert II\\spawner.xdp", "./spawner.xdp")

		// rename the temp back
		os.Rename("C:\\Program Files (x86)\\Origin Games\\Command and Conquer Red Alert II\\spawner.xdp.tmp", "C:\\Program Files (x86)\\Origin Games\\Command and Conquer Red Alert II\\spawner.xdp")

	}
}
