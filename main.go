// imgserver project
// Copyright (C) 2016  geosoft1  geosoft1@gmail.com
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func index(w http.ResponseWriter, r *http.Request) {
	//CORS response (http://www.html5rocks.com/en/tutorials/cors/)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-type", "image/jpg")

	now := time.Now()

	filename := "img/" + now.Format("0201") + ".jpg"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		filename = "img/Untitled.jpg"
	}

	log.Println(filename)

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Write(file)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}
