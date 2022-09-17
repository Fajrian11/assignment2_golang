package repositories

import (
	"assignment2_golang/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  int    `json:"username"`
	Phone     int    `json:"phone"`
	Email     int    `json:"email"`
	Uuid      int    `json:"uuid"`
}

func NewOrderPerson() Person {
	return Person{}
}

func (p *Person) GetPerson() ([]model.Person, error) {
	var response = []model.Person{}

	fmt.Println("Calling API...")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://hiyaa.site/data.php?qty=1&apikey=7f8fc96e-de1f-4aab-9c62-3dd1de365e66", nil)
	if err != nil {
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	json.Unmarshal(bodyBytes, &response)
	fmt.Printf("API Response as struct %+v\n", response)

	return response, nil
}
