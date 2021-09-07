package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/signintech/gopdf"
)

const IMAGES_PATH = "./images"
const PDFS_PATH = "./pdfs"

func main() {
	flag.Parse()

	pdfFileName := flag.Arg(0)
	if pdfFileName == "" {
		//default pdf file name
		pdfFileName = time.Now().Format("20060102150405")
	}

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	files, err := ioutil.ReadDir(IMAGES_PATH)
	if err != nil {
		log.Fatal("read dir error:", err)
	}
	//sort by file ModTime from old to new
	sort.Slice(files, func(i, j int) bool { return files[i].ModTime().UnixNano() < files[j].ModTime().UnixNano() })

	for _, f := range files {
		fileName := f.Name()
		fileExtension := filepath.Ext(fileName)
		mime := mime.TypeByExtension(fileExtension)
		//check file is image
		if strings.Contains(mime, "image") {
			pdf.AddPage()
			//set image to full page
			err := pdf.Image(IMAGES_PATH+"/"+fileName, 0, 0, gopdf.PageSizeA4) //print image
			if err != nil {
				log.Fatal("print image error:", err)
			}
		}
	}

	pdf.WritePdf(fmt.Sprintf("%v/%v.pdf", PDFS_PATH, pdfFileName))
}
