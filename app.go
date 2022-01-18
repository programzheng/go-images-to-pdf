package main

import (
	"context"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/signintech/gopdf"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (b *App) startup(ctx context.Context) {
	// Perform your setup here
	b.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
func (b *App) domReady(ctx context.Context) {
	// Add your action here
	filePaths, err := runtime.OpenMultipleFilesDialog(ctx, runtime.OpenDialogOptions{
		Title: "Select Image File",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Images (*.png;*.jpg)",
				Pattern:     "*.png;*.jpg",
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if len(filePaths) == 0 {
		return
	}
	//sort by filename
	sort.Slice(filePaths, func(i, j int) bool { return filePaths[i] < filePaths[j] })

	pdfFileName := time.Now().Format("20060102150405")
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})

	for _, path := range filePaths {
		fmt.Println(path)
		pdf.AddPage()
		err := pdf.Image(path, 0, 0, gopdf.PageSizeA4) //print image
		if err != nil {
			log.Fatal("print image error:", err)
		}
	}

	err = pdf.WritePdf(fmt.Sprintf("%v%v.pdf", "./", pdfFileName))
	if err != nil {
		log.Fatal("write pdf error:", err)
	}
	panic("close app")
}

// shutdown is called at application termination
func (b *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}
