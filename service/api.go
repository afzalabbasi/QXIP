package service

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

// CallLokiPushLogAPI method request LOKI POST Log API.
func CallLokiPushLogAPI(requestBody []byte) {
	var retryCount int
	logrus.Debugln("PushLogLoki API Function Trigger......")
	url := os.Getenv("URL")
	// Create a HTTP post request
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		logrus.Errorln("Error During Create a HTTP POST Request", err.Error())
		return
	}
	headerInfo := os.Getenv("HEADER")
	header := strings.Split(headerInfo, "|")
	for _, i := range header {
		headerInformation := strings.Split(i, ":")
		r.Header.Add(headerInformation[0], headerInformation[1])
	}
	r.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		logrus.Errorln("Error During Make The Post Request", err.Error())
		return
	}
	logrus.Infoln("Status Code....", res.StatusCode)
	if res.StatusCode != http.StatusNoContent {
		logrus.Infoln("Status Code Is Not Equal To StatusNoContent")
		logrus.Errorln("Status Code", res.StatusCode)
		logrus.Errorln("Response Body", res.Body)
		body, err := io.ReadAll(res.Body)
		if err != nil {
			logrus.Errorln("Read Response Body Error", string(body))
		}
		retryCount++
		retry(requestBody, retryCount)
	}

	defer res.Body.Close()
	logrus.Debugln("PushLogLoki API Function Complete......")
}

// retry method check retryCount and call CallLokiPushLogAPI.
func retry(requestBody []byte, retryCount int) {
	if retryCount > 10 {
		return
	}
	CallLokiPushLogAPI(requestBody)

}
