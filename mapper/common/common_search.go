package common

import (
	"gorm.io/gorm"
	"strings"
)

func FindByName(name string) func(db *gorm.DB) *gorm.DB {
	if name == "" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("1=1")
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name=?", name)
	}
}

func FindById(userId string) func(db *gorm.DB) *gorm.DB {
	if userId == "" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("1=1")
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("id=?", userId)
	}
}

func FindByDomainId(domainId string) func(db *gorm.DB) *gorm.DB {
	if domainId == "" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("1=1")
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("domain_id=?", domainId)
	}
}

func FindByEnabled(enabled string) func(db *gorm.DB) *gorm.DB {
	if strings.ToLower(enabled) == "false" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("enabled=?", 0)
		}
	}
	if strings.ToLower(enabled) == "true" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("enabled=?", 1)
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("1=1")
	}
}

func FindByRoleId(roleId string) func(db *gorm.DB) *gorm.DB {
	if roleId == "" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("1=1")
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("role_id=?", roleId)
	}
}

func FindByActorId(actorId string) func(db *gorm.DB) *gorm.DB {
	if actorId == "" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("1=1")
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("actor_id=?", actorId)
	}
}

func FindByTargetId(targetId string) func(db *gorm.DB) *gorm.DB {
	if targetId == "" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("1=1")
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("target_id=?", targetId)
	}
}
