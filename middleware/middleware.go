package middleware

import (
	"myapp/data"

	"github.com/virmos/django"
)

type Middleware struct {
	App *celeritas.Celeritas
	Models data.Models
}