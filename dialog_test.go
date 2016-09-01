package dialog

import (
	"fmt"
	// "github.com/weldpua2008/go-dialog"
	//"strconv"
	//"time"
	// "strings"
	// "reflect"
	// "github.com/stretchr/testify/mock"
	"strconv"
	"testing"
	"time"
)

// type MyMockedObject struct {
// 	mock.Mock
// }

// func (m *MyMockedObject) exec(dType string, allowLabel bool) (string, error) {
// 	// return "", fmt.Errorf(DIALOG_ERR_CANCEL)
// 	args := m.Called(dType, allowLabel)
// 	return args.String(0), args.Error(1)

//
var exec_current_error error
var exec_string_error string

type MyDialog struct {
	Dialog
}

func (d *MyDialog) exec(dType string, allowLabel bool) (string, error) {
	// fmt.Sprintf("dType:" + dType)
	return exec_string_error, exec_current_error
}

func NewTestDialog(environment string, parentId int) DialogIface {
	var res = new(Dialog)
	LastCMD = []string{}
	return res
}
func NewTestDialogExec(environment string, parentId int) DialogIface {
	var res = new(MyDialog)
	res.environment = DIALOG_TEST_ENV
	LastCMD = []string{}
	return res
}

func NewTestDialogRAW(environment string, parentId int) Dialog {
	var res = new(Dialog)
	LastCMD = []string{}
	// f := func(dType string, allowLabel bool) (string, error) {
	// 	return "", fmt.Errorf(DIALOG_ERR_CANCEL)
	// }
	// res.exec = f
	return *res
}

func tearDown() {
	LastCMD = []string{}
}

func Test_exec(t *testing.T) {
	var res = new(MyDialog)
	exec_current_error = fmt.Errorf(DIALOG_ERR_CANCEL)
	_, x := res.exec("dType", true)
	if x != exec_current_error {
		t.Errorf("Expected %v, actual %v ", exec_current_error, x)
	}

	exec_current_error = nil
	_, x1 := res.exec("dType", true)
	if x1 != exec_current_error {
		t.Errorf("Expected %v, actual %v ", exec_current_error, x)
	}
}

func TestYesNo(t *testing.T) {
	d := NewTestDialogExec(DIALOG_TEST_ENV, 0)
	exec_current_error = nil
	actual_bool := d.Yesno()
	expected_bool := true
	if expected_bool != actual_bool {
		t.Errorf("Expected %v, actual %v ", expected_bool, actual_bool)
	}
	x := LastCMD
	expected_str := "[" + DIALOG_TEST_ENV + " --no-shadow --yesno  0 0 --attach 0]"
	if fmt.Sprintf("%v", x) != expected_str {
		t.Errorf("Expected %v, actual %v ", expected_str, x)
	}

	var res = new(MyDialog)
	exec_current_error = fmt.Errorf(DIALOG_ERR_CANCEL)
	expected_bool = false
	res.environment = DIALOG_TEST_ENV
	res.exec_output = ""
	res.exec_error = fmt.Errorf(DIALOG_ERR_CANCEL)
	actual_bool1 := res.Yesno()

	if expected_bool != actual_bool1 {
		t.Errorf("Expected2 %v, actual %v ", expected_bool, actual_bool1)
	}
}

func TestSlider(t *testing.T) {
	var res = new(MyDialog)
	res.reset()
	exec_current_error = fmt.Errorf(DIALOG_ERR_CANCEL)
	expected := 1
	res.environment = DIALOG_TEST_ENV
	res.exec_output = "1"
	res.exec_error = nil //fmt.Errorf(DIALOG_ERR_CANCEL)
	min := 8
	max := 10
	step := 2
	actual_bool, err := res.Slider(min, max, step)

	if expected != actual_bool {
		t.Errorf("Expected %v, actual %v ", expected, actual_bool)
	}

	if err != nil {
		t.Errorf("Expected nil actual %v", err)

	}

	expected = 0
	res.exec_output = "2 2 2 23 "

	res.exec_error = fmt.Errorf("error d.exec(slider, true)")
	expected_err := fmt.Errorf("strconv.ParseInt: parsing \"" + res.exec_output + "\": invalid syntax")

	actual_bool, err = res.Slider(min, max, step)

	if expected != actual_bool {
		t.Errorf("Expected2 %v, actual %v ", expected, actual_bool)
	}

	if fmt.Sprintf("%v", err) != fmt.Sprintf("%v", expected_err) {
		t.Errorf("Expected %v actual '%v'", expected_err, err)
	}

}

func TestPassivepopup(t *testing.T) {
	var res = new(MyDialog)
	res.reset()
	exec_current_error = fmt.Errorf(DIALOG_ERR_CANCEL)

	var passivepopupTests = []struct {
		text    string
		timeout int
	}{
		{"[w]", 3},
		{"[ string2]", 10},
		{"[ string2  22 %$ ]", 10},
	}
	res.environment = DIALOG_TEST_ENV
	for _, tt := range passivepopupTests {
		res.Passivepopup(tt.text, tt.timeout)
		expected_str := "[" + DIALOG_TEST_ENV + " --ok-label OK --no-shadow --passivepopup 0 0 " + tt.text + " " + strconv.Itoa(tt.timeout) + " --attach 0]"
		if fmt.Sprintf("%v", LastCMD) != expected_str {
			t.Errorf("Expected %v, actual %v ", expected_str, LastCMD)
		}
	}
}

// func TestCombobox(t *testing.T) {
// 	var res = new(MyDialog)
// 	res.reset()
// 	exec_current_error = fmt.Errorf(DIALOG_ERR_CANCEL)

// 	var passivepopupTests = []struct {
// 		text    string
// 		timeout int
// 	}{
// 		{"[w]", 3},
// 		{"[ string2]", 10},
// 		{"[ string2  22 %$ ]", 10},
// 	}
// 	res.environment = DIALOG_TEST_ENV
// 	for _, tt := range passivepopupTests {
// 		res.Combobox(tt.text)
// 		expected_str := "[" + DIALOG_TEST_ENV + " --ok-label OK --no-shadow --passivepopup 0 0 " + tt.text + " " + strconv.Itoa(tt.timeout) + " --attach 0]"
// 		if fmt.Sprintf("%v", LastCMD) != expected_str {
// 			t.Errorf("Expected %v, actual %v ", expected_str, LastCMD)
// 		}
// 	}
// }

func TestGeticon(t *testing.T) {
	var res = new(MyDialog)
	res.reset()
	exec_current_error = fmt.Errorf(DIALOG_ERR_CANCEL)

	var geticonTests = []struct {
		text string
		err  error
	}{
		{"[tex1t]", nil},
		{"[text]", fmt.Errorf("xerrorx")},
		{"", nil},
	}
	res.environment = DIALOG_TEST_ENV
	for _, tt := range geticonTests {

		res.exec_output = tt.text
		res.exec_error = tt.err
		val_Getcolor := res.Geticon()
		expected_str := tt.text
		if val_Getcolor != expected_str {
			t.Errorf("Expected %v, actual '%v' error %v", expected_str, val_Getcolor, tt.err)

		}
	}
}

func TestGetcolor(t *testing.T) {
	var res = new(MyDialog)
	res.reset()
	exec_current_error = fmt.Errorf(DIALOG_ERR_CANCEL)

	var getcolorTests = []struct {
		text string
		err  error
	}{
		{"[tex1t]", nil},
		{"[text]", fmt.Errorf("xerrorx")},
		{"", nil},
	}
	res.environment = DIALOG_TEST_ENV
	for _, tt := range getcolorTests {

		res.exec_output = tt.text
		res.exec_error = tt.err
		val_Getcolor := res.Getcolor()
		expected_str := tt.text
		if val_Getcolor != expected_str {
			t.Errorf("Expected %v, actual '%v' error %v", expected_str, val_Getcolor, tt.err)

		}
	}
}

func TestCalendar(t *testing.T) {
	var res = new(MyDialog)
	res.reset()
	exec_current_error = fmt.Errorf(DIALOG_ERR_CANCEL)
	var date time.Time

	lastCmd := fmt.Sprintf("[test_env --ok-label OK --no-shadow --calendar  0 0 %v %v %v --attach 0]", date.Format(DIALOG_CALENDAR_START_YEAR), date.Format(DIALOG_CALENDAR_START_MONTH), date.Format(DIALOG_CALENDAR_START_DAY))

	var getCalendarTests = []struct {
		text    string
		err     error
		lastCmd string
	}{
		{"[tex1t]", nil, lastCmd},
		{"[text]", fmt.Errorf("xerrorx"), lastCmd},
		{"", nil, lastCmd},
	}
	var ti time.Time
	res.environment = DIALOG_TEST_ENV
	for _, tt := range getCalendarTests {

		res.exec_output = tt.text
		res.exec_error = tt.err
		val_Calendar, _ := res.Calendar(ti)
		// t.Logf("%v %v", res.lastCmd, lastCmd)
		if fmt.Sprintf("%v", res.lastCmd) != fmt.Sprintf("%v", lastCmd) {
			t.Errorf("Expected res.lastCmd %v, actual '%v' ", fmt.Sprintf("%v", res.lastCmd), fmt.Sprintf("%v", lastCmd))

		}
		expected_str := tt.text
		if val_Calendar != expected_str {
			t.Errorf("Expected %v, actual '%v' error %v", expected_str, val_Calendar, tt.err)

		}
	}
}

func TestChecklist(t *testing.T) {
	var res = new(MyDialog)
	res.reset()
	exec_current_error = fmt.Errorf(DIALOG_ERR_CANCEL)
	lastCmd := fmt.Sprintf("[test_env --ok-label OK --no-shadow --checklist  0 0 0 tag1 tag2 --attach 0]")

	var getCalendarTests = []struct {
		text    string
		err     error
		lastCmd string
	}{
		{"[[tex1t]]", nil, lastCmd},
		{"[text]", fmt.Errorf("xerrorx"), lastCmd},
		{"", nil, lastCmd},
	}
	res.environment = DIALOG_TEST_ENV
	for _, tt := range getCalendarTests {

		res.exec_output = tt.text
		res.exec_error = tt.err
		val_Calendar, _ := res.Checklist(0, "tag1", "tag2")
		// t.Logf("%v %v", res.lastCmd, lastCmd)
		if fmt.Sprintf("%v", res.lastCmd) != fmt.Sprintf("%v", lastCmd) {
			t.Errorf("Expected res.lastCmd %v, actual '%v' ", fmt.Sprintf("%v", res.lastCmd), fmt.Sprintf("%v", lastCmd))

		}
		expected_str := tt.text
		if fmt.Sprintf("%v", val_Calendar) != fmt.Sprintf("[%v]", expected_str) {
			t.Errorf("Expected %v, actual '%v' error %v", fmt.Sprintf("[%v]", expected_str), val_Calendar, tt.err)

		}
	}
}

func TestMixedform(t *testing.T) {
	var res = new(MyDialog)
	res.reset()
	exec_current_error = fmt.Errorf(DIALOG_ERR_CANCEL)
	lastCmdF := fmt.Sprintf("[test_env --ok-label OK --no-shadow --mixedform Title 0 0 0 Selection1 1 1 2046 1 10 20 0 0 Selection2 2 1 0 2 10 20 0 0 Selection3 3 1 0 3 10 20 0 0 --attach 0]")
	lastCmdT := fmt.Sprintf("[test_env --ok-label OK --no-shadow --insecure --mixedform Title 0 0 0 Selection1 1 1 2046 1 10 20 0 0 Selection2 2 1 0 2 10 20 0 0 Selection3 3 1 0 3 10 20 0 0 --attach 0]")
	var lastCmd string
	var getCalendarTests = []struct {
		text    string
		err     error
		lastCmd string
	}{
		{"[[tex1t]]", nil, lastCmd},
		{"[text]", fmt.Errorf("xerrorx"), lastCmd},
		{"", nil, lastCmd},
	}
	l := []string{"Selection1", "1", "1", "2046", "1", "10", "20", "0", "0", "Selection2", "2", "1", "0", "2", "10", "20", "0", "0", "Selection3", "3", "1", "0", "3", "10", "20", "0", "0"}

	res.environment = DIALOG_TEST_ENV
	for _, is_insecure := range []bool{true, false} {

		for _, tt := range getCalendarTests {
			lastCmd = lastCmdF
			if is_insecure {
				lastCmd = lastCmdT
			}
			res.exec_output = tt.text
			res.exec_error = tt.err
			val_Calendar, _ := res.Mixedform("Title", is_insecure, l[0:]...)
			if fmt.Sprintf("%v", res.lastCmd) != fmt.Sprintf("%v", lastCmd) {
				t.Errorf("Expected res.lastCmd %v, actual '%v' ", fmt.Sprintf("%v", res.lastCmd), fmt.Sprintf("%v", lastCmd))

			}
			expected_str := tt.text
			if fmt.Sprintf("%v", val_Calendar) != fmt.Sprintf("[%v]", expected_str) {
				t.Errorf("Expected %v, actual '%v' error %v", fmt.Sprintf("[%v]", expected_str), val_Calendar, tt.err)

			}
		}
	}
}

func TestFselect(t *testing.T) {
	var res = new(MyDialog)
	res.reset()
	exec_current_error = fmt.Errorf(DIALOG_ERR_CANCEL)
	lastCmd := fmt.Sprintf("[test_env --ok-label OK --no-shadow --fselect /tmp/test.txt 0 0 --attach 0]")

	var getCalendarTests = []struct {
		text    string
		err     error
		lastCmd string
	}{
		{"[[tex1t]]", nil, lastCmd},
		{"[text]", fmt.Errorf("xerrorx"), lastCmd},
		{"", nil, lastCmd},
	}

	res.environment = DIALOG_TEST_ENV
	for _, tt := range getCalendarTests {
		res.exec_output = tt.text
		res.exec_error = tt.err
		val_Calendar, _ := res.Fselect("/tmp/test.txt")
		if fmt.Sprintf("%v", res.lastCmd) != fmt.Sprintf("%v", lastCmd) {
			t.Errorf("Expected res.lastCmd %v, actual '%v' ", fmt.Sprintf("%v", res.lastCmd), fmt.Sprintf("%v", lastCmd))

		}
		expected_str := tt.text
		if fmt.Sprintf("%v", val_Calendar) != fmt.Sprintf("%v", expected_str) {
			t.Errorf("Expected %v, actual '%v' error %v", expected_str, val_Calendar, tt.err)

		}
	}
}

func TestInfobox(t *testing.T) {
	var res = new(MyDialog)
	res.reset()
	exec_current_error = fmt.Errorf(DIALOG_ERR_CANCEL)
	lastCmd := fmt.Sprintf("[test_env --ok-label OK --no-shadow --infobox /tmp/test.txt 0 0 --attach 0]")
	var getCalendarTests = []struct {
		text    string
		err     error
		lastCmd string
	}{
		{"[[tex1t]]", nil, lastCmd},
		{"[text]", fmt.Errorf("xerrorx"), lastCmd},
		{"", nil, lastCmd},
	}

	res.environment = DIALOG_TEST_ENV
	for _, tt := range getCalendarTests {
		res.exec_output = tt.text
		res.exec_error = tt.err
		res.Infobox("/tmp/test.txt")
		if fmt.Sprintf("%v", res.lastCmd) != fmt.Sprintf("%v", lastCmd) {
			t.Errorf("Expected res.lastCmd %v, actual '%v' ", fmt.Sprintf("%v", res.lastCmd), fmt.Sprintf("%v", lastCmd))

		}
	}
}

func TestInputbox(t *testing.T) {
	var res = new(MyDialog)
	res.reset()
	exec_current_error = fmt.Errorf(DIALOG_ERR_CANCEL)
	lastCmd := fmt.Sprintf("[test_env --ok-label OK --no-shadow --inputbox  0 0 /tmp/test.txt --attach 0]")

	var getCalendarTests = []struct {
		text    string
		err     error
		lastCmd string
	}{
		{"[[tex1t]]", nil, lastCmd},
		{"[text]", fmt.Errorf("xerrorx"), lastCmd},
		{"", nil, lastCmd},
	}

	res.environment = DIALOG_TEST_ENV
	for _, tt := range getCalendarTests {
		res.exec_output = tt.text
		res.exec_error = tt.err
		val_Calendar, _ := res.Inputbox("/tmp/test.txt")
		if fmt.Sprintf("%v", res.lastCmd) != fmt.Sprintf("%v", lastCmd) {
			t.Errorf("Expected res.lastCmd %v, actual '%v' ", fmt.Sprintf("%v", res.lastCmd), fmt.Sprintf("%v", lastCmd))

		}
		expected_str := tt.text
		if fmt.Sprintf("%v", val_Calendar) != fmt.Sprintf("%v", expected_str) {
			t.Errorf("Expected %v, actual '%v' error %v", expected_str, val_Calendar, tt.err)
		}
	}
}

func TestInputmenu(t *testing.T) {
	var res = new(MyDialog)
	res.reset()
	exec_current_error = fmt.Errorf(DIALOG_ERR_CANCEL)
	lastCmd := fmt.Sprintf("[test_env --ok-label OK --no-shadow --inputmenu  0 0 10 Tag 1 Item 1 --attach 0]")

	var getCalendarTests = []struct {
		text    string
		err     error
		lastCmd string
	}{
		{"[[tex1t]]", nil, lastCmd},
		{"[text]", fmt.Errorf("xerrorx"), lastCmd},
		{"", nil, lastCmd},
	}

	res.environment = DIALOG_TEST_ENV
	for _, tt := range getCalendarTests {
		res.exec_output = tt.text
		res.exec_error = tt.err
		val_Calendar, _ := res.Inputmenu(10, "Tag 1", "Item 1")
		if fmt.Sprintf("%v", res.lastCmd) != fmt.Sprintf("%v", lastCmd) {
			t.Errorf("Expected res.lastCmd %v, actual '%v' ", fmt.Sprintf("%v", res.lastCmd), fmt.Sprintf("%v", lastCmd))

		}
		expected_str := tt.text
		if fmt.Sprintf("%v", val_Calendar) != fmt.Sprintf("[%v]", expected_str) {
			t.Errorf("Expected [%v], actual '%v' error %v", expected_str, val_Calendar, tt.err)
		}
	}
}

func TestMenu(t *testing.T) {
	var res = new(MyDialog)
	res.reset()
	exec_current_error = fmt.Errorf(DIALOG_ERR_CANCEL)
	lastCmd := fmt.Sprintf("[test_env --ok-label OK --no-shadow --menu  0 0 10 Tag 1 Item 1 --attach 0]")

	var getCalendarTests = []struct {
		text    string
		err     error
		lastCmd string
	}{
		{"[[tex1t]]", nil, lastCmd},
		{"[text]", fmt.Errorf("xerrorx"), lastCmd},
		{"", nil, lastCmd},
	}

	res.environment = DIALOG_TEST_ENV
	for _, tt := range getCalendarTests {
		res.exec_output = tt.text
		res.exec_error = tt.err
		val_Calendar, _ := res.Menu(10, "Tag 1", "Item 1")
		if fmt.Sprintf("%v", res.lastCmd) != fmt.Sprintf("%v", lastCmd) {
			t.Errorf("Expected res.lastCmd %v, actual '%v' ", fmt.Sprintf("%v", res.lastCmd), fmt.Sprintf("%v", lastCmd))

		}
		expected_str := tt.text
		if fmt.Sprintf("%v", val_Calendar) != fmt.Sprintf("%v", expected_str) {
			t.Errorf("Expected %v, actual '%v' error %v", expected_str, val_Calendar, tt.err)
		}
	}
}

func TestMsgbox(t *testing.T) {
	var res = new(MyDialog)
	res.reset()
	exec_current_error = fmt.Errorf(DIALOG_ERR_CANCEL)
	lastCmd := fmt.Sprintf("[test_env --ok-label OK --no-shadow --msgbox /tmp/test.txt 0 0 --attach 0]")
	var getCalendarTests = []struct {
		text    string
		err     error
		lastCmd string
	}{
		{"[[tex1t]]", nil, lastCmd},
		{"[text]", fmt.Errorf("xerrorx"), lastCmd},
		{"", nil, lastCmd},
	}

	res.environment = DIALOG_TEST_ENV
	for _, tt := range getCalendarTests {
		res.exec_output = tt.text
		res.exec_error = tt.err
		res.Msgbox("/tmp/test.txt")
		if fmt.Sprintf("%v", res.lastCmd) != fmt.Sprintf("%v", lastCmd) {
			t.Errorf("Expected res.lastCmd %v, actual '%v' ", fmt.Sprintf("%v", res.lastCmd), fmt.Sprintf("%v", lastCmd))

		}
	}
}

func TestPasswordbox(t *testing.T) {
	var res = new(MyDialog)
	res.reset()
	exec_current_error = fmt.Errorf(DIALOG_ERR_CANCEL)
	lastCmdT := fmt.Sprintf("[test_env --ok-label OK --no-shadow --insecure --passwordbox  0 0  --attach 0]")
	lastCmdF := fmt.Sprintf("[test_env --ok-label OK --no-shadow --passwordbox  0 0  --attach 0]")
	var lastCmd string
	var getCalendarTests = []struct {
		text    string
		err     error
		lastCmd string
	}{
		{"[[tex1t]]", nil, lastCmd},
		{"[text]", fmt.Errorf("xerrorx"), lastCmd},
		{"", nil, lastCmd},
	}

	res.environment = DIALOG_TEST_ENV
	for _, is_insecure := range []bool{true, false} {
		for _, tt := range getCalendarTests {
			res.exec_output = tt.text
			res.exec_error = tt.err
			lastCmd = lastCmdF
			if is_insecure {
				lastCmd = lastCmdT
			}
			val_Calendar, _ := res.Passwordbox(is_insecure)
			if fmt.Sprintf("%v", res.lastCmd) != fmt.Sprintf("%v", lastCmd) {
				t.Errorf("Expected res.lastCmd %v, actual '%v' ", fmt.Sprintf("%v", res.lastCmd), fmt.Sprintf("%v", lastCmd))

			}
			expected_str := tt.text
			if fmt.Sprintf("%v", val_Calendar) != fmt.Sprintf("%v", expected_str) {
				t.Errorf("Expected %v, actual '%v' error %v", expected_str, val_Calendar, tt.err)
			}
		}
	}
}

// tests for structure changes
func TestHelpButtonTrue(t *testing.T) {
	d := NewTestDialogRAW(DIALOG_TEST_ENV, 0)
	expected_val := true
	d.HelpButton(expected_val)

	if d.helpButton != expected_val {
		t.Errorf("Expected %v, actual %v ", expected_val, d.helpButton)
	}
	// x := LastCMD
	// fmt.Sprintf("%v \n", LastCMD)
	// expected_str := "[]"
	// if fmt.Sprintf("%v", x) != expected_str {
	// 	t.Errorf("Expected %v, actual %v ", expected_str, x)
	// }
}

func TestHelpButton(t *testing.T) {
	d := NewTestDialogRAW(DIALOG_TEST_ENV, 0)

	var tests = []bool{true, false}
	for _, expected_val := range tests {

		d.HelpButton(expected_val)

		if d.helpButton != expected_val {
			t.Errorf("Expected %v, actual %v ", expected_val, d.helpButton)
		}
	}
}

func TestSetBackTitle(t *testing.T) {
	d := NewTestDialogRAW(DIALOG_TEST_ENV, 0)

	var tests = []string{"", "backtitle"}
	for _, expected_val := range tests {

		d.SetBackTitle(expected_val)

		if d.backtitle != expected_val {
			t.Errorf("Expected %v, actual %v ", expected_val, d.helpButton)
		}
	}
}

func TestSetTitle(t *testing.T) {
	d := NewTestDialogRAW(DIALOG_TEST_ENV, 0)

	var tests = []string{"", "backtitle"}
	for _, expected_val := range tests {

		d.SetTitle(expected_val)

		if d.title != expected_val {
			t.Errorf("Expected %v, actual %v ", expected_val, d.helpButton)
		}
	}
}

func TestSetLabel(t *testing.T) {
	d := NewTestDialogRAW(DIALOG_TEST_ENV, 0)

	var tests = []string{"", "backtitle"}
	for _, expected_val := range tests {

		d.SetLabel(expected_val)

		if d.label != expected_val {
			t.Errorf("Expected %v, actual %v ", expected_val, d.helpButton)
		}
	}
}

func TestSetOkLabel(t *testing.T) {
	d := NewTestDialogRAW(DIALOG_TEST_ENV, 0)

	var tests = []string{"", "backtitle"}
	for _, expected_val := range tests {

		d.SetOkLabel(expected_val)

		if d.okLabel != expected_val {
			t.Errorf("Expected %v, actual %v ", expected_val, d.helpButton)
		}
	}
}

func TestSetYesLabel(t *testing.T) {
	d := NewTestDialogRAW(DIALOG_TEST_ENV, 0)

	var tests = []string{"", "backtitle"}
	for _, expected_val := range tests {

		d.SetYesLabel(expected_val)

		if d.yesLabel != expected_val {
			t.Errorf("Expected %v, actual %v ", expected_val, d.yesLabel)
		}
	}
}

func TestSetNoLabel(t *testing.T) {
	d := NewTestDialogRAW(DIALOG_TEST_ENV, 0)

	var tests = []string{"", "backtitle"}
	for _, expected_val := range tests {

		d.SetNOLabel(expected_val)

		if d.noLabel != expected_val {
			t.Errorf("Expected %v, actual %v ", expected_val, d.noLabel)
		}
	}
}

func TestSetExtraLabel(t *testing.T) {
	d := NewTestDialogRAW(DIALOG_TEST_ENV, 0)

	var tests = []string{"", "backtitle"}
	for _, expected_val := range tests {

		d.SetExtraLabel(expected_val)

		if d.extraLabel != expected_val {
			t.Errorf("Expected %v, actual %v ", expected_val, d.helpButton)
		}
	}
}

func TestShadowFalse(t *testing.T) {
	d := NewTestDialogRAW(DIALOG_TEST_ENV, 0)

	expected_val := false
	d.Shadow(expected_val)

	if d.shadow != expected_val {
		t.Errorf("Expected %v, actual %v ", expected_val, d.helpButton)
	}
}

func TestLabel(t *testing.T) {
	d := NewTestDialogRAW(DIALOG_TEST_ENV, 0)
	expected_val := "label"
	d.SetHelpLabel(expected_val)

	if d.helpLabel != expected_val {
		t.Errorf("Expected %v, actual %v ", expected_val, d.helpButton)
	}
}

func TestSetCancelLabel(t *testing.T) {
	d := NewTestDialogRAW(DIALOG_TEST_ENV, 0)
	expected_val := "label"
	d.SetCancelLabel(expected_val)

	if d.cancelLabel != expected_val {
		t.Errorf("Expected %v, actual %v ", expected_val, d.helpButton)
	}
}
