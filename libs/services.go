/*
Package kaguya : The library for kaguya

Copyright(c) 2020 Star Inc. All Rights Reserved.
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/
package kaguya

func (handle *Handle) HandleServices() {
	switch handle.request.ActionType {
	case "authService":
		handle.AuthService()
		break
	case "talkService":
		handle.TalkService()
		break
	}
}

func (handle *Handle) AuthService() {
	switch handle.request.Action{
	case "getAccess":
		data := GetAccess(handle.request.Data)
		handle.Response(handle.request.ActionType, handle.request.Action, data)
		break
	}
}

func (handle *Handle) TalkService() {

}
