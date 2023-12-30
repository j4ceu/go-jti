package utils

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

func generateRandomNumberWithProvider(prefixes map[string][]int, length int) (string, string) {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	provider := chooseRandomProvider(prefixes)                                     // Choose a random provider
	prefix := strconv.Itoa(prefixes[provider][rand.Intn(len(prefixes[provider]))]) // Choose a random prefix from the selected provider

	rangeStart := int(math.Pow10(length - 1))
	rangeEnd := int(math.Pow10(length)) - 1
	number := strconv.Itoa(rand.Intn(rangeEnd-rangeStart+1) + rangeStart) // Generate random number within the specified length

	intl := "0"

	return intl + prefix + number, provider // Return both the generated number and the provider
}

func chooseRandomProvider(prefixes map[string][]int) string {
	providers := make([]string, 0, len(prefixes)) // Get a list of providers
	for provider := range prefixes {
		providers = append(providers, provider)
	}
	return providers[rand.Intn(len(providers))] // Choose a random provider
}

func GenerateRandomNumber() (string, string) {
	prefixes := map[string][]int{
		"Telkomsel": {811, 812, 813, 821, 822, 823, 852, 853, 851},
		"Axis":      {838, 831, 832, 833},
		"XL":        {817, 818, 819, 859, 877, 878},
		"Indosat":   {814, 815, 816, 855, 856, 857, 858},
		"3":         {895, 896, 897, 898, 899},
		"Smartfren": {881, 882, 884, 885, 886, 887, 888, 889},
	}

	randomNumber, provider := generateRandomNumberWithProvider(prefixes, 8)
	return randomNumber, provider
}
