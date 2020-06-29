/*
Package kaguya : The library for kaguya

Copyright(c) 2020 Star Inc. All Rights Reserved.
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/
package kaguya

import (
	"github.com/gorilla/websocket"
)

type Handle struct {
	identify string
	request  KaguyaRequest
	wsHandle *websocket.Conn
}

func NewHandleInterface(wsHandle *websocket.Conn) *Handle {
	handle := new(Handle)
	handle.wsHandle = wsHandle
	return handle
}

// Start :
func (handle *Handle) Start() {
	for {
		err := handle.wsHandle.ReadJSON(&handle.request)
		DeBug("WS Read", err)
		if handle.request.Version < 1 {
			handle.wsHandle.WriteJSON("Not supported")
			return
		}
		if handle.request.AuthToken != "" {
			go handle.HandleActions()
		} else {
			if handle.request.Action == "auth" {
				handle.identify = handle.request.AuthToken
				handle.wsHandle.WriteJSON("Authorized")
			} else {
				handle.wsHandle.WriteJSON("Not authorized")
			}
		}
	}
}
