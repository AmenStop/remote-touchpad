/*
 *    Copyright (c) 2018 Unrud<unrud@outlook.com>
 *
 *    This file is part of Remote-Touchpad.
 *
 *    Foobar is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU General Public License as published by
 *    the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    Remote-Touchpad is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *
 *   You should have received a copy of the GNU General Public License
 *   along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

type PluginInfo struct {
	Name string
	Init func() (Plugin, error)
}

var Plugins []PluginInfo = []PluginInfo{
	{"X11", InitX11Plugin},
	{"RemoteDesktop portal", InitPortalPlugin},
	{"Windows", InitWindowsPlugin},
}

type UnsupportedPlatformError struct {
	err error
}

func (e UnsupportedPlatformError) Error() string {
	return e.err.Error()
}

type Plugin interface {
	Close() error
	KeyboardText(text string) error
	PointerButton(button uint, press bool) error
	PointerMove(deltaX, deltaY int) error
	PointerScroll(deltaHorizontal, deltaVertical int) error
	PointerScrollFinish() error
}
