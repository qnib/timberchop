package timberchop

import (
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
			testFile := path.Join(path.Join(basePath, "tests"), f.Name())
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

			for testname, gtest := range config.Tests {
				testVals, _ := testPattern(path.Join(basePath, "patterns"), gtest)
				for expKey, expVal := range gtest.Result {
					val, ok := testVals[expKey]
					if ok {
						if val != expVal {
							t.Errorf("%v >> Exp:%v != Test:%v\n", testname, expVal, val)
						}
					}
				}

			}
		}
	}
}
