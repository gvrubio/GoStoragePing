package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/fatih/color"
)

func CheckIfFileExists(file string) bool {
	_, error := os.Stat(file)
	var value bool
	// check if error is "file not exists"
	if os.IsNotExist(error) {
		value = false
	} else {
		value = true
	}
	return value
}

func CreateFile(file string) {
	if CheckIfFileExists(file) {
		color.Set(color.FgRed)
		log.Println("File exists!")
		os.Exit(1)
	} else {
		f, e := os.Create(file)
		if e != nil {
			log.Fatal(e)
		}
		color.Set(color.FgCyan)
		log.Printf("Created file \"%v\"\n", file)
		f.Close()
	}
}

func RemoveFile(file string) {
	e := os.Remove(file)
	if e != nil {
		log.Fatal(e)
	}
	color.Set(color.FgCyan)
	log.Printf("Removed file \"%v\"\n", file)
}

func WriteStringToFile(text string, file string) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err2 := f.WriteString(text)
	if err2 != nil {
		log.Fatal(err2)
	}
}

func ReadStringFromFile(file string) {
	_, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// FLAGS
	pathPtr := flag.String("path", "empty", "Give a path for the file, this file will use to read/write. THIS FILE WILL BE OVERWRITTEN")
	// TODO: RW Options
	//readPtr := flag.Bool("r", false, "Give a path for the file, this file will use to read/write. THIS FILE WILL BE OVERWRITTEN")
	//writePtr := flag.Bool("w", false, "Give a path for the file, this file will use to read/write. THIS FILE WILL BE OVERWRITTEN")
	intervalPtr := flag.Int64("interval", 1000, "Interval between checks in milliseconds.")
	textPtr := flag.String("text", "ping", "You can change the text that is written to the test file.")
	latencyWarnPtr := flag.Int64("latencywarn", 20, "Latency in milliseconds at which the text will change color to show high latency")
	flag.Parse()

	// Vars from flags
	var path string = *pathPtr
	// > TODO: Read and write options are disabled, as they are messed by caches, I need to study it more.
	//var read bool = *readPtr
	//var write bool = *writePtr
	var read bool = false
	var write bool = false
	// TODO <

	var interval int64 = *intervalPtr
	var text string = *textPtr
	var latencyWarn int64 = *latencyWarnPtr

	// Time calculation vars
	start := time.Now()
	elapsed := time.Since(start)
	var latency float64 = float64(elapsed.Microseconds()) / 1000

	// Initial file creation
	CreateFile(path)

	//CTRL+C Handle
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			RemoveFile(path)
			log.Printf("Captured %v, byebye!", sig)
			os.Exit(0)
		}
	}()
	// Measures
	//var interval time.Duration = time.Duration(*intervalPtr)
	for true {
		if (write && !read) || (!write && !read) {
			//WRITE
			//Measurement
			start = time.Now()
			WriteStringToFile(text, path)
			elapsed = time.Since(start)

			//Conversion to milliseconds
			latency = float64(elapsed.Microseconds()) / 1000

			//Coloring
			if latency >= float64(latencyWarn) {
				color.Set(color.FgRed)
			} else {
				color.Set(color.FgGreen)
			}

			// Print!
			log.Printf("W: => P[I]NG time=%v ms", float64(elapsed.Microseconds())/1000)

			// Interval
			time.Sleep(time.Duration(interval) * time.Millisecond)
		}
		if (!write && read) || (!write && !read) {
			//READ
			//Measurement
			start = time.Now()
			ReadStringFromFile(path)
			elapsed = time.Since(start)

			//Conversion to milliseconds
			latency = float64(elapsed.Microseconds()) / 1000

			// Coloring
			if latency >= float64(latencyWarn) {
				color.Set(color.FgRed)
			} else {
				color.Set(color.FgGreen)
			}

			// Print!
			log.Printf("R: <= P[O]NG time=%v ms", float64(elapsed.Microseconds())/1000)

			// Interval
			time.Sleep(time.Duration(interval) * time.Millisecond)
		}
	}
}
