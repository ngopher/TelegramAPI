package mtprotox

type UserStatus struct {
	Status    string
	Online    bool
	Timestamp int32
}
type UserProfilePhoto struct {
	ID         int64
	PhotoSmall TL_fileLocation
	PhotoLarge TL_fileLocation
}
type User struct {
	flags                int32
	ID                   int32
	Username             string
	FirstName            string
	LastName             string
	Phone                string
	Photo                *UserProfilePhoto
	Status               *UserStatus
	Inactive             bool
	Mutual               bool
	Verified             bool
	Restricted           bool
	AccessHash           int64
	BotInfoVersion       int32
	BotInlinePlaceHolser string
	RestrictionReason    string
}

func (user *User) IsSelf() bool {
	if user.flags&1<<10 != 0 {
		return true
	}
	return false
}
func (user *User) IsContact() bool {
	if user.flags&1<<11 != 0 {
		return true
	}
	return false
}
func (user *User) IsMutualContact() bool {
	if user.flags&1<<12 != 0 {
		return true
	}
	return false
}
func (user *User) IsDeleted() bool {
	if user.flags&1<<13 != 0 {
		return true
	}
	return false
}
func (user *User) IsBot() bool {
	if user.flags&1<<14 != 0 {
		return true
	}
	return false
}
func (user *User) GetInputPeer() TL {
	if user.IsSelf() {
		return TL_inputPeerSelf{}
	} else {
		return TL_inputPeerUser{}
	}
}
func (user *User) GetPeer() TL {
	return TL_peerUser{
		user_id: user.ID,
	}
}
func NewUserStatus(userStatus TL) (s *UserStatus) {
	s = new(UserStatus)
	switch status := userStatus.(type) {
	case TL_userStatusEmpty:
		return nil
	case TL_userStatusOnline:
		s.Status = USER_STATUS_ONLINE
		s.Online = true
		s.Timestamp = status.expires
	case TL_userStatusOffline:
		s.Status = USER_STATUS_OFFLINE
		s.Online = false
		s.Timestamp = status.was_online
	case TL_userStatusRecently:
		s.Status = USER_STATUS_RECENTLY
		s.Online = false
	case TL_userStatusLastWeek:
		s.Status = USER_STATUS_LAST_WEEK
	case TL_userStatusLastMonth:
		s.Status = USER_STATUS_LAST_MONTH
	}
	return
}
func NewUserProfilePhoto(userProfilePhoto TL) (u *UserProfilePhoto) {
	u = new(UserProfilePhoto)
	switch pp := userProfilePhoto.(type) {
	case TL_userProfilePhotoEmpty:
		return nil
	case TL_userProfilePhoto:
		u.ID = pp.photo_id
		switch big := pp.photo_big.(type) {
		case TL_fileLocationUnavailable:
		case TL_fileLocation:
			u.PhotoLarge.dc_id = big.dc_id
			u.PhotoLarge.local_id = big.local_id
			u.PhotoLarge.secret = big.secret
			u.PhotoLarge.volume_id = big.volume_id
		}
		switch small := pp.photo_big.(type) {
		case TL_fileLocationUnavailable:
		case TL_fileLocation:
			u.PhotoSmall.dc_id = small.dc_id
			u.PhotoSmall.local_id = small.local_id
			u.PhotoLarge.secret = small.secret
			u.PhotoSmall.volume_id = small.volume_id
		}
	}
	return
}
func NewUser(in TL) (user *User) {
	user = new(User)
	switch u := in.(type) {
	case TL_userEmpty:
		user.ID = u.id
	case TL_user:
		user.ID = u.Id
		user.Username = u.Username
		user.FirstName = u.First_name
		user.LastName = u.Last_name
		user.AccessHash = u.Access_hash
		user.BotInfoVersion = u.Bot_info_version
		user.BotInlinePlaceHolser = u.Bot_inline_placeholder
		user.RestrictionReason = u.Restriction_reason
		user.Phone = u.Phone
		if u.Flags&1<<5 != 0 {
			user.Photo = NewUserProfilePhoto(u.Photo)
		}
		if u.Flags&1<<6 != 0 {
			user.Status = NewUserStatus(u.Status)
		}

	default:
		//fmt.Println(reflect.TypeOf(u).String())
		return nil
	}
	return
}
