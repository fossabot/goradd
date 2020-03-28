//** This file was code generated by got. DO NOT EDIT. ***

package tutorial

import (
	"bytes"
	"context"

	"github.com/goradd/goradd/pkg/page"
)

func (ctrl *IndexForm) AddHeadTags() {
	ctrl.FormBase.AddHeadTags()
	if "Tutorial" != "" {
		ctrl.Page().SetTitle("Tutorial")
	}

	// double up to deal with body attributes if they exist
	ctrl.Page().BodyAttributes = `
`
}

func (ctrl *IndexForm) DrawTemplate(ctx context.Context, buf *bytes.Buffer) (err error) {

	buf.WriteString(`
`)

	buf.WriteString(`<script>
function toggleSidebar() {
    g$('sidebar').toggleClass('open');
    g$('content').toggleClass('open');
}
</script>
`)

	buf.WriteString(`

`)
	path := page.GetContext(ctx).HttpContext.URL.Path
	buf.WriteString(`<div id="sidebar" class="open">
    <a href="javascript:void(0)" id="togglebtn" onclick="toggleSidebar();"><span id="isopen">&lt;</span><span id="isclosed">&gt;</span></a>
    <div id="sidebar_content">
        <h2><a href="/goradd/tutorial.g">Home</a></h2>
        <h2>ORM</h2>
          <ul>
        `)
	for _, pr := range pages["orm"] {
		buf.WriteString(`            <li><a href="`)

		buf.WriteString(path)

		buf.WriteString(`?pageID=orm-`)

		buf.WriteString(pr.id)

		buf.WriteString(`">`)

		buf.WriteString(pr.title)

		buf.WriteString(`</a></li>
        `)
	}

	buf.WriteString(`          </ul>
  </div>
</div>
<div id="content" class="open">
<h1>Tutorial</h1>
`)

	buf.WriteString(`
`)
	if `` == "" {
		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("viewSourceButton").Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	} else {

		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("viewSourceButton").ProcessAttributeString(``).Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	}

	buf.WriteString(`
<div id="detail_container">
	`)

	buf.WriteString(`
`)
	if `` == "" {
		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("detailPanel").Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	} else {

		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("detailPanel").ProcessAttributeString(``).Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	}

	buf.WriteString(`
</div>
</div>

`)

	buf.WriteString(`
`)

	return
}
