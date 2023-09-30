package main

import (
	"flag"
	"fmt"
	c "go-gs/configs"
	g "go-gs/globals"
	m "go-gs/models"
	"go-gs/services"
	"log"
	"net/http"
	"os"
	"time"

	mux "github.com/gorilla/mux"
)

func init() {
	parseFlags()
	g.Cfg = c.LoadConfig(g.CfgPath)
}

func parseFlags() {
	flagSet := flag.NewFlagSet("flags", flag.PanicOnError)
	flagSet.StringVar(&g.CfgPath, "config", "tmp/config.json", "config file path")
	flagSet.StringVar(&m.FOLDER_PATH, "path", "./tmp", "folder path containing pdf files")
	flagSet.IntVar(&m.WORKERS, "workers", 4, "number of workers to run")
	flagSet.IntVar(&m.THREADS, "threads", 4, "number of threads to run")
	flagSet.Parse(os.Args[1:])
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/pdfToJpeg", services.PdfToJpegHandler)
	fmt.Println(g.Cfg)
	// services.PdfToJpg(cfg)
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
