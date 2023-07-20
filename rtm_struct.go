package rtmpackage

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type List_Jbdesk struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Jbtitle_LJ   string             `bson:"jbtitle_lj,omitempty" json:"jbtitle_lj,omitempty"`
	Deskripsi_LJ string             `bson:"deskripsi_lj,omitempty" json:"deskripsi_lj,omitempty"`
	Deadline_LJ  string             `bson:"deadline_lj,omitempty" json:"deadline_lj,omitempty"`
	Priority_LJ  string             `bson:"priority_lj,omitempty" json:"priority_lj,omitempty"`
}
