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
	ChildDOB     time.Time `json:"child_dob"`
	Phone        string    `json:"phone"`
	Email        string    `json:"email"`
	Notification int       `json:"notification"`
	City         string    `json:"city"`
	Address1     string    `json:"address1"`
	Address2     string    `json:"address2"`
	Contact      string    `json:"contact"`
	AgencyId     int       `json:"agencyId"`
	Unemployed   int       `json:"unemployed"`
	Newcomer     int       `json:"newcomer"`
	Homeless     int       `json:"homeless"`
	SpecialNeeds int       `json:"special_needs"`
}

type Gear struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Referral struct {
	ID           int      `json:"id"`
	ClientID     int      `json:"clientId"`
	Appointment1 NullTime `json:"appointment1"`
	Appointment2 NullTime `json:"appointment2"`
}

type ReferralGear struct {
	ReferralID int `json:"referralId"`
	GearID     int `json:"rearId"`
	Status     int `json:"status"`
}

type TimeInterval struct {
	Time     time.Time `json:"time"`
	Interval int       `json:"interval"`
}
