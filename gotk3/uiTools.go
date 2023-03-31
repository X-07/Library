package gotk3

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var (
	progressBarPopup *gtk.Window
	ProgressBar      *gtk.ProgressBar
)

func CreateWindow() *gtk.Window {
	window, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	ErrorCheckIHM("Unable to create Window ", err)
	return window
}

func CreatePopup(window *gtk.Window, border uint, position gtk.WindowPosition) *gtk.Window {
	popup, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	ErrorCheckIHM("Unable to create Window ", err)
	popup.SetTransientFor(window)
	popup.SetPosition(position)
	popup.SetBorderWidth(border)
	popup.SetModal(true)
	return popup
}

func CreateCssProvider() *gtk.CssProvider {
	provider, err := gtk.CssProviderNew()
	ErrorCheckIHM("Unable to create CssProviderNew ", err)
	screen, err := gdk.ScreenGetDefault()
	ErrorCheckIHM("Unable to create ScreenGetDefault ", err)
	gtk.AddProviderForScreen(screen, provider, 600)
	return provider
}

func CreateHeaderBar() *gtk.HeaderBar {
	header, err := gtk.HeaderBarNew()
	ErrorCheckIHM("Unable to create HeaderBar ", err)
	return header
}

func CreatePaned() *gtk.Paned {
	panedWin, err := gtk.PanedNew(gtk.ORIENTATION_HORIZONTAL)
	ErrorCheckIHM("Unable to create Paned ", err)
	return panedWin
}

func CreateNoteBook() *gtk.Notebook {
	noteBook, err := gtk.NotebookNew()
	ErrorCheckIHM("Unable to create NoteBook ", err)
	return noteBook
}

func CreateVBox(padding int) *gtk.Box {
	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, padding)
	ErrorCheckIHM("Unable to create vBox ", err)
	return box
}

func CreateHBox(padding int) *gtk.Box {
	box, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, padding)
	ErrorCheckIHM("Unable to create hBox ", err)
	return box
}

func CreateCalendar() *gtk.Calendar {
	calendar, err := gtk.CalendarNew()
	ErrorCheckIHM("Unable to create Calendar ", err)
	return calendar
}

func CreateHBoxImg(pixbuf *gdk.Pixbuf) *gtk.Box {
	itemBox := CreateHBox(0)
	itemBox.PackStart(CreateImageFromPixbuf(pixbuf), false, false, 0)
	return itemBox
}

func CreateHBoxImgLabel(leftLabel string, pixbuf *gdk.Pixbuf, rightLabel string) *gtk.Box {
	itemBox := CreateHBox(0)
	itemBox.PackStart(CreateLabel(leftLabel), false, false, 0)
	itemBox.PackStart(CreateImageFromPixbuf(pixbuf), false, false, 0)
	itemBox.PackStart(CreateLabel(rightLabel), false, false, 0)
	return itemBox
}

func CreateHBoxImgRightLabel(pixbuf *gdk.Pixbuf, rightLabel string) *gtk.Box {
	itemBox := CreateHBox(0)
	itemBox.PackStart(CreateImageFromPixbuf(pixbuf), false, false, 0)
	itemBox.PackStart(CreateLabel(rightLabel), false, false, 0)
	return itemBox
}

func CreateHBoxImgLeftLabel(leftLabel string, pixbuf *gdk.Pixbuf) *gtk.Box {
	itemBox := CreateHBox(0)
	itemBox.PackStart(CreateLabel(leftLabel), false, false, 0)
	itemBox.PackStart(CreateImageFromPixbuf(pixbuf), false, false, 0)
	return itemBox
}

func CreateHBoxImgLeftLabelWithStyle(leftLabel string, pixbuf *gdk.Pixbuf, styles ...string) *gtk.Box {
	itemBox := CreateHBox(0)
	itemBox.PackStart(CreateLabelWithStyle(leftLabel, styles...), false, false, 0)
	itemBox.PackStart(CreateImageFromPixbuf(pixbuf), false, false, 0)
	return itemBox
}

func CreateToolbar() *gtk.Toolbar {
	toolbar, err := gtk.ToolbarNew()
	ErrorCheckIHM("Unable to create ToolBar ", err)
	//toolbar.SetStyle(gtk.TOOLBAR_TEXT)
	return toolbar
}

func CreateToolItem() *gtk.ToolItem {
	toolItem, err := gtk.ToolItemNew()
	ErrorCheckIHM("Unable to create ToolItem ", err)
	return toolItem
}

func CreateSeparatorToolItem() *gtk.SeparatorToolItem {
	separatorToolItem, err := gtk.SeparatorToolItemNew()
	ErrorCheckIHM("Unable to create SeparatorToolItem ", err)
	return separatorToolItem
}

func CreateMenu() *gtk.Menu {
	menu, err := gtk.MenuNew()
	ErrorCheckIHM("Unable to create Menu ", err)
	return menu
}

func CreateMenuBar() *gtk.MenuBar {
	menuBar, err := gtk.MenuBarNew()
	ErrorCheckIHM("Unable to create MenuBar ", err)
	return menuBar
}

func CreateMenuItemWithLabel(titre string) *gtk.MenuItem {
	menuItem, err := gtk.MenuItemNewWithLabel(titre)
	ErrorCheckIHM("Unable to create MenuItem ", err)
	return menuItem
}

func CreateMenuItem() *gtk.MenuItem {
	menuItem, err := gtk.MenuItemNew()
	ErrorCheckIHM("Unable to create MenuItem ", err)
	return menuItem
}

func CreateImageMenuItem(label string, pixbuf *gdk.Pixbuf) *gtk.MenuItem {
	menuItem := CreateMenuItem()
	itemBox := CreateHBox(0)

	itemBox.PackStart(CreateImageFromPixbuf(pixbuf), false, false, 0)

	itemLabel := CreateLabel(label)
	itemLabel.SetMarginStart(10)
	itemBox.PackStart(itemLabel, false, false, 0)
	menuItem.Add(itemBox)

	return menuItem
}

func CreateRatingMenuItem(label string, pixbufList []*gdk.Pixbuf) *gtk.MenuItem {
	menuItem := CreateMenuItem()
	itemBox := CreateHBox(0)

	itemLabel := CreateLabel(label)
	itemLabel.SetMarginEnd(10)
	itemBox.PackStart(itemLabel, false, false, 20)
	for _, pixbuf := range pixbufList {
		itemBox.PackStart(CreateImageFromPixbuf(pixbuf), false, false, 0)
	}

	//menuItem.SetLabel(label)
	menuItem.Add(itemBox)

	return menuItem
}

func CreateTreeView() *gtk.TreeView {
	treeView, err := gtk.TreeViewNew()
	ErrorCheckIHM("Unable to create TreeView ", err)
	return treeView
}

func CreateTreeStore(types ...glib.Type) *gtk.TreeStore {
	treeStore, err := gtk.TreeStoreNew(types...)
	ErrorCheckIHM("Unable to create TreeStore ", err)
	return treeStore
}

func CreateCellRendererText() *gtk.CellRendererText {
	cellRenderer, err := gtk.CellRendererTextNew()
	ErrorCheckIHM("Unable to create CellRendererText ", err)
	// cellRenderer.Set("cell-background", "orange")
	// cellRenderer.Set("cell-background-set", false)
	// cellRenderer.Set("foreground", "red")
	// cellRenderer.Set("foreground-set", false)
	return cellRenderer
}

func CreateCellRendererPixbuf() *gtk.CellRendererPixbuf {
	cellRenderer, err := gtk.CellRendererPixbufNew()
	ErrorCheckIHM("Unable to create CellRendererPixbufNew ", err)
	return cellRenderer
}

func CreateTreeViewColumnNewWithAttribute(title string, cellRenderer gtk.ICellRenderer, attribute string, idColumn int) *gtk.TreeViewColumn {
	column, err := gtk.TreeViewColumnNewWithAttribute(title, cellRenderer, attribute, idColumn)
	ErrorCheckIHM("Unable to create TreeViewColumn ", err)
	return column
}

func CreateTextColumn(title string, id int) *gtk.TreeViewColumn {
	return CreateTreeViewColumnNewWithAttribute(title, CreateCellRendererText(), "text", id)
}

func CreateImageColumn(title string, id int) *gtk.TreeViewColumn {
	return CreateTreeViewColumnNewWithAttribute(title, CreateCellRendererPixbuf(), "pixbuf", id)
}

func CreateProgressBar() *gtk.ProgressBar {
	progressBar, err := gtk.ProgressBarNew()
	ErrorCheckIHM("Unable to create ProgressBar ", err)
	return progressBar
}

func CreateSeparatorMenuItem() *gtk.SeparatorMenuItem {
	separator, err := gtk.SeparatorMenuItemNew()
	ErrorCheckIHM("Unable to create SeparatorMenuItem ", err)
	separator.SetMarginTop(10)
	separator.SetMarginBottom(10)
	return separator
}

func CreateEntry() *gtk.Entry {
	entry, err := gtk.EntryNew()
	ErrorCheckIHM("Unable to create Entry ", err)
	styleCtx, err := entry.GetStyleContext()
	ErrorCheckIHM("Unable to get StyleEntry ", err)
	styleCtx.AddClass("bgWhite")
	return entry
}

func CreateEntryWithStyle(styles ...string) *gtk.Entry {
	entry := CreateEntry()
	styleCtx, err := entry.GetStyleContext()
	ErrorCheckIHM("Unable to get StyleEntry ", err)
	for _, style := range styles {
		styleCtx.AddClass(style)
	}
	return entry
}

func CreateTextView() *gtk.TextView {
	textView, err := gtk.TextViewNew()
	ErrorCheckIHM("Unable to create TextView ", err)
	// styleCtx, err := textView.GetStyleContext()
	// ErrorCheckIHM("Unable to get StyleEntry ", err)
	// styleCtx.AddClass("bgWhite")
	return textView
}

func CreateLabel(libelle string) *gtk.Label {
	label, err := gtk.LabelNew(libelle)
	ErrorCheckIHM("Unable to create Label ", err)
	return label
}

func CreateBoldLabel(libelle string) *gtk.Label {
	label, err := gtk.LabelNew("")
	ErrorCheckIHM("Unable to create Label ", err)
	label.SetMarkup("<b>" + libelle + "</b>")
	return label
}

func CreateLabelWithStyle(libelle string, styles ...string) *gtk.Label {
	label := CreateLabel(libelle)
	styleCtx, err := label.GetStyleContext()
	ErrorCheckIHM("Unable to get StyleLabel ", err)
	for _, style := range styles {
		styleCtx.AddClass(style)
	}
	return label
}

func CreateRightLabel(libelle string) *gtk.Label {
	label := CreateLabel(libelle)
	label.SetJustify(gtk.JUSTIFY_RIGHT)
	return label
}

func CreateBlankLine() *gtk.Label {
	label, err := gtk.LabelNew("")
	ErrorCheckIHM("Unable to create Label ", err)
	label.SetMarginTop(11)
	label.SetMarginBottom(12)
	return label
}

func CreateFrame(titre string) *gtk.Frame {
	frame, err := gtk.FrameNew(titre)
	ErrorCheckIHM("Unable to create Frame ", err)
	frame.SetShadowType(gtk.SHADOW_ETCHED_IN)
	return frame
}

func CreateRadioButton(radioPrev *gtk.RadioButton, libelle string) (*gtk.RadioButton, *gtk.RadioButton) {
	radioButton, err := gtk.RadioButtonNewWithLabelFromWidget(radioPrev, libelle)
	ErrorCheckIHM("Unable to create RadioButton ", err)
	return radioButton, radioButton
}

func CreateCheckButton(libelle string) *gtk.CheckButton {
	checkButton, err := gtk.CheckButtonNewWithLabel(libelle)
	ErrorCheckIHM("Unable to create CheckButton ", err)
	return checkButton
}

func CreateToolButton(libelle string) *gtk.ToolButton {
	toolButton, err := gtk.ToolButtonNew(nil, "")
	ErrorCheckIHM("Unable to create ToolButton ", err)
	return toolButton
}

func CreateComboBoxTextWithEntry(liste []string) *gtk.ComboBoxText {
	comboBoxText, err := gtk.ComboBoxTextNewWithEntry()
	ErrorCheckIHM("Unable to create ComboBoxText ", err)
	entry, err := comboBoxText.GetEntry()
	ErrorCheckIHM("Unable to create Entry from ComboBox ", err)
	styleCtx, err := entry.GetStyleContext()
	ErrorCheckIHM("Unable to get StyleEntry ", err)
	styleCtx.AddClass("bgWhite")
	for _, txt := range liste {
		comboBoxText.AppendText(txt)
	}
	return comboBoxText
}

func CreateComboBoxText(liste []string) *gtk.ComboBoxText {
	comboBoxText, err := gtk.ComboBoxTextNew()
	ErrorCheckIHM("Unable to create ComboBoxText ", err)
	for _, txt := range liste {
		comboBoxText.AppendText(txt)
	}
	comboBoxText.SetActive(0)
	return comboBoxText
}

func CreateComboBoxWithModel(model gtk.ITreeModel) *gtk.ComboBox {
	comboBox, err := gtk.ComboBoxNewWithModel(model)
	ErrorCheckIHM("Unable to create ComboBox ", err)
	return comboBox
}

func CreateSpinButtonWithRange(min, max, step float64) *gtk.SpinButton {
	spinButton, err := gtk.SpinButtonNewWithRange(min, max, step)
	ErrorCheckIHM("Unable to create SpinButton ", err)
	entry := spinButton.Entry
	styleCtx, err := entry.GetStyleContext()
	ErrorCheckIHM("Unable to get StyleEntry ", err)
	styleCtx.AddClass("bgWhite")
	return spinButton
}

func CreateSpinButton(lower float64) *gtk.SpinButton {
	adjustment, err := gtk.AdjustmentNew(1, lower, 10000, 1, 10, 0)
	ErrorCheckIHM("Unable to create Adjustment ", err)
	spinButton, err := gtk.SpinButtonNew(adjustment, 0, 0)
	ErrorCheckIHM("Unable to create SpinButton ", err)
	//	spinButton.SetNumeric(true)
	return spinButton
}

func CreateTextButton(label string) *gtk.Button {
	button, err := gtk.ButtonNew()
	ErrorCheckIHM("Unable to create Button ", err)
	button.SetLabel(label)
	return button
}

func CreateTextButtonWithStyle(label string, styles ...string) *gtk.Button {
	button := CreateTextButton(label)
	styleCtx, err := button.GetStyleContext()
	ErrorCheckIHM("Unable to get StyleButton ", err)
	for _, style := range styles {
		styleCtx.AddClass(style)
	}
	return button
}

// func CreateImageButton(label string, imgPath string) *gtk.Button {
// 	button := CreateTextButton(label)
// 	button.SetImage(getImage(imgPath, 22, 22, true))
// 	return button
// }

func CreateImageButton(label string, icon *gdk.Pixbuf) *gtk.Button {
	button := CreateTextButton(label)
	button.SetAlwaysShowImage(true)
	button.SetImage(CreateImageFromPixbuf(icon))
	return button
}

func CreateImageButtonWithStyle(label string, icon *gdk.Pixbuf, styles ...string) *gtk.Button {
	button := CreateTextButton(label)
	button.SetAlwaysShowImage(true)
	button.SetImage(CreateImageFromPixbuf(icon))
	styleCtx, err := button.GetStyleContext()
	ErrorCheckIHM("Unable to get StyleButton ", err)
	for _, style := range styles {
		styleCtx.AddClass(style)
	}
	return button
}

func CreateIconButton(label string, icon []byte) *gtk.Button {
	button := CreateTextButton(label)
	pixbuf := CreateIconPixBuf(icon)
	img, err := gtk.ImageNewFromPixbuf(pixbuf)
	ErrorCheckIHM("Unable to create Image ", err)
	button.SetImage(img)
	return button
}

func CreateIconPixBufButton(label string, icon *gdk.Pixbuf) *gtk.Button {
	button := CreateTextButton(label)
	img, err := gtk.ImageNewFromPixbuf(icon)
	ErrorCheckIHM("Unable to create Image ", err)
	button.SetImage(img)
	return button
}

func CreateIconPixBuf(icon []byte) *gdk.Pixbuf {
	pixbuf, err := gdk.PixbufNewFromBytesOnly(icon)
	ErrorCheckIHM("Unable to create PixBuf from []byte ", err)
	return pixbuf
}

func CreateIconPixBufAnimation(imgPath string) *gdk.PixbufAnimation {
	pixbufAnimation, err := gdk.PixbufAnimationNewFromFile(imgPath)
	ErrorCheckIHM("Unable to create PixBufAnimation from file ", err)
	return pixbufAnimation
}

func GetImage(imgPath string, width, height int, ratio bool) *gtk.Image {
	pixbuf := CreatePixBuf(imgPath, width, height, ratio)
	img, err := gtk.ImageNewFromPixbuf(pixbuf)
	ErrorCheckIHM("Unable to create Image ", err)
	return img
}

func ResizeIconPixBuf(icon *gdk.Pixbuf, width int, height int) *gdk.Pixbuf {
	ico2, err := icon.ScaleSimple(width, height, gdk.INTERP_BILINEAR)
	ErrorCheckIHM("Unable to ScaleSimple PixBuf ", err)
	return ico2
}

func CreateImageFromPixbuf(pixbuf *gdk.Pixbuf) *gtk.Image {
	img, err := gtk.ImageNewFromPixbuf(pixbuf)
	ErrorCheckIHM("Unable to create Image ", err)
	return img
}

func CreateImage() *gtk.Image {
	picture, err := gtk.ImageNew()
	ErrorCheckIHM("Unable to create Image ", err)
	return picture
}

func CreatePixBuf(imgPath string, width, height int, ratio bool) *gdk.Pixbuf {
	pixBuf, err := gdk.PixbufNewFromFileAtScale(imgPath, width, height, ratio)
	ErrorCheckIHM("Unable to create PixBuf ", err)
	return pixBuf
}

func CreateGrid() *gtk.Grid {
	grid, err := gtk.GridNew()
	ErrorCheckIHM("Unable to create Grid ", err)
	grid.SetColumnHomogeneous(true)
	grid.SetRowHomogeneous(true)
	return grid
}

func CreateScrolledWindow() *gtk.ScrolledWindow {
	scrolledWindow, err := gtk.ScrolledWindowNew(nil, nil)
	ErrorCheckIHM("Unable to create ScrolledWindow ", err)
	return scrolledWindow
}

func CreateEventBox() *gtk.EventBox {
	eventBox, err := gtk.EventBoxNew()
	ErrorCheckIHM("Unable to create EventBox ", err)
	return eventBox
}

func GetComboBoxTextList(comboBox *gtk.ComboBoxText) []string {
	iModel, err := comboBox.GetModel()
	ErrorCheckIHM("Unable to create ComboBox.GetModel ", err)

	result := GetModelList(iModel)
	return result
}

// =================
func GetModelList(iModel gtk.ITreeModel) []string {
	var result []string

	model := iModel.(*gtk.ListStore)

	model.ForEach(func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
		result = append(result, GetStringValue(model, iter, 0))
		return false // pour ne pas stopper la boucle ForEach
	})

	return result
}

func SelectRow(listStore *gtk.ListStore, treeSelection *gtk.TreeSelection, ident int64, idCol int) {
	ok := false
	listStore.ForEach(func(model *gtk.TreeModel, path *gtk.TreePath, iter *gtk.TreeIter) bool {
		id := GetIntValue(model, iter, idCol)

		if id == ident {
			treeSelection.SelectIter(iter)
			ok = true
			return true // pour stopper la boucle ForEach, on a trouvé ce que l'on cherchait
		} else {
			return false // pour ne pas stopper la boucle ForEach
		}
	})
	if !ok {
		SelectFirstRow(listStore, treeSelection)
	}
}

func SelectFirstRow(listStore *gtk.ListStore, treeSelection *gtk.TreeSelection) {
	iter, ok := listStore.GetIterFirst()
	if ok {
		treeSelection.SelectIter(iter)
	}
}

// =================
func GetColumnButton(treeViewColumn *gtk.TreeViewColumn) *gtk.Button {
	iButton, err := treeViewColumn.GetButton()
	ErrorCheckIHM("Unable to GetButton from TreeViewColumn ", err)
	button := iButton.(*gtk.Button)
	return button
}

// =================
func GetIntValue(model *gtk.TreeModel, iter *gtk.TreeIter, column int) int64 {
	value, err := model.GetValue(iter, column) // int64
	ErrorCheckIHM("Unable to call GetValue from TreeModel ", err)
	typ, err := value.GoValue()
	ErrorCheckIHM("Unable to call GoValue from Value ", err)
	return typ.(int64)
}

func GetBoolValue(model *gtk.TreeModel, iter *gtk.TreeIter, column int) bool {
	value, err := model.GetValue(iter, column) // int64
	ErrorCheckIHM("Unable to call GetValue from TreeModel ", err)
	typ, err := value.GoValue()
	ErrorCheckIHM("Unable to call GoValue from Value ", err)
	return typ.(bool)
}

func GetStringValue(model *gtk.TreeModel, iter *gtk.TreeIter, column int) string {
	value, err := model.GetValue(iter, column) // string
	ErrorCheckIHM("Unable to call GetValue from TreeModel ", err)
	text, err := value.GetString()
	ErrorCheckIHM("Unable to call GetString from Value ", err)
	return text
}

func SetValue(store *gtk.TreeStore, iter *gtk.TreeIter, column int, value interface{}) {
	err := store.SetValue(iter, column, value)
	ErrorCheckIHM("Unable to SetValue to TreeStore ", err)
}

func SetLsValue(store *gtk.ListStore, iter *gtk.TreeIter, column int, value interface{}) {
	err := store.SetValue(iter, column, value)
	ErrorCheckIHM("Unable to SetValue to ListStore ", err)
}

// =================
func ErrorCheckIHM(msg string, err error) {
	if err != nil {
		appID := "fr.jlt.appError"
		app, _ := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)

		//Shows an application as soon as the app starts
		app.Connect("activate", func() {
			notify := glib.NotificationNew(">>> FATAL ERROR <<<")
			notify.SetBody(msg + "\n" + err.Error())
			notify.SetPriority(glib.NOTIFICATION_PRIORITY_HIGH)
			app.SendNotification(appID, notify)
		})

		app.Run(nil)

		fmt.Printf("%s : %s", msg, err.Error())
		panic(err)
	}
}

// =================
// #### CUSTOM #####
// =================
func CreateCalendarButton(date *gtk.Entry, winMain *gtk.Window, iconCalendar *gdk.Pixbuf) *gtk.Button {
	dateBtn := CreateImageButton("", iconCalendar)
	dateBtn.Connect("clicked", func() {
		winCalendar := CreatePopup(winMain, 10, gtk.WIN_POS_MOUSE)
		title, err := date.GetProperty("name")
		ErrorCheckIHM("Unable to create GetProperty from Entry ", err)

		winCalendar.SetTitle(title.(string))
		winCalendar.Connect("destroy", winCalendar.Close)

		calendar := CreateCalendar()
		winCalendar.Add(calendar)

		dateStr, err := date.GetText()
		ErrorCheckIHM("Unable to create GetText from Entry ", err)
		if dateStr != "" {
			dateElmt := strings.Split(dateStr, "/")
			if len(dateElmt) == 3 {
				calendar.SelectDay(uint(AtoI(dateElmt[0])))
				calendar.SelectMonth(uint(AtoI(dateElmt[1])-1), uint(AtoI(dateElmt[2])))
			}
		}

		monthChanged := false

		calendar.Connect("month-changed", func() {
			monthChanged = true
		})
		calendar.Connect("day-selected", func() {
			if monthChanged {
				monthChanged = false
				return
			}
			year, month, day := calendar.GetDate()
			jj := ItoA(int(day))
			mm := ItoA(int(month) + 1)
			aaaa := ItoA(int(year))
			if len(jj) == 1 {
				jj = "0" + jj
			}
			if len(mm) == 1 {
				mm = "0" + mm
			}
			date.SetText(jj + "/" + mm + "/" + aaaa)
			winCalendar.Close()
		})

		winCalendar.ShowAll()
	})
	return dateBtn
}

func AtoI(saisie string) int {
	value, err := strconv.Atoi(saisie)
	if err != nil {
		value = 0
	}
	return value
}

func ItoA(value int) string {
	return strconv.Itoa(value)
}

// =================
func MakeProgressBarPopup(win *gtk.Window, border uint, title string, position gtk.WindowPosition) {
	progressBarPopup = CreatePopup(win, border, position)
	progressBarPopup.SetDefaultSize(500, 100) // ($width, $height)
	progressBarPopup.SetTitle(title)

	progressBarMain := CreateVBox(0)
	progressBarPopup.Add(progressBarMain)

	progressBarLine := CreateHBox(0)
	progressBarMain.PackStart(progressBarLine, false, false, 0)
	progressBarLine.SetMarginTop(20)
	progressBarLine.SetMarginBottom(15)
	ProgressBar = CreateProgressBar()
	progressBarLine.PackStart(ProgressBar, true, true, 20)
	ProgressBar.SetShowText(true)

	progressBarPopup.Connect("destroy", func() {
		// fmt.Println("progressBarPopup DESTROY")
		progressBarPopup.Destroy()
		ProgressBar = nil
		progressBarPopup = nil
	})

	progressBarPopup.ShowAll()
}

func CloseProgressBarPopup() {
	progressBarPopup.Close()
	// progressBarPopup.Destroy()
}

// ***********************************************************
type IterStack []*gtk.TreeIter

func (stack *IterStack) IsEmpty() bool {
	return len(*stack) == 0
}

func (stack *IterStack) GetLastIter() *gtk.TreeIter {
	if stack.IsEmpty() {
		return nil
	} else {
		return (*stack)[len(*stack)-1]
	}
}

func (stack *IterStack) Push(iter *gtk.TreeIter) {
	*stack = append(*stack, iter)
}

func (stack *IterStack) Pop() {
	if !stack.IsEmpty() {
		index := len(*stack) - 1
		*stack = (*stack)[:index]
	}
}

// =================
