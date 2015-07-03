# batch
Package batch runs many funcs asynchronously with multiple workers

```go

    workers := 5
    
	batch := New(workers, func(err error) {
		log.Println(err.Error())
	})
	
	batch.Start()

	for i := 1; i <= 10; i++ {

		fn := func(i int) func() error {
			return func() error {
			    err := SomeJob(i)
			    if err != nil {
			        return err
			    }
			    return nil
			}
		}
		batch.Add(fn(i))
	}
	batch.Close()

```