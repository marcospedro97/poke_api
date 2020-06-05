# poke_api
My pokemons list api in go with gin and GORM
To use it you'll need

- Go (I used 1.14.3 on manjaro, IDK if I'll face any trouble in other versions)
- Redis (I used the docker oficial image)
- go.mod requirements

After this you'll be able to run
Only one path is locked with authentication, it's the create, wich enable
you to create a new pokemon in the DB, the list of paths is bellow

/sign_up path to create a new user

{
	"email":"alfred@gmail",
	"password":"123456789"
}

/sign_in login of a existent user

{
	"email":"alfred@gmail",
	"password":"123456789"
}

will return a token that should be placed in the headers with key "Authentication"

/ root path to list all pokemons in the DB

[
	{
	  "Name": "vasdfasd",
	  "Code": 21,
	  "Type": "f",
	  "Next_evolution": 22,
	  "Previous_evolution": 20
	}
]

/:id will return a single pokemon from the DB 

/21

{
  "Code": 21,
  "Name": "vasdfasd",
  "Type": "f",
  "Next_evolution": 22,
  "Previous_evolution": 20
}

/user/create creates a new pokemon, stores the user who created it and return
the pokemon

{
	"Name":"adsfasdf",
	"Code": 22,
	"Type": "f",
	"Next_evolution": 3,
	"Previous_evolution": 1
}

