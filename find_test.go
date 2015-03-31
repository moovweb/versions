package versions

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"runtime"
	"strings"
	"testing"
)

const TestFilesPath = "test/files"

func TestFind(t *testing.T) {
	if runtime.GOOS == "windows" {
		return
	}
	rawTests, err := ioutil.ReadFile("test/find_tests.yml")

	if err != nil {
		t.Errorf("Couldn't read tests:\n%v", err.Error())
	}

	tests := make([]interface{}, 0)

	yaml.Unmarshal(rawTests, &tests)

	globalPass := true
	errors := make([]string, 0)

	for _, testInfo := range tests {
		pass := true
		var error string
		test, err := LoadTest(testInfo)

		if err != nil {
			pass = false
			error = err.Error()
		}

		if pass {
			pass, error = test.Run()
		}

		if pass {
			print(".")
		} else {
			print("F")
			globalPass = false
			errors = append(errors, error)
		}
	}

	println("")

	if globalPass {
		fmt.Printf("All (%v) tests passed!\n", len(tests))
	} else {
		t.Errorf("Some tests failed! See error output:\n\n%v\n", strings.Join(errors, "\n\n"))
	}

}

type Test struct {
	Name           string
	Description    string
	Input          map[interface{}]interface{}
	ExpectedOutput []interface{}
}

func (t *Test) Run() (pass bool, error string) {
	paths := make([]*FilePath, 0)

	if t.Input["version"] == nil || len(t.Input["version"].(string)) == 0 {
		path, err := FindByName(TestFilesPath, t.Input["name"].(string))

		if err != nil {
			return false, err.Error()
		}

		paths = append(paths, path)
	} else {
		newPaths, err := FindByNameAndVersion(TestFilesPath, t.Input["name"].(string), t.Input["version"].(string))

		if err != nil {
			return false, err.Error()
		}

		paths = newPaths
	}

	return t.Validate(paths)
}

func (t *Test) Validate(output []*FilePath) (pass bool, error string) {
	pass = true
	errors := ""

	if len(t.ExpectedOutput) != len(output) {
		errors = fmt.Sprintf("%v\n==========\n%v\n----------\nInput:\t\tname : (%v), version : (%v)\n\nMismatched output! Differing number of output paths.\n\nExpected:\t%v\nGot:\t\t%v\n", t.Name, t.Description, t.Input["name"].(string), t.Input["version"].(string), t.ExpectedOutput, output)

		return false, errors
	}

	for index, rawExpectedPath := range t.ExpectedOutput {
		expectedPath := rawExpectedPath.(string)
		resultFilePath := output[index]
		resultPath := resultFilePath.Path

		if expectedPath != resultPath {
			pass = false
			errors += fmt.Sprintf("%v\n==========\n%v\n----------\nInput:\t\tname : (%v), version : (%v)\nExpected:\t%v\nGot:\t\t%v\n", t.Name, t.Description, t.Input["name"].(string), t.Input["version"].(string), expectedPath, resultPath)
		}
	}

	return pass, errors
}

func LoadTest(rawTest interface{}) (*Test, error) {
	test := rawTest.(map[interface{}]interface{})

	input := make(map[interface{}]interface{})
	input = test["input"].(map[interface{}]interface{})

	output := make([]interface{}, 0)

	if test["output"] != nil {
		output = test["output"].([]interface{})
	}

	return &Test{
		Name:           test["name"].(string),
		Description:    test["description"].(string),
		Input:          input,
		ExpectedOutput: output,
	}, nil
}
