package mediaqueries

type MediaQuery string

const (
	Portrait  MediaQuery = "--portrait"
	Landscape MediaQuery = "--landscape"
)

const (
	XXS_Only MediaQuery = "--xxs-only"
	XS_Only  MediaQuery = "--xs-only"
	SM_Only  MediaQuery = "--sm-only"
	MD_Only  MediaQuery = "--md-only"
	LG_Only  MediaQuery = "--lg-only"
	Xl_Only  MediaQuery = "--xl-only"
	XXL_Only MediaQuery = "--xxl-only"
)

const (
	XXS_N_Above MediaQuery = "--xxs-n-above"
	XS_N_Above  MediaQuery = "--xs-n-above"
	SM_N_Above  MediaQuery = "--sm-n-above"
	MD_N_Above  MediaQuery = "--md-n-above"
	LG_N_Above  MediaQuery = "--lg-n-above"
	Xl_N_Above  MediaQuery = "--xl-n-above"
)

const (
	XXS_N_Below MediaQuery = "--xxs-n-below"
	XS_N_Below  MediaQuery = "--xs-n-below"
	SM_N_Below  MediaQuery = "--sm-n-below"
	MD_N_Below  MediaQuery = "--md-n-below"
	LG_N_Below  MediaQuery = "--lg-n-below"
	Xl_N_Below  MediaQuery = "--xl-n-below"
)

const (
	XXS_Phone MediaQuery = "--xxs-phone"
	XS_Phone  MediaQuery = "--xs-phone"
	SM_Phone  MediaQuery = "--sm-phone"
	MD_Phone  MediaQuery = "--md-phone"
	LG_Phone  MediaQuery = "--lg-phone"
)
