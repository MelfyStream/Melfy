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

package main

import (
	"github.com/MelfyStream/Melfy/database"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	log.SetPrefix("[Melfy] ")
}

func main() {
	log.Println("[1/2] Loading env variables.")
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	log.Println("[2/2] Loading database...")
	err := database.InitDatabase(os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}
}
