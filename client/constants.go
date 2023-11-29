package aoc

type Level uint8

const (
	A = iota + 1
	B
)

type Year uint16

const (
	Y2017 Year = iota + 2017
	Y2018
	Y2019
	Y2020
	Y2021
	Y2022
	Y2023
)

type Day uint8

const (
	Day01 Day = iota + 1
	Day02
	Day03
	Day04
	Day05
	Day06
	Day07
	Day08
	Day09
	Day10
	Day11
	Day12
	Day13
	Day14
	Day15
	Day16
	Day17
	Day18
	Day19
	Day20
	Day21
	Day22
	Day23
	Day24
	Day25
)

type SubmissionResult uint8

const (
	Correct SubmissionResult = iota
	AlreadyCompleted
	Incorrect
	TooRecent
)
