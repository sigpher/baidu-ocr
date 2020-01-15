package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/chenqinghe/baidu-ai-go-sdk/vision/ocr"
)

// APIKEY 百度AI－ocr相关信息
const (
	APIKEY    = "WGsG5BdaZW28oTNsX3dn3dfN"
	APISECRET = "bLPdETnbo5XwwMc7Tyc8FZ1NOkfwt0IN"
)

// Words 结构体
type Words struct {
	Words string `json:"words"`
}

// FecthData 结构体
type FecthData struct {
	LogID          int     `json:"log_id"`
	WordsResultNum int     `json:"words_result_num"`
	WordsResult    []Words `json:"words_result"`
}

func main() {
	OCR()
	DelayClose(30)
}

//OCR ocr
func OCR() {
	var imagePath string
	fmt.Println("请输入需要识别的图片路径,如d:/img/a.jpg")
	fmt.Printf("请输入图片路径:")
	fmt.Scanln(&imagePath)

	client := ocr.NewOCRClient(APIKEY, APISECRET)
	rs, err := client.AccurateRecognizeBasic(
		ocr.MustFromFile(imagePath),
	)
	if err != nil {
		rs, err = client.GeneralRecognizeBasic(
			ocr.MustFromFile(imagePath),
		)
	}

	result, err := rs.ToString()
	if err != nil {
		log.Fatal(err)
	}
	data := &FecthData{}
	// fmt.Println(result)

	err = json.Unmarshal([]byte(result), data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-----------------------------------------")
	for _, wr := range data.WordsResult {
		fmt.Println(wr.Words)
	}
	fmt.Println("-----------------------------------------")
}

// DelayClose delay to close the window
func DelayClose(num int) {
	time.Sleep(time.Duration(num) * time.Second)
}
