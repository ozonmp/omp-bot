package product

import (
	"encoding/json"
	"github.com/ozonmp/omp-bot/internal/model/recomendation"
	"log"
)

type JsonSerializer struct {
}

func (serialize *JsonSerializer) serialize(data string) (recomendation.Product, error) {
	parsedData := JsonProductModel{}
	if err := json.Unmarshal([]byte(data), &parsedData); err != nil {
		log.Printf("JsonSerializer.serialize: %s\n Data: %s\n", err.Error(), data)
		return recomendation.Product{}, err
	} else {
		product := recomendation.Product{
			Id:          parsedData.Id,
			Title:       parsedData.Title,
			Description: parsedData.Description,
			Rating:      parsedData.Rating}
		return product, nil
	}
}

func (serialize *JsonSerializer) deserialize(product *recomendation.Product) (string, error) {
	serializedData, err := json.Marshal(JsonProductModel{
		Id:          product.Id,
		Title:       product.Title,
		Description: product.Description,
		Rating:      product.Rating,
	})
	if err != nil {
		log.Printf("JsonSerializer.deserialize: %s\n", err.Error())
		return "", err
	} else {
		return string(serializedData), nil
	}
}
