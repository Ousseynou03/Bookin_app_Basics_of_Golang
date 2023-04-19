package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conférence"
	const conferenceTickets = 50
	var remainingTickets = 50

	//Appel de la fonction greetUsers()

	greetUsers(conferenceName)

	fmt.Println("Welcome to\n", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Println("Get yours tickets here to attend")

	//Arrays in Go
	var bookings []string
	//bookings[0] = "Weuz"
	//bookings[1]  = "Dione"

	//Les Boucles

	for {
		var firtsName string
		var lastName string
		var email string
		var userTickets int

		fmt.Println("Entrer your first Name")
		fmt.Scan(&firtsName)

		fmt.Println("Entrer your last Name")
		fmt.Scan(&lastName)

		fmt.Println("Entrer your email")
		fmt.Scan(&email)

		fmt.Println("Entrer your number of tickets")
		fmt.Scan(&userTickets)

		//Dans le cas de validations des formulaires.
		isValidName := len(firtsName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets >= 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {

			remainingTickets = remainingTickets - userTickets

			bookings = append(bookings, firtsName+" "+lastName)

			fmt.Printf("The whole array:%v\n", bookings)
			fmt.Printf("The fisrt value :%v\n", bookings[0])
			fmt.Printf("Array type :%T\n", bookings)
			fmt.Printf("Array len :%v\n", len(bookings))
			/*
				Notion de pointer avec le & commercial pour désigner le nom de la variable en question
				fmt.Println(remainingTickets)
				fmt.Println(&remainingTickets)

			*/

			//Demander à un utilisateur de saisir son nom

			/*Printf est une autre façon d'imprimer un message en utilisant des mots clés comme %v
			qui permet de référencer une variable ou une constante
			*/
			fmt.Printf("Thank you %v %v for booking %v ticjets. You will receive a confirmation email at %v\n", firtsName, lastName, userTickets, email)

			firstNames := []string{}
			//Equivalence à forEach
			//Le "_" c'est pour qu'il y'a une variable qu'on décide d'ignorer
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The firstName of bookin,gs are: %v\n", firstNames)

			if userTickets > remainingTickets {
				fmt.Printf("we have only %v tickets remaining, so you cain't book %v tickets \n", remainingTickets, userTickets)
				continue
			}

			if remainingTickets == 0 {
				fmt.Println("Our Conférence is booked out.Come back next time")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("Your first name or your last name you entered is too short\n")
			}
			if !isValidEmail {
				fmt.Println("Your email address you entered doesn't contains @ sign\n")
			}
			if !isValidTicketNumber {
				fmt.Printf("Number of ticket you entered is invalid\n")
			}
		}
	}

	//Switch Statement
	city := "London"

	switch city {
	case "New York":
		//Exécution code concernant New York

	case "Singapore", "Hong Kong":
		//Exécution code concernant Singuapore
	
	case "London":
		//Exécution code concernant London
	
	case "Berlin":
		//Exécution code 

	case "Mexico City":
		//some code here
	
	
	default:
		fmt.Println("Vous n'avez pas séléctionnez une ville valide")
	}
}

func greetUsers(confName string)  {
	fmt.Printf("Welcome to %v booking application", confName)
}
