// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameGroupUserRelation = "GroupUserRelation"

// GroupUserRelation mapped from table <GroupUserRelation>
type GroupUserRelation struct {
	UserName      string `gorm:"column:userName;primaryKey" json:"userName"`
	GroupNameList string `gorm:"column:groupNameList" json:"groupNameList"`
}

// TableName GroupUserRelation's table name
func (*GroupUserRelation) TableName() string {
	return TableNameGroupUserRelation
}
