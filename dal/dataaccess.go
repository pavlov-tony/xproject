package dal

// type service string
// const (
//     serviceAWS = service("aws")
//     serviceGCP = service("gcp")
// )
//
// type storage string
// const (
//     storageHDD = storage("hdd")
//     storageSSD = storage("ssd")
// )
//
// type term string
// const (
//     termHourly = term("hourly")
//     termWeekly = term("weekly")
//     termMonthly = term("monthly")
// )

type Instance struct {
	id              int
	provider        string
	typee           string
	core            float64
	ram             float64
	disk            float64
	disk_type       string
	price_per_month float64
	price_per_hour  float64
	lease_type      string
	location        string
}

type DataAccesser interface {
	getInstanceById(id string) (i *Instance)
}
