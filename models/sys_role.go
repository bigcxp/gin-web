package models

const (
	// 用户状态
	SysRoleStatusDisabled    uint   = 0    // 禁用
	SysRoleStatusNormal      uint   = 1    // 正常
	SysRoleStatusDisabledStr string = "禁用" // 禁用
	SysRoleStatusNormalStr   string = "正常" // 正常
)

// 定义map方便取值
var SysRoleStatusConst = map[uint]string{
	SysRoleStatusDisabled: SysRoleStatusDisabledStr,
	SysRoleStatusNormal:   SysRoleStatusNormalStr,
}

// 系统角色表
type SysRole struct {
	Model
	Name    string    `gorm:"comment:'角色名称'" json:"name"`
	Keyword string    `gorm:"unique;comment:'角色关键词'" json:"keyword"`
	Desc    string    `gorm:"comment:'角色说明'" json:"desc"`
	Status  *uint     `gorm:"type:tinyint(1);default:1;comment:'角色状态(正常/禁用, 默认正常)'" json:"status"` // 由于设置了默认值, 这里使用ptr, 可避免赋值失败
	Creator string    `gorm:"comment:'创建人'" json:"creator"`
	Menus   []SysMenu `gorm:"many2many:relation_role_menu;" json:"menus"` // 角色菜单多对多关系
	Users   []SysUser `gorm:"foreignkey:RoleId"`                          // 一个角色有多个user
}

func (m SysRole) TableName() string {
	return m.Model.TableName("sys_role")
}
