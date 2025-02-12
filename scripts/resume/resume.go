package resume

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Certification struct {
	Name         string `yaml:"name"`
	Organization string `yaml:"organization"`
	Date         string `yaml:"date"`
	Link         string `yaml:"link"`
}

type Education struct {
	Degree      string   `yaml:"degree"`
	Institution string   `yaml:"institution"`
	Description string   `yaml:"description"`
	Subjects    []string `yaml:"subjects"`
	Date        string   `yaml:"date"`
}

type Language struct {
	Name  string `yaml:"name"`
	Level string `yaml:"level"`
}

type Skill struct {
	Name  string `yaml:"name"`
	Level int    `yaml:"level"`
}

type SkillCategory struct {
	Category string  `yaml:"category"`
	Items    []Skill `yaml:"items"`
}

type WorkExperience struct {
	Title            string   `yaml:"title"`
	Company          string   `yaml:"company"`
	Date             string   `yaml:"date"`
	Description      string   `yaml:"description"`
	Responsibilities []string `yaml:"responsibilities"`
}

type Resume struct {
	Name           string           `yaml:"name"`
	Location       string           `yaml:"location"`
	Email          string           `yaml:"email"`
	Website        string           `yaml:"website"`
	ProfileImage   string           `yaml:"profile_image"`
	Profile        string           `yaml:"profile"`
	WorkExperience []WorkExperience `yaml:"work_experience"`
	Education      []Education      `yaml:"education"`
	Languages      []Language       `yaml:"languages"`
	Skills         []SkillCategory  `yaml:"skills"`
	Certifications []Certification  `yaml:"certifications"`
}

type htmlBufferWriter struct {
	buffer *[]byte
}

func loadResume() (*Resume, error) {
	customFile := os.Getenv("RESUME_FILE")
	if customFile == "" {
		customFile = "data/resume.yaml"
	}

	if _, err := os.Stat(customFile); os.IsNotExist(err) {
		log.Println("No resume file found. Displaying placeholder template.")
		return nil, nil
	}

	data, err := os.ReadFile(customFile)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file: %v", err)
	}

	var resume Resume
	err = yaml.Unmarshal(data, &resume)
	if err != nil {
		return nil, fmt.Errorf("error parsing YAML: %v", err)
	}

	profileImage := os.Getenv("PROFILE_IMAGE")
	if profileImage == "" {
		profileImage = "assets/img/profile.jpg"
	}
	resume.ProfileImage = profileImage

	return &resume, nil
}

func RenderResumeHTML() ([]byte, error) {
	resumeData, err := loadResume()
	if err != nil {
		return nil, err
	}

	funcMap := template.FuncMap{
		"seq": func(n int) []int {
			out := make([]int, n)
			for i := range out {
				out[i] = i
			}
			return out
		},
		"sub": func(a, b int) int {
			return a - b
		},
	}

	tmpl, err := template.New("index.html").Funcs(funcMap).ParseFiles("index.html")
	if err != nil {
		return nil, fmt.Errorf("error loading HTML template: %v", err)
	}

	var htmlBuffer []byte
	writer := &htmlBufferWriter{buffer: &htmlBuffer}
	if err := tmpl.Execute(writer, resumeData); err != nil {
		return nil, fmt.Errorf("error rendering template: %v", err)
	}

	return htmlBuffer, nil
}

func (w *htmlBufferWriter) Write(p []byte) (int, error) {
	*w.buffer = append(*w.buffer, p...)
	return len(p), nil
}
