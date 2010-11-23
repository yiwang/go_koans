package koans

import "runtime"
import "os"

var Int__ int = -10000

type T struct {
	errors string
	failed bool
	ch     chan *T
}

func (t *T) Fail() { t.failed = true }

func (t *T) FailNow() {
	t.Fail()
	t.ch <- t
	runtime.Goexit()
}

func (t *T) AssertTrue(value bool) {
	if value == false {
		t.FailNow()
	}
}

func (t *T) AssertEqualInt(expected int, received int) {
	if expected != received {
		t.FailNow()
	}
}

func (t *T) AssertTrueWithMessage(value bool, message string) {
	if value == false {
		t.FailNow()
	}
}

func (t *T) Failed() bool { return t.failed }

type Test struct {
	Name string
	F    func(*T)
}

func tRunner(t *T, test *Test) {
	test.F(t)
	t.ch <- t
}

func Main(tests []Test) {
	ok := true
	for i := 0; i < len(tests); i++ {
		t := new(T)
		t.ch = make(chan *T)
		go tRunner(t, &tests[i])
		<-t.ch
		if t.failed {
			println(tests[i].Name, " has damaged your karma.")
			println("Please keep meditating")
			println(t.errors)
			ok = false
			break
		} else {
			println(tests[i].Name, " has expanded your awareness.")
			println(t.errors)
		}
	}
	if !ok {
		println("You have not yet reached enlightnment ...")
		os.Exit(1)
	}
	println("Satori!!!")
}