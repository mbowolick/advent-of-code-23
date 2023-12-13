// mbowolick advent-of-code-2023/day3 attempt
//
// https://adventofcode.com/2023/day/3
//
// --- Day 3: Gear Ratios ---
//
// An engine part seems to be missing from the engine, but nobody
// can figure out which one.
//
// If you can add up all the part numbers in the `engine schematic`,
// the puzzle input, it should be easy to work out which part is missing.
//
// Engine schematic = visual representation of the engine.
// Adjacent numbers in the schematic to symbols (even diagonally)
// should be considered a part number which should included in your sum.
// Periods (".") do not count as a symbol.
//
// Example input:
// 467..114..
// ...*......
// ..35..633.
// ......#...
// 617*......
// .....+.58.
// ..592.....
// ......755.
// ...$.*....
// .664.598..

// All numbers should be considered except for:
// 114 (top right) and 58 (middle right) are not adjacent to a symbol.
// So all numbers (minus the ones above) should sum to... 4361.
// 467 + 35 + 633 + 617 + 592 + 755 + 664 + 598 = 4361

//     | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 |
//     |---|---|---|---|---|---|---|---|---|---|
//  0: | 4 | 6 | 7 | . | . | 1 | 1 | 4 | . | . |
//  1: | . | . | . | * | . | . | . | . | . | . |
//  2: | . | . | 3 | 5 | . | . | 6 | 3 | 3 | . |
//  3: | . | . | . | . | . | . | # | . | . | . |
//  4: | 6 | 1 | 7 | * | . | . | . | . | . | . |
//  5: | . | . | . | . | . | + | . | 5 | 8 | . |
//  6: | . | . | 5 | 9 | 2 | . | . | . | . | . |
//  7: | . | . | . | . | . | . | 7 | 5 | 5 | . |
//  8: | . | . | . | $ | . | * | . | . | . | . |
//  9: | . | 6 | 6 | 4 | . | 5 | 9 | 8 | . | . |

// LOGIC:
// Scan and detect symbols then traverse back through the matrix for surrounding numbers
// [Part 1 - Detect] If symbol is found at N, M (X, Y) then detect number by:
// - If row is not first row, check row above: X: [(N - 1), (N), (N + 1)] Y: (M - 1)
// - Check current row : X: [(N - 1), (N + 1)] Y: M
// - If row is not last row, check row below: X: [(N - 1), (N), (N + 1)] Y: (M + 1)
// - If Column is first or last column then don't detect on (N - 1) or (N + 1)
// [Part 2 - Traverse for number] When a number is detected continue in the direction
// until a period or a wall then convert to number and add to sum.

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readInput(path string) ([]string, error) {
	input, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer input.Close()

	var lines []string

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	
	return lines, scanner.Err()
}

// buildMatrix returns matrix of [][]string with int of rows and cols
func buildMatrix(lines []string) ([][]string, int, int) {
	var matrix [][]string
	for _, row := range lines {
		line_split := strings.Split(row, "")
		matrix = append(matrix, line_split)
	}
	return matrix, len(lines), len(matrix[0])
}

func returnNumberWithRange(str string, start int, end int) (int) {
	substring := str[start:end+1]
	number, err := strconv.Atoi(substring)
	if err != nil {
		return 0
	}
	return number
}

func main() {
	fmt.Print("Starting day 3\n\n")
	// Step 1 - read input schematic
	inputDocLines, err := readInput("real_input.txt")

	if err != nil {
		fmt.Println("Issue reading input:", err.Error())
	}
	
	// Step 2 - build matrix
	SCHEMATIC_MATRIX, ROWS_MAX, COLS_MAX  := buildMatrix(inputDocLines)
	

	IS_SYMBOL_REGEX := `[^.|\d|a-zA-Z| |\n]+`
	IS_DIGIT_REGEX := `\d+`

	// [Part 1 - Detect] If symbol is found at N, M (X, Y) then detect number by:
	// - If row is not first row, check row above: X: [(N - 1), (N), (N + 1)] Y: (M - 1)
	// - Check current row : X: [(N - 1), (N + 1)] Y: M
	// - If row is not last row, check row below: X: [(N - 1), (N), (N + 1)] Y: (M + 1)
	// - If Column is first or last column then don't detect on (N - 1) or (N + 1)
	// [Part 2 - Traverse for number] When a number is detected continue in the direction
	// until a period or a wall then convert to number and add to sum.

	// Step 3 - detect numbers and tranverse back through the matrix to build the sum
	SCHEMATIC_SUM := 0
	var min_pos int
	var max_pos int

	for row_m, row := range SCHEMATIC_MATRIX {
		
		// Simplify the search by checking if the row contains a symbol, otherwise skip row
		row_str := strings.Join(row, "")
		row_contains_symbol, _ := regexp.MatchString(IS_SYMBOL_REGEX, row_str)
		if(row_contains_symbol) {
			// For each of the cells in the row find the symbol and then transveration around the symbol to find numbers
			for col_n, cell := range row {
				cell_matches_symbol, _ := regexp.MatchString(IS_SYMBOL_REGEX, cell)
				
				if(cell_matches_symbol) {
					// ROW 'm' has a symbol, now move up (if not first row), below (if not last row), left (if not first colum), right (if not last column)
					
					// Part 1 - search vertically (up) --------------------------------------
					if(row_m != 0) {
						// Not first row, move up and search directly above and diagonall left (if not first column) and right (if not last column)
						above := SCHEMATIC_MATRIX[row_m-1][col_n]
						above_is_a_digit, _ := regexp.MatchString(IS_DIGIT_REGEX, above)

						if(above_is_a_digit) {
							// expand to find the number
							min_pos = col_n
							max_pos = col_n

							// if(col_n==0) {
							// 	// Move right until not number
							for i := col_n; i < COLS_MAX; i++ {
								new_cell_is_a_digit, _ := regexp.MatchString(IS_DIGIT_REGEX, SCHEMATIC_MATRIX[row_m-1][i])
								if(new_cell_is_a_digit) {
									max_pos = i
								} else {
									break
								}
							}
							// } else {
							// 	// Move left until not number
							for i := col_n; i >= 0; i-- {
								new_cell_is_a_digit, _ := regexp.MatchString(IS_DIGIT_REGEX, SCHEMATIC_MATRIX[row_m-1][i])
								if(new_cell_is_a_digit) {
									min_pos = i
								} else {
									break
								}
							}
							// }
							
							num := returnNumberWithRange(strings.Join(SCHEMATIC_MATRIX[row_m-1], ""),min_pos,max_pos)
							fmt.Printf("Num found: %d\n", num)

							SCHEMATIC_SUM += num


						} else {
							
							// check diagonally
							if(col_n > 0) {
								// Not first col, move left check if number, if yes, continue until not
								left_is_a_digit, _ := regexp.MatchString(IS_DIGIT_REGEX, SCHEMATIC_MATRIX[row_m-1][col_n-1])
								if(left_is_a_digit) {
									// continue left until not number
									min_pos = col_n-1
									max_pos = col_n-1

									for i := col_n-1; i >= 0; i-- {
										new_cell_is_a_digit, _ := regexp.MatchString(IS_DIGIT_REGEX, SCHEMATIC_MATRIX[row_m-1][i])
										if(new_cell_is_a_digit) {
											min_pos = i
										} else {
											break
										}
									}
									num := returnNumberWithRange(strings.Join(SCHEMATIC_MATRIX[row_m-1], ""),min_pos,max_pos)
									fmt.Printf("Num found: %d\n", num)

									SCHEMATIC_SUM += num
								}
							} 
								
							if(col_n < COLS_MAX){
								// Not last col, move right check if number, if yes, continue until not
								right_is_a_digit, _ := regexp.MatchString(IS_DIGIT_REGEX, SCHEMATIC_MATRIX[row_m-1][col_n+1])
								if(right_is_a_digit){
									// continue left until not number
									min_pos = col_n+1
									max_pos = col_n+1
									for i := col_n+1; i < COLS_MAX; i++ {
										new_cell_is_a_digit, _ := regexp.MatchString(IS_DIGIT_REGEX, SCHEMATIC_MATRIX[row_m-1][i])
										if(new_cell_is_a_digit) {
											max_pos = i
										} else {
											break
										}
									}

									// SCHEMATIC_SUM += returnNumberWithRange(strings.Join(SCHEMATIC_MATRIX[row_m-1], ""),min_pos,max_pos)
									num := returnNumberWithRange(strings.Join(SCHEMATIC_MATRIX[row_m-1], ""),min_pos,max_pos)

									fmt.Printf("Num found: %d\n", num)
									SCHEMATIC_SUM += num
								}
							}

						}


					}

					// Part 2 - search vertically (down) ---------------------------------------
					if( row_m != ROWS_MAX-1) {
						// Not last row, move down and search directly above, left (if not first column and right if not last column)
						above := SCHEMATIC_MATRIX[row_m+1][col_n]
						above_is_a_digit, _ := regexp.MatchString(IS_DIGIT_REGEX, above)

						if(above_is_a_digit) {
							// expand to find the number
							min_pos = col_n
							max_pos = col_n

							// if(col_n==0) {
							// 	// Move right until not number
							for i := col_n; i < COLS_MAX; i++ {
								new_cell_is_a_digit, _ := regexp.MatchString(IS_DIGIT_REGEX, SCHEMATIC_MATRIX[row_m+1][i])
								if(new_cell_is_a_digit) {
									max_pos = i
								} else {
									break
								}
							}
							// } else {
							// 	// Move left until not number
							for i := col_n; i >= 0; i-- {
								new_cell_is_a_digit, _ := regexp.MatchString(IS_DIGIT_REGEX, SCHEMATIC_MATRIX[row_m+1][i])
								if(new_cell_is_a_digit) {
									min_pos = i
								} else {
									break
								}
							}
							// }

							num := returnNumberWithRange(strings.Join(SCHEMATIC_MATRIX[row_m+1], ""),min_pos,max_pos)
							fmt.Printf("Num found: %d\n", num)
							SCHEMATIC_SUM += num
							// SCHEMATIC_SUM += returnNumberWithRange(strings.Join(SCHEMATIC_MATRIX[row_m+1], ""),min_pos,max_pos)
						} else {
							
							// check diagonally
							if(col_n > 0) {
								// Not first col, move left check if number, if yes, continue until not
								left_is_a_digit, _ := regexp.MatchString(IS_DIGIT_REGEX, SCHEMATIC_MATRIX[row_m+1][col_n-1])
								if(left_is_a_digit) {
									// continue left until not number
									min_pos = col_n-1
									max_pos = col_n-1

									for i := col_n-1; i >= 0; i-- {
										new_cell_is_a_digit, _ := regexp.MatchString(IS_DIGIT_REGEX, SCHEMATIC_MATRIX[row_m+1][i])
										if(new_cell_is_a_digit) {
											min_pos = i
										} else {
											break
										}
									}

									// SCHEMATIC_SUM += returnNumberWithRange(strings.Join(SCHEMATIC_MATRIX[row_m+1], ""),min_pos,max_pos)
									num := returnNumberWithRange(strings.Join(SCHEMATIC_MATRIX[row_m+1], ""),min_pos,max_pos)
									fmt.Printf("Num found: %d\n", num)
									SCHEMATIC_SUM += num
								}
							} 
								
							if(col_n < COLS_MAX){
								// Not last col, move right check if number, if yes, continue until not
								right_is_a_digit, _ := regexp.MatchString(IS_DIGIT_REGEX, SCHEMATIC_MATRIX[row_m+1][col_n+1])
								if(right_is_a_digit){
									// continue left until not number
									min_pos = col_n+1
									max_pos = col_n+1
									for i := col_n+1; i < COLS_MAX; i++ {
										new_cell_is_a_digit, _ := regexp.MatchString(IS_DIGIT_REGEX, SCHEMATIC_MATRIX[row_m+1][i])
										if(new_cell_is_a_digit) {
											max_pos = i
										} else {
											break
										}
									}

								// SCHEMATIC_SUM += returnNumberWithRange(strings.Join(SCHEMATIC_MATRIX[row_m+1], ""),min_pos,max_pos)
								num := returnNumberWithRange(strings.Join(SCHEMATIC_MATRIX[row_m+1], ""),min_pos,max_pos)
								fmt.Printf("Num found: %d\n", num)
								SCHEMATIC_SUM += num
								}
							}

						}
						
					}
					
					// Part 3 - search horizontally (left) --------------------------------------
					
					if( col_n != 0 ) {
						// Not first column, move left and check for numbers
						left_is_a_digit, _ := regexp.MatchString(IS_DIGIT_REGEX, SCHEMATIC_MATRIX[row_m][col_n-1])
						if(left_is_a_digit) {
							// continue left until not number
							min_pos = col_n-1
							max_pos = col_n-1

							for i := col_n-1; i >= 0; i-- {
								new_cell_is_a_digit, _ := regexp.MatchString(IS_DIGIT_REGEX, SCHEMATIC_MATRIX[row_m][i])
								if(new_cell_is_a_digit) {
									min_pos = i
								} else {
									break
								}
							}

							// SCHEMATIC_SUM += returnNumberWithRange(strings.Join(SCHEMATIC_MATRIX[row_m], ""),min_pos,max_pos)
							num := returnNumberWithRange(strings.Join(SCHEMATIC_MATRIX[row_m], ""),min_pos,max_pos)
							fmt.Printf("Num found: %d\n", num)
							SCHEMATIC_SUM += num
						}
					}
					
					// Part 4 - search horizontally (right) --------------------------------------
					if( col_n < COLS_MAX ) {
						// Not last column, move right and check for numbers
						right_is_a_digit, _ := regexp.MatchString(IS_DIGIT_REGEX, SCHEMATIC_MATRIX[row_m][col_n+1])
						if(right_is_a_digit){
							// continue left until not number
							min_pos = col_n+1
							max_pos = col_n+1
							for i := col_n+1; i < COLS_MAX; i++ {
								new_cell_is_a_digit, _ := regexp.MatchString(IS_DIGIT_REGEX, SCHEMATIC_MATRIX[row_m][i])
								if(new_cell_is_a_digit) {
									max_pos = i
								} else {
									break
								}
							}

							// SCHEMATIC_SUM += returnNumberWithRange(strings.Join(SCHEMATIC_MATRIX[row_m], ""),min_pos,max_pos)
							num := returnNumberWithRange(strings.Join(SCHEMATIC_MATRIX[row_m], ""),min_pos,max_pos)
							fmt.Printf("Num found: %d\n", num)
							SCHEMATIC_SUM += num
						}
					}

				}
			}
		} 
	}

	fmt.Printf("Schematic sum is: %d\n", SCHEMATIC_SUM)
		
}
