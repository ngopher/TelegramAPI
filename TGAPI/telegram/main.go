package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mtproto"
	"os"
	"strconv"
)

func main() {
	var err error
	var m *mtproto.MTProto
	var PhoneNumber string
	router := gin.Default()
	m, err = mtproto.NewMTProto(os.Getenv("HOME")+"/.telegram_go", "149.154.167.91:443", 64)
	if err != nil {
		fmt.Printf("Create failed: %s\n", err)
		os.Exit(2)
	}
	//Connecting to telegram Nearest DataCenter
	err = m.Connect()
	if err != nil {
		fmt.Printf("Connect failed: %s\n", err)
		os.Exit(2)
	}

	router.GET("/import/:phone", func(c *gin.Context) {
		firstName := "NewContact"
		lastName := "New Contact"
		PhoneNumber = c.Param("phone")
		contact := mtproto.Contact{
			Firstname: firstName,
			Lastname:  lastName,
			Phone:     PhoneNumber,
		}
		m.Contacts_ImportContacts([]mtproto.TL{contact.GetInputContact()})
		c.JSON(200, gin.H{"status": "successFully Added New Contact"})
	})
	router.GET("/verification", func(c *gin.Context) {
		user_id, _ := strconv.Atoi(c.Request.URL.Query().Get("user_id"))

		text := c.Request.URL.Query().Get("msg")
		err = m.SendMsg(int32(user_id), text)
		if err != nil {
			c.JSON(400, err.Error())
		} else {
			c.JSON(200, gin.H{"Status": "Successful"})
		}
	})
	router.Run(":8989") // listen and serve on 0.0.0.0:8080
}
