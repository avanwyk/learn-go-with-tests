package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3
const finishedWord = "Go!"

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct {
	duration time.Duration
	sleep    func(duration time.Duration)
}

func (d *DefaultSleeper) Sleep() {
	d.sleep(d.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprintln(out, finishedWord)
}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{1 * time.Second, time.Sleep})
}
