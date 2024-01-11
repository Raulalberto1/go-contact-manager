package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// estructura contactos
type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}

	return nil
}

// Cargar contactos desde un archivo json
func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	//slice de contactos
	var contacts []Contact

	//Cargar contactos existentes desde el archivo
	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar los contactos:", err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("=== Gestor de Contactos ===\n",
			"1. Agregar un contacto\n",
			"2. Mostrar todos los contactos\n",
			"3. Salir\n",
			"Elige una opción: ")

		//leer el
		var opc int
		_, err = fmt.Scanln(&opc)
		if err != nil {
			fmt.Println("Error al leer la opcion:", err)
			return
		}

		//manejar la opcion del usuario
		switch opc {
		case 1:
			//ingresar y crear contacto
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Teléfono: ")
			c.Phone, _ = reader.ReadString('\n')

			//agregar un contacto a Slice
			contacts = append(contacts, c)

			//guardar en un archivo json
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error al uardar el contacto:", err)
			}

		case 2:
			//mostrar todos los contactos
			fmt.Println("==========================================")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s Email: %s Telefono: %s\n",
					index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("==========================================")

		case 3:
			//salir del programa
			return

		default:
			fmt.Println("Opción inválida")
		}
	}
}
