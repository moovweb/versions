package versions

import(
	"os"
	"path/filepath"
	)

type FilePath struct{
	Path string
	Name string
	Version *Version
}


func FindByName(rootPath string, name string) (*FilePath, os.Error){
	filePaths, err := FindByNameAndVersion(rootPath, name, "")
	
	if err != nil {
		return nil, err
	}

	if len(filePaths) == 0 {
		return nil, os.NewError("Found no files named " + name + " in path: " + rootPath)
	}

	newestFilePath := filePaths[0]
	

	for _, thisFilePath := range(filePaths) {
		pattern, _ := NewPattern("> " + newestFilePath.Version.String() )

		if pattern.Match(thisFilePath.Version) {
			newestFilePath = thisFilePath
		}
	}

	return newestFilePath, nil
}

func FindByNameAndVersion(rootPath string, name string, versionPattern string) (paths []*FilePath, err os.Error) {

	query := filepath.Join(rootPath, name) + "*"
	results, _ := filepath.Glob(query)
	
	for _, result := range(results) {
		version, err := GetVersion(result)

		if err != nil {
			println("Didnt get valid version from:" + result)
			println(err.String())
			continue
		}

		matched, err := version.Matches(versionPattern)
		
		if err != nil {
			return nil, err
		}

		if matched {
			matchingPath := &FilePath{
			Path: result,
			Name: name,
			Version: version,
			}
			paths = append(paths, matchingPath)
		}
	}

	return
}

func (fp *FilePath) String() string {
	// TODO(SJ) -- support any extension and save the actual filename!
	return fp.Path
}