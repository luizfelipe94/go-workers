package send

import (
	"fmt"
	"time"

	"github.com/luizfelipe94/worker-poc/data"
)

func Send(workerId int, record data.Record) {
	time.Sleep(time.Duration(3000) * time.Millisecond)
	fmt.Println("worker id", workerId, "message sent", record.ID)
}
