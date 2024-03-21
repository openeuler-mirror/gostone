package model

import (
	"database/sql"
	"time"
)

// Assignment [...]
type Assignment struct {
	Type      string `gorm:"primary_key;column:type;type:enum('UserProject','GroupProject','UserDomain','GroupDomain');not null"`
	ActorId   string `gorm:"primary_key;index:ix_actor_id;column:actor_id;type:varchar(64);not null" json:"actor_id"`
	TargetId  string `gorm:"primary_key;column:target_id;type:varchar(64);not null" json:"target_id"`
	RoleId    string `gorm:"primary_key;column:role_id;type:varchar(64);not null" json:"role_id"`
	Inherited bool   `gorm:"primary_key;column:inherited;type:tinyint(1);not null" json:"inherited"`
}

// Endpoint [...]
type Endpoint struct {
	Id               string `gorm:"primary_key;column:id;type:varchar(64);not null"`
	LegacyEndpointId string `gorm:"column:legacy_endpoint_id;type:varchar(64)" json:"legacy_endpoint_id"`
	Interface        string `gorm:"column:interface;type:varchar(8);not null"`
	ServiceId        string `gorm:"index:service_id;column:service_id;type:varchar(64);not null" json:"service_id"`
	URL              string `gorm:"column:url;type:text;not null"`
	Extra            string `gorm:"column:extra;type:text"`
	Enabled          int    `gorm:"column:enabled;type:tinyint(1);not null"`
	RegionId         string `gorm:"index:fk_endpoint_region_id;column:region_id;type:varchar(255)" json:"region_id"`
}

type EndpointResponse struct {
	Id               string `json:"id"`
	LegacyEndpointId string `json:"legacy_endpoint_id"`
	Interface        string `json:"interface"`
	ServiceId        string `json:"service_id"`
	Url              string `json:"url"`
	Extra            string `json:"extra"`
	Enabled          bool   `json:"enabled"`
	RegionId         string `json:"region_id"`
	Region           string `json:"region"`
}

// LocalUser [...]
type LocalUser struct {
	Id              int       `gorm:"primary_key;column:id;type:int(11);not null;"`
	UserId          string    `gorm:"unique;index:local_user_user_id_fkey;column:user_id;type:varchar(64);not null" json:"user_id"`
	DomainId        string    `gorm:"unique_index:domain_id;index:local_user_user_id_fkey;column:domain_id;type:varchar(64);not null" json:"domain_id"`
	Name            string    `gorm:"unique_index:domain_id;column:name;type:varchar(255);not null"`
	FailedAuthCount int       `gorm:"column:failed_auth_count;type:int(11)"`
	FailedAuthAt    time.Time `gorm:"column:failed_auth_at;type:datetime"`
}

// Password [...]
type Password struct {
	Id           int       `gorm:"primary_key;column:id;type:int(11);not null"`
	LocalUserId  int       `gorm:"index:local_user_id;column:local_user_id;type:int(11);not null" json:"local_user_id"`
	Password     string    `gorm:"column:password;type:varchar(128)" json:"password"`
	ExpiresAt    time.Time `gorm:"column:expires_at;type:datetime" json:"expires_at"`
	SelfService  bool      `gorm:"column:self_service;type:tinyint(1);not null" json:"self_service"`
	PasswordHash string    `gorm:"column:password_hash;type:varchar(255)"`
	CreatedAtInt int64     `gorm:"column:created_at_int;type:bigint(20);not null" json:"created_at_int"`
	ExpiresAtInt int64     `gorm:"column:expires_at_int;type:bigint(20)" json:"expires_at_int"`
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
	PasswordSm3  string    `gorm:"column:password_sm3;type:varchar(255)" json:"password_sm_3"`
}

// Project
type Project struct {
	Id          string         `gorm:"primary_key;column:id;type:varchar(64);not null" json:"id"`
	Name        string         `gorm:"unique_index:ixu_project_name_domain_id;column:name;type:varchar(64);not null"`
	Extra       string         `gorm:"column:extra;type:text"`
	Description string         `gorm:"column:description;type:text"`
	Enabled     int            `gorm:"column:enabled;type:tinyint(1)"`
	DomainId    string         `gorm:"unique_index:ixu_project_name_domain_id;column:domain_id;type:varchar(64);not null js" json:"domain_id"`
	ParentId    sql.NullString `gorm:"index:project_parent_id_fkey;column:parent_id;type:varchar(64)" json:"parent_id"`
	IsDomain    int            `gorm:"column:is_domain;type:tinyint(1);not null" json:"is_domain"`
}

// Region [...]
type Region struct {
	Id             string `gorm:"primary_key;column:id;type:varchar(255);not null"`
	Description    string `gorm:"column:description;type:varchar(255);not null" json:"description"`
	ParentRegionId string `gorm:"column:parent_region_id;type:varchar(255)" json:"parent_region_id"`
	Extra          string `gorm:"column:extra;type:text"`
}

// Role [...]
type Role struct {
	Id       string `gorm:"primary_key;column:id;type:varchar(64);not null"`
	Name     string `gorm:"unique_index:ixu_role_name_domain_id;column:name;type:varchar(255);not null"`
	Extra    string `gorm:"column:extra;type:text"`
	DomainId string `gorm:"unique_index:ixu_role_name_domain_id;column:domain_id;type:varchar(64);not null" json:"domain_id"`
}

// Service [...]
type Service struct {
	Id      string `gorm:"primary_key;column:id;type:varchar(64);not null"`
	Type    string `gorm:"column:type;type:varchar(255)"`
	Enabled int    `gorm:"column:enabled;type:tinyint(1);not null"`
	Extra   string `gorm:"column:extra;type:text"`
}

type ServiceResponse struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	Enabled     bool   `json:"enabled"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// User [...]
type User struct {
	Id               string    `gorm:"primary_key;unique_index:ixu_user_id_domain_id;column:id;type:varchar(64);not null"`
	Extra            string    `gorm:"column:extra;type:text"`
	Enabled          bool      `gorm:"column:enabled;type:tinyint(1)"`
	DefaultProjectId string    `gorm:"index:ix_default_project_id;column:default_project_id;type:varchar(64)" json:"default_project_id"`
	CreatedAt        time.Time `gorm:"column:created_at;type:datetime"`
	LastActiveAt     time.Time `gorm:"column:last_active_at;type:date"`
	DomainId         string    `gorm:"unique_index:ixu_user_id_domain_id;index:domain_id;column:domain_id;type:varchar(64);not null" json:"domain_id"`
}

type UserInfo struct {
	Id               string    `json:"id"`
	Name             string    `json:"name"`
	Enabled          int       `json:"enabled"`
	DomainId         string    `json:"domain_id"`
	DefaultProjectId string    `json:"default_project_id"`
	CreateAt         time.Time `json:"create_at"`
	LastActiveAt     time.Time `json:"last_active_at" qorm:"column:last_active_at"`
	Password         string    `json:"password"`
	PasswordSm3      string    `json:"password_sm3" qorm:"column:password_sm3"`
	LocalUserId      int       `json:"local_user_id"`
	Extra            string    `json:"extra"`
}
