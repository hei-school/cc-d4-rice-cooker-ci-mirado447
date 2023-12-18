package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	powerConnected bool
	lidOpen        bool
	foodAdded      bool
	waterAdded     bool
	lidClosed      bool
	cooking        bool
	maxCupLimit    int
	foodQuantity   int
	waterQuantity  int
)

func main() {
	for {
		displayMenu()
		var choice int
		fmt.Print("Choix : ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Une erreur s'est produite :", err)
			os.Exit(1)
		}
		handleUserChoice(choice)
	}
}

func displayMenu() {
	fmt.Println("Voici les étapes à suivre :")
	fmt.Println("1. Brancher la prise")
	fmt.Println("2. Ouvrir le couvercle")
	fmt.Println("3. Verser l'aliment")
	fmt.Println("4. Verser de l'eau")
	fmt.Println("5. Fermer le couvercle")
	fmt.Println("6. Mettre le bouton sur cook")
	fmt.Println("7. Faire cuire")
	fmt.Println("8. Débrancher la prise")
	fmt.Println("0. Quitter")
}

func handleUserChoice(choice int) {
	switch choice {
	case 1:
		connectPower()
	case 2:
		openLid()
	case 3:
		pourFood()
	case 4:
		pourWater()
	case 5:
		closeLid()
	case 6:
		startCooking()
	case 7:
		cook()
	case 8:
		disconnectPower()
	case 0:
		fmt.Println("Programme terminé.")
		os.Exit(0)
	default:
		fmt.Println("Option invalide. Veuillez choisir à nouveau.")
	}
}

func connectPower() {
	if !powerConnected {
		powerConnected = true
		fmt.Println("La prise est branchée, le courant passe")
		fmt.Println("Le bouton warm s'allume")
	} else {
		fmt.Println("La prise est déjà branchée")
	}
}

func openLid() {
	if !lidOpen {
		lidOpen = true
		fmt.Println("Le couvercle est ouvert")
	} else {
		fmt.Println("Le couvercle est déjà ouvert")
	}
}

func pourFood() {
	if lidOpen {
		fmt.Print("Quantité d'aliment (en tasses) : ")
		var input string
		fmt.Scan(&input)
		foodQuantity, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Erreur lors de la saisie de la quantité d'aliment :", err)
			return
		}

		if foodQuantity > 0 && foodQuantity <= maxCupLimit {
			foodAdded = true
			fmt.Println("Aliment ajouté avec succès")
		} else {
			fmt.Println("Quantité invalide d'aliment")
		}
	} else {
		fmt.Println("Le couvercle doit être ouvert pour verser l'aliment")
	}
}

func pourWater() {
	if lidOpen && foodAdded {
		fmt.Print("Quantité d'eau par tasses : ")
		var input string
		fmt.Scan(&input)
		waterQuantity, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Erreur lors de la saisie de la quantité d'eau :", err)
			return
		}

		if waterQuantity > 0 && waterQuantity <= maxCupLimit {
			waterAdded = true
			fmt.Println("Eau ajoutée avec succès")
		} else {
			fmt.Println("Quantité invalide d'eau")
		}
	} else {
		fmt.Println("Le couvercle doit être ouvert et l'aliment doit être ajouté avant de verser de l'eau")
	}
}

func closeLid() {
	if lidOpen {
		lidClosed = true
		fmt.Println("Le couvercle est fermé")
	} else {
		fmt.Println("Le couvercle est déjà fermé")
	}
}

func startCooking() {
	if lidClosed {
		cooking = true
		fmt.Println("La cuisson a commencé")
	} else {
		fmt.Println("Le couvercle doit être fermé pour démarrer la cuisson")
	}
}

func cook() {
	if cooking && waterQuantity >= foodQuantity {
		fmt.Println("L'aliment est cuit")
		switchToWarm()
	} else {
		fmt.Println("La cuisson n'est pas encore prête")
	}
}

func switchToWarm() {
	if cooking {
		cooking = false
		fmt.Println("Le bouton est passé sur warm")
	} else {
		fmt.Println("La cuisson doit être terminée pour passer sur warm")
	}
}

func disconnectPower() {
	if powerConnected {
		powerConnected = false
		fmt.Println("Le courant ne passe plus")
	} else {
		fmt.Println("Le rice cooker est déjà débranché")
	}
}
