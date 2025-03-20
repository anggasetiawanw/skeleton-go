package dao

import (
	"akatsuki/skeleton-go/src/datasources/connections"
	"akatsuki/skeleton-go/src/datasources/pqsql/entity"
	"fmt"

	"gorm.io/gorm"
)

type DeviceDAO interface {
	FindAll() ([]entity.Devices, error)
}
type devicedao struct {
	db *gorm.DB
}

var db, _ = connections.NewPGSQLConnection()

func NewDevicesDAO() (*devicedao, error) {
	dts := db.PGSQLConnection()

	return &devicedao{dts}, nil
}

func (r *devicedao) Save(device entity.Devices) (entity.Devices, error) {
	err := r.db.Create(&device).Error
	if err != nil {
		return device, err
	}

	return device, nil
}

func (r *devicedao) FindByID(ID string) (entity.Devices, error) {
	var device entity.Devices

	err := r.db.Where("id = ?", ID).Find(&device).Error
	if err != nil {
		return device, err
	}

	return device, nil
}

func (r *devicedao) Update(device entity.Devices) (entity.Devices, error) {
	err := r.db.Save(&device).Error

	if err != nil {
		return device, err
	}

	return device, nil
}

func (r *devicedao) FindAll() ([]entity.Devices, error) {
	var devices []entity.Devices

	err := r.db.Find(&devices).Error
	if err != nil {
		return devices, err
	}
	defer func() {
		db.PGSQLClose()
		fmt.Println("Closed to PGSQL")

	}()
	return devices, nil
}
