package main

import (
	"context"
	"fmt"

	"github.com/SamHennessy/hlive"
	"github.com/SamHennessy/hlive/hlivekit"
	"github.com/SamHennessy/isles"
)

func app(ctxWails context.Context) *hlive.Page {
	page := hlive.NewPage()

	inputVal := hlive.Box("")
	greetMsg := hlive.Box("Please enter your name below 👇")

	themeAttr, themeSelect := isles.ThemePicker("select-bordered", "")

	page.DOM().HTML().Add(themeAttr)
	page.DOM().Head().Add(hlive.T("style", hlive.HTML(`
.logo {
	font-family: "Nunito", -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto",
	"Oxygen", "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue",
	sans-serif;
}

@font-face {
	font-family: "Nunito";
	font-style: normal;
	font-weight: 400;
	src: local(""),
	url("/assets/fonts/nunito-v16-latin-regular.woff2") format("woff2");
}`)))

	page.DOM().Body().Add(
		isles.Navbar("bg-neutral mb-4",
			isles.NavbarStart("", hlive.T("div", hlive.Class("logo btn btn-ghost normal-case text-xl"), "Isles Demo")),
			isles.NavbarCenter(""),
			isles.NavbarEnd("",
				hlive.T("div", hlive.Class("text-neutral-content"), "Theme"),
				hlive.T("div", hlive.Class("m-2"), themeSelect),
			),
		),
		hlive.T("img", hlive.Class("w-48 h-48 mx-auto"), hlive.Attrs{"src": "assets/images/logo-universal.png"}),

		isles.Container("mx-auto w-96",
			isles.Div("m-4", greetMsg),
			isles.Form(
				"",
				func(_ context.Context, _ hlive.Event) {
					greetMsg.Set(fmt.Sprintf("Hello %s, It's show time!", inputVal.Get()))
				},
				isles.Join(
					isles.InputText(
						"join-item input-bordered ",
						"",
						func(_ context.Context, e hlive.Event) {
							inputVal.Set(e.Value)
						},
						hlive.Attrs{"autocomplete": "off"},
						hlivekit.Focus(),
						hlive.OnOnce("focus", func(ctx context.Context, e hlive.Event) {
							if c, ok := e.Binding.Component.(hlive.Adder); ok {
								hlivekit.FocusRemove(c)
							}
						}),
					),
					isles.Button("join-item btn-primary rounded-r-full", nil, "Greet"),
				),
			),
		),
	)

	return page
}
