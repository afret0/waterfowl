package err

func InfrErrTem() string {
	t := `
	package err

	import "errors"

	type Item struct {
		Code    int  
		Message string 
		Err     error  
	}

	type Manager struct {
		errs map[error]*Item
	}

	var m *Manager

	func getManager() *Manager {
		if m != nil {
		return m
	}
		m = new(Manager)
		m.errs = make(map[error]*Item, 0)

		return m
	}

	func NewErr(code int, msg string) error {
		m := getManager()
		err := errors.New(msg)
		m.errs[err] = &Item{
		Code:    code,
		Message: msg,
		Err:     err,
	}

		return err
	}

	func NewDefaultErr(msg string) error {
		m := getManager()
		err := errors.New(msg)
		m.errs[err] = &Item{
		Code:    0,
		Message: msg,
		Err:     err,
	}
		return err
	}

	func GetErrs(err error) *Item {
		return m.errs[err]
	}
`
	return t

}
