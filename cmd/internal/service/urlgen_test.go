package service

import (
	"errors"
	"fmt"
	"testing"
)

var testStringLength int = 10
var testChars []rune = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_")
var getError error = errors.New("get some error")

type testUrlList struct {
	shortUrl  string
	longUrl   string
	available bool
	err       error
}

func (t *testUrlList) GetUrl(shortUrl string) (string, error) {
	return t.longUrl, t.err
}

func (t *testUrlList) GetShortUrl(longUrl string) (string, error) {
	return t.shortUrl, t.err
}

func (t *testUrlList) PostUrl(shortUrl string, longUrl string) error {
	return t.err
}

func (t *testUrlList) Available(shortUrl string) (bool, error) {
	return t.available, t.err
}

var urlList testUrlList

func TestGetUrl(t *testing.T) {
	type getTestData struct {
		shortUrl  string
		available bool
		err       bool
	}

	var getTests = []getTestData{
		{"val1", true, false},
		{"val2", true, false},
		{"val3", true, false},
		{"val3", true, true},
		{"val4", true, false},
		{"val5", true, false},
		{"gewghh^jl:/ncdsad", true, false},
		{"", true, true},
	}

	service := NewUrlListService(&urlList, testStringLength, testChars)
	for _, getTest := range getTests {
		if getTest.err {
			urlList.err = getError
		} else {
			urlList.err = nil
		}
		longUrl, err := service.GetUrl(getTest.shortUrl)

		if (err != nil) != getTest.err {
			t.Error(fmt.Sprintf("GetUrl(\"%s\") = \"%s\", %t, expected \"uniquestr\", %t", getTest.shortUrl, longUrl, err != nil, getTest.err))
		}

	}
}

func TestPostUrl(t *testing.T) {
	type postTestData struct {
		longUrl   string
		available bool
		err       bool
	}

	var postTests = []postTestData{
		{"val1", true, false},
		{"val2", true, false},
		{"val3", true, false},
		{"val5", true, false},
		{"val5", true, true},
		{"val6", true, false},
		{"val7", true, false},
		{"val8", true, false},
	}
	service := NewUrlListService(&urlList, testStringLength, testChars)

	for _, postTest := range postTests {
		if postTest.err {
			urlList.err = getError
		} else {
			urlList.err = nil
		}
		_, err := service.GetUrl(postTest.longUrl)

		if (err != nil) != postTest.err {
			t.Error(fmt.Sprintf("PostUrl(\"%s\") = \"uniquestr\", %t, expected \"uniquestr\", %t", postTest.longUrl, err != nil, postTest.err))
		}
	}
}

func TestGetUniqueString(t *testing.T) {
	uniqueUrlPair := make(map[string]bool)
	count := 100
	service := NewUrlListService(&urlList, testStringLength, testChars)

	for i := 0; i < count; i++ {
		uniqueStr := service.getUniqueString()
		if ok := uniqueUrlPair[uniqueStr]; ok {
			t.Error(fmt.Sprintf("Not unique \"%s\" #:%d", uniqueStr, i+1))
		}
		uniqueUrlPair[uniqueStr] = true
	}
}
