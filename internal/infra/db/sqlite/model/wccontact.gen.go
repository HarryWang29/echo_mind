// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameWCContact = "WCContact"

// WCContact mapped from table <WCContact>
type WCContact struct {
	MNsUsrName           string `gorm:"column:m_nsUsrName;primaryKey" json:"m_nsUsrName"`
	MUIConType           int32  `gorm:"column:m_uiConType" json:"m_uiConType"`
	Nickname             string `gorm:"column:nickname" json:"nickname"`
	MNsFullPY            string `gorm:"column:m_nsFullPY" json:"m_nsFullPY"`
	MNsShortPY           string `gorm:"column:m_nsShortPY" json:"m_nsShortPY"`
	MNsRemark            string `gorm:"column:m_nsRemark" json:"m_nsRemark"`
	MNsRemarkPYFull      string `gorm:"column:m_nsRemarkPYFull" json:"m_nsRemarkPYFull"`
	MNsRemarkPYShort     string `gorm:"column:m_nsRemarkPYShort" json:"m_nsRemarkPYShort"`
	MUICertificationFlag int32  `gorm:"column:m_uiCertificationFlag" json:"m_uiCertificationFlag"`
	MUISex               int32  `gorm:"column:m_uiSex" json:"m_uiSex"`
	MUIType              int32  `gorm:"column:m_uiType" json:"m_uiType"`
	MNsImgStatus         string `gorm:"column:m_nsImgStatus" json:"m_nsImgStatus"`
	MUIImgKey            int32  `gorm:"column:m_uiImgKey" json:"m_uiImgKey"`
	MNsHeadImgURL        string `gorm:"column:m_nsHeadImgUrl" json:"m_nsHeadImgUrl"`
	MNsHeadHDImgURL      string `gorm:"column:m_nsHeadHDImgUrl" json:"m_nsHeadHDImgUrl"`
	MNsHeadHDMd5         string `gorm:"column:m_nsHeadHDMd5" json:"m_nsHeadHDMd5"`
	MNsChatRoomMemList   string `gorm:"column:m_nsChatRoomMemList" json:"m_nsChatRoomMemList"`
	MNsChatRoomAdminList string `gorm:"column:m_nsChatRoomAdminList" json:"m_nsChatRoomAdminList"`
	MUIChatRoomStatus    int32  `gorm:"column:m_uiChatRoomStatus" json:"m_uiChatRoomStatus"`
	MNsChatRoomDesc      string `gorm:"column:m_nsChatRoomDesc" json:"m_nsChatRoomDesc"`
	MNsDraft             string `gorm:"column:m_nsDraft" json:"m_nsDraft"`
	MNsBrandIconURL      string `gorm:"column:m_nsBrandIconUrl" json:"m_nsBrandIconUrl"`
	MNsGoogleContactName string `gorm:"column:m_nsGoogleContactName" json:"m_nsGoogleContactName"`
	MNsAliasName         string `gorm:"column:m_nsAliasName" json:"m_nsAliasName"`
	MNsEncodeUserName    string `gorm:"column:m_nsEncodeUserName" json:"m_nsEncodeUserName"`
	MUIChatRoomVersion   int32  `gorm:"column:m_uiChatRoomVersion" json:"m_uiChatRoomVersion"`
	MUIChatRoomMaxCount  int32  `gorm:"column:m_uiChatRoomMaxCount" json:"m_uiChatRoomMaxCount"`
	MUIChatRoomType      int32  `gorm:"column:m_uiChatRoomType" json:"m_uiChatRoomType"`
	MPatSuffix           string `gorm:"column:m_patSuffix" json:"m_patSuffix"`
	RichChatRoomDesc     string `gorm:"column:richChatRoomDesc" json:"richChatRoomDesc"`
	PackedWCContactData  []byte `gorm:"column:_packed_WCContactData" json:"_packed_WCContactData"`
	OpenIMInfo           []byte `gorm:"column:openIMInfo" json:"openIMInfo"`
}

// TableName WCContact's table name
func (*WCContact) TableName() string {
	return TableNameWCContact
}
