package tree

type Enumerator func(children Children, index int) string

func DefaultEnumerator(children Children, index int) string { _ = "STUB: not implemented"; return "" }

func RoundedEnumerator(children Children, index int) string { _ = "STUB: not implemented"; return "" }

type Indenter func(children Children, index int) string

func DefaultIndenter(children Children, index int) string { _ = "STUB: not implemented"; return "" }
