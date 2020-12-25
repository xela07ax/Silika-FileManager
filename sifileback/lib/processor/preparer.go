package processor

import (
	"fmt"
	"github.com/xela07ax/toolsXela/archiver"
	"github.com/xela07ax/toolsXela/tp"
	"path/filepath"
)

// Распакуем пришедший проект, развернем все

func ectractor(kodConfigZip, workMinion string, gophere int, loger chan <-[4]string) {
	// Распакуем конфиг
	z := archiver.Zip{
		CompressionLevel:       -1,
		MkdirAll:               true,
		SelectiveCompression:   true,
		ContinueOnError:        false,
		OverwriteExisting:      true,
		ImplicitTopLevelFolder: false,  //Неявная папка верхнего уровня
	}
	workMinionKod := filepath.Join(workMinion, "kod")
	err := tp.CheckMkdir(workMinion)
	if err != nil {
		loger <- [4]string{fn,fmt.Sprintf("%d",gophere),fmt.Sprintf("CheckMkdir:%s",err),"1"}
		return
	}

	loger <- [4]string{fn,fmt.Sprintf("%d",gophere),fmt.Sprintf("Unarchive src:%s|dst:%s",kodConfigZip,workMinionKod)}
	err = z.Unarchive(kodConfigZip,workMinionKod)
	if err != nil {
		loger <- [4]string{fn,fmt.Sprintf("%d",gophere),fmt.Sprintf("Unarchive:%s",err),"1"}
		return
	}
	loger <- [4]string{fn,fmt.Sprintf("%d",gophere),"Unrchive: SUCCESS"}
}