package services

import (
	"akatsuki/skeleton-go/src/dao"
	"akatsuki/skeleton-go/src/datasources/pqsql/entity"
)

func GetAllDevice() ([]entity.Devices, error) {
	deviceDAO, _ := dao.NewDevicesDAO()
	allDevices, err := deviceDAO.FindAll()
	if err != nil {
		return nil, err
	}
	return allDevices, nil
}
