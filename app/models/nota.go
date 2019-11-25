package models

import (
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Nota struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"id"`
	Titulo    string    `gorm:"size:255;not null;" json:"titulo"`
	Conteudo  string    `gorm:"size:2000;not null;" json:"conteudo"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (n *Nota) Prepare() {
	n.ID = 0
	n.Titulo = html.EscapeString(strings.TrimSpace(n.Titulo))
	n.Conteudo = html.EscapeString(strings.TrimSpace(n.Conteudo))
	n.CreatedAt = time.Now()
	n.UpdatedAt = time.Now()
}

func (n *Nota) SalvaNota(db *gorm.DB) (*Nota, error) {
	var err error
	db.AutoMigrate(Nota{})
	err = db.Debug().Model(&Nota{}).Create(&n).Error
	if err != nil {
		return &Nota{}, err
	}
	return n, nil
}

func (p *Nota) AllNotas(db *gorm.DB) (*[]Nota, error) {
	var err error
	notas := []Nota{}
	err = db.Debug().Model(&Nota{}).Limit(100).Find(&notas).Error
	if err != nil {
		return &[]Nota{}, err
	}
	return &notas, nil
}
