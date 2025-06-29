package types

import (
	"fmt"
	"net/url"
	"strings"
)

type CreateClientParams struct {
	Name         string   `json:"name"`
	RedirectURLs []string `json:"redirectUrls" schema:"redirectUrls"`
}

type ValidationError struct {
	Key        string
	Message    string
	Validation string
}

type ValidationErrors []ValidationError

func (c *CreateClientParams) TrimValues() {
	c.Name = strings.TrimSpace(c.Name)

	var urls []string
	for _, url := range c.RedirectURLs {
		url = strings.TrimSpace(url)
		if url != "" {
			urls = append(urls, url)
		}
	}
}

func (c *CreateClientParams) Validate() (ok bool, errors ValidationErrors) {
	c.TrimValues()

	if c.Name == "" {
		errors = append(errors, ValidationError{
			Key:        "name",
			Message:    "can't be blank",
			Validation: "required",
		})
	}

	if len(c.RedirectURLs) == 0 {
		errors = append(errors, ValidationError{
			Key:        "redirectUrls",
			Message:    "please provide at least one redirect URL.",
			Validation: "required",
		})
	}

	for i := range c.RedirectURLs {
		parsed, err := url.Parse(c.RedirectURLs[i])
		if err != nil {
			errors = append(errors, ValidationError{
				Key:        fmt.Sprintf("redirectUrls.%d", i),
				Message:    fmt.Sprintf("invalid URL: %s", err),
				Validation: "url",
			})
		}

		if parsed.Scheme != "https" || (parsed.Scheme == "http" && parsed.Hostname() != "localhost") {
			errors = append(errors, ValidationError{
				Key:        fmt.Sprintf("redirectUrls.%d", i),
				Message:    "redirect URLs must use HTTPS",
				Validation: "url",
			})
		}
	}

	return len(errors) == 0, errors
}
