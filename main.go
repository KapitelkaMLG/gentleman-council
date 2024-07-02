package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/kapitelkamlg/gentleman-council/templ/base"
	"github.com/kapitelkamlg/gentleman-council/templ/footer"
	"github.com/kapitelkamlg/gentleman-council/templ/index"
	"github.com/kapitelkamlg/gentleman-council/templ/gentlemen"
	"github.com/kapitelkamlg/gentleman-council/templ/shame"
	"github.com/kapitelkamlg/gentleman-council/templ/sheikh"
	"github.com/kapitelkamlg/gentleman-council/templ/movies"
	"github.com/kapitelkamlg/gentleman-council/templ/games"
)

func main() {
	outDir := "public/"
	gentlemenFile := "data/gentlemen.json"
	shameFile := "data/shame.json"
	sheikhFile := "data/sheikh.json"
	moviesFile := "data/movies.json"
	gamesFile := "data/games.json"
	
	pages := []footer.Page {
		{
			Name: "index",
			Title: "Джентельменское соглашение",
		},
		{
			Name: "gentlemen",
			Title: "Джентельмены",
		},
		{
			Name: "shame",
			Title: "Доска Позора",
		},
		{
			Name: "sheikh",
			Title: "Шейхи",
		},
		{
			Name: "movies",
			Title: "Киноаук",
		},
		{
			Name: "games",
			Title: "Игроаук",
		},
	}	
	RenderIndexPage(outDir, pages)
	RenderGentlemenPage(outDir, pages, gentlemenFile)
	RenderShamePage(outDir, pages, shameFile)
	RenderSheikhPage(outDir, pages, sheikhFile)
	RenderMoviesPage(outDir, pages, moviesFile)
	RenderGamesPage(outDir, pages, gamesFile)
}

func RenderIndexPage(outDir string, pages []footer.Page) {
	name := "index"
	title := GetPageByName(pages, name).Title
	f := CreateFile(outDir + name + ".html")
	index := index.IndexComp()
	footer := footer.FooterComp(pages, name)
	err := base.BaseComp(name, title, index, footer).Render(context.Background(), f)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}
}

func RenderGentlemenPage(outDir string, pages []footer.Page, gentlemenFile string) {
	gentlemenJson, err := os.ReadFile(gentlemenFile)
	if err != nil {
		log.Fatalf("Failed to read gentlemen data: %v", err)
	}
	var gentlemenData []gentlemen.Gentleman
	err = json.Unmarshal(gentlemenJson, &gentlemenData)
	if err != nil {
		log.Fatalf("Failed to decode gentlemen data: %v", err)
	}
	name := "gentlemen"
	title := GetPageByName(pages, name).Title
	f := CreateFile(outDir + name + ".html")
	gentlemen := gentlemen.GentlemenComp(gentlemenData)
	footer := footer.FooterComp(pages, name)
	err = base.BaseComp(name, title, gentlemen, footer).Render(context.Background(), f)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}
}

func RenderShamePage(outDir string, pages []footer.Page, shameFile string) {
	shameJson, err := os.ReadFile(shameFile)
	if err != nil {
		log.Fatalf("Failed to read shame data: %v", err)
	}
	var shameData []shame.Shame
	err = json.Unmarshal(shameJson, &shameData)
	if err != nil {
		log.Fatalf("Failed to decode shame data: %v", err)
	}
	name := "shame"
	title := GetPageByName(pages, name).Title
	f := CreateFile(outDir + name + ".html")
	shame := shame.ShameComp(shameData)
	footer := footer.FooterComp(pages, name)
	err = base.BaseComp(name, title, shame, footer).Render(context.Background(), f)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}
}

func RenderSheikhPage(outDir string, pages []footer.Page, sheikhFile string) {
	sheikhJson, err := os.ReadFile(sheikhFile)
	if err != nil {
		log.Fatalf("Failed to read gentlemen data: %v", err)
	}
	var sheikhData []string
	err = json.Unmarshal(sheikhJson, &sheikhData)
	if err != nil {
		log.Fatalf("Failed to decode gentlemen data: %v", err)
	}
	name := "sheikh"
	title := GetPageByName(pages, name).Title
	f := CreateFile(outDir + name + ".html")
	sheikh := sheikh.SheikhComp(sheikhData)
	footer := footer.FooterComp(pages, name)
	err = base.BaseComp(name, title, sheikh, footer).Render(context.Background(), f)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}
}

func RenderMoviesPage(outDir string, pages []footer.Page, moviesFile string) {
	moviesJson, err := os.ReadFile(moviesFile)
	if err != nil {
		log.Fatalf("Failed to read movies data: %v", err)
	}
	var moviesData []movies.Movie
	err = json.Unmarshal(moviesJson, &moviesData)
	if err != nil {
		log.Fatalf("Failed to decode movies data: %v", err)
	}
	name := "movies"
	title := GetPageByName(pages, name).Title
	f := CreateFile(outDir + name + ".html")
	movies := movies.MoviesComp(moviesData)
	footer := footer.FooterComp(pages, name)
	err = base.BaseComp(name, title, movies, footer).Render(context.Background(), f)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}
}

func RenderGamesPage(outDir string, pages []footer.Page, gamesFile string) {
	gamesJson, err := os.ReadFile(gamesFile)
	if err != nil {
		log.Fatalf("Failed to read games data: %v", err)
	}
	var gamesData []games.Game
	err = json.Unmarshal(gamesJson, &gamesData)
	if err != nil {
		log.Fatalf("Failed to decode games data: %v", err)
	}
	name := "games"
	title := GetPageByName(pages, name).Title
	f := CreateFile(outDir + name + ".html")
	g := games.GamesComp(gamesData)
	footer := footer.FooterComp(pages, name)
	err = base.BaseComp(name, title, g, footer).Render(context.Background(), f)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}
}

func CreateFile(name string) *os.File {
	f, err := os.Create(name)
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	return f
}

func GetPageByName(pages []footer.Page, name string) *footer.Page {
	for _, page := range pages {
		if page.Name == name {
			return &page
		}
	}
	return nil
}

