// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMessage = "Chat_c26eca0d2bea3169cfc5b550f4e6b040"

// Message mapped from table <Chat_c26eca0d2bea3169cfc5b550f4e6b040>
type Message struct {
	MesLocalID      int32  `gorm:"column:mesLocalID;primaryKey" json:"mesLocalID"`
	MesSvrID        int64  `gorm:"column:mesSvrID" json:"mesSvrID"`
	MsgCreateTime   int32  `gorm:"column:msgCreateTime" json:"msgCreateTime"`
	MsgContent      string `gorm:"column:msgContent" json:"msgContent"`
	MsgStatus       int32  `gorm:"column:msgStatus" json:"msgStatus"`
	MsgImgStatus    int32  `gorm:"column:msgImgStatus" json:"msgImgStatus"`
	MessageType     int32  `gorm:"column:messageType" json:"messageType"`
	MesDes          int32  `gorm:"column:mesDes" json:"mesDes"`
	MsgSource       string `gorm:"column:msgSource" json:"msgSource"`
	IntRes1         int32  `gorm:"column:IntRes1" json:"IntRes1"`
	IntRes2         int32  `gorm:"column:IntRes2" json:"IntRes2"`
	StrRes1         string `gorm:"column:StrRes1" json:"StrRes1"`
	StrRes2         string `gorm:"column:StrRes2" json:"StrRes2"`
	MsgVoiceText    string `gorm:"column:msgVoiceText" json:"msgVoiceText"`
	MsgSeq          int32  `gorm:"column:msgSeq" json:"msgSeq"`
	CompressContent []byte `gorm:"column:CompressContent" json:"CompressContent"`
	ConBlob         []byte `gorm:"column:ConBlob" json:"ConBlob"`
}

// TableName Message's table name
func (*Message) TableName() string {
	return TableNameMessage
}
