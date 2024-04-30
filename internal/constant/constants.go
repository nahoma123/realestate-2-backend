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
	DbValuation           = "valuations"
	DbRealEstate          = "real_estates"
	DbProperties          = "properties"
	DbContracts           = "contracts"
	DbInspectionResults   = "inspection_results"
	DbMaintenanceRequests = "maintenance_requests"
	DbCompliances         = "compliances"
)

const (
	ActiveRealEstateStatus   = "Active"
	CanceledRealEstateStatus = "Cancelled"
)

const (
	ActivePropertyStatus   = "Active"
	CanceledPropertyStatus = "Cancelled"
)
