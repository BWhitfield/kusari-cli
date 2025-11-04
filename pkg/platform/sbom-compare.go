package platform

import (
	"fmt"
	"strconv"

	"github.com/kusaridev/kusari-cli/pkg/auth"
	"github.com/kusaridev/kusari-cli/pkg/url"
)

func SbomCompare(tenantUrl string, outputFormat string, softwareId int, sbomidA int, sbomidB int) error {

	//parth.api.dev.kusari.cloud/app/pico/v1/software/131/sbom/131/contents

	u0, e := url.Build(tenantUrl, "pico/v1/software/", strconv.Itoa(softwareId), "sbom", strconv.Itoa(sbomidA), "contents")
	if e != nil {
		return e
	}

	u1, e := url.Build(tenantUrl, "pico/v1/software/", strconv.Itoa(softwareId), "sbom", strconv.Itoa(sbomidB), "contents")
	if e != nil {
		return e
	}

	fmt.Printf("urlA: %s\n", *u0)
	fmt.Printf("urlB: %s\n", *u1)

	token, err := auth.LoadToken("kusari")
	if err != nil {
		return fmt.Errorf("failed to load auth token: %w", err)
	}

	err = auth.CheckTokenExpiry(token)
	if err != nil {
		return err
	}

	accessToken := token.AccessToken

	var packageListA SbomPackages
	var packageListB SbomPackages
	e = get(*u0, accessToken, &packageListA)
	if e != nil {
		return e
	}
	e = get(*u1, accessToken, &packageListB)
	if e != nil {
		return e
	}

	// fmt.Printf("Package List A: %+v\n", packageListA)
	// fmt.Printf("Package List B: %+v\n", packageListB)

	diff := ComparePackages(packageListA, packageListB)

	fmt.Printf("--- Comparison Results ---\n")
	fmt.Printf("Added (%d):\n", len(diff.Added))
	for _, p := range diff.Added {
		fmt.Printf("  + %s\n", p.Key())
	}

	fmt.Printf("\nRemoved (%d):\n", len(diff.Removed))
	for _, p := range diff.Removed {
		fmt.Printf("  - %s\n", p.Key())
	}

	fmt.Printf("\nSame (%d):\n", len(diff.Same))
	for _, p := range diff.Same {
		fmt.Printf("  = %s\n", p.Key())
	}

	return nil
}
