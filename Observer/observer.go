package main

type Observer interface {
    update(string, int)
}