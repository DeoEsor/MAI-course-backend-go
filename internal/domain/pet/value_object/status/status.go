package status

type Status string

const (
	NONE                 Status = "NONE"
	OnMedicalExamination Status = "ON_MEDICAL_EXAMINATION"
	OnRecover            Status = "ON_RECOVER"
	WaitingOwner         Status = "WAITING_OWNER"
	Adopted              Status = "ADOPTED"
)
