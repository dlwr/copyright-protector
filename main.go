package main

import (
	"fmt"
	"github.com/gographics/imagick/imagick"
	// "github.com/k0kubun/pp"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/protected.png", handler)
	http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}


func handler(w http.ResponseWriter, r *http.Request) {
	wand,err := imageFromUrl(w, r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	q := r.URL.Query()
	if q.Get("mozaic") == "true" {
		wand = resizeImage(wand, 640, true)
	} else {
		wand = resizeImage(wand, 640, false)
	}
	if q.Get("tile") == "true" {
		tileLineImage(wand)
	}
	var data []byte
	if q.Get("glitch") == "true" {
		data, err = glitchImage(wand, r.URL.Query())
	} else {
		data = wand.GetImage().GetImageBlob()
	}
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// data := wand.GetImage().GetImageBlob()
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))
	w.Header().Set("Cache-Control", "no-transform,public,max-age=86400ms-maxage=259000")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func imageFromUrl(w http.ResponseWriter, r *http.Request) (*imagick.MagickWand, error) {
	queryUrl := r.URL.Query().Get("url")
	url, _ := url.QueryUnescape(queryUrl)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error retrieving url", 500)
		return nil, err
	}
	wand := imagick.NewMagickWand()
	err = wand.ReadImageBlob(data)
	if err != nil {
		http.Error(w, "Error retrieving url", 500)
		return nil, err
	}
	if err = wand.SetImageFormat("JPG"); err != nil {
		http.Error(w, "Error retrieving url", 500)
		return nil, err
	}
	wand.AutoLevelImage()
	return wand, nil
}

func resizeImage(wand *imagick.MagickWand, size int, mozaic bool) *imagick.MagickWand {
	width := float32(wand.GetImageWidth())
	height := float32(wand.GetImageHeight())
	var rate float32
	if width > height {
		rate = float32(size) / width
	} else {
		rate = float32(size) / height
	}
	if mozaic {
		wand.ResizeImage(uint(width*rate/20), uint(height*rate/20), imagick.FILTER_LANCZOS, 1)
		wand.ResizeImage(uint(width*rate), uint(height*rate), imagick.FILTER_POINT, 1)
	} else {
		wand.ResizeImage(uint(width*rate), uint(height*rate), imagick.FILTER_LANCZOS, 1)
	}
	return wand.GetImage()
}

func tileLineImage(wand *imagick.MagickWand) {
	it := wand.NewPixelIterator()
	it.SetLastIteratorRow()
	width := float32(wand.GetImageWidth())
	height := float32(wand.GetImageHeight())
	heightRemain := float32(int(height / 20))
	if heightRemain != 0 {
		heightRemain = height / float32(heightRemain) - 20
	}
	widthRemain := float32(int(width / 20))
	if widthRemain != 0 {
		widthRemain = width / float32(widthRemain) - 20
	}
	cnt := it.GetIteratorRow()
	it.SetFirstIteratorRow()
	heightCount := 0
	nextHeight := 21
	hLineCount := 0
	hRemainCount := heightRemain
	widthCount := 0
	nextWidth := 21
	wLineCount := 0
	wRemainCount := widthRemain
	for i := 0; i <= cnt; i++ {
		heightCount++
		it.SetIteratorRow(i)
		pws := it.GetCurrentIteratorRow()
		if nextHeight - heightCount < 2 {
			for j := 0; j < len(pws); j++ {
				pws[j].SetColor("#ffffff")
			}
			if nextHeight == heightCount {
				heightCount = 0
				hLineCount++
				hRemainCount += heightRemain
				if hRemainCount > 1 {
					hRemainCount -= 1
					nextHeight = 21
				} else {
					nextHeight = 20
				}
			}
		} else {
			for j := 0; j < len(pws); j++ {
				widthCount++
				if nextWidth - widthCount < 2 {
					pws[j].SetColor("#ffffff")
					if nextWidth == widthCount {
						widthCount = 0
						wLineCount++
						wRemainCount += widthRemain
						if wRemainCount > 1 {
							wRemainCount -= 1
							nextWidth = 21
						} else {
							nextWidth = 20
						}
					}
				}
				// if j % 20 < 2 {
				// 	pws[j].SetColor("#ffffff")
				// }
			}
		}
		it.SyncIterator()
	}
}


func glitchImage(wand *imagick.MagickWand, q url.Values) ([]byte, error) {
	data := wand.GetImage().GetImageBlob()
	jpgHeaderLength := getJpegHeaderSize(data)
	maxIndex := len(data) -jpgHeaderLength - 4
	params := getParams(q)
	length := int(params["iterations"])
	for i := 0; i < length; i++ {
		pxMin := math.Floor(float64(maxIndex) / params["iterations"] * float64(i))
		pxMax := math.Floor(float64(maxIndex) / params["iterations"] * float64((i + 1)))
		delta := pxMax - pxMin
		pxI := math.Floor(pxMin + delta*params["seed"])
		if int(pxI) > maxIndex {
			pxI = float64(maxIndex)
		}
		index := math.Floor(float64(jpgHeaderLength) + pxI)
		data[int(index)] = byte(math.Floor(params["amount"] * float64(256)))
	}
	wand2 := imagick.NewMagickWand()
	if err := wand2.ReadImageBlob(data); err != nil {
		return nil, err
	}
	if err := wand2.SetImageFormat("PNG"); err != nil {
		return nil, err
	}
	return wand2.GetImage().GetImageBlob(), nil
}

func getParams(q url.Values) map[string]float64 {
	params := make(map[string]float64)
	rand.Seed(time.Now().UnixNano())
	if seed, _ := strconv.Atoi(q.Get("seed")); seed != 0 {
		params["seed"] = float64(seed) / 100
	} else {
		params["seed"] = float64(rand.Intn(100)) / 100
	}
	if amount, _ := strconv.Atoi(q.Get("amount")); amount != 0 {
		params["amount"] = float64(amount) / 100
	} else {
		params["amount"] = float64(rand.Intn(99)) / 100
	}
	if iterations, _ := strconv.Atoi(q.Get("iterations")); iterations != 0 {
		params["iterations"] = float64(iterations % 51)
		if params["iterations"] == 0 {
			params["iterations"] = 1
		}
	} else {
		params["iterations"] = float64(rand.Intn(29)) + 21
	}
	return params
}

func getJpegHeaderSize(data []byte) int {
	var result = 417
	len := len(data)
	for i := 0; i < len; i++ {
		if data[i] == 0xFF && data[i+1] == 0xDA {
			result = i + 2
			break
		}
	}
	return result
}
