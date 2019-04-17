# Gocap
> Calculate permutations and combinations in go

## Installation

With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```sh
go get -u github.com/jaumecapdevila/gocap
```

## Usage

First of all you need to create a new instance of the Operation struct using the `NewOperation` constructor:

```go
operation := NewOperation(3, 1, false)
```

This constructor receives three parameters:

1. N: the number of different types
2. R: the number of types to choose
3. Repetition: Indicates if is possible or not to repeat a type

Now, you can use the `CombinationCalculator` and the `PermutationCalculator` to get the result of the operation:

```go
calculator := NewCombinationCalculator()

result := calculator.Calculate(operation)
```
## Meta

Jaume Capdevila – [@otherjaume](https://twitter.com/otherjaume) – contact@jaumecapdevila.net

Distributed under the MIT license.

[https://github.com/jaumecapdevila/gocap](https://github.com/jaumecapdevila/gocap)

## Contributing

1. Fork it (<https://github.com/jaumecapdevila/gocap/fork>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request
