// 手順0(ファイル名はxxxx_test.go)

package test

import (
	"github.com/veloyamigiku/gosamplelib"
	// 先頭のドットは、パッケージ名を省略してアクセスするためのもの。
	// 手順1
	. "gopkg.in/check.v1"
	"testing"
)

// 手順2
type GoSampleLibSuite struct {}

// 手順3
func init() {
	Suite(&GoSampleLibSuite{})
}

// 手順4
func Test(t *testing.T) {
	TestingT(t)
}

// 手順5
// Testで始まる関数を、構造体(手順2で定義した構造体)に定義する。
func (s *GoSampleLibSuite) TestSay(c *C) {
	person := gosamplelib.Person { "Sample", 10 }
	c.Check(person.Say2(), Equals, "My name is Sample, I'm 10")
}
