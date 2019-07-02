/*
 * Copyright 2018 Foolin.  All rights reserved.
 *
 * Use of this source code is governed by a MIT style
 * license that can be found in the LICENSE file.
 *
 */

package main

import (
	"angeldm.echoview/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//conn, err := pop.Connect("development")
	//if err!= nil {
	//	panic(err)
	//}
	//users := models.Users{}
	//err = conn.All(&users)
	//if err!=nil {
	//	panic(err)
	//}else {
	//	fmt.Println(users)
	//}

	app := application.NewApplication()
	app.Start()
}
