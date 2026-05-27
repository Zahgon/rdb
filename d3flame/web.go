package d3flame

import (
	"html/template"
	"net/http"
)

var flameTmplData = &struct {
	D3Css        template.CSS
	D3Js         template.JS
	D3Flame      template.JS
	D3Tip        template.JS
	BootstrapCSS template.CSS
}{
	D3Css:        template.CSS(d3Css),
	D3Js:         template.JS(d3Js),
	D3Flame:      template.JS(d3FlameGraphJs),
	D3Tip:        template.JS(d3TipJs),
	BootstrapCSS: template.CSS(bootstrapCSS),
}

func flamegraph(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// FlameItem is an Element in flamegraph
type FlameItem struct {
	Name     string   `json:"n"`
	Value    int      `json:"v"`
	Children children `json:"c,omitempty"`
}

func (ch children) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// AddChild add a child node into FlameItem
func (node *FlameItem) AddChild(n *FlameItem) { _ = "STUB: not implemented"; return }

type children map[string]*FlameItem

// Web starts a web server to render flamegraph
func Web(data []byte, port int) chan<- struct{} { _ = "STUB: not implemented"; return nil }
