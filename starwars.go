package main

import (
	"context"
	"fmt"
	"strings"
)

var characters = []*Character{
	{
		Id:         1000,
		Name:       "Luke Skywalker",
		Friends:    []string{"Han Solo", "Leia Organa", "C-3PO", "R2-D2"},
		AppearsIn:  []Episode{Episode_NEW_HOPE, Episode_EMPIRE, Episode_JEDI},
		Height:     1.72,
		HomePlanet: "Tatooine",
	},
	{
		Id:         1001,
		Name:       "Darth Vader",
		Friends:    []string{"Wilhuff Tarkin"},
		AppearsIn:  []Episode{Episode_NEW_HOPE, Episode_EMPIRE, Episode_JEDI},
		Height:     2.02,
		HomePlanet: "Tatooine",
	},
	{
		Id:        1002,
		Name:      "Han Solo",
		Friends:   []string{"Luke Skywalker", "Leia Organa", "C-3PO"},
		AppearsIn: []Episode{Episode_NEW_HOPE, Episode_EMPIRE, Episode_JEDI},
		Height:    1.8,
	},
	{
		Id:         1003,
		Name:       "Leia Organa",
		Friends:    []string{"Han Solo", "Luke Skywalker", "C-3PO", "R2-D2"},
		AppearsIn:  []Episode{Episode_NEW_HOPE, Episode_EMPIRE, Episode_JEDI},
		Height:     1.5,
		HomePlanet: "Alderaan",
	},
	{
		Id:        1004,
		Name:      "Wilhuff Tarkin",
		Friends:   []string{"Darth Vader"},
		AppearsIn: []Episode{Episode_NEW_HOPE},
		Height:    1.8,
	},
	{
		Id:              2000,
		Name:            "C-3PO",
		Friends:         []string{"Han Solo", "Luke Skywalker", "Leia Organa", "R2-D2"},
		AppearsIn:       []Episode{Episode_NEW_HOPE, Episode_EMPIRE, Episode_JEDI},
		PrimaryFunction: "Protocol",
	},
	{
		Id:              2001,
		Name:            "R2-D2",
		Friends:         []string{"Han Solo", "Luke Skywalker", "C-3PO"},
		AppearsIn:       []Episode{Episode_NEW_HOPE, Episode_EMPIRE, Episode_JEDI},
		PrimaryFunction: "Astromech",
	},
}

var characterData = make(map[int32]*Character)

func init() {
	for _, c := range characters {
		characterData[c.Id] = c
	}
}

var reviewData = make(map[Episode][]*Review)

type starwars struct{}

func (g *starwars) SearchCharacter(ctx context.Context, req *SearchCharacterRequest) (*SearchCharacterResponse, error) {
	fmt.Println("Incoming search character request: " + req.Name)
	var l []*Character
	for _, c := range characters {
		if strings.Contains(c.Name, req.Name) {
			l = append(l, c)
		}
	}
	resp := SearchCharacterResponse{
		Characters: l,
	}
	return &resp, nil
}

func (g *starwars) GetHero(ctx context.Context, req *GetHeroRequest) (*Character, error) {
	fmt.Println("Incoming get hero request: " + req.Episode.String())
	if req.Episode == Episode_EMPIRE {
		return characterData[1000], nil
	}

	return characterData[2001], nil
}

func (g *starwars) AddReview(ctx context.Context, req *Review) (*Review, error) {
	fmt.Println("Incoming add review request : " + req.Episode.String())
	val := reviewData[req.Episode]
	val = append(val, req)
	reviewData[req.Episode] = val

	return req, nil
}

func (g *starwars) GetReviews(ctx context.Context, req *GetReviewsRequest) (*GetReviewsResponse, error) {
	fmt.Println("Incoming get reviews request: " + req.Episode.String())
	val := reviewData[req.Episode]
	resp := GetReviewsResponse{
		Reviews: val,
	}
	return &resp, nil
}
