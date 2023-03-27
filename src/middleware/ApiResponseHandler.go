package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	routing "github.com/qiangxue/fasthttp-routing"
)

type ResultStatus struct {
	Status  string      `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"data" bson:"data"`
	Code    string      `json:"code" bson:"code"`
}

func BuildJsonResponse(reqId string, resp ResultStatus, statusCode int, ctx *routing.Context) error {

	log.Println(reqId + " - Time of Request: " + ctx.RequestCtx.Time().String() + ", Request Body: " + string(ctx.PostBody()))

	bytes, err := json.Marshal(resp)
	if err != nil {
		ctx.SetStatusCode(500)
		log.Println(reqId + ": " + err.Error())
		return err
	}
	ctx.SetStatusCode(statusCode)
	ctx.SetContentType("application/json")
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,POST,DELETE,OPTIONS")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "*")
	fmt.Fprintf(ctx, "%+v", string(bytes))

	log.Println(reqId + " - Time of Response: " + time.Now().String() + ", Response Body: " + string(bytes))

	return nil
}

type ConstantDescription struct {
	Id   string `json:"id" bson:"id"`
	En   string `json:"en" bson:"en"`
	Code string `json:"code" bson:"code"`
}

func GetResponseMessage(status ConstantDescription, lang string) string {
	switch lang {
	case "en":
		return status.En
	case "id":
		return status.Id
	default:
		return status.En
	}
}

func GetResponseMessageWithId(status ConstantDescription, lang string, id string) string {
	switch lang {
	case "en":
		return status.En + " for piece id : " + id
	case "id":
		return status.Id + " for piece id : " + id
	default:
		return status.En + " for piece id : " + id
	}
}
