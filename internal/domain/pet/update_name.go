package pet

import "time"

func (pet *Pet) UpdateName(newName string) {
	pet.Name = newName
	pet.UpdatedAt = time.Now().UTC()
}
