package db

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/marcoshuck/go-distributed/challenge/pkg/sum"
)

func Add(tx *gorm.DB, ip, user string, value int64) (interface{}, error) {
	var s sum.Sum
	if notFound := tx.Model(&sum.Sum{}).Where("user = ?", user).First(&s).RecordNotFound(); notFound {
		s.User = user
	}
	s.IP = ip
	s.Total += value

	if err := tx.Model(&sum.Sum{}).Save(&s).Error; err != nil{
		return nil, err
	}

	return s, nil
}

func Get(tx *gorm.DB, user string) (interface{}, error) {
	var s sum.Sum
	if notFound := tx.Model(&sum.Sum{}).Where("user = ?", user).First(&s).RecordNotFound(); notFound {
		return nil, errors.New("record not found")
	}
	return s, nil
}