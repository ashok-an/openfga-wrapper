package openfga

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func (e *MissingEnv) Error() string {
	return fmt.Sprintf("Environment variable '%s' is not set", e.Name)
}

func checkEnvVars(requiredVars []string) error {
	var missingVars []string
	for _, varName := range requiredVars {
		value := os.Getenv(varName)
		if value == "" {
			missingVars = append(missingVars, varName)
		}
	}

	if len(missingVars) > 0 {
		return &MissingEnv{Name: strings.Join(missingVars, ", ")}
	}
	return nil
}

func getClient() Client {
	envVars := []string{"OPENFGA_API_URL", "OPENFGA_STORE_ID", "OPENFGA_MODEL_ID"}
	if err := checkEnvVars(envVars); err != nil {
		log.Printf("Error: env lookup failed: %s\n", err)
		return Client{}
	}

	c := Client{
		Url:     os.Getenv("OPENFGA_API_URL"),
		StoreID: os.Getenv("OPENFGA_STORE_ID"),
		ModelID: os.Getenv("OPENFGA_MODEL_ID"),
	}
	return c
}

func IsHealthy() bool {
	c := getClient()
	url := fmt.Sprintf("%s/healthz", c.Url)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error: http Get failed for %s: %s\n", url, err)
		return false
	}
	defer resp.Body.Close()
	return (resp.StatusCode == http.StatusOK)
}

func GetStores() Stores {
	c := getClient()
	url := fmt.Sprintf("%s/stores", c.Url)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error: http Get failed for %s: %s\n", url, err)
		return Stores{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error: ReadAll failed for %s: %s\n", url, err)
		return Stores{}
	}

	store := Stores{}
	if err := json.Unmarshal(body, &store); err != nil {
		fmt.Println("Error unmarshalling response:", err)
	}
	return store
}

func GetStore(storeID string) Store {
	c := getClient()
	url := fmt.Sprintf("%s/stores/%s", c.Url, storeID)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error: http Get failed for %s: %s\n", url, err)
		return Store{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error: ReadAll failed for %s: %s\n", url, err)
		return Store{}
	}

	store := Store{}
	if err := json.Unmarshal(body, &store); err != nil {
		fmt.Println("Error unmarshalling response:", err)
	}
	return store
}

func GetModels(storeID string) Models {
	c := getClient()
	url := fmt.Sprintf("%s/stores/%s/authorization-models", c.Url, storeID)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error: http Get failed for %s: %s\n", url, err)
		return Models{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error: ReadAll failed for %s: %s\n", url, err)
		return Models{}
	}

	models := Models{}
	if err := json.Unmarshal(body, &models); err != nil {
		fmt.Println("Error unmarshalling response:", err)
	}
	return models
}

func GetModel(storeID string, modelID string) ModelResponse {
	c := getClient()
	url := fmt.Sprintf("%s/stores/%s/authorization-models/%s", c.Url, storeID, modelID)
	resp, err := http.Get(url)

	if err != nil {
		log.Printf("Error: http Get failed for %s: %s\n", url, err)
		return ModelResponse{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error: ReadAll failed for %s: %s\n", url, err)
		return ModelResponse{}
	}

	model := ModelResponse{}
	if err := json.Unmarshal(body, &model); err != nil {
		fmt.Println("Error unmarshalling response:", err)
	}

	return model
}
