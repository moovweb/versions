package versions

import "fmt"

func (p *Pattern) Less(version *Version) bool {
	if version.Major < p.Version.Major {
		return true
	} else if version.Major == p.Version.Major {
		if version.Minor < p.Version.Minor {
			return true
		} else if version.Minor == p.Version.Minor {
			if version.Patch < p.Version.Patch {
				return true
			}
		}
	}
	
	return false
}

func (p *Pattern) LessEqual(version *Version) bool {
	if p.Equal(version) {
		return true
	}

	return p.Less(version)
}

func (p *Pattern) Equal(version *Version) bool {
	return version.Major == p.Version.Major && version.Minor == p.Version.Minor && version.Patch == p.Version.Patch
}

func (p *Pattern) Pessimistic(version *Version) bool {
	// TODO(SJ) !! When versions like `0.4` are zero'd out to `0.4.0` this may not behave as expected
	
	if version.Major > p.Version.Major || version.Minor > p.Version.Minor {
		return false
	}
	
	if version.Patch < p.Version.Patch {
		return false
	}

	return true
}

func (p *Pattern) Greater(version *Version) bool {
	fmt.Printf("%v > %v ? ", version.String(), p.Version.String())


	if version.Major > p.Version.Major {
		println("yes")
		return true
	} else if version.Major == p.Version.Major {
		if version.Minor > p.Version.Minor {
			println("yes")
			return true
		} else if version.Minor == p.Version.Minor {
			if version.Patch > p.Version.Patch {
				println("yes")
				return true
			}
		}
	}

	println("nope")	
	return false
}

func (p *Pattern) GreaterEqual(version *Version) bool {
	if p.Equal(version) {
		return true
	}

	return p.Greater(version)
}