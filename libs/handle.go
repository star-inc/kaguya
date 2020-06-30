/*
Package kaguya : The library for kaguya

Copyright(c) 2020 Star Inc. All Rights Reserved.
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/
package kaguya

import (
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Handle struct {
	identify      string
	request       KaguyaRequest
	wsHandle      *websocket.Conn
	dataInterface *DataInterface
}

func NewHandleInterface(wsHandle *websocket.Conn) *Handle {
	handle := new(Handle)
	handle.wsHandle = wsHandle
	handle.dataInterface = NewDataInterface()
	return handle
}

// Start :
func (handle *Handle) Start() {
	for {
		err := handle.wsHandle.ReadJSON(&handle.request)
		DeBug("WS Read", err)
		if handle.request.Version < 1 {
			handle.Response(false, "core", "End of Support", nil)
			return
		}
		if handle.request.AuthToken != "" || handle.request.ActionType == "authService" {
			go handle.HandleServices()
		} else {
			go handle.Response(false, "core", "Unauthorized", nil)
		}
	}
}

// Response :
func (handle *Handle) Response(initiative bool, serviceCode string, actionCode string, data interface{}) {
	var actionID string
	now := time.Now().Unix()
	if initiative {
		actionID = uuid.New().String()
	} else {
		actionID = handle.request.ActionID
	}
	handle.wsHandle.WriteJSON(
		&KaguyaResponse{
			Time:       now,
			ActionID:   actionID,
			ActionType: serviceCode,
			Action:     actionCode,
			Data:       data,
		},
	)
}
