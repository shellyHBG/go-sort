# Sorting by golang

## Introduction
In go-sort, it implements two poker sorting functions. 
One is sorting in DESC order: Sort first by suits (spade(1), heart(2), diamond(3), club(4)) and then sort by number/face (K, Q, J,  10, ... , 1).
The other is sorting in ASC order: Sort first by suits (club(4), diamond(3), heart(2), spade(1)) and then sort by number/face (1, 2, ..., 10, J, Q, K).

The file structure in this project:
```
- main.go
- deck
  - card.go
  - sortDeck.go
```

A local package `deck` is created. It supports two functions, `SortDeckByDESC()` and `SortDeckByASC()`.
```
func (d Deck) SortDeckByDESC()
```
```
func (d Deck) SortDeckByASC()
```
```
type Card struct {
	suit int // 1: spade, 2: heart, 3: diamond, 4: club
	face int // 1-10, 11: Jack, 12: Queen, 13: King
}

type Deck []Card
```
  
For bonus, users can call `shuffle()` to rearrange the order of Cards in a random serial.
```
func (d Deck) Shuffle()
```

## Requirements
Following are standard library (external package supported by go) used in this project:
* fmt
* math/rand
* sort
* time

## Installation
No, just use it :D

## Usage
Compile and run this code by typing the command line:
```
> go run 
```
Build an executable file using:
```
> go build 
```

## Authors
By shellyHBG
