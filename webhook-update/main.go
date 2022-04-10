package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

func update(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "updating...\n")
	if _, err := os.Stat("../hugo.sh"); err == nil {
		fmt.Println("File exists :-)")
		cmd := exec.Command("../hugo.sh")
		cmd.Run()
	}

}

func main() {
	fmt.Println("Up and running!")
	http.HandleFunc("/update", update)

	http.ListenAndServe(":8080", nil)
}
