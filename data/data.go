package data

import (
	"time"

	"github.com/go-faker/faker/v4"
)

type Record struct {
	ID    string `json:"id" faker:"uuid_digit"`
	Name  string `json:"name" faker:"name"`
	Email string `json:"email" faker:"email"`
}

func getRecord() Record {
	record := Record{}
	faker.FakeData(&record)
	return record
}

func Generator(maxItems uint) <-chan Record {
	channel := make(chan Record)
	if maxItems > 0 {
		go func() {
			defer close(channel)
			for i := 0; i < int(maxItems); i++ {
				channel <- getRecord()
				time.Sleep(time.Millisecond * time.Duration(100))
			}
		}()
	} else {
		go func() {
			for {
				channel <- getRecord()
				time.Sleep(time.Millisecond * time.Duration(100))
			}
		}()
	}
	return channel
}
