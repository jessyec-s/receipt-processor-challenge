package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"

	"gopkg.in/validator.v2"
)

type Receipt struct {
	Retailer     string         `json:"retailer" binding:"required" validate:"nonzero,nonnil,regexp=^[\\w\\s\\-&]+$"`
	PurchaseDate string         `json:"purchaseDate" binding:"required" validate:"nonzero,nonnil,regexp=^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\\d|3[01])$"`
	PurchaseTime string         `json:"purchaseTime" binding:"required" validate:"nonzero,nonnil,regexp=^([01]\\d|2[0-3]):[0-5]\\d$"`
	Items        []*ReceiptItem `json:"items" binding:"required" validate:"min=1"`
	Total        string         `json:"total" binding:"required" validate:"regexp=^\\d+\\.\\d{2}$"`
}

type ReceiptItem struct {
	ShortDescription string `json:"shortDescription" binding:"required" svalidate:"nonzero,nonnil,regexp=^[\\w\\s\\-]+$"`
	Price            string `json:"price" binding:"required" validate:"nonzero,nonnil,regexp=^\\d+\\.\\d{2}$"`
}

func validateReceipt(receipt Receipt) bool {

	errs := validator.Validate(receipt)

	fmt.Printf("%+v\n\n\n", errs)
	return errs == nil
}

func calculateRetailerPoints(retailer string) (points int) {
	for _, char := range retailer {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			points++
		}
	}
	return points
}

func calculateTotalPoints(total string) (points int) {
	if total, err := strconv.ParseFloat(total, 64); err == nil {
		if float64(int64(total)) == total {
			points += 50
		}

		if float64(int64(total/0.25)) == total/0.25 {
			points += 25
		}
	}
	return points
}

func calculateItemsPoints(items []*ReceiptItem) (points int) {
	itemCount := len(items)
	points += itemCount / 2 * 5

	for _, item := range items {
		if length := len(strings.Trim(item.ShortDescription, " ")); length%3 == 0 {
			if price, err := strconv.ParseFloat(item.Price, 64); err == nil {
				points += int(math.Ceil(price * 0.2))
			}
		}

	}
	return points
}

func calculateDatePoints(date string) (points int) {
	day := strings.Split(date, "-")[2]
	if dayNumber, err := strconv.ParseInt(day, 10, 64); err == nil && dayNumber%2 != 0 {
		points += 6
	}
	return points
}

func calculateTimePoints(time string) (points int) {
	hour := strings.Split(time, ":")[0]
	if timeNumber, err := strconv.ParseInt(hour, 10, 64); err == nil && timeNumber >= 14 && timeNumber < 16 && time != "14:00" {
		points += 10
	}

	return points
}

func calculatePoints(receipt Receipt) (points int) {

	points += calculateRetailerPoints(receipt.Retailer)
	points += calculateTotalPoints(receipt.Total)
	points += calculateItemsPoints(receipt.Items)
	points += calculateDatePoints(receipt.PurchaseDate)
	points += calculateTimePoints(receipt.PurchaseTime)

	return points

}
