package versions

import "testing"
import "exec"
import "bufio"
import "fmt"
//import "rand"
//import "time"

func evalRuby(script string, t *testing.T) string {
	// Find and create the command
	cmdpath, err := exec.LookPath("ruby")
	if err != nil {
		t.Error(err)
		t.Fatal("Failed to find ruby")
	}
	r := exec.Command(cmdpath)

	// Get the stdin and stdout pipes
	stdin, err := r.StdinPipe()
	if err != nil {
		t.Error(err)
		t.Fatal("Failed to create stdin pipe")
	}
	stdout, err := r.StdoutPipe()
	if err != nil {
		t.Error(err)
		t.Fatal("Failed to create stdout pipe")
	}

	// Start the process
	r.Start()

	// Write the script to stdin
	n, err := stdin.Write([]byte("puts " + script))
	if err != nil {
		t.Error(err)
		t.Fatal("Failed to write to stdin", n, "bytes to stdin")
	}
	stdin.Close()

	// Read the output from stdout
	outbr := bufio.NewReader(stdout)
	line, _, err := outbr.ReadLine()
	if err != nil {
		t.Error(err)
		t.Fatal("Failed to read from pipe")
	}

	// Wait for the process to terminate normally
	err = r.Wait()
	if err != nil {
		t.Error(err)
		t.Fatal("Process failed")
	}

	return string(line)
}

func checkNewVersion(version string, t *testing.T) {
	// Test the creation and marshaling of a version to/from string
	ruby := fmt.Sprintf("Gem::Version.new(\"%s\")", version)
	ruby_output := evalRuby(ruby, t)
	go_output, err := NewVersion(version)
	if err != nil {
		t.Error("Failed to create Go version", "ruby got", ruby_output)
	} else {
		if go_output.String() != ruby_output {
			t.Error("Failed to match:", "expected", ruby_output, "got", go_output.String())
		}
	}
}

func checkMatcher(version string, matcher string, t *testing.T) {
	// Test the truth of a matcher
	ruby := "Gem::Requirement.new(\""+matcher+"\").satisfied_by?Gem::Version.new(\""+version+"\")"
	ruby_output := evalRuby(ruby, t)
	go_output, err := NewVersion(version)
	if err != nil {
		t.Error("Failed to create Go version", "ruby got", ruby_output)
	} else {
		match, err := go_output.Matches(matcher)
		if err != nil {
			t.Error("Failed to execute match", "ruby got", ruby_output)
		}
		if (match == true && ruby_output == "true") || (match == false && ruby_output == "false") {
		} else {
			t.Error("Failed match", version, matcher, "Go:", match, "Ruby:", ruby_output)
		}
	}
}

func disableTestNewVersion(t *testing.T) {
	/*checkNewVersion("0.0.0", t)

	rand.Seed(time.Nanoseconds())

	for i := 0; i < 20; i++ {
		major := rand.Int() % 65535
		minor := rand.Int() % 65535
		build := rand.Int() % 65535
		version := fmt.Sprintf("%d.%d.%d", major, minor, build)
		checkNewVersion(version, t)
	}

	for i := 0; i < 20; i++ {
		major := rand.Int() % 10
		minor := rand.Int() % 10
		build := rand.Int() % 10
		version := fmt.Sprintf("%d.%d.%d", major, minor, build)
		major = rand.Int() % 10
		minor = rand.Int() % 10
		build = rand.Int() % 10
		match := fmt.Sprintf("%d.%d.%d", major, minor, build)

		checkMatcher(version, "~> " + match, t)
	}

	for i := 0; i < 20; i++ {
		major := rand.Int() % 10
		minor := rand.Int() % 10
		build := rand.Int() % 10
		version := fmt.Sprintf("%d.%d.%d", major, minor, build)
		major = rand.Int() % 10
		minor = rand.Int() % 10
		build = rand.Int() % 10
		match := fmt.Sprintf("%d.%d", major, minor)

		checkMatcher(version, "~> " + match, t)
	}*/

	/*checkMatcher("0", "~> 0", t)
	checkMatcher("1", "~> 0", t)
	checkMatcher("5", "~> 0", t)
	checkMatcher("9", "~> 0", t)*/

	//checkMatcher("0.1", "~> 0.0", t)
	//checkMatcher("0.9", "~> 0.0", t)
	/*checkMatcher("1.0", "~> 0.0", t)
	checkMatcher("1.1", "~> 0.0", t)
	checkMatcher("1.9", "~> 0.0", t)

	checkMatcher("0.0", "~> 0.0", t)
	//checkMatcher("0.1", "~> 0.0", t)
	checkMatcher("1.0", "~> 0.0", t)

	checkMatcher("0.0.0", "~> 0.0", t)
	checkMatcher("0.0.1", "~> 0.0", t)
	//checkMatcher("0.1.0", "~> 0.0", t)
	checkMatcher("1.0.0", "~> 0.0", t)

	checkMatcher("0.0.0", "~> 1.1", t)
	checkMatcher("0.0.1", "~> 1.1", t)
	checkMatcher("0.1.0", "~> 1.1", t)
	checkMatcher("1.0.0", "~> 1.1", t)*/

	// Nice to have
	//checkNewVersion("0.0.src", t)
	//checkNewVersion("0.src.0", t)
}

