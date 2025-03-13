package main

import (
	"bufio"
	"context"
	"log"
	"os"

	"github.com/Gustrb/kramer/internal/infra"
	"github.com/Gustrb/kramer/internal/provider"
	"github.com/Gustrb/kramer/internal/repository"
	"github.com/Gustrb/kramer/models"
	"github.com/google/uuid"
)

func main() {
	if err := provider.SetupLogger(); err != nil {
		log.Fatalf("error setting up logger: %v", err)
		return
	}

	db, err := infra.SetupDatabase()
	if err != nil {
		log.Fatalf("error setting up database: %v", err)
		return
	}
	defer db.Close()

	store := repository.StoreFactory(db)

	p, err := provider.ProviderFactory(store, provider.ChatGPT)
	if err != nil {
		log.Fatalf("error creating provider: %v", err)
		return
	}

	var c models.Context
	// If the program has command line arguments, use the first argument as the context name
	// and create a context with that name
	if len(os.Args) > 1 {
		c, err = store.Context().FindContextByName(context.Background(), os.Args[1])

		if err == repository.ErrContextNotFound {
			// Create
			c, err = store.Context().Create(context.Background(), models.CreateContext{Name: os.Args[1]})
			if err != nil {
				log.Fatalf("error creating context: %v", err)
				return
			}
		}

		if err != nil {
			log.Fatalf("error finding context by name: %v", err)
			return
		}

		err = p.LoadContext(c.ID)
		if err != nil {
			log.Fatalf("error loading context: %v", err)
			return
		}
	} else {
		c, err = store.Context().Create(context.Background(), models.CreateContext{Name: uuid.New().String()})
		if err != nil {
			log.Fatalf("error creating context: %v", err)
			return
		}
	}

	for {
		scanner := bufio.NewScanner(os.Stdin)

		// read a question from stdin
		log.Printf("Ask me a question: ")
		scanner.Scan()
		question := scanner.Text()
		// get the answer
		answer, err := p.GetResponse(c.ID, question)
		if err != nil {
			log.Fatalf("error getting response: %v", err)
			return
		}

		// print the answer
		log.Printf("Answer: %s", answer)
	}
}
