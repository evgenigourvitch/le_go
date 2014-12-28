package le_go

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestConnectOpensConnection(t *testing.T) {
	le, err := Connect("")

	if err != nil {
		t.Error(err)
	}

	if le.conn == nil {
		t.Fail()
	}

	if le.isOpenConnection() == false {
		t.Fail()
	}
}

func TestConnectSetsToken(t *testing.T) {
	le, _ := Connect("myToken")

	if le.token != "myToken" {
		t.Fail()
	}
}

func TestCloseClosesConnection(t *testing.T) {
	le, _ := Connect("")

	le.Close()

	if le.isOpenConnection() == true {
		t.Fail()
	}
}

func TestOpenConnectionOpensConnection(t *testing.T) {
	le, _ := Connect("")

	le.Close()

	le.openConnection()

	if le.isOpenConnection() == false {
		t.Fail()
	}
}

func TestEnsureOpenConnectionDoesNothingOnOpenConnection(t *testing.T) {
	le, _ := Connect("")

	old := &le.conn

	le.openConnection()

	if old != &le.conn {
		t.Fail()
	}
}

func TestEnsureOpenConnectionCreatesNewConnection(t *testing.T) {
	le, _ := Connect("")

	le.Close()

	le.openConnection()

	if le.isOpenConnection() == false {
		t.Fail()
	}
}

func TestFlagsReturnsFlag(t *testing.T) {
	le := Logger{flag: 2}

	if le.Flags() != 2 {
		t.Fail()
	}
}

func TestSetFlagsSetsFlag(t *testing.T) {
	le := Logger{flag: 2}

	le.SetFlags(1)

	if le.flag != 1 {
		t.Fail()
	}
}

func TestPrefixReturnsPrefix(t *testing.T) {
	le := Logger{prefix: "myPrefix"}

	if le.Prefix() != "myPrefix" {
		t.Fail()
	}
}

func TestSetPrefixSetsPrefix(t *testing.T) {
	le := Logger{prefix: "myPrefix"}

	le.SetPrefix("myNewPrefix")

	if le.prefix != "myNewPrefix" {
		t.Fail()
	}
}

func TestLoggerImplementsWriterInterface(t *testing.T) {
	le, _ := Connect("myToken")

	// the test will fail to compile if Logger doesn't implement io.Writer
	func(w io.Writer) {}(le)
}

func TestReplaceNewline(t *testing.T) {
	le, _ := Connect("myToken")
	le.Println("1\n2\n3")

	if strings.Count(string(le.buf), "\u2028") != 2 {
		t.Fail()
	}
}

func TestAddNewline(t *testing.T) {
	le, _ := Connect("myToken")
	le.Print("123")

	if !strings.HasSuffix(string(le.buf), "\n") {
		t.Fail()
	}

	le.Printf("%s", "123")

	if !strings.HasSuffix(string(le.buf), "\n") {
		t.Fail()
	}
}

func ExampleLogger() {
	le, err := Connect("XXXX-XXXX-XXXX-XXXX") // replace with token
	if err != nil {
		panic(err)
	}

	defer le.Close()

	le.Println("another test message")
}

func ExampleLogger_write() {
	le, err := Connect("XXXX-XXXX-XXXX-XXXX") // replace with token
	if err != nil {
		panic(err)
	}

	defer le.Close()

	fmt.Fprintln(le, "another test message")
}
