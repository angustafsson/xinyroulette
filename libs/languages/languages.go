package languages

import (
	"math/rand"
	"time"
)

type (
	// Language : Struct for each language
	Language struct {
		LanguageName   string
		LanguageDocURL string
	}
)

func GetRandomLanguage() (lang Language) {
	slice := languages()
	rand.Seed(time.Now().Unix())
	return slice[rand.Intn(len(slice))]
}

func languages() (langs []Language) {
	langs = []Language{
		Language{
			LanguageName:   "PHP",
			LanguageDocURL: "https://secure.php.net/docs.php",
		},
		Language{
			LanguageName:   "Python",
			LanguageDocURL: "https://www.python.org/doc/",
		},
		Language{
			LanguageName:   "Ruby",
			LanguageDocURL: "https://www.ruby-lang.org/en/documentation/",
		},
		Language{
			LanguageName:   "Rust",
			LanguageDocURL: "https://www.rust-lang.org/documentation.html",
		},
		Language{
			LanguageName:   "Elixir",
			LanguageDocURL: "http://elixir-lang.org/docs.html",
		},
		Language{
			LanguageName:   "Haskell",
			LanguageDocURL: "https://www.haskell.org/documentation",
		},
		Language{
			LanguageName:   "Crystal",
			LanguageDocURL: "http://crystal-lang.org/docs/",
		},
		Language{
			LanguageName:   "C",
			LanguageDocURL: "https://www.gnu.org/software/gnu-c-manual/gnu-c-manual.html",
		},
		Language{
			LanguageName:   "C++",
			LanguageDocURL: "http://www.cplusplus.com/doc/",
		},
		Language{
			LanguageName:   "Bash",
			LanguageDocURL: "http://www.tldp.org/LDP/abs/html/",
		},
		Language{
			LanguageName:   "Erlang",
			LanguageDocURL: "https://www.erlang.org/docs",
		},
		Language{
			LanguageName:   "NodeJS",
			LanguageDocURL: "https://nodejs.org/en/docs/",
		},
		Language{
			LanguageName:   "Dart",
			LanguageDocURL: "https://www.dartlang.org/docs/",
		},
		Language{
			LanguageName:   "Scala",
			LanguageDocURL: "http://www.scala-lang.org/",
		},
		Language{
			LanguageName:   "F#",
			LanguageDocURL: "http://fsharp.org/",
		},
	}
	return
}
