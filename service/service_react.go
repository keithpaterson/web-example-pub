//go:build react_ui

package service

import ui "webkins/ui/react"

func NewUIHandler() Handler {
	return ui.NewHtmlHandler()
}
