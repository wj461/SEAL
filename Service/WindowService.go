package Service

import (
	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
	"golang.design/x/hotkey"
)

type WindowService struct {
	app        *application.App
	petService *PetService
	shortcuts  []*hotkey.Hotkey // 存儲快捷鍵實例，防止被垃圾回收
}

func NewWindowService(app *application.App, petService *PetService) *WindowService {
	return &WindowService{
		app:        app,
		petService: petService,
		shortcuts:  make([]*hotkey.Hotkey, 0),
	}
}

func (w *WindowService) SetApp(app *application.App) {
	w.app = app
}

func (w *WindowService) CreateNewWindow() string {
	id := w.petService.IdForNewWindow()
	w.NewWindow(id)
	return id
}

func (w *WindowService) NewWindow(name string) {
	w.app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: name,
		Name:  name,
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
		},
		URL:            "/",
		Frameless:      true,
		Width:          1024,
		Height:         768,
		DisableResize:  true,
		BackgroundType: application.BackgroundTypeTransparent,
		AlwaysOnTop:    true,
		Windows: application.WindowsWindow{
			DisableFramelessWindowDecorations: true,
		},
	})
}

func (w *WindowService) RegisterGlobalShortcuts() {
	// 註冊 Ctrl+N 或 Command+N 快捷鍵
	hk := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl}, hotkey.KeyN)
	err := hk.Register()
	if err != nil {
		log.Printf("Failed to register hotkey: %v", err)
		return
	}

	// 保存快捷鍵實例，防止被垃圾回收
	w.shortcuts = append(w.shortcuts, hk)

	// 創建一個協程來監聽快捷鍵事件
	go func() {
		for range hk.Keydown() {
			// 當按下快捷鍵時創建新視窗
			w.CreateNewWindow()
		}
	}()
}
