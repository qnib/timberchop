package timberchop

import (
	"fmt"
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

func testPattern(patternsDir string, gtest GrokTest) (res map[string]string, err error) {
	g := grok.New()
	//fmt.Printf("\nPATTERNS: %v || COMPARE:%v || INPUT: %v \n", patternsDir, gtest.Compare, gtest.Input)
	g.AddPatternsFromPath(patternsDir)
	values, _ := g.Parse(gtest.Compare, gtest.Input)
	if values == nil {
		fmt.Printf("\nPATTERNS: %v || COMPARE:%v || INPUT: %v \n", patternsDir, gtest.Compare, gtest.Input)
		fmt.Printf("OUTPUT:%v\n", values)
	}
	//fmt.Printf("OUTPUT:%v\n", values)
	return values, nil
}
