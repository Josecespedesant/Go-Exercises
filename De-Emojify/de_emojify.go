package de_emojify

/*
EXERCISE:
You are given a string made up of chains of emotes separated by 1 space each, with chains having 2 spaces in-between each.
Each emote represents a digit.
Each emote chain represents the digits of the ASCII/Unicode code for a character, e.g. :( ;) is 97, which is the ASCII code for 'a'.
Given a such string of emotes, find the string it represents. Example:
':D :) :/  :D :) :|' is 2 chains: ':D :) :/' and ':D :) :|'.
These represent ASCII codes 104 and 105 respectively, translating to 'hi'.
Input will always be valid. Chains may start with leading zeroes; these are valid and do not change the chain's value.

REFERENCE: https://www.codewars.com/kata/6627696c86b953001280675e
KYU: 7
*/

import (
	"strings"
)

var map_emoji_num = map[string]int{
	":)":  0,
	":D":  1,
	">(":  2,
	">:C": 3,
	":/":  4,
	":|":  5,
	":O":  6,
	";)":  7,
	"^.^": 8,
	":(":  9,
}

func De_Emojify(emote_string string) string {
	separated_chain := strings.Split(emote_string, "  ")
	var resulting_converted_chain []byte
	for _, value := range separated_chain {
		separated_emojis := strings.Split(value, " ")
		resulting_ascii_code := 0
		for index, emoji := range separated_emojis {
			resulting_ascii_code += map_emoji_num[emoji]
			if index < len(separated_emojis)-1 {
				resulting_ascii_code *= 10
			}
		}
		ascii_conversion := byte(resulting_ascii_code)
		resulting_converted_chain = append(resulting_converted_chain, ascii_conversion)
	}
	return string(resulting_converted_chain)
}
