package main

import (
    "context"
    "fmt"
    "github.com/aws/aws-lambda-go/lambda"
)

// Request estrutura o input da função Lambda
type Request struct {
    N int `json:"n"`
}

// Response estrutura o output da função Lambda
type Response struct {
    Result []string `json:"result"`
}

// Handler é a função Lambda que implementa o algoritmo FizzBuzz
func Handler(ctx context.Context, req Request) (Response, error) {
    var result []string

    for i := 1; i <= req.N; i++ {
        if i%15 == 0 {
            result = append(result, "FizzBuzz")
        } else if i%3 == 0 {
            result = append(result, "Fizz")
        } else if i%5 == 0 {
            result = append(result, "Buzz")
        } else {
            result = append(result, fmt.Sprintf("%d", i))
        }
    }

    return Response{Result: result}, nil
}

func main() {
    lambda.Start(Handler)
}
