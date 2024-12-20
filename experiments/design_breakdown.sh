#!/bin/bash

go run ../main.go --indicator 5 --ratio 0.8 --blockNum 1000 --repair=true  --perceptron=false --fast=false
go run ../main.go --indicator 5 --ratio 0.8 --blockNum 1000 --repair=true  --perceptron=true  --fast=false
go run ../main.go --indicator 5 --ratio 0.8 --blockNum 1000 --repair=false --perceptron=false --fast=false

go run ../main.go --indicator 5 --ratio 0.6 --blockNum 1000 --repair=true  --perceptron=false --fast=false
go run ../main.go --indicator 5 --ratio 0.6 --blockNum 1000 --repair=true  --perceptron=true  --fast=false
go run ../main.go --indicator 5 --ratio 0.6 --blockNum 1000 --repair=false --perceptron=false --fast=false

go run ../main.go --indicator 5 --ratio 0.4 --blockNum 1000 --repair=true  --perceptron=false --fast=false
go run ../main.go --indicator 5 --ratio 0.4 --blockNum 1000 --repair=true  --perceptron=true  --fast=false
go run ../main.go --indicator 5 --ratio 0.4 --blockNum 1000 --repair=false --perceptron=false --fast=false

go run ../main.go --indicator 5 --ratio 0.2 --blockNum 1000 --repair=true  --perceptron=false --fast=false
go run ../main.go --indicator 5 --ratio 0.2 --blockNum 1000 --repair=true  --perceptron=true  --fast=false
go run ../main.go --indicator 5 --ratio 0.2 --blockNum 1000 --repair=false --perceptron=false --fast=false
