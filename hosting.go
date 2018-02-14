package main 

type Hosting struct {
	Id int			`json:"id"`	
	Name string		`json:"name"`
	Cores int		`json:"cores"`
	Memory int		`json:"memory"`	
	Disc int		`json:"disc"`
}

type Hostings []Hosting