package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"os/exec"
	"sync"
	"strings"
	"bufio"
	"log"
	"reflect"
	"math/rand"
	"time"
)
//import "syscall"


func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}


func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}






func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}


func exe_cmd(cmd string, wg *sync.WaitGroup) {
	fmt.Println("command is ",cmd)
	// splitting head => g++ parts => rest of the command
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head,parts...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)
	wg.Done() // Need to signal to waitgroup that this goroutine is done
}


func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return (rand.Intn(max - min) + min)
}

func main() {

	voFlag := flag.String("VO", "", "VO for report")
	uuidFlag := flag.String("UUID", "", "please set a valid uuid")
	flag.Parse()
	//_=voFlag
	//_=uuidFlag
	if *voFlag =="" || IsValidUUID(*uuidFlag)==false{
		fmt.Println("one of mandatory options missing or wrong a sample execution: ./run -VO sdger -UUID 550e8400-e29b-41d4-a716-446655440000")
		os.Exit(1)
	}

	var dwhp_report_name string= fmt.Sprintf("Reporting_%s_%s.csv", *voFlag, *uuidFlag)  
	fmt.Println("the new report name: ", dwhp_report_name)

		
	//fmt.Println(flag.Args())



	lines, err := readLines("./sample_dwhp_report/Reporting.csv")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println(reflect.TypeOf(lines))
	fmt.Println(len(lines))

	var line int = random(90921, 181842)


//	for i := 1; i <= line; i++ {
//		fmt.Println(i, lines[i])
//	}

	if err := writeLines(lines[:line], dwhp_report_name); err != nil {
		log.Fatalf("writeLines: %s", err)
	}
	
	wg := new(sync.WaitGroup)
	wg.Add(2)
	//var cmd string=fmt.Sprintf("scp ./%s tvpp@10.1.234.189:/svc/tmv1684/moip/mo30501/dwhp", dwhp_report_name)
	var cmd string=fmt.Sprintf("scp -P 6189 ./%s tvpp@localhost:/svc/tmv1684/moip/mo30501/dwhp", dwhp_report_name)
	exe_cmd(cmd, wg)
	
	exe_cmd(fmt.Sprintf("rm -rf ./%s",dwhp_report_name), wg)
	wg.Wait()


}
