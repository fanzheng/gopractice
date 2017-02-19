package main

/*
a simple package to try to listen on any port you specify, to test if the firewall has blocked it.
usage: 
    ./listen [a number,like 80]
and try to visit that server using your browser like http://ip-to-your-host:port/
*/

import "os"
import "fmt"
import "strconv"
import "net/http"
import "log"

func main() {
	if len(os.Args) == 1 {
		fmt.Println("please pass me a port to listen to.")
		fmt.Println("now I am listening on port 80")
		port := 80
		http.HandleFunc("/", home)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
	} else {
		port, err := strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			fmt.Println(err)
			panic(fmt.Sprintf("fail to convert %v to a number.", os.Args[1]))
		} else {
			http.HandleFunc("/", home)
			log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got a request from %v.\n", r.Host)
	fmt.Fprint(w, "Welcome to listen, All is go.")
}