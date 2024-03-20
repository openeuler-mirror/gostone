package utils

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
	"work.ctyun.cn/git/GoStack/gostone/connect"
	"work.ctyun.cn/git/GoStack/gostone/model"
)

func InitTable(db *gorm.DB, initEndpoint string, initRegion string) {
	if !db.Migrator().HasTable(&model.User{}) {
		db.Exec(" SET FOREIGN_KEY_CHECKS = ?", 0)
		err := db.Migrator().AutoMigrate(&model.Assignment{}, &model.Endpoint{},
			&model.Project{}, &model.Region{}, &model.Role{}, &model.Service{}, &model.User{}, &model.LocalUser{}, &model.Password{})
		if err != nil {
			panic(err)
		}
		db.Exec("ALTER TABLE local_user MODIFY id int(11)  not null  AUTO_INCREMENT")
		db.Exec("ALTER TABLE password MODIFY id int(11)  not null  AUTO_INCREMENT")
		rootProject := model.Project{
			Id:       "<<keystone.domain.root>>",
			Name:     "<<keystone.domain.root>>",
			Extra:    "{}",
			Enabled:  0,
			DomainId: "<<keystone.domain.root>>",
			IsDomain: 1,
		}
		defaultProject := model.Project{
			Id:          "default",
			Name:        "default",
			Extra:       "{}",
			Description: "The default domain",
			Enabled:     1,
			DomainId:    "<<keystone.domain.root>>",
			IsDomain:    1,
		}
		projectId := GenerateUUID()
		adminProject := model.Project{
			Id:          projectId,
			Name:        "admin",
			Description: "Bootstrap project for initializing the cloud.",
			Enabled:     1,
			Extra:       "{}",
			DomainId:    "default",
			ParentId: sql.NullString{
				String: "default",
				Valid:  true,
			},
			IsDomain: 0,
		}
		serviceProject := model.Project{
			Id:          GenerateUUID(),
			Name:        "service",
			Description: "Service Project",
			Enabled:     1,
			Extra:       "{}",
			DomainId:    "default",
			ParentId: sql.NullString{
				String: "default",
				Valid:  true,
			},
			IsDomain: 0,
		}
		db.Table("project").Create(rootProject)
		db.Table("project").Create(defaultProject)
		db.Table("project").Create(adminProject)
		db.Table("project").Create(serviceProject)
		userId := GenerateUUID()
		err = db.Transaction(func(tx *gorm.DB) error {
			adminUser := model.User{
				Id:           userId,
				Extra:        "{}",
				Enabled:      true,
				CreatedAt:    time.Now(),
				DomainId:     "default",
				LastActiveAt: time.Now(),
			}
			ok := tx.Table("user").Omit("LastActiveAt", "DefaultProjectId").Create(&adminUser)
			if ok.Error != nil {
				return ok.Error
			}
			adminLocalUser := model.LocalUser{
				Id:           1,
				UserId:       userId,
				DomainId:     "default",
				Name:         "admin",
				FailedAuthAt: time.Now(),
			}
			ok = tx.Table("local_user").Omit("FailedAuthAt").Create(&adminLocalUser)
			if ok.Error != nil {
				return ok.Error
			}
			adminPassword := model.Password{
				Id:           1,
				LocalUserId:  1,
				PasswordHash: GetPwd(connect.AppConf.GoStone.AdminPassword),
				CreatedAtInt: time.Now().UnixNano() / 1e6,
				CreatedAt:    time.Now(),
				PasswordSm3:  GenSM3Pwd(connect.AppConf.GoStone.AdminPassword),
				ExpiresAt:    time.Now(),
			}
			ok = tx.Table("password").Omit("ExpiresAt").Create(&adminPassword)
			if ok.Error != nil {
				return ok.Error
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
		roleId := GenerateUUID()
		adminRole := model.Role{
			Id:       roleId,
			Name:     "admin",
			Extra:    "{}",
			DomainId: "<<null>>",
		}
		db.Table("role").Create(adminRole)
		assignment := model.Assignment{
			Type:     "UserProject",
			ActorId:  userId,
			TargetId: projectId,
			RoleId:   roleId,
		}
		db.Table("assignment").Create(assignment)
		region := model.Region{
			Id:    initRegion,
			Extra: "{}",
		}
		db.Table("region").Create(region)
		serviceId := GenerateUUID()
		gostoneService := model.Service{
			Id:      serviceId,
			Type:    "identity",
			Enabled: 1,
			Extra:   "{\"name\": \"keystone\"}",
		}
		db.Table("service").Create(gostoneService)
		internal := model.Endpoint{
			Id:        GenerateUUID(),
			Interface: "internal",
			ServiceId: serviceId,
			URL:       initEndpoint,
			Extra:     "{}",
			Enabled:   1,
			RegionId:  initRegion,
		}
		admin := model.Endpoint{
			Id:        GenerateUUID(),
			Interface: "admin",
			ServiceId: serviceId,
			URL:       initEndpoint,
			Extra:     "{}",
			Enabled:   1,
			RegionId:  initRegion,
		}
		public := model.Endpoint{
			Id:        GenerateUUID(),
			Interface: "public",
			ServiceId: serviceId,
			URL:       initEndpoint,
			Extra:     "{}",
			Enabled:   1,
			RegionId:  initRegion,
		}
		db.Table("endpoint").Create(internal)
		db.Table("endpoint").Create(admin)
		db.Table("endpoint").Create(public)
		db.Exec(" SET FOREIGN_KEY_CHECKS = ?", 1)
		log.Infof("init gostone success userId:[%s],projectId:[%s] domainId:[%s]", userId, projectId, "default")
	}
	if !db.Migrator().HasColumn(&model.Password{}, "password_sm3") {
		err := db.Migrator().AddColumn(&model.Password{}, "password_sm3")
		if err != nil {
			panic(err)
		}
	}
}
