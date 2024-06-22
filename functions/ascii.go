package functions

import "fmt"

// Ascii generates ASCII art for each string in txtslice using patterns from slice.
func Ascii(slice []string, txtslice []string) {
	for j, txt := range txtslice {
		if txt != "" {
			// Iterate over 8 lines to construct ASCII art
			for i := 0; i < 8; i++ {
				for _, v := range txt {
					// Calculate and print the corresponding line of ASCII art
					firstLine := int(v-32)*9 + 1 + i
					fmt.Print(slice[firstLine])
				}
				fmt.Println()
			}
		} else if j != len(txtslice)-1 {
			// Print an empty line if the current text is empty and it's not the last element
			fmt.Println("")
		}
	}
}
