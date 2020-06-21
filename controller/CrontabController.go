package controller

import (
	"database/sql"
	"goploy/core"
	"goploy/model"
	"net/http"
)

// Crontab struct
type Crontab Controller

// GetList crontab list
func (crontab Crontab) GetList(w http.ResponseWriter, gp *core.Goploy) *core.Response {
	type RespData struct {
		Crontab    model.Crontabs   `json:"crontabList"`
		Pagination model.Pagination `json:"pagination"`
	}
	pagination, err := model.PaginationFrom(gp.URLQuery)
	if err != nil {
		return &core.Response{Code: core.Error, Message: err.Error()}
	}

	crontabList, pagination, err := model.Crontab{}.GetList(pagination)

	if err != nil {
		return &core.Response{Code: core.Error, Message: err.Error()}
	}
	return &core.Response{Data: RespData{Crontab: crontabList, Pagination: pagination}}
}

// Add one crontab
func (crontab Crontab) Add(w http.ResponseWriter, gp *core.Goploy) *core.Response {
	type ReqData struct {
		Command string `json:"command" validate:"required"`
	}
	type RespData struct {
		ID int64 `json:"id"`
	}
	var reqData ReqData
	if err := verify(gp.Body, &reqData); err != nil {
		return &core.Response{Code: core.Error, Message: err.Error()}
	}

	crontabInfo, err := model.Crontab{Command: reqData.Command}.GetDataByCommand()
	if err != nil && err != sql.ErrNoRows {
		return &core.Response{Code: core.Error, Message: err.Error()}
	} else if crontabInfo != (model.Crontab{}) {
		return &core.Response{Code: core.Error, Message: "Command is already exist"}
	}

	id, err := model.Crontab{
		Command:   reqData.Command,
		Creator:   gp.UserInfo.Name,
		CreatorID: gp.UserInfo.ID,
	}.AddRow()

	if err != nil {
		return &core.Response{Code: core.Error, Message: err.Error()}

	}
	return &core.Response{Data: RespData{ID: id}}
}

// Edit one crontab
func (crontab Crontab) Edit(w http.ResponseWriter, gp *core.Goploy) *core.Response {
	type ReqData struct {
		ID      int64  `json:"id" validate:"gt=0"`
		Command string `json:"command" validate:"required"`
	}
	var reqData ReqData
	if err := verify(gp.Body, &reqData); err != nil {
		return &core.Response{Code: core.Error, Message: err.Error()}
	}

	crontabInfo, err := model.Crontab{Command: reqData.Command}.GetDataByCommand()
	if err != nil {
		if err != sql.ErrNoRows {
			return &core.Response{Code: core.Error, Message: err.Error()}
		}
	} else if crontabInfo.ID != reqData.ID {
		return &core.Response{Code: core.Error, Message: "Command is already exist"}
	}
	
	err = model.Crontab{
		ID:       reqData.ID,
		Command:  reqData.Command,
		Editor:   gp.UserInfo.Name,
		EditorID: gp.UserInfo.ID,
	}.EditRow()

	if err != nil {
		response := &core.Response{Code: core.Deny, Message: err.Error()}
		response.JSON(w)
		return &core.Response{Code: core.Error, Message: err.Error()}
	}
	return &core.Response{}
}

// Remove one Crontab
func (crontab Crontab) Remove(w http.ResponseWriter, gp *core.Goploy) *core.Response {
	type ReqData struct {
		ID int64 `json:"id" validate:"gt=0"`
	}
	var reqData ReqData
	if err := verify(gp.Body, &reqData); err != nil {
		return &core.Response{Code: core.Error, Message: err.Error()}
	}
	err := model.Crontab{
		ID: reqData.ID,
	}.Remove()

	if err != nil {
		return &core.Response{Code: core.Error, Message: err.Error()}
	}
	return &core.Response{}
}
