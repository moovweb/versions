package versions

import (
	"errors"
	"strconv"
	"strings"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func NewVersion(rawVersions string) (*Version, error) {
	rawVersions = strings.Trim(rawVersions, "\r\n ")
	rawVersions = trimExtension(rawVersions)

	versions := strings.Split(rawVersions, ".")

	if len(versions) < 2 {
		return nil, errors.New("Invalid version string(" + rawVersions + "). Must be of the form x.x or x.x.x")
	}

	majorVersion, err := strconv.Atoi(versions[0])

	if err != nil {
		return nil, err
	}

	minorVersion, err := strconv.Atoi(versions[1])

	if err != nil {
		return nil, err
	}

	var patchVersion int

	if len(versions) == 3 {
		patchVersion, err = strconv.Atoi(versions[2])

		if err != nil {
			return nil, err
		}
	}

	return &Version{
		Major: majorVersion,
		Minor: minorVersion,
		Patch: patchVersion,
	}, nil
}

func GetVersion(fullName string) (*Version, error) {

	segments := strings.Split(fullName, "-")

	if len(segments) < 2 {
		return nil, errors.New("Invalid fullname. No version suffix found")
	}

	versions := strings.SplitN(segments[len(segments)-1], ".", 4)

	if len(versions) < 3 {
		return nil, errors.New("Invalid fullname. No minor version or patch version found")
	}

	return NewVersion(strings.Join(versions[:3], "."))
}

// TODO(SJ) : Support a slice of patterns
func (v *Version) Matches(pattern string) (match bool, err error) {
	if len(pattern) == 0 {
		return true, nil
	}

	p, err := NewPattern(pattern)

	if err != nil {
		return false, err
	}

	match = p.Match(v)

	return
}

func (v *Version) String() string {
	var output string

	output += strconv.Itoa(v.Major)
	output += "." + strconv.Itoa(v.Minor)

	// Patch level is auto-populated for now
	//	if v.Patch != nil {
	output += "." + strconv.Itoa(v.Patch)
	//	}

	return output
}
