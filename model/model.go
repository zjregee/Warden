package model

type OldInformation struct {
	ID             int64  `gorm:"size:20"`
	Name           string `gorm:"type:varchar(100)"`
	Age            string `gorm:"type:varchar(100)"`
	Sex            string `gorm:"type:varchar(100)"`
	Identification string `gorm:"type:varchar(100)"`
	Room           string `gorm:"type:varchar(100)"`
	Building       string `gorm:"type:varchar(100)"`
	FileId         string `gorm:"type:varchar(100)"`
	OldType        string `gorm:"type:varchar(100)"`
	CheckInTime    string `gorm:"type:varchar(100)"`
	ExpirationTime string `gorm:"type:varchar(100)"`
	Telephone      string `gorm:"type:varchar(100)"`
	Birthday       string `gorm:"type:varchar(100)"`
	Height         string `gorm:"type:varchar(100)"`
	Marriage       string `gorm:"type:varchar(100)"`
	Weight         string `gorm:"type:varchar(100)"`
	BloodType      string `gorm:"type:varchar(100)"`
	HeadPortrait   string `gorm:"type:varchar(100)"`
	Remark         string `gorm:"type:varchar(100)"`
}

type CheckOutInformation struct {
	ID           int64  `gorm:"size:20"`
	Name         string `gorm:"type:varchar(100)"`
	CheckOutType string `gorm:"type:varchar(100)"`
	Reason       string `gorm:"type:varchar(100)"`
	OutTime      string `gorm:"type:varchar(100)"`
	ApplyTime    string `gorm:"type:varchar(100)"`
	Remark       string `gorm:"type:varchar(100)"`
	IsPass       string `gorm:"type:varchar(100)"`
}

type OutInformation struct {
	ID              int64  `gorm:"size:20"`
	Name            string `gorm:"type:varchar(100)"`
	Reason          string `gorm:"type:varchar(100)"`
	OutTime         string `gorm:"type:varchar(100)"`
	PredictBackTime string `gorm:"type:varchar(100)"`
	People          string `gorm:"type:varchar(100)"`
	RelationShip    string `gorm:"type:varchar(100)"`
	Telephone       string `gorm:"type:varchar(100)"`
	Remark          string `gorm:"type:varchar(100)"`
	FactBackTime    string `gorm:"type:varchar(100)"`
	IsPass          string `gorm:"type:varchar(100)"`
}

type EmployeeInformation struct {
	ID           int64  `gorm:"size:20"`
	Name         string `gorm:"type:varchar(100)"`
	Sex          string `gorm:"type:varchar(100)"`
	Age          string `gorm:"type:varchar(100)"`
	Type         string `gorm:"type:varchar(100)"`
	HireDate     string `gorm:"type:varchar(100)"`
	Duty         string `gorm:"type:varchar(100)"`
	JobTitle     string `gorm:"type:varchar(100)"`
	Telephone    string `gorm:"type:varchar(100)"`
	Introduction string `gorm:"type:varchar(100)"`
	HeadPortrait string `gorm:"type:varchar(100)"`
	Remark       string `gorm:"type:varchar(100)"`
}

type UserInformation struct {
	ID       int64  `gorm:"size:20"`
	Name     string `gorm:"type:varchar(100)"`
	Password string `gorm:"type:varchar(100)"`
	Token    string `gorm:"type:varchar(100)"`
}