package ghcheck

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"context"
	"io"
	"log"
	"net/http"
	"time"
	"context"
	"github.com/google/go-github/v60/github"
	
)

func GithubCall(endpoint string) {
	client := github.NewClient(nil)
	context := context.Background()
// list all organizations for user "willnorris"
	repo, _, err := client.Repositories.Get(context, "subsurface", "subsurface")

	fmt.Println(orgs)
}
