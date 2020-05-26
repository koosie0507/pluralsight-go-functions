package semver

import "fmt"

type SemanticVersion struct {
	major, minor, patch int
}

func NewSemanticVersion(major, minor, patch int) SemanticVersion {
	return SemanticVersion{
		major: major,
		minor: minor,
		patch: patch,
	}
}

func (sv SemanticVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)
}

func (sv *SemanticVersion) IncMajor() {
	sv.major+=1
}

func (sv *SemanticVersion) IncMinor() {
	sv.minor+=1
}

func (sv *SemanticVersion) IncPatch() {
	sv.patch+=1
}