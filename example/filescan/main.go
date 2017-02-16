package main

import(
	"path/filepath"
	"os"
	"flag"
	"fmt"
)

/*func visit(path string, f os.FileInfo, err error) error{
	fmt.Printf("Visited: %s\n", path)
	return nil
}*/


func main(){
	flag.Parse()
	root:=flag.Arg(0)
	fmt.Println(root)
	err:= filepath.Walk(root, func(path string, _ os.FileInfo, _ error) error{
		fmt.Printf("Visited: %s\n", path)
		return nil
	})
	fmt.Printf("filepath.Walk returned %v\n", err)

	/*dirname := string(root) // + string(filepath.Separator)
	fmt.Println(dirname)
	
	d, err := os.Open(dirname)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer d.Close()
	
	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Reading "+ dirname)
	for _, file := range files {
		//if file.Mode().IsRegular() {
		fmt.Println(file.Name(), file.Size(), "bytes")
		//if filepath.Ext(file.Name()) == ".png" {
			//	os.Remove("file.Name()")
			//	fmt.Println("Deleted ", file.Name())
			//}
	//	}
	}*/

}

