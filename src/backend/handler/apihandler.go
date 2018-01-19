package handler

import (
	restful "github.com/emicklei/go-restful"
	"net/http"
	"wharf/wharf-server/src/backend/res"

	"log"
)

type APIHandler struct {

}

func CreateHTTPAPIHandler() (http.Handler, error) {

	wsContainer := restful.NewContainer()
	wsContainer.EnableContentEncoding(true)
	apiHandler := APIHandler{}
	apiV1Ws := new(restful.WebService)

	apiV1Ws.Path("/api/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	wsContainer.Add(apiV1Ws)

	apiV1Ws.Route(
		apiV1Ws.GET("/token").To(apiHandler.handleGetToken).Writes(res.Token{}))

	return wsContainer,nil
}

func (apiHandler *APIHandler) handleGetToken(request *restful.Request,response *restful.Response){


	log.Print("Request to get token...",request)
	result,err := res.GetToken()
	if err != nil {
		log.Fatalf("Get Token Error",err)
		return
	}

	response.WriteHeaderAndEntity(http.StatusOK,result)
}