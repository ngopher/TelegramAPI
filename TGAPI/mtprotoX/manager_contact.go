package mtprotox

import (
	"fmt"
	"log"
	_ "strconv"
)

// Contact
type Contact struct {
	ClientID  int64
	Firstname string
	Lastname  string
	Phone     string
	UserID    int32
	Mutual    bool
}

func (c *Contact) GetInputContact() TL {
	return TL_inputPhoneContact{
		client_id:  c.ClientID,
		first_name: c.Firstname,
		last_name:  c.Lastname,
		phone:      c.Phone,
	}
}

func NewContact(in TL) (contact *Contact) {
	contact = new(Contact)
	switch c := in.(type) {
	case TL_contact:
		contact.UserID = c.user_id
		contact.Mutual = toBool(c.mutual)
	case TL_importedContact:
		contact.UserID = c.user_id
		contact.ClientID = c.client_id
	case TL_inputPhoneContact:
		contact.ClientID = c.client_id
		contact.Firstname = c.first_name
		contact.Lastname = c.last_name
		contact.Phone = c.phone
	default:
		log.Println("GetContact::Error::Invalid Type")
		return nil
	}
	return
}

func (m *MTProto) Contacts_GetContacts(hash string) ([]Contact, []User) {
	resp := make(chan TL, 1)
	m.queueSend <- packetToSend{TL_contacts_getContacts{
		hash},
		resp,
	}
	x := <-resp
	list, ok := x.(TL_contacts_contacts)
	if !ok {
		log.Println(fmt.Sprintf("RPC: %#v", x))
		return []Contact{}, []User{}
	}
	TContacts := make([]Contact, 0, len(list.contacts))
	TUsers := make([]User, 0, len(list.users))
	for _, v := range list.contacts {
		TContacts = append(
			TContacts,
			*NewContact(v),
		)
	}
	for _, v := range list.users {
		switch u := v.(type) {
		case TL_userEmpty:
			TUsers = append(TUsers, *NewUser(u))
		case TL_userProfilePhoto:
			TUsers[len(TUsers)-1].Photo = NewUserProfilePhoto(u)
		case TL_userStatusRecently, TL_userStatusOffline, TL_userStatusOnline, TL_userStatusLastWeek, TL_userStatusLastMonth:
			TUsers[len(TUsers)-1].Status = NewUserStatus(u)
		}
	}
	return TContacts, TUsers
}

func (m *MTProto) Contacts_ImportContacts(contacts []TL) {
	resp := make(chan TL, 1)
	m.queueSend <- packetToSend{
		TL_contacts_importContacts{
			contacts,
		},
		resp,
	}
	x := <-resp
	switch r := x.(type) {
	case TL_contacts_importedContacts:
		//TODO:: must do something with response
		log.Println(r)
	default:
		log.Println(fmt.Sprintf("RPC: %#v", x))
		return

	}

}
