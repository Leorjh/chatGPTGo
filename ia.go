package main

type IADoer interface {
	GetHistory(text string) (string, error)
}
