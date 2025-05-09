package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//easyjson:json
type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	domain = "." + strings.ToLower(domain)
	var user User
	result := make(DomainStat)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()
		if err := user.UnmarshalJSON(line); err != nil {
			return nil, fmt.Errorf("failed to unmarshal user: %w", err)
		}

		email := strings.ToLower(user.Email)
		if strings.HasSuffix(email, domain) {
			parts := strings.SplitN(email, "@", 2)
			if len(parts) == 2 {
				result[parts[1]]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan failed: %w", err)
	}

	return result, nil
}
