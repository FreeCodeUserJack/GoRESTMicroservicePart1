package github

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)


func TestCreateRepoRequestAsJson(t *testing.T) {
	t.Run("Marshal CreateRepoRequest", func(t *testing.T) {
		request := CreateRepoRequest{
			Name:"Hello-World",
			Description: "This is first repo",
			Homepage: "https://github.com",
			Private: false,
			HasIssues: true,
			HasProjects: true,
			HasWiki: true,
		}
		bytes, err := json.Marshal(request)
	
		want := `{"name":"Hello-World","description":"This is first repo","homepage":"https://github.com","private":false,"has_issues":true,"has_projects":true,"has_wiki":true}`
	
		if err != nil {
			t.Errorf("didn't want error but got %v", err)
		}
	
		if bytes == nil || len(bytes) <= 0 {
			t.Error("expected to get json []byte but didn't get it")
		}
	
		if string(bytes) != want {
			t.Errorf("got %v, want %v", string(bytes), want)
		}
	
		fmt.Println(string(bytes))
	})

	t.Run("Unmarshal CreateRepoRequest", func(t *testing.T) {
		jsonString := `{"name":"Hello-World","description":"This is first repo","homepage":"https://github.com","private":false,"has_issues":true,"has_projects":true,"has_wiki":true}`

		var unmarshalledRequest CreateRepoRequest

		err := json.Unmarshal([]byte(jsonString), &unmarshalledRequest)

		want := CreateRepoRequest{
			Name:"Hello-World",
			Description: "This is first repo",
			Homepage: "https://github.com",
			Private: false,
			HasIssues: true,
			HasProjects: true,
			HasWiki: true,
		}

		if err != nil {
			t.Errorf("expected no error but got %v", err)
		}

		if !reflect.DeepEqual(unmarshalledRequest, want) {
			t.Errorf("got %v, want %v", unmarshalledRequest, want)
		}
	})
}

// should create a test for CreateRepoReponse obj marshal/unmarshal