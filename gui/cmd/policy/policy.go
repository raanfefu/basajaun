package policy

import (
	"log"
	"os"

	"github.com/kluctl/go-jinja2"
	"github.com/raanfefu/basajaun/gui/pkg"
)

const TEMPLATEPATH = "/app/template/policy.tpl"

func PostPolicy(bodyPolicy pkg.Rule) *pkg.TemplateResponse {

	bfile, err := os.ReadFile(TEMPLATEPATH)
	if err != nil {
		log.Panic(err)
	}
	j2, err := jinja2.NewJinja2("policy.tpl", 1)
	if err != nil {
		log.Panic(err)
	}
	data := jinja2.WithGlobal("data", bodyPolicy)
	signature := jinja2.WithGlobal("signature", "B41BD5F462719C6D6118E673A2389")
	s, err := j2.RenderString(string(bfile), data, signature)
	if err != nil {
		log.Panic(err)
	}

	response := &pkg.TemplateResponse{
		Rule:   &bodyPolicy,
		Render: &s,
	}

	return response

}
