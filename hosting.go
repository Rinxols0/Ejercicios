package main 

type Hosting struct {
	id int			`json:"id"`	
	name string		`json:"name"`
	cores int		`json:"cores"`
	memory int		`json:"memory"`	
	disc int		`json:"disc"`
}

type Hostings []Hosting