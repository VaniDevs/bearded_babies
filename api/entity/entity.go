package entity

import "time"

type Agency struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     int    `json:"role"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	City     string `json:"city"`
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	Contact  string `json:"contact"`
}

type Client struct {
	ID           int       `json:"id"`
	Status       int       `json:"status"`
	Name         string    `json:"name"`
	DOB          time.Time `json:"dob"`
	ChildDOB     time.Time `json:"childDob"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	City         string    `json:"city"`
	Address1     string    `json:"address1"`
	Address2     string    `json:"address2"`
	AgencyId     int       `json:"agencyId"`
	Unemployed   int       `json:"unemployed"`
	Newcomer     int       `json:"newcomer"`
	Homeless     int       `json:"homeless"`
	SpecialNeeds int       `json:"specialNeeds"`
}

type Gear struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Referral struct {
	ID           int      `json:"id"`
	ClientID     int      `json:"clientId"`
	Appointment1 time.Time `json:"appointment1"`
	Appointment2 time.Time `json:"appointment2"`
	Requested    []int    `json:"requested"`
	Unavailable  []int    `json:"unavailable"`
}

type ReferralGear struct {
	ReferralID int `json:"referralId"`
	GearID     int `json:"gearId"`
	Status     int `json:"status"`
}

type TimeInterval struct {
	Time     time.Time `json:"time"`
	Interval int       `json:"interval"`
}
