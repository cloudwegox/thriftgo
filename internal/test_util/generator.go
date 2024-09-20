// Copyright 2023 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test_util

import (
	"github.com/cloudwegox/thriftgo/generator"
	"github.com/cloudwegox/thriftgo/generator/backend"
	"github.com/cloudwegox/thriftgo/generator/golang"
	"github.com/cloudwegox/thriftgo/parser"
	"github.com/cloudwegox/thriftgo/plugin"
	"github.com/cloudwegox/thriftgo/semantic"
)

func GenerateGolang(idl string, output string, genOpts []plugin.Option, pluginOpts []*plugin.Desc) (generator.Generator, *plugin.Response) {
	ast, err := parser.ParseFile(idl, nil, true)
	if err != nil {
		panic(err)
	}

	checker := semantic.NewChecker(semantic.Options{FixWarnings: true})
	resolver, ok := checker.(interface {
		ResolveSymbols(t *parser.Thrift) error
	})
	if ok {
		if err = resolver.ResolveSymbols(ast); err != nil {
			panic(err)
		}
	}

	var gen generator.Generator
	if err := gen.RegisterBackend(new(golang.GoBackend)); err != nil {
		panic(err)
	}

	log := backend.LogFunc{
		Info:      func(v ...interface{}) {},
		Warn:      func(v ...interface{}) {},
		MultiWarn: func(warns []string) {},
	}
	out := &generator.LangSpec{
		Language:    "go",
		Options:     genOpts,
		UsedPlugins: pluginOpts,
	}
	req := &plugin.Request{
		Language:   out.Language,
		Version:    "?",
		OutputPath: output,
		Recursive:  true,
		AST:        ast,
	}
	arg := &generator.Arguments{Out: out, Req: req, Log: log}

	res := gen.Generate(arg)

	if v := res.GetError(); v != "" {
		panic(v)
	}

	return gen, res
}
