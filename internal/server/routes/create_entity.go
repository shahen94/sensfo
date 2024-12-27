package routes

import (
	"github.com/sensfo/server/domain"
	"github.com/sensfo/server/internal/data"
)

// ─────────────────────────────────────────────────────────────────────────────

type CreateEntityRoute struct {
	ds data.DataSource
}

func (r *CreateEntityRoute) Bind(engine domain.Router) {
	engine.POST("/entity", r.createEntityHandler)
}

func (r *CreateEntityRoute) createEntityHandler(ctx domain.Context) {
	result, err := r.ds.Entity().Store(ctx.Query("key"), ctx.Query("value"))

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, result.Content())
}

// ─────────────────────────────────────────────────────────────────────────────

func NewCreateRoute(datasource data.DataSource) domain.Route {
	return &CreateEntityRoute{
		ds: datasource,
	}
}
