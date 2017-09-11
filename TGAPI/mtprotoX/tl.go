package mtprotox

import (
	"math/big"
)

type TL interface {
	encode() []byte
}

type TL_msg_container struct {
	items []TL_MT_message
}

type TL_MT_message struct {
	msg_id int64
	seq_no int32
	size   int32
	data   interface{}
}

type TL_req_pq struct {
	nonce []byte
}

type TL_p_q_inner_data struct {
	pq           *big.Int
	p            *big.Int
	q            *big.Int
	nonce        []byte
	server_nonce []byte
	new_nonce    []byte
}
type TL_req_DH_params struct {
	nonce        []byte
	server_nonce []byte
	p            *big.Int
	q            *big.Int
	fp           uint64
	encdata      []byte
}
type TL_client_DH_inner_data struct {
	nonce        []byte
	server_nonce []byte
	retry        int64
	g_b          *big.Int
}
type TL_set_client_DH_params struct {
	nonce        []byte
	server_nonce []byte
	encdata      []byte
}
type TL_resPQ struct {
	nonce        []byte
	server_nonce []byte
	pq           *big.Int
	fingerprints []int64
}

type TL_server_DH_params_ok struct {
	nonce            []byte
	server_nonce     []byte
	encrypted_answer []byte
}

type TL_server_DH_inner_data struct {
	nonce        []byte
	server_nonce []byte
	g            int32
	dh_prime     *big.Int
	g_a          *big.Int
	server_time  int32
}
type TL_user struct {
	Flags int32
	// Self	bool // flags_10?true
	// Contact	bool // flags_11?true
	// Mutual_contact	bool // flags_12?true
	// Deleted	bool // flags_13?true
	// Bot	bool // flags_14?true
	// Bot_chat_history	bool // flags_15?true
	// Bot_nochats	bool // flags_16?true
	// Verified	bool // flags_17?true
	// Restricted	bool // flags_18?true
	// Min	bool // flags_20?true
	// Bot_inline_geo	bool // flags_21?true
	Id                     int32
	Access_hash            int64
	First_name             string
	Last_name              string
	Username               string
	Phone                  string
	Photo                  TL // flags_5?UserProfilePhoto
	Status                 TL // flags_6?UserStatus
	Bot_info_version       int32
	Restriction_reason     string
	Bot_inline_placeholder string
	Lang_code              string
}
type TL_new_session_created struct {
	first_msg_id int64
	unique_id    int64
	server_salt  []byte
}

type TL_bad_server_salt struct {
	bad_msg_id      int64
	bad_msg_seqno   int32
	error_code      int32
	new_server_salt []byte
}

type TL_crc_bad_msg_notification struct {
	bad_msg_id    int64
	bad_msg_seqno int32
	error_code    int32
}

type TL_msgs_ack struct {
	msgIds []int64
}

type TL_rpc_result struct {
	req_msg_id int64
	obj        interface{}
}

type TL_rpc_error struct {
	error_code    int32
	error_message string
}

type TL_dh_gen_ok struct {
	nonce           []byte
	server_nonce    []byte
	new_nonce_hash1 []byte
}

type TL_ping struct {
	ping_id int64
}

type TL_pong struct {
	msg_id  int64
	ping_id int64
}
