package phonetics

import (
	"regexp"

	"github.com/njnest/go-nlang/distance"
	"github.com/njnest/go-nlang/tokenizer"
)

// Metaphone phonetics model
type Metaphone struct {
	textA, textB string
}

// SoundTheSame - the function check if the texts contained the same sounds
func (m *Metaphone) SoundTheSame() bool {
	t := tokenizer.NewTokenizer(`([a-zA-ZА-Яа-я])\w+`)
	tokensA := t.Tokenize(m.textA)
	tokensB := t.Tokenize(m.textB)
	return distance.Compare(m.textA, m.textB) > 0.9 && len(tokensA) == len(tokensB)
}

// Compare strings
func Compare(textA, textB string) Metaphone {

	phTextA := phoneticProcessing(textA)
	phTextB := phoneticProcessing(textB)

	return Metaphone{textA: phTextA, textB: phTextB}
}

//bring to a common phonetic kind
func phoneticProcessing(token string) string {
	token = transformCK(token)
	token = cTransform(token)
	token = dTransform(token)
	token = dropG(token)
	token = transformG(token)
	token = dropH(token)
	token = transformPH(token)
	token = transformQ(token)
	token = transformS(token)
	token = transformX(token)
	token = transformT(token)
	token = dropT(token)
	token = transformV(token)
	token = transformWH(token)
	token = dropW(token)
	token = dropY(token)
	token = transformZ(token)
	return token
}

func transformCK(token string) string {
	return regexp.MustCompile("ck").ReplaceAllLiteralString(token, "k")
}

func cTransform(token string) string {
	token = regexp.MustCompile("([^s]|^)(c)(h)").ReplaceAllLiteralString(token, "")
	token = regexp.MustCompile("cia").ReplaceAllLiteralString(token, "xia")
	token = regexp.MustCompile("c(i|e|y)").ReplaceAllLiteralString(token, "")
	token = regexp.MustCompile("c").ReplaceAllLiteralString(token, "k")

	return token
}

func dTransform(token string) string {
	return regexp.MustCompile("d").ReplaceAllLiteralString(token, "t")
}

func dropG(token string) string {
	token = regexp.MustCompile("(g+h)(^$|[^aeiou])").ReplaceAllLiteralString(token, "")
	token = regexp.MustCompile("g(n|ned)$").ReplaceAllLiteralString(token, "")
	return token
}

func transformG(token string) string {
	token = regexp.MustCompile("gh").ReplaceAllLiteralString(token, "f")
	token = regexp.MustCompile("gg").ReplaceAllLiteralString(token, "g")
	token = regexp.MustCompile("g").ReplaceAllLiteralString(token, "k")
	return token
}

func dropH(token string) string {
	return regexp.MustCompile("([aeiou])h([^aeiou]|$)").ReplaceAllLiteralString(token, "")
}

func transformPH(token string) string {
	return regexp.MustCompile("ph").ReplaceAllLiteralString(token, "f")
}

func transformQ(token string) string {
	return regexp.MustCompile("q").ReplaceAllLiteralString(token, "k")
}

func transformS(token string) string {
	return regexp.MustCompile("s(h|io|ia)").ReplaceAllLiteralString(token, "")
}

func transformT(token string) string {
	token = regexp.MustCompile("t(ia|io)").ReplaceAllLiteralString(token, "")
	token = regexp.MustCompile("th").ReplaceAllLiteralString(token, "z")
	return token
}

func dropT(token string) string {
	return regexp.MustCompile("tch").ReplaceAllLiteralString(token, "ch")
}

func transformV(token string) string {
	return regexp.MustCompile("v").ReplaceAllLiteralString(token, "f")
}

func transformWH(token string) string {
	return regexp.MustCompile("wh").ReplaceAllLiteralString(token, "w")
}

func dropW(token string) string {
	return regexp.MustCompile("w([^aeiou]|$)").ReplaceAllLiteralString(token, "")
}

func transformX(token string) string {
	token = regexp.MustCompile("^x").ReplaceAllLiteralString(token, "s")
	token = regexp.MustCompile("x").ReplaceAllLiteralString(token, "ks")
	return token
}

func dropY(token string) string {
	return regexp.MustCompile("y([^aeiou]|$)").ReplaceAllLiteralString(token, "")
}

func transformZ(token string) string {
	return regexp.MustCompile("z").ReplaceAllLiteralString(token, "s")
}
