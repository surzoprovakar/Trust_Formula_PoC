package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func WriteDataInFile(filename string) {

	buf, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		buf, _ = os.Create(filename)
	}

	for i := 0; i < 30; i++ {
		buf.WriteString(strconv.Itoa(rand.Intn(30)) + " ")
		time.Sleep(time.Second * time.Duration(rand.Intn(3-1)+1))
		buf.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
		buf.WriteString("\n")
	}
}

func GenerateTrust(fileIn string, fileOut string) {

	bufIn, _ := os.Open(fileIn)

	bufout, err := os.OpenFile(fileOut, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		bufout, _ = os.Create(fileOut)
	}

	snl := bufio.NewScanner(bufIn)

	prev_time := 0
	prev_trust := 0.5
	total_request := 1
	accepted_merge := 0
	result := ""
	trust := 0.0
	alpha := 0.1
	for snl.Scan() {

		var delta int
		var time int

		line := strings.Split(snl.Text(), " ")

		delta, _ = strconv.Atoi(line[0])
		time, _ = strconv.Atoi(line[1])

		err = snl.Err()
		if err != nil {
			log.Fatal(err)
		}

		if delta == 0 {
			delta = 1
		}
		if total_request == 1 {
			result = "Yes"
			trust = ((1 - alpha) * prev_trust) / (alpha * float64(delta))
			accepted_merge++
		} else {
			numerator := (1 - alpha) * prev_trust * float64(accepted_merge)
			if time == prev_time {
				time = prev_time + 1
			}
			denominator := float64(delta*(time-prev_time)*total_request) * alpha

			bufout.WriteString(fmt.Sprintf("num "+"%.4v", numerator) + " ")
			bufout.WriteString(fmt.Sprintf("den "+"%.4v", denominator) + " ")

			trust = numerator / denominator

			if trust >= 0.5 {
				result = "Yes"
				accepted_merge++
			} else {
				result = "No"
			}
		}

		bufout.WriteString(fmt.Sprintf("%.4v", trust) + " ")
		bufout.WriteString(result + "\n")

		total_request++
		prev_trust = trust
		prev_time = time

	}
}

func main() {
	WriteDataInFile("counter.txt")
	GenerateTrust("counter.txt", "trust.txt")

	// alpha := 0.1
	// prev_trust := 0.5
	// prev_time := 0.0
	// time := 1649901344
	// delta := 11
	// numerator := (1 - alpha) * prev_trust * (float64(time) - prev_time)
	// denominator := float64(delta)

	// trust := numerator / denominator
	// s := fmt.Sprintf("%.4f", trust)
	// fmt.Println(s)

}
