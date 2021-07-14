package fakedatabase

import ( 
	"encoding/json"
	"log"
)

// Dog properties json
// Type attributes must be uppercase!
type Dog struct {
 // id		string  // is bad
	Id 		string `json:"id"`  
	Name 	string `json:"name"` 
	Photo	string `json:"photo"`
	Age 	int8   `json:"age"`
} 

func GetDogById(id int) string {
	dogs :=	[]Dog{
		Dog{"0", "kongo", "http://images.com/kongo.png", 4},
		Dog{"1", "firulais", "http://images.com/firu.jpg", 21},
		Dog{"2", "stacy", "http://images.com/stacy_smile.png", 21},
		Dog{"3", "firulais", "http://images.com/firulais.jpeg", 24},
	}	
	// id exceds dogs len	
	if id > len(dogs)-1{
		return "{}"
	}

	jdog, err := json.Marshal(dogs[id])
	if err != nil {
		log.Fatal("Error to convert Json")
		return "error"
	}

	return string(jdog)
}

