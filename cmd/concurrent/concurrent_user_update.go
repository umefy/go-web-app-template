package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func updateUser(id int, email string, age int) error {
	req, err := http.NewRequest(http.MethodPatch, "http://localhost:8082/api/v1/users/1", nil)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Body = io.NopCloser(bytes.NewBufferString(fmt.Sprintf(`{"email": "%s", "age": %d}`, email, age)))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update user: %s", resp.Status)
	}
	return nil
}

func main() {
	ctx := context.Background()
	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		return updateUser(1, "test11@test.com", 11)
	})

	g.Go(func() error {
		return updateUser(1, "test22@test.com", 22)
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
