package user

import (
	"fmt"
	"time"
)

type User struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

func (this *User) SetUser(id int, nombre string, fechaAlta time.Time, status bool) {
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaAlta
	this.Status = status
}

func (this *User) GetUser() {
	fmt.Println("Id: ", this.Id)
	fmt.Println("Nombre: ", this.Nombre)
	fmt.Println("Fecha Alta: ", this.FechaAlta)
	fmt.Println("Status: ", this.Status)
}
