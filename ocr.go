package main

import (
	"encoding/json"
	"fmt"
	"log"

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
	client := ocr.NewOCRClient(APIKEY, APISECRET)
	rs, err := client.AccurateRecognizeBasic(
		ocr.MustFromFile("ocr2.png"),
	)
	if err != nil {
		rs, err = client.GeneralRecognizeBasic(
			ocr.MustFromFile("ocr2.png"),
		)
	}

	result, err := rs.ToString()
	if err != nil {
		log.Fatal(err)
	}
	data := &FecthData{}
	fmt.Println(result)
	fmt.Println("-----------------------------------------")
	err = json.Unmarshal([]byte(result), data)
	if err != nil {
		log.Fatal(err)
	}

	for _, wr := range data.WordsResult {
		fmt.Println(wr.Words)
	}
}
