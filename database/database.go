//    Melfy  Copyright (C) 2018  MelfyStream
//
//    This program is free software: you can redistribute it and/or modify
//    it under the terms of the GNU General Public License as published by
//    the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.
//
//    This program is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU General Public License for more details.
//
//    You should have received a copy of the GNU General Public License
//    along with this program.  If not, see <http://www.gnu.org/licenses/>.

package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `gorm:"Default:null" sql:"index" json:"deleted_at"`
}

var Instance *gorm.DB

func InitDatabase(username, password, name string) error {
	db, err := gorm.Open("mysql",
		username+":"+password+"@/"+name+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return err
	}

	Instance = db

	return nil
}
