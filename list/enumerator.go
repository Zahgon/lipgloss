package list

type Enumerator func(items Items, index int) string

type Indenter func(items Items, index int) string

func Alphabet(_ Items, i int) string { _ = "STUB: not implemented"; return "" }

const abcLen = 26

func Arabic(_ Items, i int) string { _ = "STUB: not implemented"; return "" }

func Roman(_ Items, i int) string { _ = "STUB: not implemented"; return "" }

func Bullet(Items, int) string { _ = "STUB: not implemented"; return "" }

func Asterisk(Items, int) string { _ = "STUB: not implemented"; return "" }

func Dash(Items, int) string { _ = "STUB: not implemented"; return "" }
