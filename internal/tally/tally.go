package tally

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type TallyApi struct {
	Key string
	Url string
}

const proposalQuery = `
query Proposals($chainId: ChainID!, $pagination: Pagination, $governors: [Address!], $sort: ProposalSort) {
    proposals(chainId: $chainId, pagination: $pagination, governors: $governors, sort: $sort) {
        id
        title
        eta
        governor {
            name
        }
        block {
            id
            number
            timestamp
            ts
        }
    }
}`

type Payload struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type Proposal struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	ETA      string `json:"eta"`
	Governor struct {
		Name string `json:"name"`
	} `json:"governor"`
	Block struct {
		ID        string `json:"id"`
		Number    int    `json:"number"`
		Timestamp string `json:"timestamp"`
		TS        string `json:"ts"`
	} `json:"block"`
}

type ProposalsResponse struct {
	Data struct {
		Proposals []Proposal `json:"proposals"`
	} `json:"data"`
}

func (t *TallyApi) GetLastProposal(ctx context.Context, governorAddress string) (*Proposal, error) {
	variables := map[string]interface{}{
		"chainId": "eip155:1",
		"governors": []string{
			governorAddress,
		},
		"sort": map[string]string{
			"field": "START_BLOCK",
			"order": "DESC",
		},
		"pagination": map[string]int{
			"limit":  1,
			"offset": 0,
		},
	}

	payload := Payload{
		Query:     proposalQuery,
		Variables: variables,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error marshalling the payload:", err)
		return &Proposal{}, err
	}

	req, err := http.NewRequest("POST", t.Url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Println("Error creating the request:", err)
		return &Proposal{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Api-key", t.Key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error executing the request:", err)
		return &Proposal{}, err

	}
	defer resp.Body.Close()

	var proposalsResponse ProposalsResponse

	err = json.NewDecoder(resp.Body).Decode(&proposalsResponse)
	if err != nil {
		log.Println("Error decoding proposal response:", err)
		return &Proposal{}, err
	}

	return &proposalsResponse.Data.Proposals[0], nil
}
