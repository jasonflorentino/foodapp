package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	http     *http.Client
	port     string
	PageSize int
}

type FoodName struct {
	FoodId            string `json:"id"`
	Code              int    `json:"code"`
	FoodGroupId       int    `json:"food_group_id"`
	FoodSourceId      int    `json:"food_source_id"`
	Description       string `json:"description"`
	DescriptionF      string `json:"description_f"`
	DateOfEntry       string `json:"date_of_entry"`
	DateOfPublication string `json:"date_of_publication"`
	CountryCode       string `json:"country_code"`
	ScientificName    string `json:"scientific_name"`
}

type Results struct {
	Status       string     `json:"status"`
	TotalResults int        `json:"totalResults`
	FoodNames    []FoodName `json:"foodNames"`
}

func NewClient(httpClient *http.Client, port string, pageSize int) *Client {
	if pageSize > 100 {
		pageSize = 100
	}

	return &Client{httpClient, port, pageSize}
}

func (c *Client) FetchEverything(query string, page string) (*[]FoodName, error) {
	endpoint := fmt.Sprintf("http://192.168.1.105:%s/foods?q=%s&limit=%d&offset=%s", c.port, url.QueryEscape(query), c.PageSize, page)
	resp, err := c.http.Get(endpoint)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	res := &[]FoodName{}
	return res, json.Unmarshal(body, res)
}
