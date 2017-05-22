package kraken

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

// C global Configuration.
var C *Configuration

// E global Engine.
var E *Engine

//Start the service.
func Start(blocking bool) {
	shutdown := false
	C = UseConfiguration()

	log.Println("Starting " + C.ApplicationName + " Version " + C.ApplicationVersion)
	E = NewEngine()
	log.Println("Engine online.")

	E.LoadDirectory(C.DefaultStore)
	log.Println("Loaded " + strconv.Itoa(E.CountGraphs()) + " graph(s).")

	go autoSave(&shutdown)
	log.Println("Auto-Saving online.")

	hostConfig := C.Host + ":" + strconv.Itoa(C.Port)
	log.Println("Booting HTTP-API at " + hostConfig)
	go serve(hostConfig)
	log.Println("HTTP-API online.")

	log.Println("Boot completed.")

	listenOnShutDownEvent(&shutdown)
	if blocking {
		for {
			// Boot complete, sleep forever.
			time.Sleep(time.Hour * 1000)
		}
	}
}

func listenOnShutDownEvent(shutdown *bool) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Shutdown initiated. Waiting for processes to finish...")
		*shutdown = true
		time.Sleep(time.Hour * 1000)
	}()
	fmt.Println("Hit Ctrl-c to initiate shutdown.")
}

func serve(conf string) {
	router := mux.NewRouter().StrictSlash(C.StrictSlashesInURLs)
	router.HandleFunc("/", ServeEngine)
	router.HandleFunc("/{graph}/", ServeGraph)
	router.HandleFunc("/{graph}/{node}/", ServeNode)
	log.Fatal(http.ListenAndServe(conf, router))
}

func autoSave(shutdown *bool) {
	time.Sleep(C.AutoWriteInterval)
	for {
		num, err := E.WriteAllToDisk()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Wrote " + strconv.Itoa(num) + " graph(s) to disk.")
		if *shutdown {
			// autosave should shutdown to avoid data loss
			os.Exit(0)
		}
		time.Sleep(C.AutoWriteInterval)
	}
}
