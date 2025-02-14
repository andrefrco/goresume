# goresume

Goresume is a simple tool that generates a resume from a YAML file, creating both a web page and a downloadable PDF.

## Features

- **Web View:** The resume is served as a web page to build and test.
- **PDF Generator:** Generate a PDF file of the resume with a single command.
- **Dynamic Structure:** Resume information is loaded from a YAML file, making maintenance easy.

## Dependencies
- **Golang:** Required to run the script.
- **Google Chrome:** Needed to generate the PDF.

## How to Use

1. **Clone the repository:**

   ```bash
   git clone https://github.com/andrefrco/goresume.git
   cd goresume
   ```

2. **Configure resume data:**
   - Edit the `data/resume.yaml` file with your information.

3. **Add your profile picture:**
   - Paste your photo inside `/assets/img/profile.jpg`.

4. **Run the web server:**

   ```bash
   go run scripts/main.go -mode=serve
   ```

   Access http://localhost:8080 to view the resume in your browser.

5. **Generate the PDF:**

   ```bash
   go run scripts/main.go -mode=pdf
   ```

   The `resume.pdf` file will be generated in the project root.
