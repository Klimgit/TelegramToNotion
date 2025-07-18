package notion

import (
	"TelegramToNotion/internal/config"
	"TelegramToNotion/internal/state"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PageRequest struct {
	Parent     Parent     `json:"parent"`
	Properties Properties `json:"properties"`
}

type Parent struct {
	DatabaseID string `json:"database_id"`
}

type Properties struct {
	Title    TitleProperty `json:"Title"`
	Status   Checkbox      `json:"Status"`
	DueDate  DateProperty  `json:"Сроки"`
	Priority SelectOption  `json:"Приоритет"`
	Urgency  SelectOption  `json:"Срочность"`
}

type TitleProperty struct {
	Title []TextContent `json:"title"`
}

type TextContent struct {
	Text Text `json:"text"`
}

type Text struct {
	Content string `json:"content"`
}

type Checkbox struct {
	Checkbox bool `json:"checkbox"`
}

type DateProperty struct {
	Date DateValue `json:"date"`
}

type DateValue struct {
	Start string `json:"start"`
}

type SelectOption struct {
	Select Option `json:"select"`
}

type Option struct {
	Name string `json:"name"`
}

func CreatePage(cfg *config.Config, state *state.State) error {
	requestBody := PageRequest{
		Parent: Parent{
			DatabaseID: cfg.NotionDatabaseID,
		},
		Properties: Properties{
			Title: TitleProperty{
				Title: []TextContent{
					{Text: Text{Content: state.Title}},
				},
			},
			Status: Checkbox{
				Checkbox: state.Status,
			},
			DueDate: DateProperty{
				Date: DateValue{Start: state.DueDate},
			},
			Priority: SelectOption{
				Select: Option{Name: state.Priority},
			},
			Urgency: SelectOption{
				Select: Option{Name: state.Urgency},
			},
		},
	}

	jsonBody, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest(
		"POST",
		"https://api.notion.com/v1/pages",
		bytes.NewBuffer(jsonBody),
	)

	req.Header.Add("Authorization", "Bearer "+cfg.NotionToken)
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("Notion API error: %d - %s", resp.StatusCode, string(body))
	}

	return nil
}
