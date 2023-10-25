package main

import (
	"log"

	"github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/pkg/config"
	"github.com/nikhilnarayanan623/random-school-generator-grpc/school-service/pkg/di"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	srv, err := di.InitializeAPI(cfg)
	if err != nil {
		log.Fatalf("failed to initialize api: %v", err)
	}

	if err = srv.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}

// func main() {

// 	// 	cfg, err := config.LoadConfig()
// 	// if err != nil {
// 	// 	log.Fatalf("failed to load config: %v", err)
// 	// }

// 	schoolRepo, err := repository.NewSchoolUseCase()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	schoolUseCase := usecase.NewSchoolUseCase(schoolRepo)

// 	school := schoolUseCase.Create("hii")
// 	jsonData, err := json.Marshal(school)

// 	if err != nil {
// 		log.Fatalf("failed to marshal school to json: %v", err)
// 	}
// 	fmt.Println(len(jsonData))
// 	// fmt.Printf("shoo: %+v\n", school)

// }
