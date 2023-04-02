package util

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

func stringInSlice(str string, list []string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}
	return false
}

func ExtractSID(cookie string) string {
	var apiSIDHash string

	validCookieNames := []string{"SAPISID", "__Secure-1PAPISID", "__Secure-3PAPISID"}

	for _, cookies := range strings.Split(cookie, "; ") {
		cookieParts := strings.SplitN(cookies, "=", 2)
		if stringInSlice(cookieParts[0], validCookieNames) {
			// Calculate SHA-1 hash
			hasher := sha1.New()
			now := time.Now().Unix()
			hasher.Write([]byte(fmt.Sprintf("%d %s %s", now, cookieParts[1], "https://voice.google.com")))
			shaHash := hex.EncodeToString(hasher.Sum(nil))
			apiSIDHash = fmt.Sprintf("%sHASH %d_%s", cookieParts[0], now, shaHash)
			break
		}
	}

	return apiSIDHash
}
