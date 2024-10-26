package make

type Manga struct {
	Name string `json:"name"`
	Author string `json:"author"`
	Chapters int `json:"chapters"`
}

func main() {

	// JSON = Javascript Object Notation
	mangaDataJson := `
[
	{
		"name": "Ajin",
		"author": "Unknow",
		"chapters": 83
	},
	{
		"name": "One Punch Man",
		"author": "One, Yusuke Murata",
		"chapters": 200
	}	
]`

	var x 


}
