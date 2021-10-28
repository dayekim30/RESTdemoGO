package model

type Member struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	FavoriteTeam string `json:"team"`
}

var members = []Member{
	{ID: "1", Name: "Rowland", FavoriteTeam: "Mumbai Indians"},
	{ID: "2", Name: "Neel", FavoriteTeam: "Chennai Super Kings"},
	{ID: "3", Name: "Kunal", FavoriteTeam: "Rajasthan Royals"},
	{ID: "4", Name: "Tanuj", FavoriteTeam: "Sunrisers Hyderabad"},
	{ID: "5", Name: "Daye", FavoriteTeam: "Chennai super Kings"},
}

func Getall() []Member {
	memberList := members
	return memberList
}

func Getbyid(ID string) Member {
	var i = 0
	for _, a := range members {
		if a.ID == ID {
			return members[i]
		}
		i++
	}
	return Member{}
}

func Editmember(editmember Member) []Member {
	var i = 0
	for _, a := range members {
		if a.ID == editmember.ID {
			members[i] = editmember
			return members
		}
	}
	return []Member{}
}

func Addmember(newmember Member) []Member {
	members = append(members, newmember)
	return members
}

func Deletemember(ID string) []Member {
	var i = 0
	for _, a := range members {
		if a.ID == ID {
			members = RemoveIndex(members, i)
			return members
		}
		i++
	}
	return []Member{}
}
func RemoveIndex(s []Member, index int) []Member {
	return append(s[:index], s[index+1:]...)
}
