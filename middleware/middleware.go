package middleware

import (
	"myapp/data"

	"github.com/yamagit01/celeritas"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
