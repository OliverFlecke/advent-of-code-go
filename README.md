# Advent of Code in Go!

Repository containing solutions for [Advent of Code](https://adventofcode.com). This follows the same pattern as [my previous year's solutions in Rust](https://github.com/oliverflecke/advent-of-code-rust).

The code contains functonality to fetch the input (with caching) and submit answers to the AoC website programatically with authentication. All of this is contained in the `client` directory through the `GetInput` and `SubmitAnswer` functions.

For examples of how to use it, see solution for year 2022 day 1 in `solutions/2022/1.go`.
