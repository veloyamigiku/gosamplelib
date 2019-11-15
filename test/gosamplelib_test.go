// 手順0(ファイル名はxxxx_test.go)

package test

import (
	"github.com/veloyamigiku/gosamplelib"
	// 先頭のドットは、パッケージ名を省略してアクセスするためのもの。
	// 手順1
	. "gopkg.in/check.v1"
	"testing"
)

// 手順2 テストスイート（構造体）を定義する。
type GoSampleLibSuite struct {
	// テストフィクスチャの変数を定義する。
	name string
	age int
}

// 手順3 テストスイートを登録する。
func init() {
	Suite(&GoSampleLibSuite{})
}

// 手順4 パッケージtestingと統合する。
func Test(t *testing.T) {
	TestingT(t)
}

// 各テスト実行前の処理を定義する。
func (s *GoSampleLibSuite) SetUpTest(c *C) {
	// テストフィクスチャを作成する。
	s.name = "Sample"
	s.age = 10
}

// 各テスト実行後の処理を定義する。
func (s *GoSampleLibSuite) TearDownTest(c *C) {
	c.Log("Finished test - ", c.TestName())
}

// テスト開始前の処理を定義する。
func (s *GoSampleLibSuite) SetUpSuite(c *C) {
	c.Log("Starting Test Suite")
}

// テスト終了前の処理を定義する。
func (s *GoSampleLibSuite) TearDownSuite(c *C) {
	c.Log("Finished Test Suite")
}

// 手順5
// テストスイートのメソッド（Testで始まる）を定義する。
func (s *GoSampleLibSuite) TestSay(c *C) {
	person := gosamplelib.Person { s.name, s.age }
	c.Check(person.Say2(), Equals, "My name is Sample, I'm 10")
}
