package utils

import (
	"fmt"
	"strings"
)

type ProgressBar struct {
	Title string
	last_value interface{}
}

func (pb *ProgressBar) clean () {
	fmt.Print("\r")
}

func (pb *ProgressBar) SetText(text string) {
	pb.clean()
	fmt.Print(pb.Title, ": ", text)
}

func (pb *ProgressBar) Init() {
	pb.SetText("Waiting...")
}

func (pb *ProgressBar) Update (total, done uint64) {
	percent := int((float32(done)/float32(total))*100)
	if pb.last_value != nil && percent != pb.last_value.(int) && percent > 1 {
		bar := fmt.Sprintf("[%-49v] %2v%% %v/%v", strings.Repeat("=", int(float32(percent)/2.0)-1)+">", percent, done, total)
		pb.SetText(bar)
	}
	pb.last_value = percent
}

func (pb *ProgressBar) Finish () {
	fmt.Print("\n")
}
