package versions

import(
	"testing"
	yaml "goyaml"
	"io/ioutil"
	"path/filepath"
	"strings"
	"os"
	"fmt"
	)

const TestFilesPath = "test/files"

func TestFind(t *testing.T) {
	testPath := "test/tests"
	tests, err := ioutil.ReadDir(testPath)

	if err != nil {
		t.Errorf("Couldn't read tests directory:\n%v", err.String() )
	}

	globalPass := true
	errors := make([]string,0)
	
	for _, testInfo := range(tests) {
		pass := true
		var error string
		test, err := LoadTest(filepath.Join(testPath, testInfo.Name))

		if err != nil {
			pass = false
			error = err.String()
		}

		if pass {
			pass, error = test.Run()
		}

		if pass {
			println(".")
		} else {
			println("F")
			globalPass = false
			errors = append(errors, error)
		}
	}


	if globalPass {
		println("All tests passed!")
	} else {
		t.Errorf("Some tests failed! See error output:\n\n%v\n", strings.Join(errors,"--\n"))
	}
	
}


type Test struct {
	Name string
	Input map[string]string
	ExpectedOutput []string
}

func (t *Test) Run() (pass bool, error string) {
	paths := make([]*FilePath, 0)

	if len(t.Input["versions"]) == 0 {
		path, err := FindByName(TestFilesPath, t.Input["name"])

		if err != nil {
			return false, err.String()
		}

		paths = append(paths, path)
	} else {
		newPaths, err := FindByNameAndVersion(TestFilesPath, t.Input["name"], t.Input["version"])

		if err != nil {
			return false, err.String()
		}

		paths = newPaths
	}

	return t.Validate(paths)
}

func (t *Test) Validate(output []*FilePath) (pass bool, error string){
	pass = true
	errors := ""

	for index, expectedPath := range(t.ExpectedOutput) {
		resultFilePath := output[index]
		resultPath := resultFilePath.Path

		if expectedPath != resultPath {
			pass = false
			errors += fmt.Sprintf("%v\n==========\nInput:\t\tname : (%v), version : (%v)\nExpected:\t%v\nGot:\t\t%v\n", t.Name, t.Input["name"], t.Input["version"], expectedPath, resultPath)
		}
	}

	return pass, errors
}


func LoadTest(path string) (*Test, os.Error) {
	data, err := ioutil.ReadFile( filepath.Join(path, "input.yml") )

	if err != nil {
		return nil, err
	}

	input := make( map[string]string )
	yaml.Unmarshal(data, &input)

	data, err = ioutil.ReadFile( filepath.Join(path, "output.yml") )

	if err != nil {
		return nil, err
	}

	output := make( []string, 0)
	yaml.Unmarshal(data, &output)
	
	_, name := filepath.Split(path)

	return &Test{
	Name: name,
	Input: input,
	ExpectedOutput: output,
	}, nil
}