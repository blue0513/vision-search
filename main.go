package main

import (
	"context"
	"fmt"
	"log"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
	pb "cloud.google.com/go/vision/v2/apiv1/visionpb"
)

func checkError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

func openImageFile(path string) (*os.File, error) {
	return os.Open(path)
}

func detectWebEntities(ctx context.Context, client *vision.ImageAnnotatorClient, file *os.File) (*pb.WebDetection, error) {
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		return nil, err
	}
	return client.DetectWeb(ctx, image, nil)
}

func printWebDetectionResults(webDetection *pb.WebDetection) {
	if len(webDetection.WebEntities) == 0 {
		log.Fatalf("No web entities found.")
	}
	for _, entity := range webDetection.WebEntities {
		fmt.Printf("Description: %v\n", entity.Description)
	}
	if webDetection.FullMatchingImages != nil {
		for _, image := range webDetection.FullMatchingImages {
			fmt.Printf("Full matching image URL: %s\n", image.Url)
		}
	}
	if webDetection.PagesWithMatchingImages != nil {
		for _, page := range webDetection.PagesWithMatchingImages {
			fmt.Printf("Page with matching image: %s\n", page.Url)
		}
	}
	if webDetection.PartialMatchingImages != nil {
		for _, page := range webDetection.PartialMatchingImages {
			fmt.Printf("Page with partial matching image: %s\n", page.Url)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: <image_path>")
	}
	imagePath := os.Args[1]

	file, err := openImageFile(imagePath)
	checkError(err, "Failed to read image file")
	defer file.Close()

	ctx := context.Background()

	client, err := vision.NewImageAnnotatorClient(ctx)
	checkError(err, "Failed to create client")

	webDetection, err := detectWebEntities(ctx, client, file)
	checkError(err, "DetectWebEntities")

	printWebDetectionResults(webDetection)
}
