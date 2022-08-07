package models

//go:generate toolkit gen model Applet --database Demo
// Applet database model demo
// @def primary                     ID
// @def unique_index UI_applet_name Name
// @def unique_index UI_applet_id   AppletID
type Applet struct {
	PrimaryID
	RefApplet
	AppletInfo
	OperationTimes
}

type RefApplet struct {
	AppletID string `db:"f_applet_id" json:"appletID"`
}

type AppletInfo struct {
	Name string `db:"f_name" json:"name"`
}
