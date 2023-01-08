package controller

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/afzalabbasi/QXIP/service"
	"github.com/sirupsen/logrus"
)

// StreamInformation contain parent stream object.
type StreamInformation struct {
	Streams []MainStream `json:"streams"`
}

// MainStream contain stream information.
type MainStream struct {
	Stream Stream     `json:"stream"`
	Values [][]string `json:"values"`
}

// Stream contain stream data.
type Stream struct {
	Foo string `json:"foo"`
}

// QXIPJob handle all the main logic for making the value data.
func QXIPJob() {

	speed := os.Getenv("SPEED")
	speedInformation := strings.Split(speed, "l")
	logrus.Debugln("Speed Details....", speedInformation[0])
	speedInfo, _ := strconv.Atoi(speedInformation[0])
	intervalValue := float64(1) / float64(speedInfo)
	duration := time.Duration(intervalValue * float64(time.Second))
	currentTime := time.Now().UTC()
	intervalTime := currentTime.Add(time.Second * 1)
	var value [][]string
	for {
		if currentTime.After(intervalTime) {
			logrus.Debugln("Loop Break.......")
			break
		} else {
			currentTime = currentTime.Add(time.Nanosecond * duration)
			s := strconv.Itoa(int(currentTime.UnixNano()))
			data := []string{s, "foo bar test"}
			value = append(value, data)

		}
	}
	logrus.Debugln("Length of Values....", len(value))
	streamInformation := MainStream{
		Stream: Stream{
			Foo: "BAR",
		},
		Values: value,
	}
	m := []MainStream{
		streamInformation,
	}
	log := StreamInformation{
		m,
	}

	jsonData, _ := json.Marshal(log)
	// call CallLokiPushLogAPI method for POST request to LOKI Push Log API.
	service.CallLokiPushLogAPI(jsonData)

}
