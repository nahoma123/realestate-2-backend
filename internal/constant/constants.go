package constant

type Context string

const (
	Active   = "ACTIVE"
	Inactive = "INACTIVE"
)

const (
	AdminRole       = "ADMIN_ROLE"
	RegularUserRole = "REGULAR_USER_ROLE"
)

const (
	DbValuation         = "valuations"
	DbRealEstate        = "real_estates"
	DbProperties        = "properties"
	DbInspectionResults = "inspection_results"
)

const (
	ActiveRealEstateStatus   = "Active"
	CanceledRealEstateStatus = "Cancelled"
)

const (
	ActivePropertyStatus   = "Active"
	CanceledPropertyStatus = "Cancelled"
)
