package gcpcln

// // Object represents csv objects (files) from gcp
// type Object struct {
// 	Name    string
// 	Bucket  string
// 	Created time.Time
// }
//
// // Objects is a list of Object structures
// type Objects []Object
//
// // after filters object, select only objects which after t
// func (objs Objects) after(t time.Time) (res Objects) {
// 	for _, o := range objs {
// 		if o.Created.After(t) {
// 			res = append(res, o)
// 		}
// 	}
//
// 	return res
// }
