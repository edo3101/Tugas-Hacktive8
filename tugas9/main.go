package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type Data struct {
	Water int64 `json:"water"`
	Wind  int64 `json:"wind"`
}

func main() {
	var (
		waterstats string
		windstats  string
	)

	for {
		water := rand.Intn(100)
		wind := rand.Intn(100)

		if water < 5 {
			waterstats = "aman"
		} else if water >= 5 && water <= 8 {
			waterstats = "siaga"
		} else {
			waterstats = "bahaya"
		}

		if wind < 6 {
			windstats = "aman"
		} else if wind >= 6 && wind <= 15 {
			windstats = "siaga"
		} else {
			windstats = "bahaya"
		}

		data := Data{}
		data.Water = int64(water)
		data.Wind = int64(wind)
		reqJson, err := json.Marshal(data)
		client := &http.Client{}
		if err != nil {
			fmt.Print("Error bro")
		}

		req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(reqJson))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			fmt.Print("Error bro")
		}

		res, err := client.Do(req)
		if err != nil {
			fmt.Print("Error bro")
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Print("Error bro")
		}

		if err := json.Unmarshal([]byte(string(body)), &data); err != nil {
			fmt.Println(err)
			return
		}

		postjson, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			fmt.Print("Error bro")
		}

		fmt.Println(string(postjson))
		fmt.Printf("status water : %s \n", waterstats)
		fmt.Printf("status wind : %s \n", windstats)
		time.Sleep(15 * time.Second)
	}
}
