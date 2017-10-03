package util

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"reflect"

	"gopkg.in/mgo.v2/bson"
)

func GetEnv(environmentVarName, defaultValue string) string {
	value := os.Getenv(environmentVarName)
	if value != "" {
		return value
	}

	return defaultValue
}

func RandomString() string {
	return bson.NewObjectId().Hex()
}

func RandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func StringToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return i
}

func Int64ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}

func WriteStringToFile(contents string, fileLocation string) error {

	fi, err := os.Create(fileLocation)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fi.Name(), []byte(contents), 0644)
	return err
}

func ReadStringFromFile(fileLocation string) (string, error) {
	fi, err := os.Open(fileLocation)
	if err != nil {
		return "", err
	}

	byts, err := ioutil.ReadFile(fi.Name())
	if err != nil {
		return "", err
	}

	return string(byts), nil
}

func TrimInnerSpace(val string) string {
	re := regexp.MustCompile("[\n\r]")
	re2 := regexp.MustCompile("[\t]")

	return re2.ReplaceAllString(re.ReplaceAllString(val, ""), " ")
}

func FirstOrNil(t interface{}) interface{} {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)

		if s.Len() < 1 {
			return nil
		}

		return s.Index(0).Interface()
	}

	return nil
}

func DbDialect(connectionString string) (string, error) {
	if len(connectionString) == 0 {
		return "", errors.New("Must have db connection string")
	}

	tmp := strings.Split(connectionString, ":")
	if len(tmp) < 2 {
		return "", errors.New("Malformed db connection string")
	}

	dialect := tmp[0]

	if dialect == "postgres" {
		return "postgres", nil
	} else if dialect == "root" {
		return "mysql", nil
	}

	return "", errors.New("Unable to parse db connection string")
}
