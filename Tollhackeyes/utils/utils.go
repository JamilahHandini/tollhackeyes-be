package utils

import (
	"regexp"
	"strings"
	"time"
)

// * Validasi untuk input data berat
func ValidateInputData(tanggal string, min int, max int) (bool, string) {
	res := ValidateDateFormat(tanggal)
	if res == false {
		return false, "Error: Format tanggal tidak sesuai [yyyy-mm-dd]"
	}

	if min == 0 {
		return false, "Error: Nilai minimal harus diisi"
	}

	if max == 0 {
		return false, "Error: Nilai maximal harus diisi"
	}

	if min >= max {
		return false, "Error: Nilai minimal harus kurang dari / sama dengan maximal"
	}

	return true, ""
}

// * Validasi untuk input string alphanumeric
func ValidateString(str string) bool {
	//* String harus alphanumeric
	match, _ := regexp.MatchString("^[a-zA-Z0-9]+$", str)
	return match
}

// * Match antara dua string
func MatchString(str1 string, str2 string) bool {
	//* Cek apakah str1 dan str 2 sama
	match := strings.EqualFold(str1, str2)
	return match
}

// * Validasi string dengan format date
func ValidateDateFormat(dateString string) bool {
	layout := "2006-01-02" // Format yang diharapkan: yyyy-mm-dd

	_, err := time.Parse(layout, dateString)
	if err != nil {
		return false // Format tanggal tidak sesuai
	}
	return true // Format tanggal sesuai
}
