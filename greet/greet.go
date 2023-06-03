package greet

/*
Para declarar una variable a nivel de paquete es necesario 
utilizar la palabra "var".
Para exportar una variable o función a través del paquete 
es necesario que empiece su nombre con mayúsculas, de lo 
contrario será visible sólo a nivel de archivo.
*/
var emoji = "🙋🏿‍♂️"

// English retorna saludo en ingles
func English() string {
	return "Hi " + emoji
}

// Italian retorna saludo en italiano
func Italian() string {
	return "Ciao " + emoji
}