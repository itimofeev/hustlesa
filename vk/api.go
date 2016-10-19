package vk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func NewVkClient() *VkClient {
	return &VkClient{
		client: http.DefaultClient,
	}
}

type VkClient struct {
	client *http.Client
}

type VkUser struct {
	Uid       int    `json:"uid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type VkUserResponse struct {
	Count int      `json:"count"`
	Users []VkUser `json:"users"`
}

type VkResponse struct {
	Response VkUserResponse `json:"response"`
}

// https://vk.com/dev/groups.getMembers?params[group_id]=dwizhenie&params[offset]=0&params[count]=10&params[fields]=bdate&params[v]=5.59
// https://vk.com/dev/database.getCitiesById?params[city_ids]=0%2C1%2C2&params[v]=5.59
//
// Друзья Макуся: https://api.vk.com/method/friends.get?user_id=3956992&fields=sex,bdate&count=10&offset=0
//
// disco_dubrovka
// dwizhenie
// discoswing
// discofoxan
// club3688533 (Динамика)
// svitelev_hustle
//
//https://api.vk.com/method/groups.getMembers?group_id=dwizhenie&fields=sex,bdate,photo_200_orig&count=1&offset=1
func (vc *VkClient) GetGroupMembers(count, offset int, groupId string) (userResponse VkUserResponse, err error) {
	fields := "sex,bdate,city,country,status"
	// https://api.vk.com/method/groups.getMembers?group_id=dwizhenie&fields=sex,bdate&count=10&offset=0
	reqPath := fmt.Sprintf("https://api.vk.com/method/groups.getMembers?group_id=%s&fields=%s&count=%d&offset=%d", groupId, fields, count, offset)

	req, _ := http.NewRequest("GET", reqPath, nil)
	resp, err := vc.client.Do(req)

	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("Vk returned non 200 code: %d", resp.StatusCode))
	}

	data, err := ioutil.ReadAll(resp.Body)

	var response VkResponse

	err = json.Unmarshal(data, &response)

	userResponse = response.Response
	return
}
