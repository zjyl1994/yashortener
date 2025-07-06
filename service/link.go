package service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/zjyl1994/yashortener/infra/model"
	"github.com/zjyl1994/yashortener/infra/utils"
	"github.com/zjyl1994/yashortener/infra/vars"
	"gorm.io/gorm"
)

func CreateLink(code, link string) (string, error) {
	var m model.Link
	m.Link = strings.TrimSpace(link)
	m.CreateAt = time.Now().Unix()

	shortLen := 5
	for {
		if code != "" {
			m.Code = strings.TrimSpace(code)
		} else {
			m.Code = utils.RandChars(shortLen)
		}
		err := vars.DB.Create(&m).Error
		if err == nil {
			return m.Code, nil
		}
		if !errors.Is(err, gorm.ErrDuplicatedKey) {
			return "", err
		}
		if code != "" {
			return "", fmt.Errorf("duplicate code %s", code)
		}
		shortLen++
	}
}

func UpdateLink(code, link string) error {
	return vars.DB.Model(&model.Link{}).Where("code = ?", strings.TrimSpace(code)).Update("link", link).Error
}

func GetLink(code string) (*model.Link, error) {
	var link model.Link
	err := vars.DB.Where("code = ?", code).First(&link).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &link, nil
}

func RecordAccess(id uint, ip, userAgent string) error {
	access := model.Access{
		LinkID:    id,
		UserAgent: userAgent,
		IP:        ip,
		CreateAt:  time.Now().Unix(),
	}
	return vars.DB.Create(&access).Error
}

func DeleteLink(code string) error {
	link, err := GetLink(code)
	if err != nil {
		return err
	}
	if link == nil {
		return nil
	}
	return vars.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Delete(link).Error
		if err != nil {
			return err
		}
		return tx.Delete(&model.Access{}, "link_id = ?", link.ID).Error
	})
}

func ListAccessRecord(linkID uint, page, size int) ([]model.Access, int64, error) {
	var count int64
	query := vars.DB.Model(&model.Access{}).Where("link_id = ?", linkID)
	err := query.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return []model.Access{}, 0, nil
	}

	var access []model.Access
	err = query.Order("create_at desc").Offset((page - 1) * size).Limit(size).Find(&access).Error
	if err != nil {
		return nil, 0, err
	}
	return access, count, nil
}

func ListLink(keyword string, page, size int) ([]model.Link, int64, error) {
	var count int64
	query := vars.DB.Model(&model.Link{})
	if keyword != "" {
		query = query.Where("code LIKE ?", "%"+keyword+"%").Or("link LIKE ?", "%"+keyword+"%")
	}
	err := query.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	if count == 0 {
		return []model.Link{}, 0, nil
	}
	var links []model.Link
	err = query.Order("create_at desc").Offset((page - 1) * size).Limit(size).Find(&links).Error
	if err != nil {
		return nil, 0, err
	}
	return links, count, nil
}
