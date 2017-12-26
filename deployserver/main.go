package webserver

import (
	"os/exec"
	"log"
	"net/http"
	"io"
)



func reLaunch() {
	cmd := exec.Command("sh","./deploy.sh")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func firstPage (w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "deploy")
	reLaunch()
}



func main() {

	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":5000", nil)
}
