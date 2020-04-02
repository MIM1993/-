package compr

type Compr interface {
	Compress([]byte) ([]byte, error)
	Name() string
}

/*
	param := os.Args[1]
	width, _ := strconv.Atoi(param)

	// open "test.jpg"
	file, err := os.Open("test.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(uint(width), 0, img, resize.Lanczos3)

	out, err := os.Create("test_resized.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
*/