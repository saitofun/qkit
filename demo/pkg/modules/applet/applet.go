package applet

import (
	"context"

	"github.com/google/uuid"

	"github.com/saitofun/qkit/demo/cmd/demo/global"
	"github.com/saitofun/qkit/demo/pkg/errors/status"
	"github.com/saitofun/qkit/demo/pkg/models"
	"github.com/saitofun/qkit/kit/sqlx"
	"github.com/saitofun/qkit/kit/sqlx/builder"
)

type CreateAppletByNameReq struct {
	Name string `json:"name"`
}

func CreateAppletByName(ctx context.Context, req *CreateAppletByNameReq) (*models.Applet, error) {
	applet := &models.Applet{
		RefApplet:  models.RefApplet{AppletID: uuid.New().String()},
		AppletInfo: models.AppletInfo{Name: req.Name},
	}

	d := global.DBExecutorFromContext(ctx)
	l := global.LoggerFromContext(ctx)

	l.Start(ctx, "CreateAppletByName")
	defer l.End()

	err := sqlx.NewTasks(d).With(
		func(db sqlx.DBExecutor) error {
			return applet.Create(db)
		},
		func(db sqlx.DBExecutor) error {
			return applet.FetchByAppletID(db)
		},
	).Do()
	if err != nil {
		l.Error(err)
		if sqlx.DBErr(err).IsConflict() {
			return nil, status.Conflict.StatusErr().WithMsg("create applet conflict")
		}
		return nil, err
	}

	return applet, nil
}

type ListAppletReq struct {
	IDs       []string `in:"query" name:"id,omitempty"`
	AppletIDs []string `in:"query" name:"appletID,omitempty"`
	Names     []string `in:"query" name:"name,omitempty"`
	builder.Pager
}

func (r *ListAppletReq) Condition() builder.SqlCondition {
	var (
		m  = &models.Applet{}
		cs []builder.SqlCondition
	)
	if len(r.IDs) > 0 {
		cs = append(cs, m.ColID().In(r.IDs))
	}
	if len(r.AppletIDs) > 0 {
		cs = append(cs, m.ColAppletID().In(r.AppletIDs))
	}
	if len(r.Names) > 0 {
		cs = append(cs, m.ColName().In(r.Names))
	}
	return builder.And(cs...)
}

func (r *ListAppletReq) Additions() builder.Additions {
	m := &models.Applet{}
	return builder.Additions{
		builder.OrderBy(builder.DescOrder(m.ColCreatedAt())),
		r.Pager.Addition(),
	}
}

func ListApplets(ctx context.Context, r *ListAppletReq) ([]models.Applet, error) {
	applet := &models.Applet{}

	d := global.DBExecutorFromContext(ctx)
	l := global.LoggerFromContext(ctx)

	l.Start(ctx, "ListApplets")
	defer l.End()

	applets, err := applet.List(d, r.Condition(), r.Additions()...)
	if err != nil {
		l.Error(err)
		return nil, err
	}
	return applets, nil
}
