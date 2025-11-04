package platform

import "fmt"

type ComparisonResult struct {
	Added   []Package
	Removed []Package
	Same    []Package
}

func (p Package) Key() string {
	return fmt.Sprintf("%s::%s::%s::%s", p.Type, p.Namespace, p.Name, p.Version)
}

// ComparePackages takes two SbomPackages instances (previous and current) and returns the difference.
func ComparePackages(prev SbomPackages, curr SbomPackages) ComparisonResult {
	result := ComparisonResult{}

	// Create a map of keys for the PREVIOUS packages for fast lookups.
	prevMap := make(map[string]Package)
	for _, pkg := range prev.Contents {
		prevMap[pkg.Key()] = pkg
	}

	// Determine Added and Same packages by iterating through CURRENT.
	currMap := make(map[string]Package) // We also build a map for the current for the next step.
	for _, currPkg := range curr.Contents {
		key := currPkg.Key()
		currMap[key] = currPkg // Build map for next step

		if _, found := prevMap[key]; found {
			// Found in both
			result.Same = append(result.Same, currPkg)
		} else {
			// Found only in current (it was added)
			result.Added = append(result.Added, currPkg)
		}
	}

	// Determine Removed packages by iterating through PREVIOUS.
	for _, prevPkg := range prev.Contents {
		key := prevPkg.Key()
		if _, found := currMap[key]; !found {
			// Found only in previous (it was removed)
			result.Removed = append(result.Removed, prevPkg)
		}
	}

	return result
}
