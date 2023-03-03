package user

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const FemasLogin = "https://femascloud.com/freedomsystems/fsapi/V3/login.json"
const FemasPunch = "https://fsapi.femascloud.com/freedomsystems/fsapi/V3/punch_card.json"

type User struct {
	DomainName string `json:"domainName"`
	Account    string `json:"account"`
	Password   string `json:"password"`
	IsAccount  string `json:"isAccount"`
}

type LoginResponse struct {
	Response struct {
		Status            string      `json:"status"`
		ErrMsg            string      `json:"err_msg"`
		LoginType         string      `json:"login_type"`
		Token             string      `json:"token"`
		IsSso             bool        `json:"is_sso"`
		DomainName        string      `json:"domain_name"`
		OrgName           string      `json:"org_name"`
		DeptName          string      `json:"dept_name"`
		UserSn            string      `json:"user_sn"`
		UserName          string      `json:"user_name"`
		UserType          string      `json:"user_type"`
		UserEnName        string      `json:"user_en_name"`
		UserTitle         string      `json:"user_title"`
		UserPhoto         string      `json:"user_photo"`
		NotificationToken interface{} `json:"notification_token"`
		AppMenu           struct {
			Home                       int `json:"home"`
			Clock                      int `json:"clock"`
			FormSign                   int `json:"form_sign"`
			Schedule                   int `json:"schedule"`
			GroupSchedule              int `json:"group_schedule"`
			Eboard                     int `json:"eboard"`
			Inquiry                    int `json:"inquiry"`
			InquiryAttendance          int `json:"inquiry_attendance"`
			InquiryOvertime            int `json:"inquiry_overtime"`
			InquiryLeave               int `json:"inquiry_leave"`
			InquirySalary              int `json:"inquiry_salary"`
			InquiryUnderlingAttendance int `json:"inquiry_underling_attendance"`
			PersonalLvdaysLimit        int `json:"personal_lvdays_limit"`
		} `json:"app_menu"`
	} `json:"response"`
}

type PunchPayload struct {
	ClockData string `json:"clockData"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func newPunchPayload(punchType string) *PunchPayload {
	punchData := fmt.Sprintf("%s,%s,%s", "2", "1", punchType)
	return &PunchPayload{
		ClockData: punchData,
		Latitude:  "24.9916698",
		Longitude: "121.5431419",
	}
}

func NewUser(userName, userPassword string) *User {
	return &User{
		DomainName: "freedomsystems",
		Account:    userName,
		Password:   userPassword,
		IsAccount:  "1", // using account name
	}
}

func (u *User) getToken() string {
	userData := fmt.Sprintf("%s:%s", u.Account, u.Password)
	encodedUserData := base64.StdEncoding.EncodeToString([]byte(userData))
	authorizationToken := fmt.Sprintf("Basic %s", encodedUserData)

	currentUserPayload, err := json.Marshal(NewUser(u.Account, u.Password))
	if err != nil {
		panic(err)
	}

	// POST request
	request, err := http.NewRequest(http.MethodPost, FemasLogin, bytes.NewBuffer(currentUserPayload))
	if err != nil {
		panic(err)
	}
	request.Header.Set("Authorization", authorizationToken)

	client := http.Client{}
	httpResponse, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	// close the response after use
	defer func(response *http.Response) {
		if err := response.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}(httpResponse)

	bytesResponse, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		panic(err)
	}

	var response LoginResponse
	if err = json.Unmarshal(bytesResponse, &response); err != nil {
		log.Fatal(err)
	}

	return response.Response.Token
}

func (u *User) Punch(date, punchType string) bool {
	newPunchPayload := newPunchPayload(punchType)

	punchJson, err := json.Marshal(newPunchPayload)
	if err != nil {
		panic(err)
	}

	var client http.Client

	punchRequest, err := http.NewRequest(http.MethodPost, FemasPunch, bytes.NewBuffer(punchJson))
	punchRequest.Header.Set("Authorization", u.getToken())

	httpResponse, err := client.Do(punchRequest)
	if err != nil {
		panic(err)
	}

	defer func(response *http.Response) {
		if err := response.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}(httpResponse)

	if httpResponse.StatusCode != http.StatusOK {
		return false
	}
	return true
}
