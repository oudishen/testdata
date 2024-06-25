package database

import "github.com/oudishen/testdata/pkg"

// App is a struct that represents the app table.
//
//	@TABLE
//	@UX(UxCode=Code)
//	@IDX(IdxOrgID=OrgID)
type App struct {
	pkg.CommonField

	AppCode string `db:"app_code" json:"app_code"`
	AppName string `db:"app_name" json:"app_name"`
	OrgID   int64  `db:"org_id" json:"org_id"`
}
