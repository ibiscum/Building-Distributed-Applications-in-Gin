package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	models "github.com/ibiscum/Building-Distributed-Applications-in-Gin/chapter07/api-with-db/models"
	"github.com/stretchr/testify/assert"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

func TestListRecipesHandler(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/recipes", ts.URL))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	data, _ := io.ReadAll(resp.Body)

	var recipes []models.Recipe
	json.Unmarshal(data, &recipes)
	assert.Equal(t, len(recipes), 10)
}

func TestUpdateRecipeHandler(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	id, err := primitive.ObjectIDFromHex("c0283p3d0cvuglq85log")
	if err != nil {
		log.Fatal(err)
	}

	recipe := models.Recipe{
		ID:   id,
		Name: "Oregano Marinated Chicken",
	}

	raw, _ := json.Marshal(recipe)
	resp, err := http.NewRequest("PUT", fmt.Sprintf("%s/recipes/%s", ts.URL, recipe.ID), bytes.NewBuffer(raw))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	assert.Nil(t, err)
	//assert.Equal(t, http.StatusOK, resp.StatusCode)
	data, _ := io.ReadAll(resp.Body)

	var payload map[string]string
	json.Unmarshal(data, &payload)

	assert.Equal(t, payload["message"], "Recipe has been updated")
}

func TestDeleteRecipeHandler(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	id, err := primitive.ObjectIDFromHex("c0283p3d0cvuglq85log")
	if err != nil {
		log.Fatal(err)
	}

	recipe := models.Recipe{
		ID:   id,
		Name: "Oregano Marinated Chicken",
	}

	raw, _ := json.Marshal(recipe)
	resp, err := http.NewRequest("DELETE", fmt.Sprintf("%s/recipes/%s", ts.URL, recipe.ID), bytes.NewBuffer(raw))
	if err != nil {
		log.Fatal(err)
	}

	// resp, err := http.Delete(fmt.Sprintf("%s/recipes/c0283p3d0cvuglq85log", ts.URL))
	defer resp.Body.Close()
	assert.Nil(t, err)
	//assert.Equal(t, http.StatusOK, resp.StatusCode)
	data, _ := io.ReadAll(resp.Body)

	var payload map[string]string
	json.Unmarshal(data, &payload)

	assert.Equal(t, payload["message"], "Recipe has been deleted")
}

func TestFindRecipeHandler(t *testing.T) {
	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	id, err := primitive.ObjectIDFromHex("c0283p3d0cvuglq85log")
	if err != nil {
		log.Fatal(err)
	}
	expectedRecipe := models.Recipe{
		ID:   id,
		Name: "Oregano Marinated Chicken",
		Tags: []string{"main", "chicken"},
	}

	resp, err := http.Get(fmt.Sprintf("%s/recipes/%s", ts.URL, expectedRecipe.ID))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	data, _ := io.ReadAll(resp.Body)

	var actualRecipe models.Recipe
	json.Unmarshal(data, &actualRecipe)

	assert.Equal(t, expectedRecipe.Name, actualRecipe.Name)
	assert.Equal(t, len(expectedRecipe.Tags), len(actualRecipe.Tags))
}
