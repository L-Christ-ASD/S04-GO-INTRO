package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

// Structure pour la requête à l'API OpenAI
type OpenAIRequest struct {
	Model    string   `json:"model"`
	Messages []Message `json:"messages"`
}

// Structure pour un message dans la requête
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Structure pour la réponse de l'API OpenAI
type OpenAIResponse struct {
	Choices []Choice `json:"choices"`
}

// Structure pour un choix dans la réponse
type Choice struct {
	Message Message `json:"message"`
}

func main() {
	// Charger le fichier .env
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Erreur lors du chargement du fichier .env : %v", err)
	}

	// Obtenir les tokens
	token := os.Getenv("TOKEN_DISCORD")
	openaiKey := os.Getenv("OPENAI_API_KEY")

	// Créer une session Discord
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Erreur lors de la création de la session Discord : %v", err)
	}

	// Ajouter un gestionnaire d'événements pour les messages
	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		// Appeler l'API OpenAI pour les messages qui ne commencent pas par "!"
		if len(m.Content) > 0 && m.Content[0] != '!' {
			response, err := getAIResponse(openaiKey, m.Content)
			if err != nil {
				log.Printf("Erreur lors de l'appel à l'API OpenAI : %v", err)
				s.ChannelMessageSend(m.ChannelID, "Erreur lors de la génération de la réponse.")
				return
			}
			s.ChannelMessageSend(m.ChannelID, response)
		}
	})

	// Ouvrir la connexion
	err = dg.Open()
	if err != nil {
		log.Fatalf("Erreur lors de l'ouverture de la connexion : %v", err)
	}
	defer dg.Close()

	fmt.Println("Bot en cours d'exécution. Appuyez sur CTRL+C pour quitter.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

// Fonction pour obtenir une réponse de l'API OpenAI
func getAIResponse(apiKey, message string) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"
	reqBody := OpenAIRequest{
		Model: "gpt-4",
		Messages: []Message{
			{Role: "user", Content: message},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("erreur lors de l'encodage JSON : %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("erreur lors de la création de la requête : %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("erreur lors de l'appel à l'API OpenAI : %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("erreur lors de la lecture de la réponse : %v", err)
	}

	var openAIResp OpenAIResponse
	err = json.Unmarshal(body, &openAIResp)
	if err != nil {
		return "", fmt.Errorf("erreur lors du décodage JSON : %v", err)
	}

	if len(openAIResp.Choices) > 0 {
		return openAIResp.Choices[0].Message.Content, nil
	}

	return "Je n'ai pas pu générer de réponse.", nil
}
