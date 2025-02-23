package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/zaki031/shortner/database"
	"github.com/zaki031/shortner/models"
)

func AddUrl(w http.ResponseWriter, r *http.Request) {
	var url models.Url
	json.NewDecoder(r.Body).Decode(&url)

	err := verifyUrl(&url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = generateShortLink(&url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(url.ShortUrl)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(url.ShortUrl)

}

func verifyUrl(url *models.Url) error {
	regex := regexp.MustCompile(`^https?://`)
	newUrl := url.LongUrl
	if !regex.MatchString(newUrl) {
		newUrl = "https://" + newUrl
	}
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(newUrl)
	if err != nil {
		return err
	}
	if resp.StatusCode == http.StatusOK {
		url.LongUrl = newUrl
		return nil
	}
	return err
}

func generateShortLink(url *models.Url) error {
	hash := sha256.New()
	hash.Write([]byte(url.LongUrl))
	url.ShortUrl = hex.EncodeToString(hash.Sum(nil))[:8]

	err := addToRedis(url.LongUrl, url.ShortUrl)
	if err != nil {
		return err
	}
	return nil
}
func addToRedis(url string, shortUrl string) error {
	client := database.Client
	err := client.Set(ctx, shortUrl, url, 0).Err()
	if err != nil {
		return err
	}
	return nil
}
