package server

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
	"net/http"
	fServer "github.com/itimofeev/hustledb/server/forum"
	"github.com/itimofeev/hustledb/forum"
)

var db *runner.DB

func InitRouter(conn *runner.DB) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")

	api.GET("/competitions", ListCompetitions)
	api.GET("/competitions/:id", ListCompetition)
	api.GET("/dancers", ListDancers)
	api.GET("/dancers/:dancerId", GetDancerInfo)


	fHandlers := fServer.NewForumHandlers(forum.NewCompDao(conn))
	forumApi := api.Group("/forum")
	forumApi.GET("/competitions", fHandlers.ListCompetitions)
	forumApi.POST("/competitions", fHandlers.ParseCompetitions)

	db = conn

	return r
}

func ListCompetitions(c *gin.Context) {
	var params PageParams
	parseParamsGet(c, &params)
	params.fix()

	t := RepoListCompetitions(params)

	WriteJSONStatus(c, t, http.StatusOK)
}

func ListCompetition(c *gin.Context) {
	compId := GetPathInt64Param(c, "id")

	t := RepoGetCompetitionInfo(compId)

	WriteJSONStatus(c, t, http.StatusOK)
}

func ListDancers(c *gin.Context) {
	var params ListDancerParams
	parseParamsGet(c, &params)
	params.fix()

	t := RepoListDancers(params)

	WriteJSONStatus(c, t, http.StatusOK)
}

func GetDancerInfo(c *gin.Context) {
	dancerId := GetPathInt64Param(c, "dancerId")

	t := RepoGetDancerInfo(dancerId)

	WriteJSONStatus(c, t, http.StatusOK)
}
