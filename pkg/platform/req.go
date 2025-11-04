package platform

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func get(apiEndpoint string, jwtToken string, result any) (error) {

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", apiEndpoint, nil)
	if err != nil {
		return fmt.Errorf("Request failed with error: %w", err)
	}

	// todo make this api compatible
	req.Header.Set("Authorization", "Bearer "+jwtToken)
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Request failed with error: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		switch resp.StatusCode {
		case http.StatusUnauthorized:
			return fmt.Errorf("Not Authorized: %d", resp.StatusCode)
		case http.StatusForbidden:
			return fmt.Errorf("Forbidden: %d. Try `kusari auth login`", resp.StatusCode)
		default:
			return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body with error: %w", err)
	}


	err = json.Unmarshal(body, &result)
	if err != nil {
		return fmt.Errorf("failed to unmarshal the results with body: %s with error: %w", string(body), err)
	}

	return nil
}


