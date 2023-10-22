// Code generated by templ@v0.2.364 DO NOT EDIT.

package public

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func RedirectLogin(marshaledSession string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		err = DocType().Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = Head().Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		err = Body(pageRedirectLogin(marshaledSession)).Render(ctx, templBuffer)
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}

func pageRedirectLogin(marshaledSession string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templBuffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_2 := templ.GetChildren(ctx)
		if var_2 == nil {
			var_2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, err = templBuffer.WriteString("<p>")
		if err != nil {
			return err
		}
		var_3 := `This is the session information after Login`
		_, err = templBuffer.WriteString(var_3)
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p><p>")
		if err != nil {
			return err
		}
		var var_4 string = marshaledSession
		_, err = templBuffer.WriteString(templ.EscapeString(var_4))
		if err != nil {
			return err
		}
		_, err = templBuffer.WriteString("</p>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = templBuffer.WriteTo(w)
		}
		return err
	})
}
