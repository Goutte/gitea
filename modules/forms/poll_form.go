package forms

import (
	"net/http"

	"code.gitea.io/gitea/modules/context"
	"code.gitea.io/gitea/modules/web/middleware"

	"gitea.com/go-chi/binding"
)

// Form for creating a poll
type CreatePollForm struct {
	Subject     string `binding:"Required;MaxSize(128)"` // 128 is duplicated in the template
	Description string
}

// Validate the form fields
func (f *CreatePollForm) Validate(req *http.Request, errs binding.Errors) binding.Errors {
	ctx := context.GetContext(req)
	return middleware.Validate(errs, ctx.Data, f, ctx.Locale)
}
