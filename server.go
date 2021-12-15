package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kaitoimai01/gqlgen_auth0_sample/internal/auth0"
	"github.com/kaitoimai01/gqlgen_auth0_sample/internal/graph/generated"
	"github.com/kaitoimai01/gqlgen_auth0_sample/internal/graph/resolver"
	"github.com/kaitoimai01/gqlgen_auth0_sample/internal/mailer"
	"github.com/sendgrid/sendgrid-go"
	"gopkg.in/auth0.v5/management"
)

func main() {
	ac, err := initAuth0Client()
	if err != nil {
		log.Fatalf("could not init auth0 client: %v", err)
	}

	mc, err := initMailerClient()
	if err != nil {
		log.Fatalf("could not init sendgrid client: %v", err)
	}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &resolver.Resolver{
					Auth0Client: ac,
					MailClient:  mc,
				},
			},
		),
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	port := os.Getenv("PORT")
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// クライアント ID 、クライアントシークレットを指定して Auth0 クライアントを作成します。
func initAuth0Client() (*auth0.Client, error) {
	var (
		domain       = os.Getenv("AUTH0_DOMAIN")
		clientID     = os.Getenv("AUTH0_CLIENT_ID")
		clientSecret = os.Getenv("AUTH0_CLIENT_SECRET")
	)

	if domain == "" || clientID == "" || clientSecret == "" {
		return nil, fmt.Errorf("plese set env vars for auth0")
	}

	m, err := management.New(domain, management.WithClientCredentials(clientID, clientSecret))
	if err != nil {
		return nil, fmt.Errorf("could not create an auth0 management: %w", err)
	}

	return auth0.NewClient(m), nil
}

// API キーを指定して SendGrid クライアントを作成します。
func initMailerClient() (*mailer.Client, error) {
	var (
		apiKey = os.Getenv("SENDGRID_API_KEY")
	)
	if apiKey == "" {
		return nil, fmt.Errorf("please set env vars for sendgrid")
	}

	sg := sendgrid.NewSendClient(apiKey)
	return mailer.NewClient(sg), nil
}
