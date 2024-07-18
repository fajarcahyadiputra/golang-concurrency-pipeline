package main

import (
	imageprocessing "concurency-pipeline/image_processing"
	"fmt"
	"image"
	"strings"
)

type Job struct {
	Inputpath  string
	Image      image.Image
	outputPath string
}

func loadImage(paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		//For each input path create a job and add it to
		//the out channel
		for _, p := range paths {
			job := Job{
				Inputpath:  p,
				outputPath: strings.Replace(p, "images/", "images/output/", 1),
			}

			job.Image = imageprocessing.ReadImage(p)
			out <- job
		}
		close(out)
	}()
	return out
}

func resize(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		//For each input job, create a new job after resize and add it to
		//the out channel
		for job := range input { // Read from teh channel
			job.Image = imageprocessing.Resize(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func convertToGrayScale(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		//For each input job, create a new job after resize and add it to
		//the out channel
		for job := range input { // Read from teh channel
			job.Image = imageprocessing.GrayScale(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func saveImage(input <-chan Job) <-chan bool {
	out := make(chan bool)
	go func() {
		//For each input job, create a new job after resize and add it to
		//the out channel
		for job := range input { // Read from teh channel
			imageprocessing.WriteImage(job.outputPath, job.Image)
			out <- true
		}
		close(out)
	}()
	return out
}

func main() {
	imagePaths := []string{"images/image1.jpeg", "images/image2.jpeg", "images/image3.jpeg"}

	channel1 := loadImage(imagePaths)
	channel2 := resize(channel1)
	channel3 := convertToGrayScale(channel2)
	writeResult := saveImage(channel3)

	for success := range writeResult {
		if success {
			fmt.Println("Success!")
		} else {
			fmt.Println("Failed!")
		}
	}
}
