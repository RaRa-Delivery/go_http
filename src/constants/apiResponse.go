package constants

import "github.com/RaRa-Delivery/rara-ms-notification/src/middleware"

var SUCCESS = middleware.ConstantDescription{
	Id:   "Kesuksesan",
	En:   "Success",
	Code: "Suc20000",
}

var FAILED = middleware.ConstantDescription{
	Id:   "Gagal",
	En:   "Failed",
	Code: "Err40000",
}
