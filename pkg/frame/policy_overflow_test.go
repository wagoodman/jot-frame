package frame

import (
	"fmt"
	"sort"
	"testing"
)

var floatFreeDrawTestCases = map[string]drawTestParams{
	"FloatFree_goCase": {3, 0, 0, 10, PolicyOverflow, 40,
		[]ScreenEvent{
			// create the frame (pave a blank spot)
			{row: 10, value: []byte("")},
			{row: 11, value: []byte("")},
			{row: 12, value: []byte("")},
			// draw the first update
			{row: 10, value: []byte("LineIdx:0")},
			{row: 11, value: []byte("LineIdx:1")},
			{row: 12, value: []byte("LineIdx:2")},
		},
		[]string{},
	},
	"FloatFree_Header": {3, 1, 0, 10, PolicyOverflow, 40,
		[]ScreenEvent{
			// create the frame (pave a blank spot)
			{row: 10, value: []byte("")},
			{row: 11, value: []byte("")},
			{row: 12, value: []byte("")},
			{row: 13, value: []byte("")},
			// draw the first update
			{row: 10, value: []byte("theHeader")},
			{row: 11, value: []byte("LineIdx:0")},
			{row: 12, value: []byte("LineIdx:1")},
			{row: 13, value: []byte("LineIdx:2")},
		},
		[]string{},
	},
	"FloatFree_Footer": {3, 0, 1, 10, PolicyOverflow, 40,
		[]ScreenEvent{
			// create the frame (pave a blank spot)
			{row: 10, value: []byte("")},
			{row: 11, value: []byte("")},
			{row: 12, value: []byte("")},
			{row: 13, value: []byte("")},
			// draw the first update
			{row: 10, value: []byte("LineIdx:0")},
			{row: 11, value: []byte("LineIdx:1")},
			{row: 12, value: []byte("LineIdx:2")},
			{row: 13, value: []byte("theFooter")},
		},
		[]string{},
	},
	"FloatFree_HeaderFooter": {3, 1, 1, 10, PolicyOverflow, 40,
		[]ScreenEvent{
			// create the frame (pave a blank spot)
			{row: 10, value: []byte("")},
			{row: 11, value: []byte("")},
			{row: 12, value: []byte("")},
			{row: 13, value: []byte("")},
			{row: 14, value: []byte("")},
			// draw the first update
			{row: 10, value: []byte("theHeader")},
			{row: 11, value: []byte("LineIdx:0")},
			{row: 12, value: []byte("LineIdx:1")},
			{row: 13, value: []byte("LineIdx:2")},
			{row: 14, value: []byte("theFooter")},
		},
		[]string{},
	},
	"FloatFree_TermHeightSmall_AtTop": {3, 0, 0, 1, PolicyOverflow, 2,
		[]ScreenEvent{
			// create the frame (pave a blank spot)
			{row: 1, value: []byte("")},
			{row: 2, value: []byte("")},
			// draw the first update
			{row: 1, value: []byte("LineIdx:0")},
			{row: 2, value: []byte("LineIdx:1")},
		},
		[]string{
			"line is out of bounds (row=3)",
		},
	},
	"FloatFree_TermHeightSmall_AtTop_Header": {3, 1, 0, 1, PolicyOverflow, 2,
		[]ScreenEvent{
			// create the frame (pave a blank spot)
			{row: 1, value: []byte("")},
			{row: 2, value: []byte("")},
			// draw the first update
			{row: 1, value: []byte("theHeader")},
			{row: 2, value: []byte("LineIdx:0")},
		},
		[]string{
			"line is out of bounds (row=3)",
			"line is out of bounds (row=4)",
		},
	},
	"FloatFree_TermHeightSmall_AtTop_Footer": {3, 0, 1, 1, PolicyOverflow, 2,
		[]ScreenEvent{
			// create the frame (pave a blank spot)
			{row: 1, value: []byte("")},
			{row: 2, value: []byte("")},
			// draw the first update
			{row: 1, value: []byte("LineIdx:0")},
			{row: 2, value: []byte("LineIdx:1")},
		},
		[]string{
			"line is out of bounds (row=3)",
			"line is out of bounds (row=4)",
		},
	},
	"FloatFree_TermHeightSmall_AtTop_HeaderFooter": {3, 1, 1, 1, PolicyOverflow, 2,
		[]ScreenEvent{
			// create the frame (pave a blank spot)
			{row: 1, value: []byte("")},
			{row: 2, value: []byte("")},
			// draw the first update
			{row: 1, value: []byte("theHeader")},
			{row: 2, value: []byte("LineIdx:0")},
		},
		[]string{
			"line is out of bounds (row=3)",
			"line is out of bounds (row=4)",
			"line is out of bounds (row=5)",
		},
	},
	"FloatFree_TermHeightSmall_AtBottom": {3, 0, 0, 49, PolicyOverflow, 50,
		[]ScreenEvent{
			// create the frame (pave a blank spot)
			{row: 49, value: []byte("")},
			{row: 50, value: []byte("")},
			// draw the first update
			{row: 49, value: []byte("LineIdx:0")},
			{row: 50, value: []byte("LineIdx:1")},
		},
		[]string{
			"line is out of bounds (row=51)",
		},
	},
	"FloatFree_TermHeightSmall_AtBottom_Header": {3, 1, 0, 49, PolicyOverflow, 50,
		[]ScreenEvent{
			// create the frame (pave a blank spot)
			{row: 49, value: []byte("")},
			{row: 50, value: []byte("")},
			// draw the first update
			{row: 49, value: []byte("theHeader")},
			{row: 50, value: []byte("LineIdx:0")},
		},
		[]string{
			"line is out of bounds (row=51)",
			"line is out of bounds (row=52)",
		},
	},
	"FloatFree_termHeightSmall_AtBottom_Footer": {3, 0, 1, 49, PolicyOverflow, 50,
		[]ScreenEvent{
			// create the frame (pave a blank spot)
			{row: 49, value: []byte("")},
			{row: 50, value: []byte("")},
			// draw the first update
			{row: 49, value: []byte("LineIdx:0")},
			{row: 50, value: []byte("LineIdx:1")},
		},
		[]string{
			"line is out of bounds (row=51)",
			"line is out of bounds (row=52)",
		},
	},
	"FloatFree_TermHeightSmall_AtBottom_HeaderFooter": {3, 1, 1, 49, PolicyOverflow, 50,
		[]ScreenEvent{
			// create the frame (pave a blank spot)
			{row: 49, value: []byte("")},
			{row: 50, value: []byte("")},
			// draw the first update
			{row: 49, value: []byte("theHeader")},
			{row: 50, value: []byte("LineIdx:0")},
		},
		[]string{
			"line is out of bounds (row=51)",
			"line is out of bounds (row=52)",
			"line is out of bounds (row=53)",
		},
	},
}

func Test_FloatFreePolicy_Frame_Draw(t *testing.T) {

	names := make([]string, 0, len(floatFreeDrawTestCases))
	for name := range floatFreeDrawTestCases {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, test := range names {
		getScreen().reset()
		table := floatFreeDrawTestCases[test]

		// setup...
		terminalHeight = table.terminalHeight
		handler := NewTestEventHandler(t)
		scr := getScreen()
		scr.handlers = make([]EventHandler, 0)
		scr.addScreenHandler(handler)

		// run test...
		var errs []error
		frame, _ := New(Config{
			test:           true,
			Lines:          table.rows,
			HeaderRows:     table.headers,
			FooterRows:     table.footers,
			startRow:       table.startRow,
			PositionPolicy: table.policy,
		})
		if table.headers > 0 {
			frame.HeaderLines[0].buffer = []byte("theHeader")
		}
		for idx, line := range frame.BodyLines {
			line.buffer = []byte(fmt.Sprintf("LineIdx:%d", idx))
		}
		if table.footers > 0 {
			frame.FooterLines[0].buffer = []byte("theFooter")
		}
		errs = frame.Draw()

		// assert results...
		validateEvents(t, test, table, errs, frame, handler)

	}

}

func Test_FloatFreePolicy_Frame_AdhocDraw(t *testing.T) {

	names := make([]string, 0, len(floatFreeDrawTestCases))
	for name := range floatFreeDrawTestCases {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, test := range names {
		getScreen().reset()
		table := floatFreeDrawTestCases[test]

		// setup...
		terminalHeight = table.terminalHeight
		handler := NewTestEventHandler(t)
		scr := getScreen()
		scr.handlers = make([]EventHandler, 0)
		scr.addScreenHandler(handler)

		// run test...
		var err error
		var errs = make([]error, 0)
		frame, _ := New(Config{
			test:           true,
			Lines:          table.rows,
			HeaderRows:     table.headers,
			FooterRows:     table.footers,
			startRow:       table.startRow,
			PositionPolicy: table.policy,
		})
		if table.headers > 0 {
			err = frame.HeaderLines[0].WriteString("theHeader")
			if err != nil {
				errs = append(errs, err)
			}
		}
		for idx, line := range frame.BodyLines {
			err = line.WriteString(fmt.Sprintf("LineIdx:%d", idx))
			if err != nil {
				errs = append(errs, err)
			}
		}
		if table.footers > 0 {
			err = frame.FooterLines[0].WriteString("theFooter")
			if err != nil {
				errs = append(errs, err)
			}
		}

		// assert results...
		validateEvents(t, test, table, errs, frame, handler)

	}

}
