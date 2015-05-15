package timberchop

import (
	//"fmt"
	"github.com/gemsi/grok"
)

type GrokYaml struct {
	Description string
	Owner       string
	Tests       map[string]GrokTest
}

type GrokTest struct {
	Compare string
	Input   string
	Result  map[string]string
}

func testPattern(patterns string, gtest GrokTest) (res map[string]string, err error) {
	g := grok.New()
	g.AddPatternsFromPath(patterns)
	values, _ := g.Parse(gtest.Compare, gtest.Input)
	//fmt.Printf("COMPARE:%v || INPUT: %v || %v\n", gtest.Compare, gtest.Input, values)
	return values, nil
}
