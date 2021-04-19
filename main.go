package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/kyeett/openbanking/models"

	"github.com/go-chi/chi"

	"github.com/kyeett/openbanking/seb"

	uuid "github.com/satori/go.uuid"
)

const (

	// Config
	baseURL = "https://api-sandbox.sebgroup.com/mga/sps/oauth/oauth20/"

	// Hardcoded
	scope = "psd2_accounts psd2_payments"

	// Hardcoded
	responseTypeCode = "code"
)

func getToken(code string) (seb.TokenResponse, error) {
	u, err := url.Parse(baseURL + "token")
	if err != nil {
		return seb.TokenResponse{}, errors.New("failed to parse baseURL")
	}

	params := url.Values{}
	params.Add("client_id", myService.ClientID)
	params.Add("client_secret", myService.ClientSecret)
	params.Add("code", code)
	params.Add("grant_type", "authorization_code")
	// params.Add("authorization_code", "authorization_code")
	params.Add("redirect_uri", myService.redirectURI)
	payload := params.Encode()

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(payload))
	if err != nil {
		return seb.TokenResponse{}, errors.New("failed to create POST request")
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return seb.TokenResponse{}, errors.New("Request to SEB failed")
	}
	defer res.Body.Close()

	b, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(b))

	decoder := json.NewDecoder(bytes.NewReader(b))
	decoder.DisallowUnknownFields()
	var response seb.TokenResponse
	if err = decoder.Decode(&response); err != nil {
		return seb.TokenResponse{}, errors.New("failed to read response body")
	}

	return response, nil
}

func getAccounts(token string) (seb.Account, error) {
	url := "https://api-sandbox.sebgroup.com/ais/v5/identified2/accounts?withBalance=true"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return seb.Account{}, errors.New("failed to create GET request")
	}

	id := uuid.NewV1().String()
	req.Header.Add("content-type", "application/json")
	req.Header.Add("x-request-id", id)
	req.Header.Add("PSU-IP-Address", "127.0.0.1")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return seb.Account{}, fmt.Errorf("request to SEB failed: %w", err)
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return seb.Account{}, fmt.Errorf("read body failed: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(b))
	var response seb.Account
	if err = decoder.Decode(&response); err != nil {
		return seb.Account{}, errors.New("failed to read response body")
	}

	return response, nil
}

func handleGetAccounts(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("'token' missing"))
		return
	}

	// token = "RRG99XwwlSrTeGpfIpgU"
	fmt.Println("token", token)

	accounts, err := getAccounts(token)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong!"))
		return
	}

	var accountsResponse models.AccountsResponse
	for _, account := range accounts.Accounts {

		var b []byte
		b, err = json.MarshalIndent(account, "<br>", "  ")
		if err != nil {
			b = []byte("unknown")
		}

		accountsResponse.Accounts = append(accountsResponse.Accounts, models.Account{
			Iban:             account.Iban,
			Bban:             account.Bban,
			Type:             account.Product,
			AvailableBalance: availableBalance(account.Balances),
			Metadata:         string(b),
		})
	}

	if err := json.NewEncoder(w).Encode(&accountsResponse); err != nil {
		fmt.Printf("err: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong!"))
		return
	}
}

func availableBalance(balances []seb.Balance) float64 {
	for _, balance := range balances {
		if balance.BalanceType == "interimAvailable" {
			f, err := strconv.ParseFloat(balance.BalanceAmount.Amount, 64)
			if err != nil {
				return -1.0
			}
			return f
		}
	}
	return -2.0
}

func handleGetToken(w http.ResponseWriter, r *http.Request) {
	authorizationCode := r.URL.Query().Get("authorization_code")
	if authorizationCode == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("'authorization_code' missing"))
		return
	}

	token, err := getToken(authorizationCode)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong!"))
		return
	}

	if err := json.NewEncoder(w).Encode(&token); err != nil {
		fmt.Printf("err: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong!"))
		return
	}
}

func handleGenerateURL(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(baseURL + "authorize")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	params := url.Values{}
	params.Add("client_id", myService.ClientID)
	params.Add("scope", scope)
	params.Add("redirect_uri", myService.redirectURI)
	params.Add("response_type", responseTypeCode)
	u.RawQuery = params.Encode()

	w.Write([]byte(u.String()))
}

type Service struct {
	ClientID     string
	ClientSecret string
	redirectURI  string
}

var myService = Service{
	ClientSecret: os.Getenv("CLIENT_SECRET"),
	ClientID:     os.Getenv("CLIENT_ID"),
	redirectURI:  os.Getenv("REDIRECT_URL"),
}

func main() {
	r := chi.NewRouter()

	r.Get("/generate_url", handleGenerateURL)
	r.Get("/token", handleGetToken)
	r.Get("/accounts", handleGetAccounts)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/index.html", http.StatusSeeOther)
	})

	r.Get("/{filename}", func(w http.ResponseWriter, r *http.Request) {
		filename := chi.URLParam(r, "filename")

		b, err := ioutil.ReadFile(filename)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No such file:" + filename))
			return
		}

		w.Write(b)
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
