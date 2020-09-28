package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func hello(w http.ResponseWriter, req *http.Request) {
	app := "microk8s"

	arg0 := "add-node"

	cmd := exec.Command(app, arg0)
	stdout, err := cmd.Output()

	if err != nil {
		_, _ = fmt.Fprintf(w, err.Error())
		return
	}

	_, _ = fmt.Fprintf(w, string(stdout))
}

func main() {
	http.HandleFunc("/getToken", hello)

	_ = http.ListenAndServe(":8090", nil)
}
