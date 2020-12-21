package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var allergens = map[string][]string{}

func indexOf(a string, arr []string) int {
	for i := range arr {
		if a == arr[i] {
			return i
		}
	}

	return -1
}

func intersect(arr1, arr2 []string) []string {
	var result []string

	if len(arr1) == 0 {
		result = arr2
	}

	if len(arr2) == 0 {
		result = arr1
	}

	if len(arr1) < len(arr2) {
		arr1, arr2 = arr2, arr1
	}

	for _, a1 := range arr1 {
		if indexOf(a1, arr2) != -1 {
			result = append(result, a1)
		}
	}

	return result
}

func removeElement(r string, arr []string) []string {
	var result []string

	for _, a := range arr {
		if a != r {
			result = append(result, a)
		}
	}

	return result
}

func findMaxLength() int {
	max := 1

	for _, ing := range allergens {
		if max < len(ing) {
			max = len(ing)
		}
	}

	return max
}

var ingredientsChecked []string

func removeDuplicates() {
	ingredient, allergen := "", ""

	for all, ing := range allergens {
		if len(ing) == 1 && indexOf(ing[0], ingredientsChecked) == -1 {
			ingredient = ing[0]
			allergen = all
			break
		}
	}

	if ingredient == "" {
		return
	}

	for all, ing := range allergens {
		if all != allergen {
			allergens[all] = removeElement(ingredient, ing)
		}
	}

	ingredientsChecked = append(ingredientsChecked, ingredient)
}

func getCanonicalDangerousIngredientList(allergens map[string][]string) string {
	var allergensList, ingredientsList []string

	for all := range allergens {
		allergensList = append(allergensList, all)
	}

	sort.Strings(allergensList)

	for _, all := range allergensList {
		ingredientsList = append(ingredientsList, allergens[all][0])
	}

	return strings.Join(ingredientsList, ",")
}

func main() {
	file, _ := os.Open("./input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var ingredients []string

	for scanner.Scan() {
		line := scanner.Text()
		lineSplitted := strings.Split(line, " (contains ")

		for _, all := range strings.Split(strings.Trim(lineSplitted[1], ")"), ", ") {
			for _, ing := range strings.Split(lineSplitted[0], " ") {
				if indexOf(ing, ingredients) == -1 {
					ingredients = append(ingredients, ing)
				}
			}

			allergens[all] = intersect(allergens[all], strings.Split(lineSplitted[0], " "))
		}
	}

	for findMaxLength() > 1 {
		removeDuplicates()
	}

	healthyIngredients := ingredients

	for _, ing := range allergens {
		healthyIngredients = removeElement(ing[0], healthyIngredients)
	}

	file, _ = os.Open("./input.txt")

	defer file.Close()

	scanner = bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		for _, h := range healthyIngredients {
			if strings.Contains(line, " "+h+" ") || strings.HasPrefix(line, h+" ") {
				count++
			}
		}
	}

	fmt.Printf("Number of times healthy ingredients appear (1): %d\n", count)

	fmt.Print("Canonical dangerous ingredients list (2): ")
	fmt.Println(getCanonicalDangerousIngredientList(allergens))
}
