package main

import (
		"bufio"
		"os"
		"fmt"
		"path/filepath"
		)

func read_lines(shortest string, longest string, filename string) (string, string) {

	file, e := os.Open(filename)
	
	if e != nil {
		panic(e)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		current := scanner.Text()

		if len(current) > len(longest) {
			longest = current
		}

		if len(current) < len(shortest) || len(shortest) == 0 {
			shortest = current
		}
			
	}

	return shortest, longest

}

func go_through_files (dir string) map[string]string {

	shortest	:= ""
	longest		:= ""
	
	file_line := map[string]string{}

	files, e := os.ReadDir(dir) 

	if e != nil {
		panic(e)
	}
	
	for _, file := range files {

		file_path := filepath.Join(dir, file.Name())

		current_shortest, current_longest := read_lines(shortest, longest, file_path) 

		if shortest != current_shortest {

			delete(file_line, shortest)
			file_line[current_shortest] = file.Name()
			shortest = current_shortest	

		}
	
		if longest != current_longest {
		
			delete(file_line, longest)
			file_line[current_longest] = file.Name() 
			longest = current_longest	

		}
	} 

	return file_line
	
}

func _print(mapping map[string]string) {

	for k, v := range mapping {
		fmt.Println(v, k)
	}
		
}

func main() {

	_print(go_through_files(os.Args[1]))	
	
}
