package main

import "flag"
import "fmt"
import "os"
import "regexp"
import "os/exec"
import "sync"
import "strings"
//import "syscall"


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


	wg := new(sync.WaitGroup)
	wg.Add(1)
	exe_cmd("ls -la", wg)
	wg.Wait()
	
	fmt.Println(*uuidFlag)
	fmt.Println(flag.Args())
}
