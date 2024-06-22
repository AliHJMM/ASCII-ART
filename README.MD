# ASCII ART

## DESCRIPTION
ASCII Art is a program that allows users to generate ASCII ART from text. Users can enter text, and get the ASCII ART representation of the text.

### AUTHORS
- Ali Hasan Jasim (https://github.com/AliHJMM)
- Habib Mansoor (https://github.com/7abib04)
- Husain Ali (https://github.com/hujaafar)


## Usage

### How to Run
1. Clone the Repo
2. Navigate to the project directory
3. Run `go run main.go text`
4. Example "go run main.go Ali"

### Implementation Details

#### Algorithm

The ASCII ART generation is done by reading banner templates from text files. Each banner template file contains ASCII representations for characteristics. The server reads these files, maps characters to their ASCII ART representation, and then constructs the ASCII ART for the given input text.