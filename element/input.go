package element

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

type InputType string

const (
	ButtonInputType        InputType = "button"         //	Defines a clickable button (mostly used with a JavaScript to activate a script)
	CheckboxInputType      InputType = "checkbox"       //	Defines a checkbox
	ColorInputType         InputType = "color"          //	Defines a color picker
	DateInputType          InputType = "date"           //	Defines a date control (year, month, day (no time))
	DateTimeLocalInputType InputType = "datetime-local" //	Defines a date and time control (year, month, day, time (no timezone)
	EmailInputType         InputType = "email"          //	Defines a field for an e-mail address
	FileInputType          InputType = "file"           //	Defines a file-select field and a "Browse" button (for file uploads)
	HiddenInputType        InputType = "hidden"         //	Defines a hidden input field
	ImageInputType         InputType = "image"          //	Defines an image as the submit button
	MonthInputType         InputType = "month"          //	Defines a month and year control (no timezone)
	NumberInputType        InputType = "number"         //	Defines a field for entering a number
	PasswordInputType      InputType = "password"       //	Defines a password field
	RadioInputType         InputType = "radio"          //	Defines a radio button
	RangeInputType         InputType = "range"          //	Defines a range control (like a slider control)
	ResetInputType         InputType = "reset"          //	Defines a reset button
	SearchInputType        InputType = "search"         //	Defines a text field for entering a search string
	SubmitInputType        InputType = "submit"         //	Defines a submit button
	TelInputType           InputType = "tel"            //	Defines a field for entering a telephone number
	TextInputType          InputType = "text"           //	Default. Defines a single-line text field
	TimeInputType          InputType = "time"           //	Defines a control for entering a time (no timezone)
	UrlInputType           InputType = "url"            //	Defines a field for entering a URL
	WeekInputType          InputType = "week"           //	Defines a week and year control (no timezone)
)

type input struct {
	el        *HtmlElement
	inputType InputType
}

type isInputAccept interface {
	IsInputAccept()
}

type InputFileExtension string

func (InputFileExtension) IsInputAccept() {}

type InputAccept string

func (InputAccept) IsInputAccept() {}

const (
	InputAudio InputAccept = "audio/*"
	InputVideo InputAccept = "video/*"
	InputImage InputAccept = "image/*"
)

type InputMediaType string

func (InputMediaType) IsInputAccept() {}

func (i *input) Accept(file isInputAccept) *input {
	if i.inputType != FileInputType {
		return i
	}
	return i
}

func (i *input) Alt(text string) *input {
	if i.inputType != ImageInputType {
		return i
	}
	i.el.AddAttribute("alt", text)
	return i
}

type InputAutoCompleteValue string

const (
	InputAutoCompleteOn                       InputAutoCompleteValue = "on"                   //	Default. Autocomplete is on (enabled)
	InputAutoCompleteOff                      InputAutoCompleteValue = "off"                  //	Autocomplete is off (disabled)
	InputAutoCompleteAddressLine1             InputAutoCompleteValue = "address-line1"        //Expects the first line of the street address
	InputAutoCompleteAddressLine2             InputAutoCompleteValue = "address-line2"        // Expects the second line of the street address
	InputAutoCompleteAddressLine3             InputAutoCompleteValue = "address-line3"        // Expects the third line of the street address
	InputAutoCompleteAddressLevel1            InputAutoCompleteValue = "address-level1"       // Expects the first level of the address, e.g. the county
	InputAutoCompleteAddressLevel2            InputAutoCompleteValue = "address-level2"       // Expects the second level of the address, e.g. the city
	InputAutoCompleteAddressLevel3            InputAutoCompleteValue = "address-level3"       // Expects the third level of the address
	InputAutoCompleteAddressLevel4            InputAutoCompleteValue = "address-level4"       // Expects the fourth level of the address
	InputAutoCompleteStreetAddress            InputAutoCompleteValue = "street-address"       // Expects the full street address
	InputAutoCompleteCountry                  InputAutoCompleteValue = "country"              // Expects the country code
	InputAutoCompleteCountryName              InputAutoCompleteValue = "country-name"         // Expects the country name
	InputAutoCompletePostalCode               InputAutoCompleteValue = "postal-code"          // Expects the post code
	InputAutoCompleteName                     InputAutoCompleteValue = "name"                 // Expects the full name
	InputAutoCompleteAdditionalName           InputAutoCompleteValue = "additional-name"      // Expects the middle name
	InputAutoCompleteFamilyName               InputAutoCompleteValue = "family-name"          // Expects the last name
	InputAutoCompleteGivenName                InputAutoCompleteValue = "given-name"           // Expects the first name
	InputAutoCompleteHonoricPrefix            InputAutoCompleteValue = "honoric-prefix"       // Expects the title, like "Mr", "Ms" etc.
	InputAutoCompleteHonoricSuffix            InputAutoCompleteValue = "honoric-suffix"       // Expects the suffix, like "5", "Jr." etc.
	InputAutoCompleteNickname                 InputAutoCompleteValue = "nickname"             // Expects the nickname
	InputAutoCompleteOrganizationTitle        InputAutoCompleteValue = "organization-title"   // Expects the job title
	InputAutoCompleteUsername                 InputAutoCompleteValue = "username"             // Expects the username
	InputAutoCompleteNewPassword              InputAutoCompleteValue = "new-password"         // Expects a new password
	InputAutoCompleteCurrentPassword          InputAutoCompleteValue = "current-password"     // Expects the current password
	InputAutoCompleteBirthday                 InputAutoCompleteValue = "bday"                 // Expects the full birthday date
	InputAutoCompleteBirthdayDay              InputAutoCompleteValue = "bday-day"             // Expects the day of the birthday date
	InputAutoCompleteBirthdayMonth            InputAutoCompleteValue = "bday-month"           // Expects the month of the birthday date
	InputAutoCompleteBirthdayYear             InputAutoCompleteValue = "bday-year"            // Expects the year of the birthday date
	InputAutoCompleteSex                      InputAutoCompleteValue = "sex"                  // Expects the gender
	InputAutoCompleteOneTimeCode              InputAutoCompleteValue = "one-time-code"        // Expects a one time code for verification etc.
	InputAutoCompleteOrganization             InputAutoCompleteValue = "organization"         // Expects the company name
	InputAutoCompleteCreditCardName           InputAutoCompleteValue = "cc-name"              // Expects the credit card owner's full name
	InputAutoCompleteCreditCardGivenName      InputAutoCompleteValue = "cc-given-name"        // Expects the credit card owner's first name
	InputAutoCompleteCreditCardAdditionalName InputAutoCompleteValue = "cc-additional-name"   // Expects the credit card owner's middle name
	InputAutoCompleteCreditCardFamilyName     InputAutoCompleteValue = "cc-family-name"       // Expects the credit card owner's family name
	InputAutoCompleteCreditCardNumber         InputAutoCompleteValue = "cc-number"            // Expects the credit card's number
	InputAutoCompleteCreditCardExpired        InputAutoCompleteValue = "cc-exp"               // Expects the credit card's expiration date
	InputAutoCompleteCreditCardExpiredMonth   InputAutoCompleteValue = "cc-exp-month"         // Expects the credit card's expiration month
	InputAutoCompleteCreditCardExpiredYear    InputAutoCompleteValue = "cc-exp-year"          // Expects the credit card's expiration year
	InputAutoCompleteCreditCardCvc            InputAutoCompleteValue = "cc-csc"               // Expects the CVC code
	InputAutoCompleteCreditCardType           InputAutoCompleteValue = "cc-type"              // Expects the credit card's type of payment
	InputAutoCompleteTransactionCurrency      InputAutoCompleteValue = "transaction-currency" // Expects the currency
	InputAutoCompleteTransactionAmount        InputAutoCompleteValue = "transaction-amount"   // Expects a number, the amount
	InputAutoCompleteLanguage                 InputAutoCompleteValue = "language"             // Expects the preferred language
	InputAutoCompleteUrl                      InputAutoCompleteValue = "url"                  // Expects a URL for home page or company website address
	InputAutoCompleteEmail                    InputAutoCompleteValue = "email"                // Expects the email address
	InputAutoCompletePhoto                    InputAutoCompleteValue = "photo"                // Expects an image
	InputAutoCompleteTelephone                InputAutoCompleteValue = "tel"                  // Expects the full phone number
	InputAutoCompleteTelephoneCountryCode     InputAutoCompleteValue = "tel-country-code"     // Expects the country code of the phone number
	InputAutoCompleteTelephoneNational        InputAutoCompleteValue = "tel-national"         // Expects the phone number with no country code
	InputAutoCompleteTelephoneAreaCode        InputAutoCompleteValue = "tel-area-code"        // Expects the area code of the phone number
	InputAutoCompleteTelephoneLocal           InputAutoCompleteValue = "tel-local"            // Expects the phone number with no country code and no area code
	InputAutoCompleteTelephoneLocalPrefix     InputAutoCompleteValue = "tel-local-prefix"     // Expects the local prefix of the phone number
	InputAutoCompleteTelephoneLocalSuffix     InputAutoCompleteValue = "tel-local-suffix"     // Expects the local suffix of the phone number
	InputAutoCompleteTelephoneExtension       InputAutoCompleteValue = "tel-extension"        // Expects the extension code of the phone number
	InputAutoCompleteInstantMessagingProtocol InputAutoCompleteValue = "impp"                 // Expects the url of an instant messaging protocol endpoint
)

func (i *input) AutoComplete(value InputAutoCompleteValue) *input {
	i.el.AddAttribute("autocomplete", string(value))
	return i
}

func (i *input) AutoFocus() *input {
	i.el.AddAttributeBool("autofocus")
	return i
}

func (i *input) Checked() *input {
	if i.inputType != CheckboxInputType && i.inputType != RadioInputType {
		return i
	}
	i.el.AddAttributeBool("checked")
	return i
}

func (i *input) Dirname(fieldName string) *input {
	i.el.AddAttribute("dirname", strings.Join([]string{fieldName, "dir"}, "."))
	return i
}

func (i *input) Disabled() *input {
	i.el.AddAttributeBool("disabled")
	return i
}

func (i *input) Form(formName string) *input {
	i.el.AddAttribute("form", formName)
	return i
}

func (i *input) FormAction(url string) *input {
	if i.inputType != SubmitInputType && i.inputType != ImageInputType {
		return i
	}
	i.el.AddAttribute("formaction", url)
	return i
}

func (i *input) FormEncType(enctype EncType) *input {
	if i.inputType != SubmitInputType && i.inputType != ImageInputType {
		return i
	}

	i.el.AddAttribute("formenctype", string(enctype))
	return i
}

func (i *input) FormMethod(method FormMethod) *input {
	if i.inputType != SubmitInputType && i.inputType != ImageInputType {
		return i
	}
	i.el.AddAttribute("formmethod", string(method))
	return i
}

func (i *input) FormNoValidate() *input {
	if i.inputType != SubmitInputType {
		return i
	}
	i.el.AddAttributeBool("formnovalidate")
	return i
}

func (i *input) FormTarget(target FormTarget) *input {
	if i.inputType != SubmitInputType && i.inputType != ImageInputType {
		return i
	}
	i.el.AddAttribute("formtarget", string(target))
	return i
}

func (i *input) Height(value int) *input {
	if i.inputType != ImageInputType {
		return i
	}
	i.el.AddAttribute("height", strconv.Itoa(value))
	return i
}

func (i *input) List(datalistId string) *input {
	i.el.AddAttribute("list", datalistId)
	return i
}

type FormMinMaxValue int

func (i FormMinMaxValue) String() string {
	return strconv.Itoa(int(i))
}

func (i *input) Max(value fmt.Stringer) *input {
	if i.inputType != NumberInputType &&
		i.inputType != RangeInputType &&
		i.inputType != DateInputType &&
		i.inputType != MonthInputType &&
		i.inputType != WeekInputType &&
		i.inputType != TimeInputType &&
		i.inputType != DateTimeLocalInputType {
		return i
	}

	i.el.AddAttribute("max", value.String())
	return i
}

func (i *input) Maxlength(value int) *input {
	i.el.AddAttribute("maxlength", strconv.Itoa(value))
	return i
}

func (i *input) Multiple() *input {
	i.el.AddAttributeBool("multiple")
	return i
}

func (i *input) Name(name string) *input {
	i.el.AddAttribute("name", name)
	return i
}

func (i *input) Pattern(regexp string) *input {
	i.el.AddAttribute("pattern", regexp)
	return i
}

func (i *input) Placeholder(text string) *input {
	i.el.AddAttribute("placeholder", text)
	return i
}

func (i *input) PopoverTarget(elementId string) *input {
	i.el.AddAttribute("popovertarget", elementId)
	return i
}

type PopoverTargetAction string

const (
	PopoverTargetActionHide   PopoverTargetAction = "hide"
	PopoverTargetActionShow   PopoverTargetAction = "show"
	PopoverTargetActionToggle PopoverTargetAction = "toggle"
)

func (i *input) PopoverTargetAction(action PopoverTargetAction) *input {
	if i.inputType != ButtonInputType {
		return i
	}
	i.el.AddAttribute("popovertargetaction", string(action))
	return i
}

func (i *input) ReadOnly() *input {
	i.el.AddAttributeBool("readonly")
	return i
}

func (i *input) Required() *input {
	i.el.AddAttributeBool("required")
	return i
}

func (i *input) Size(size int) *input {
	i.el.AddAttribute("size", strconv.Itoa(size))
	return i
}

func (i *input) Src(url string) *input {
	if i.inputType != ImageInputType {
		return i
	}
	i.el.AddAttribute("src", url)
	return i
}

func (i *input) Step(step int) *input {
	if i.inputType != NumberInputType && i.inputType != RangeInputType && i.inputType != DateInputType &&
		i.inputType != MonthInputType && i.inputType != WeekInputType && i.inputType != TimeInputType &&
		i.inputType != DateTimeLocalInputType {
		return i
	}
	i.el.AddAttribute("step", strconv.Itoa(step))
	return i
}

func (i *input) Type(inputType InputType) *input {
	i.inputType = inputType
	i.el.AddAttribute("type", string(inputType))
	return i
}

func (i *input) Value(value string) *input {
	i.el.AddAttribute("value", value)
	return i
}

func (i *input) Width(value int) *input {
	if i.inputType != ImageInputType {
		return i
	}
	i.el.AddAttribute("width", strconv.Itoa(value))
	return i
}

func (i *input) Element() *HtmlElement { return i.el }
func (i *input) Render() string        { return i.el.Render() }

func Input() *input {
	el := &HtmlElement{
		node: &html.Node{
			Type: html.ElementNode,
			Data: "input",
		},
	}

	i := &input{el: el}

	return i
}
