package communicator

import (
	"encoding/json"
	"fmt"
	"github.com/xela07ax/Silika-FileManager/sifileback/lib/processor"
)

type Box struct {
	ZipKodConfigDefault string
}
type Coming struct {
	Bag *Box
	Loger chan <- [4]string
	InputClients chan []byte
	OutputTxDock chan processor.ToDocker
}

type Notyfy struct {
	Id string
	Type string
	Name string
	Text string
	Data []byte
}
func (coming *Coming)ProcessRun()  {
	for {
		msg := <- coming.InputClients
		// Если команду поняли, то будем транслировать по ней информацию
		//fmt.Printf("%s",msg)
		// {"Name":"RunKod","Text":"Запуск проекта с вложением"}
		// Проблема в том, что отправленное сообщение немедленно возвращается
		// А потому надо принимать только команды, будем парсить формат
		// Сюда должна прийти структура новой команды, или игноррируем
		var command Notyfy
		err := json.Unmarshal(msg, &command)
		if err != nil {
			fmt.Printf("Errtx:%s",err)
			continue
		}
		//
		// {"Name":"RunRootFs","Text":"Запустим root kod fs"}
		coming.Loger <- [4]string{"Anonimouse","nil",fmt.Sprintf("%s",msg)}
		// Команда для запуска RootFs
		switch command.Name{
		case "RunRootFs":
			coming.Loger <- [4]string{"Anonimouse","nil",fmt.Sprintf("%s",command.Text)}
		// Запустим пустой kod config


		}
	}
}
