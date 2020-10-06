package client

import (
	"github.com/jarcoal/httpmock"
	"net/http"
)

func (s *Suite) TestCreateTeam() {
	// Setup API mock
	u := apiUrl + pathTeamCreate
	expected := Team{
		ID:          676974,
		AccountID:   317418,
		Name:        "foobar",
		AccessLevel: TeamAccessStandard,
	}
	// FIXME: currently API returns `200 OK` on successful create; but it should
	//  instead return `201 Created`.
	//  https://github.com/rollbar/terraform-provider-rollbar/issues/8
	sr := httpmock.NewStringResponse(http.StatusOK, teamCreateResponse)
	sr.Header.Add("Content-Type", "application/json")
	r := httpmock.ResponderFromResponse(sr)
	httpmock.RegisterResponder("POST", u, r)

	// Successful create
	actual, err := s.client.CreateTeam("foobar", TeamAccessStandard)
	s.Nil(err)
	s.Equal(expected, actual)

	// Invalid name
	_, err = s.client.CreateTeam("", TeamAccessStandard)
	s.NotNil(err)

	// Internal server error
	r = httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, errResult500)
	httpmock.RegisterResponder("POST", u, r)
	_, err = s.client.CreateTeam("foobar", TeamAccessStandard)
	s.NotNil(err)

	// Server unreachable
	httpmock.Reset()
	_, err = s.client.CreateTeam("foobar", TeamAccessStandard)
	s.NotNil(err)
}

func (s *Suite) TestListTeams() {
	// Setup API mock
	u := apiUrl + pathTeamList
	expected := []Team{
		{
			AccessLevel: "everyone",
			AccountID:   317418,
			ID:          662037,
			Name:        "Everyone",
		},
		{
			ID:          676974,
			AccountID:   317418,
			Name:        "foobar",
			AccessLevel: TeamAccessStandard,
		},
		{
			AccessLevel: "owner",
			AccountID:   317418,
			ID:          662036,
			Name:        "Owners",
		},
	}
	sr := httpmock.NewStringResponse(http.StatusOK, teamListResponse)
	sr.Header.Add("Content-Type", "application/json")
	r := httpmock.ResponderFromResponse(sr)
	httpmock.RegisterResponder("GET", u, r)

	// Successful list
	actual, err := s.client.ListTeams()
	s.Nil(err)
	s.Equal(expected, actual)

	// Internal server error
	r = httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, errResult500)
	httpmock.RegisterResponder("GET", u, r)
	_, err = s.client.ListTeams()
	s.NotNil(err)

	// Server unreachable
	httpmock.Reset()
	_, err = s.client.ListTeams()
	s.NotNil(err)
}

const teamCreateResponse = `
{
    "err": 0,
    "result": {
        "access_level": "standard",
        "account_id": 317418,
        "id": 676974,
        "name": "foobar"
    }
}
`

const teamListResponse = `
{
    "err": 0,
    "result": [
        {
            "access_level": "everyone",
            "account_id": 317418,
            "id": 662037,
            "name": "Everyone"
        },
        {
            "access_level": "standard",
            "account_id": 317418,
            "id": 676974,
            "name": "foobar"
        },
        {
            "access_level": "owner",
            "account_id": 317418,
            "id": 662036,
            "name": "Owners"
        }
    ]
}


`
