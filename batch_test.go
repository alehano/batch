package batch

import (
	"errors"
	"fmt"
	"log"
	"testing"
	"time"
)

func job(text string) {
	time.Sleep(time.Millisecond * 100)
	log.Println(text)
}

func TestBatch(t *testing.T) {

	batch := New(5, func(err error) {
		log.Println(err.Error())
	})
	batch.Start()

	for i := 1; i <= 10; i++ {
		batch.Add(func() error {
			job(fmt.Sprintf("Job #: %v", i))

			if i == 9 {
				return errors.New("Intentional error")
			}
			return nil
		})

	}

	batch.Close()
}
