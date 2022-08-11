package main

import (
	"fmt"
	"math"

	"github.com/spf13/viper"
)

/*
Aperture : sans unité, l’ouverture de l’objectif
PixelPitch : en µm, la taille d’un photo-site de l’appareil photo
FocalLength : en mm, la focale de l’objectif
δ : la déclinaison minimale du champ photographié
Accuracy : un nombre entre 1 et 3 donnant la précision du temps de pose (trainée des étoiles)
CropFactor : le facteur de crop ou de focale équivalente du capteur (1 pour plein format, ~1.6 pour APS-C Canon, ~1.5 pour les autres APS-C, 2 pour les m4/3…).
*/

func NPFSimple(Aperture, PixelPitch, FocalLength float64) float64 {
	var t = (35*Aperture + 30*PixelPitch) / FocalLength
	return t
}

func NPFFull(Accuracy, Aperture, FocalLength, PixelPitch, δ float64) float64 {
	var t = (Accuracy * (16.9*Aperture + 0.10*FocalLength + 13.7*PixelPitch) / (FocalLength * math.Cos(δ)))
	return t
}

func FourCrop(CropFactor, FocalLength float64) float64 {
	var t = 100 * (4 - CropFactor) / FocalLength
	return t
}

func Rule500(CropFactor, FocalLength float64) float64 {
	var t = 500 / (CropFactor * FocalLength)
	return t
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	Aperture := viper.GetFloat64("parameters.Aperture")
	PixelPitch := viper.GetFloat64("parameters.PixelPitch")
	FocalLength := viper.GetFloat64("parameters.FocalLength")
	δ := viper.GetFloat64("parameters.δ")
	Accuracy := viper.GetFloat64("parameters.Accuracy")
	CropFactor := viper.GetFloat64("parameters.CropFactor")

	fmt.Println("NPF Simple    : ", fmt.Sprintf("temps de pause de %.1f sec.", NPFSimple(Aperture, PixelPitch, FocalLength)))
	fmt.Println("NPF Complet   : ", fmt.Sprintf("temps de pause de %.1f sec.", NPFFull(Accuracy, Aperture, FocalLength, PixelPitch, δ)))
	fmt.Println("4 Crops       : ", fmt.Sprintf("temps de pause de %.1f sec.", FourCrop(CropFactor, FocalLength)))
	fmt.Println("Règle des 500 : ", fmt.Sprintf("temps de pause de %.1f sec.", Rule500(CropFactor, FocalLength)))
}
