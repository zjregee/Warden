package model

type OldInformation struct {
	ID             int64
	Name           string
	Age            string
	Sex            string
	Identification string
	Room           string
	Building       string
	FileId         string
	OldType        string
	CheckInTime    string
	ExpirationTime string
	Telephone      string
	Birthday       string
	Height         string
	Marriage       string
	Weight         string
	BloodType      string
	HeadPortrait   string
	Remark         string
}

type CheckOutInformation struct {
	ID           int64
	Name         string
	CheckOutType string
	Reason       string
	OutTime      string
	ApplyTime    string
	Remark       string
	IsPass       string
}

type OutInformation struct {
	ID              int64
	Name            string
	Reason          string
	OutTime         string
	PredictBackTime string
	People          string
	RelationShip    string
	Telephone       string
	Remark          string
	FactBackTime    string
	IsPass          string
}

type EmployeeInformation struct {
	ID           int64
	Name         string
	Sex          string
	Age          string
	Type         string
	HireDate     string
	Duty         string
	JobTitle     string
	Telephone    string
	Introduction string
	HeadPortrait string
	Remark       string
}

type UserInformation struct {
	ID       int64
	Name     string
	Password string
	Token    string
}