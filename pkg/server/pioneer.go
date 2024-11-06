package server

import(
	"fmt"
	"net/http"
	"log/slog"
	"github.com/gin-gonic/gin"
	"github.com/influenzanet/study-service/pkg/studyengine"
	"github.com/grippenet/badge-service/pkg/types"
)

func serviceResponse(c *gin.Context, value string) {
	c.JSON(http.StatusOK, map[string]string{"value": value})
}

func errorResponse(c *gin.Context, msg string) {
	c.JSON(http.StatusBadRequest, map[string]string{"error": msg})
}

type pioneerHandler struct {
	service types.PioneerService
}

func (h *pioneerHandler) Handle(c *gin.Context) {
	var data studyengine.ExternalEventPayload
	err := c.ShouldBind(&data) 
	if(err != nil) {
		errorResponse(c, fmt.Sprintf("%s", err))
		return	
	}

	if(data.EventType != "SUBMIT") {
		errorResponse(c, "Only usable with SUBMIT event")
		return
	}

	studyKey := data.StudyKey
	instanceID := data.InstanceID


	if(data.Response.Key != "intake") {
		serviceResponse(c, "")
		return
	}
	postalCodeItemResponse, err := findSurveyItemResponse(data.Response.Responses, "intake.main.Q3")
	if err != nil {
		serviceResponse(c, "0")
		return
	}
	postalCodeResponse, err := findResponseObject(postalCodeItemResponse, "rg.0")
	if(err != nil) {
		slog.Error(fmt.Sprintf("postalcode response not found: %s", err), "participant", data.ParticipantState.ParticipantID)
		serviceResponse(c, "0")
		return
	}

	postalCode := postalCodeResponse.Value
	
	r, err := h.service.Check(instanceID, studyKey, postalCode)
	if(err != nil) {
		slog.Error(fmt.Sprintf("Unable to check postal code: %s", err), "participant", data.ParticipantState.ParticipantID)
		serviceResponse(c, "0")
		return
	}
	var v string 
	if r {
		v = "1"
	} else {
		v = "0"
	}
	serviceResponse(c, v)
	//c.JSON(http.StatusOK, data)
}