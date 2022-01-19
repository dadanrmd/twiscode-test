package NIKDataHelper

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func ExtractBirthDate(nik string) (*time.Time, error) {
	if len(nik) < 16 {
		return nil, errors.New("invalid nik length")
	}

	yearStr := nik[10:12]
	monthStr := nik[8:10]
	dayStr := nik[6:8]

	//check if int of dayStr is more than 31, than minus it by 40 (female birthdate case)
	dayInt, err := strconv.Atoi(dayStr)
	if err != nil {
		return nil, err
	}

	if dayInt > 31 {
		dayStr = strconv.Itoa(dayInt - 40)
		dayStr = fmt.Sprintf("%02s", dayStr)
	}

	birthDate, err := time.Parse("06-01-02", fmt.Sprintf("%s-%s-%s", yearStr, monthStr, dayStr))
	if err != nil {
		return nil, err
	}

	//check if birthdate year is more than 40 from NIK, then means that
	//1940 instead of 2040, so subtract by 100 years
	if birthDate.Year() >= 2040 {
		birthDate = birthDate.AddDate(-100, 0, 0)
	}

	return &birthDate, nil
}

func GetBirthDateByNIK(nik string) (string, error) {
	if len(nik) < 16 {
		return "", errors.New("incorrect RequestedNIK")
	}

	yearStr := nik[10:12]
	monthStr := nik[8:10]
	dayStr := nik[6:8]

	//check if int of dayStr is more than 31, than minus it by 40 (female birthdate case)
	dayInt, err := strconv.Atoi(dayStr)
	if err != nil {
		return "", err
	}
	if dayInt > 31 {
		dayInt -= 40
		dayStr = strconv.Itoa(dayInt)
	}
	//add 0 padding for 1 digit month or day
	if dayInt < 10 {
		dayStr = fmt.Sprintf("%02s", dayStr)
	}

	yearInt, err := strconv.Atoi(yearStr)
	if err != nil {
		return "", err
	}

	yearAddition := 1900
	if yearInt < 15 {
		yearAddition = 2000
	}
	yearInt += yearAddition

	yearStr = strconv.Itoa(yearInt)

	birthDate, err := time.Parse("2006-01-02", fmt.Sprintf("%s-%s-%s", yearStr, monthStr, dayStr))
	if err != nil {
		return "", err
	}

	return birthDate.Format("20060102"), nil
}
