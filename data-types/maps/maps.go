package maps

import "fmt"

func main() {
	// map
	// tipo de dado de chave e valor
	// map é sempre dinâmico
	// definir o tipo da "chave" dentro de [] e o tipo do "valor" após o []
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Google"])

	// adicionando novo item a um map
	websites["Linkedin"] = "https://linkedin.com"
	fmt.Println(websites)

	// deletando item de um map
	delete(websites, "Linkedin")

	// Diferenças de Map para Struct
	// pode ser usado qualquer tipo de dado nas chaves
	// struct já tem sua estrutura predefinida
}
