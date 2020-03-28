//** This file was code generated by got. DO NOT EDIT. ***

package panels

import (
	"bytes"
	"context"
)

func (ctrl *ImageCapturePanel) DrawTemplate(ctx context.Context, buf *bytes.Buffer) (err error) {

	buf.WriteString(`
<h2>Image Capture Widget</h2>
<p>
The ImageCapture widget captures images from devices that have a camera, and is an example of a GoRADD control
that has additional javascript to help it work. The example below clips the image with a circle, which is a common
thing to do when capturing images for chat groups and messaging applications.
</p>
<p>
After capturing the image once, reload the page. You will see that the image is preserved, and makes a round-trip
from the original capture, to the server, and then when you reload the page, it is sent back to the browser.
</p>

`)

	buf.WriteString(`
`)
	if `` == "" {
		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("ic1-ff").Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	} else {

		buf.WriteString(`    `)

		{
			err := ctrl.Page().GetControl("ic1-ff").ProcessAttributeString(``).Draw(ctx, buf)
			if err != nil {
				return err
			}
		}
	}

	buf.WriteString(`
`)

	buf.WriteString(`
`)

	return
}
