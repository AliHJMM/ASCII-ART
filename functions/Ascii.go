package functions

func ascii(slice []string, txtslice []string) {
	for j, txt := range txtslice {
		if txt != "" {
			for i := 0; i < 8; i++ {
				for _, v := range txt {
					firstLine := int(v-32)*9 + 1 + i
					fmt.Print(slice[firstLine])
				}
				fmt.Println()
			}
		} else if j != len(txtslice)-1 {
			fmt.Println("")
		}
	}
}