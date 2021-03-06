package smn_analysis

func iserr(err error) bool {
	return err != nil
}

func noerr(err error) bool {
	return err == nil
}

const (
	ErrNoMatchStateNode                = "ErrNoMatchStateNode: err list: [%s]" //没有满足的
	ErrTooMuchMatchStateNode           = "ErrTooMuchMatchStateNode"            //太多满足条件的
	ErrTooMuchMatchStateNodeWhenHasEnd = "ErrTooMuchMatchStateNodeWhenHasEnd"  //太多满足条件的且有结束
)
