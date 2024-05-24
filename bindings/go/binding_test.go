package tree_sitter_ini_test

import (
"context"
"fmt"
"testing"

sitter "github.com/smacker/go-tree-sitter"
tree_sitter_ini "github.com/tree-sitter/tree-sitter-ini"
)

func TestCanLoadGrammar(t *testing.T) {
	language := sitter.NewLanguage(tree_sitter_ini.Language())
	if language == nil {
		t.Errorf("Error loading Ini grammar")
	}

sourceCode := []byte(`
[default]
region = us-west-2
output = json

[profile dev-user]
region = us-east-1
output = text

[profile developers]
role_arn = arn:aws:iam::123456789012:role/developers
source_profile = dev-user
region = us-west-2
output = json
`)

	parser := sitter.NewParser()
	parser.SetLanguage(language)

	tree, _ := parser.ParseCtx(context.Background(), nil, sourceCode)

	n := tree.RootNode()

	fmt.Println(n) // (program (lexical_declaration (variable_declarator (identifier) (number))))

}