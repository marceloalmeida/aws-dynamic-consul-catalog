package config

import (
	"fmt"
	"log"
	"strings"
)

// Convert a CLI string slice into key/value
func ProcessFilters(userFilters []string) Filters {
	results := make(Filters)

	for _, filter := range userFilters {
		split := strings.Split(filter, "=")
		if len(split) != 2 {
			log.Fatal("Invalid filter, must be key=value format")
		}

		name := split[0]

		if _, ok := results[name]; ok {
			results[name] = fmt.Sprintf("%s,%s", results[name], split[1])
			continue
		}

		results[name] = split[1]
	}

	return results
}
