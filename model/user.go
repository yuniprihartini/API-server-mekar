package model

import (
	"regexp"	
	"strings"
)


type User struct {
	UserID      string            `json:"idUser"`
	IdCard      string            `json:"idCardNumber"`
	Username    string            `json:"username"`
	DateOfBirth string            `json:"dateOfBirth"`
	Job         Job               `json:"job,omitempty"`
	Education   Education         `json:"education"`
	UserStatus  int               `json:"userStatus"`
	CreatedDate string            `json:"createdDate"`
	UpdatedDate string            `json:"updatedDate"`
	Errors      map[string]string `json:"errors"`
}

type UserList struct {
	Users    []*User  `json:"users"`
	Metadata Metadata `json:"metadata"`
}

type Metadata struct {
	CurrentPage int `json:"currentPage"`
	FirstPage   int `json:"firstPage"`
	LastPage    int `json:"lastPage"`
	TotalData   int `json:"totalData"`
}

type Job struct {
	JobId    string `json:"jobId"`
	JobLabel string `json:"jobLabel"`
}

type Education struct {
	EducationId    string `json:"educationId"`
	EducationLabel string `json:"educationLabel"`
}

func (user *User) Validate() bool {
	user.Errors = make(map[string]string)
	rg := regexp.MustCompile("((19|20)\\d\\d)-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])")
	if len(strings.TrimSpace(user.IdCard)) > 16 || strings.TrimSpace(user.IdCard) == "" {
		user.Errors["IdCardNumber"] = "Please enter a valid id card number"
	}
	if strings.TrimSpace(user.Username) == "" {
		user.Errors["Username"] = "Please enter a username"
	}
	if strings.TrimSpace(user.DateOfBirth) == "" || !rg.MatchString(user.DateOfBirth) {
		user.Errors["DateOfBirth"] = "Please enter a valid date of birth"
	}
	if strings.TrimSpace(user.Job.JobId) == "" {
		user.Errors["JobId"] = "Please enter a job id"
	}
	if strings.TrimSpace(user.Education.EducationId) == "" {
		user.Errors["EducationId"] = "Please enter a education id"
	}
	return len(user.Errors) == 0
}
