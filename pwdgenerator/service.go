package pwdgenerator

//PwdIService PwdIService
type PwdIService interface {
	GetSHA256Hash(pwd string) (string, error)
}
