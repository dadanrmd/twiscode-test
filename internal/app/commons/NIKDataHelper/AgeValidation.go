package NIKDataHelper

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

const (
	GenderMale   = 1
	GenderFemale = 2
)

func ValidateAge(nik string, ageMin, ageMax int) (isAgeValid bool, err error) {
	if len(nik) < 16 {
		return false, errors.New("invalid nik length")
	}

	yearStr := nik[10:12]
	monthStr := nik[8:10]
	dayStr := nik[6:8]

	//check if int of dayStr is more than 31, than minus it by 40 (female birthdate case)
	dayInt, err := strconv.Atoi(dayStr)
	if err != nil {
		return false, err
	}

	if dayInt > 31 {
		dayStr = strconv.Itoa(dayInt - 40)
		dayStr = fmt.Sprintf("%02s", dayStr)
	}

	birthDate, err := time.Parse("06-01-02", fmt.Sprintf("%s-%s-%s", yearStr, monthStr, dayStr))
	if err != nil {
		return false, err
	}

	//check if birthdate year is more than 40 from NIK, then means that
	//1940 instead of 2040, so subtract by 100 years
	if birthDate.Year() >= 2040 {
		birthDate = birthDate.AddDate(-100, 0, 0)
	}

	currentTime := time.Now()
	age := currentTime.Year() - birthDate.Year()

	//check if birthdate month already passed in this year
	if currentTime.Month() < birthDate.Month() {
		age--
	}

	//age is less than 22 y.o or more than 55
	if age < ageMin || age > ageMax {
		return false, nil
	}

	return true, nil
}

func GetAgeFromNIK(nik string) (int, error) {
	if len(nik) < 16 {
		return 0, errors.New("invalid nik length")
	}

	yearStr := nik[10:12]
	monthStr := nik[8:10]
	dayStr := nik[6:8]

	//check if int of dayStr is more than 31, than minus it by 40 (female birthdate case)
	dayInt, err := strconv.Atoi(dayStr)
	if err != nil {
		return 0, err
	}

	if dayInt > 31 {
		dayStr = strconv.Itoa(dayInt - 40)
		dayStr = fmt.Sprintf("%02s", dayStr)
	}

	birthDate, err := time.Parse("06-01-02", fmt.Sprintf("%s-%s-%s", yearStr, monthStr, dayStr))
	if err != nil {
		return 0, err
	}

	//check if birthdate year is more than 40 from NIK, then means that
	//1940 instead of 2040, so subtract by 100 years
	if birthDate.Year() >= 2040 {
		birthDate = birthDate.AddDate(-100, 0, 0)
	}

	currentTime := time.Now()
	age := currentTime.Year() - birthDate.Year()

	//check if birthdate month already passed in this year
	if currentTime.Month() < birthDate.Month() {
		age--
	}

	return age, nil
}

func GetAgeAt(nik string, at time.Time) (int, error) {
	if len(nik) < 16 {
		return 0, errors.New("invalid nik length")
	}

	yearStr := nik[10:12]
	monthStr := nik[8:10]
	dayStr := nik[6:8]

	//check if int of dayStr is more than 31, than minus it by 40 (female birthdate case)
	dayInt, err := strconv.Atoi(dayStr)
	if err != nil {
		return 0, err
	}

	if dayInt > 31 {
		dayStr = strconv.Itoa(dayInt - 40)
		dayStr = fmt.Sprintf("%02s", dayStr)
	}

	currentDate := at

	//get year
	custBornYear, err := strconv.Atoi(yearStr)
	if err != nil {
		return 0, err
	}

	custBornMo, err := strconv.Atoi(monthStr)
	if err != nil {
		return 0, err
	}

	custBornDay, err := strconv.Atoi(dayStr)
	if err != nil {
		return 0, err
	}

	ageYear := currentDate.Year() - custBornYear
	ageMonth := int(currentDate.Month()) - custBornMo

	if ageMonth < 0 {
		ageYear = ageYear - 1
	} else if ageMonth == 0 { //this is the same month
		ageDay := currentDate.Day() - custBornDay

		if ageDay < 0 {
			ageYear = ageYear - 1
		}
	}

	return ageYear, nil
}

func GetGenderFromNIK(nik string) (int, error) {
	if len(nik) < 16 {
		return 0, errors.New("invalid nik length")
	}

	dayStr := nik[6:8]

	//check if int of dayStr is more than 31
	dayInt, err := strconv.Atoi(dayStr)
	if err != nil {
		return 0, err
	}

	if dayInt > 31 {
		return GenderFemale, nil
	}

	return GenderMale, nil
}
