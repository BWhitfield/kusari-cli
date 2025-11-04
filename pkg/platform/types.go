package platform

type SbomPackages struct {
	Contents []Package `json:"contents"`
}

type LicenseType string
type PackageEndOfLife string

type Package struct {
	ID                int64            `json:"id"`
	Type              string           `json:"type"`
	Namespace         string           `json:"namespace"`
	Name              string           `json:"name"`
	Version           string           `json:"version"`
	GuacID            string           `json:"guac_id"`
	License           string           `json:"license"`
	DiscoveredLicense string           `json:"discovered_license"`
	LicenseType       LicenseType      `json:"license_type"`
	Purls             []string         `json:"purls"`
	VulnCount         int              `json:"vuln_count"`
	EndOfLife         PackageEndOfLife `json:"end_of_life"`
}
