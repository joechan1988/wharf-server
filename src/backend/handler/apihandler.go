package handler

import (
	restful "github.com/emicklei/go-restful"
	"net/http"
	"wharf/wharf-server/src/backend/res"
	"wharf/wharf-server/src/backend/common"
	"github.com/golang/glog"
)

type APIHandler struct {
	StoreHandler *common.StoreHandler
}

func CreateHTTPAPIHandler(storeHandler *common.StoreHandler) (http.Handler, error) {

	wsContainer := restful.NewContainer()
	wsContainer.EnableContentEncoding(true)
	apiHandler := APIHandler{StoreHandler: storeHandler}
	apiV1Ws := new(restful.WebService)

	apiV1Ws.Path("/api/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	wsContainer.Add(apiV1Ws)

	apiV1Ws.Route(
		apiV1Ws.GET("/token").To(apiHandler.handleGetToken).Writes(res.Token{}))

	return wsContainer, nil
}

func (apiHandler *APIHandler) handleGetToken(request *restful.Request, response *restful.Response) {

	glog.Infof("Request to get token...", request)

	token := request.QueryParameter("token")
	result, err := res.GetToken(apiHandler.StoreHandler,token)
	if err != nil {
		glog.Fatalf("Get Token Error", err)
		return
	}

	response.WriteHeaderAndEntity(http.StatusOK, result)
}
