package main

import "fmt"

func main() {

	menu_principal := `
	Menú principal
	1.Cambiar su contraseña
	2.Editar perfil
	3.Foro
	4.Finalizar programa
	`

	menu_perfil := `
	Menú de Perfil
	1.Cambiar foto de perfil
	2.Cambiar nombre de usuario
	3.Editar biografía
	4.Volver al menú principal
	`

	menu_foro := `
	Menú de foro
	1.Publicar nueva entrada en el foro
	2.Editar una entrada del foro
	3.Eliminar una entrada del foro
	4.Volver al menú principal
	`

	var respuesta_menu int
	var respuesta_perfil int
	var respuesta_foro int
	menu := 0

	for {
		if menu == 0 {
			fmt.Println(menu_principal)
			fmt.Scanf("%d", &respuesta_menu)
		}

		switch respuesta_menu {
		case 1:
			fmt.Println("Su contraseña ha sido cambiada exitosamente")

		case 2:
			fmt.Println(menu_perfil)
			fmt.Scanf("%d", &respuesta_perfil)

			switch respuesta_perfil {

			case 1:
				fmt.Println("Su foto de perfil ha sido cambiada exitosamente")
				menu = 2

			case 2:
				fmt.Println("Su nombre de usuario ha sido cambiada exitosamente")
				menu = 2

			case 3:
				fmt.Println("Su biografia ha sido editada exitosamente")
				menu = 2

			case 4:
				menu = 0
				break

			default:
				fmt.Println("Opción no valida")
			}

		case 3:
			fmt.Println(menu_foro)
			fmt.Scanf("%d", &respuesta_foro)

			switch respuesta_foro {

			case 1:
				fmt.Println("Su nueva entrada ha sido publicada exitosamente")
				menu = 2

			case 2:
				fmt.Println("Entrada del foro ha sido editada exitosamente")
				menu = 2

			case 3:
				fmt.Println("Entrada del foro eliminada exitosamente")
				menu = 2

			case 4:
				menu = 0
				break

			default:
				fmt.Println("Opción no valida")
			}
		case 4:
			menu = 4
			break

		default:
			fmt.Println("Opción no valida")

		}

		if menu == 4 {
			fmt.Println("Hasta luego, esperamos verte pronto")
			break
		}
	}

}
