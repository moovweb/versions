package versions

import(
	"os"
	"strings"
	"strconv"
	)

type Version struct{
	Major *int
	Minor *int
	Patch *int
}

func NewVersion(rawVersions string) (*Version, os.Error) {
	rawVersions = strings.Trim(rawVersions, "\r\n ")

	versions := strings.Split(rawVersions, ".")
	
	if len(versions) < 2 {
		return nil, os.NewError("Invalid version string(" + rawVersions + "). Must be of the form x.x or x.x.x")
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
	Major: &majorVersion,
	Minor: &minorVersion,
	Patch: &patchVersion,
	}, nil
}

func GetVersion(fullName string) (*Version, os.Error){
	
	segments := strings.Split(fullName, "-")

	if len(segments) < 2 {
		return nil, os.NewError("Invalid fullname. No version suffix found")
	}
	
	versions := strings.SplitN(segments[len(segments)-1], ".", 4)

	if len(versions) < 3 {
		return nil, os.NewError("Invalid fullname. No minor version or patch version found")
	}	
		
	return NewVersion(strings.Join(versions[:3], "."))
}

func (v *Version) Matches(pattern string) (bool) {
	if len(pattern) == 0 {
		return true
	}
	
	// TODO(SJ) : Put actual version pattern matching here
	
	return false
}

func (v *Version) NewerThan(otherVersion *Version) (bool) {
	if *v.Major > *otherVersion.Major {
		return true
	} else if *v.Major == *otherVersion.Major {
		if *v.Minor > *otherVersion.Minor {
			return true
		} else if *v.Minor == *otherVersion.Minor {
			if *v.Patch > *otherVersion.Patch {
				return true
			}
		}
	}
	
	return false
}

func (v *Version) Equals(otherVersion *Version) (bool) {
	return *v.Major == *otherVersion.Major && *v.Minor == *otherVersion.Minor && *v.Patch == *otherVersion.Patch
}

func (v *Version) String() (string) {
	var output string

	output += strconv.Itoa(*v.Major)
	output += "." + strconv.Itoa(*v.Minor)

	if v.Patch != nil {
		output += "." + strconv.Itoa(*v.Patch)
	}

	return output
}