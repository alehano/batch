package batch

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

func job(text string) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
	log.Println(text)
}

func TestBatch(t *testing.T) {

	batch := New(2, func(err error) {
		log.Println("ERROR: " + err.Error())
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
