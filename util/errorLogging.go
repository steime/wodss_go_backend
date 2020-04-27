// Error Logger with similar output format as logging middleware
package util

import (
	"log"
	"os"
	"time"
)

func LogError(errString string, method string, uri string, proto string, httpCode string) {
	f,err := os.OpenFile("errorLog.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY,0644)
	if err != nil {
		log.Print(err)
	} else {
		t := time.Now()
		timeStamp := t.Format("02/Jan/2006:15:04:05 -0700")
		errString = "[" + timeStamp + "] \"" +method+ " "+ uri + " " + proto + "\" " + httpCode + " " + errString + " \n"
		if _, err = f.WriteString(errString); err != nil {
			log.Print(err.Error())
			return
		} else {
			if err = f.Close(); err != nil {
				log.Print(err.Error())
				return
			}
		}
	}
}
