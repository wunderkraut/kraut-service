package main

import (
	"http"
	"json"
)

type Projects struct {
	projects map[string]Project
	order []string
}
func (projects *Projects) safe() {
	if *projects.project == nil {
		projects.projects = map[string]Project{}
		projects.order = []string{}
	}
}
func (projects *Projects) Add(name string, project Project) error {
	projects.safe()
	if _, exists := projects.projects[name]; !exists {
		projects.order = append(projects.order, name)
	}	
	projects.projects[name] = project
}
func (projects *Projects) Get(name string) (Project, error) {
	projects.safe()
	if project, found := projects.projects[name] {
		return project, nil
	} else {
		return project, errors.New("Project of that name ["+name+"] not found")
	}
}
func (projects *Projects) Order() []string {
	projects.safe()
	return projects.order
}

// List all projects
func handler_Projects(w http.ResponseWriter, r *http.Request) {
	projects
}
