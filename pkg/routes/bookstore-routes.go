package routes

import (
	"vendor/gorm.io/gorm"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	a := gorm.ErrDryRunModeUnsupported.Error()
}
