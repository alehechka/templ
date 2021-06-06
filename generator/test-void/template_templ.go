// Code generated by templ DO NOT EDIT.

package testvoid

import "github.com/a-h/templ"
import "context"
import "io"

func render() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, "<br/>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<img")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, " src=\"https://example.com/image.png\"")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "/>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<br/>")
		if err != nil {
			return err
		}
		_, err = io.WriteString(w, "<br/>")
		if err != nil {
			return err
		}
		return err
	})
}

