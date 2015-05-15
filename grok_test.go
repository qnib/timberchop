package timberchop

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
)

func TestPatterns(t *testing.T) {
	basePath := os.Getenv("GROK_BASE")
	if basePath == "" {
		basePath = "/etc/qnib/grok/"
	}
	files, _ := ioutil.ReadDir(path.Join(basePath, "tests"))
	for _, f := range files {
		if strings.HasSuffix(f.Name(), "yml") {
			testFile := path.Join(basePath, "tests", f.Name())
			fmt.Printf("> Using YAML test %s... ", f.Name())
			filename, _ := filepath.Abs(testFile)
			yamlFile, err := ioutil.ReadFile(filename)
			if err != nil {
				panic(err)
			}

			var config GrokYaml

			err = yaml.Unmarshal(yamlFile, &config)
			if err != nil {
				panic(err)
			}

			all_ok := true
			first_test := true
			for testname, gtest := range config.Tests {
				if first_test {
					fmt.Printf("%s", testname)
					first_test = false
				} else {
					fmt.Printf(", %s", testname)
				}
				testVals, _ := testPattern(path.Join(basePath, "patterns"), gtest)
				for expKey, expVal := range gtest.Result {
					val, ok := testVals[expKey]
					if ok {
						if val != expVal {
							all_ok = false
							t.Errorf("Exp:%v != Test:%v\n", expVal, val)
						}
					} else {
						all_ok = false
						t.Errorf("%s:%v not in result\n", expKey, expVal)
					}
				}
			}
			if all_ok {
				fmt.Printf(" [OK]\n")
			} else {
				fmt.Printf(" [NOK]\n")
			}
		}
	}
}
