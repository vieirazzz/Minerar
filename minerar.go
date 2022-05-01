package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/toqueteos/webbrowser"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Você deve passar como argumento uma palavra ou frase para traduzir. \n\nExemplo: 'minerar Nice to meet you'")
		os.Exit(0)
	}

	frase := strings.Join(args, "+")

	// Google Search
	webbrowser.Open("https://www.google.com.br/search?q=" + frase)
	// Google Images
	webbrowser.Open("https://www.google.com/search?tbm=isch&q=" + frase)
	// Tatoeba
	webbrowser.Open("https://tatoeba.org/en/sentences/search?query=" + frase)
	// Word Reference
	webbrowser.Open("https://www.wordreference.com/enpt/" + frase)
	// Linguee
	webbrowser.Open("https://www.linguee.com.br/portugues-ingles/search?source=ingles&query=" + frase)
	// Forvo
	webbrowser.Open("https://forvo.com/word/" + frase + "/#en_usa")
	// Cambridge
	webbrowser.Open("https://dictionary.cambridge.org/pt/dicionario/ingles-portugues/" + frase)
	// Michaelis
	webbrowser.Open("https://michaelis.uol.com.br/moderno-ingles/busca/ingles-portugues-moderno/" + frase)
	// Context Reverso
	webbrowser.Open("https://context.reverso.net/traducao/ingles-portugues/" + frase)
	// Google Translate
	webbrowser.Open("https://translate.google.com.br/?hl=pt-BR&sl=en&tl=pt&text=" + frase)
}
