package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/skratchdot/open-golang/open"
	"strconv"
)

const (
	delimiter = "/"
)

type ticket struct {
	TicketBaseURL string   `json:"ticketBaseURL"`
	Activity      string   `json:"activity"`
	Sku           string   `json:"sku"`
	TicketURI     []string `json:"ticketURI"`
	AreaURI       string   `json:"areaURI"`
	Identity      string   `json:"identity"`
	SeatId        string   `json:"seatId"`
	MagicalNumber string   `json:"magicalNumber"`
}

func newTicketConfig() (*ticket, error) {
	file, e := ioutil.ReadFile("ticket.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		return nil, e
	}

	t := &ticket{}
	_ = json.Unmarshal(file, &t)
	return t, nil
}

func main() {
	ticket, err := newTicketConfig()
	if err != nil {
		log.Fatal("failed to load config")
	}
	fmt.Println("load config:", ticket)

	url := ticket.TicketBaseURL
	if ticket.Identity == "" {
		log.Fatal("The identity field should be used")
		return
	}
	if ticket.SeatId != "" && ticket.MagicalNumber != "" {
		identity := delimiter + ticket.Identity
		seatId := delimiter + ticket.SeatId
		magicalNum := delimiter + ticket.MagicalNumber
		url += ticket.Activity + ticket.TicketURI + ticket.Sku + identity + seatId + magicalNum
		open.Run(url)
		return
	}

	tr := &http.Transport{
		TLSClientConfig:    nil,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	identity := delimiter + ticket.Identity
	url += ticket.Activity + ticket.AreaURI + ticket.Sku + identity
	res, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	bodyStr := string(robots)
	// split front
	sp1 := strings.Split(bodyStr, "var areaUrlList = ")
	// split tail
	sp2 := strings.Split(sp1[1], ";")
	// replace '\' with ''
	replace1 := strings.Replace(sp2[0], "\\", "", -1)

	// fill to map
	urls := make(map[string]string)
	err = json.Unmarshal([]byte(replace1), &urls)

	// open url
	for key, url := range urls {
		k := strings.Split(key, "_")
		seatId, _ := strconv.Atoi(k[1])

		if seatId >= 1 && seatId <= 5 {

		}
		fmt.Println(ticket.TicketBaseURL + url)
		open.Run(ticket.TicketBaseURL + url)
	}
}
