package main

import (
	"github.com/furu2revival/musicbox/cmd/protoc-gen-musicbox-server/handler"
	"github.com/furu2revival/musicbox/protobuf/custom_option"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	generator := handler.NewGenerator[*custom_option.MethodOption](handler.Config{
		MethodOptExt:      custom_option.E_MethodOption,
		MethodOptIdent:    protogen.GoImportPath("github.com/furu2revival/musicbox/app/infrastructure/connect/aop").Ident("MethodOption"),
		MethodOptExtIdent: protogen.GoImportPath("github.com/furu2revival/musicbox/protobuf/custom_option").Ident("E_MethodOption"),
		MethodErrDefIdent: protogen.GoImportPath("github.com/furu2revival/musicbox/app/infrastructure/connect/aop").Ident("MethodErrDefinition"),
		ProxyIdent:        protogen.GoImportPath("github.com/furu2revival/musicbox/app/infrastructure/connect/aop").Ident("Proxy"),
	})
	protogen.Options{}.Run(func(plugin *protogen.Plugin) error {
		for _, file := range plugin.Files {
			if file.Desc.Package() != "api" && file.Desc.Package() != "api.debug" {
				continue
			}
			err := generator.Generate(plugin, file)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
