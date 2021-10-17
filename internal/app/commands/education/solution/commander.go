package solution

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/education"
	"github.com/ozonmp/omp-bot/internal/service/education/solution"
	"github.com/ozonmp/omp-bot/internal/servicedata"
	"log"
	"strconv"
	"strings"
)

type SolutionCommander struct {
	bot              *tgbotapi.BotAPI
	SolutionService *solution.DummySolutionService
}

type Solution_Commander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type CallbackListData struct {
	Start uint64  `json:"start"`
	Offset uint64 `json:"offset"`
}

func OneCommandName(Begin string) string{
	return fmt.Sprintf("/%s__education__solution", Begin)
}
func NewSolutionCommander(bot *tgbotapi.BotAPI, ) *SolutionCommander {
	service := solution.NewDummySolutionService()
	return &SolutionCommander{
		bot:             bot,
		SolutionService: service,
	}
}

func (p *SolutionCommander) SendMessage(inputMsg *tgbotapi.Message, msgtext string){
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, msgtext)
	p.bot.Send(msg)
}

func GetArgument(inputMsg *tgbotapi.Message) (uint64, string){
	args := inputMsg.CommandArguments()

	idx, err := strconv.ParseUint(args, 0, 64)
	if err != nil {
		TextMsg := fmt.Sprintf("wrong args %s", args)
		log.Println(TextMsg)
		return 0, TextMsg
	}
	return idx, ""
}

func (p *SolutionCommander) Help(inputMsg *tgbotapi.Message){
	p.SendMessage(inputMsg, OneCommandName("help") + " - Эта справка\n"+
		OneCommandName("get") + " - получить элемент по номеру\n"+
		OneCommandName("list") + " - получить список элементов\n"+
		OneCommandName("delete") + " - удалить элемент по номеру\n"+
		OneCommandName("new") + " - создать новый элемент\n"+
		OneCommandName("edit") + " - отредактировать элемент\n")
}

func (p *SolutionCommander) Get(inputMsg *tgbotapi.Message){
	TextMsg := ""
	defer func() {
		p.SendMessage(inputMsg, TextMsg)
	}()
	idx, TextMsg := GetArgument(inputMsg)
	if TextMsg != "" { return}

	product, err := p.SolutionService.Describe(idx)
	if err != nil {
		TextMsg = fmt.Sprintf("fail to get product with idx %d: %v", idx, err)
		log.Println(TextMsg)
		return
	}
	TextMsg = product.String()
}
func (p *SolutionCommander) List(inputMsg *tgbotapi.Message){
	cb := CallbackListData{ 1, 1}
	data, _ := json.Marshal(cb)

	callbackPath := path.CallbackPath{
		Domain:       "education",
		Subdomain:    "solution",
		CallbackName: "list",
		CallbackData: string(data),
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, strings.Join(p.SolutionService.List(cb.Start-1, cb.Offset), "\n"))
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next", callbackPath.String(),
			)))
	p.bot.Send(msg)
}
func (p *SolutionCommander) New(inputMsg *tgbotapi.Message){
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"/help - help\n"+
			"/list - list products",
	)
	p.bot.Send(msg)
}
func (p *SolutionCommander) Delete(inputMsg *tgbotapi.Message){
	TextMsg := ""
	defer func() {
		p.SendMessage(inputMsg, TextMsg)
	}()
	idx, TextMsg := GetArgument(inputMsg)
	if TextMsg != "" { return}

	_, err := p.SolutionService.Remove(idx)
	if err != nil {
		TextMsg = fmt.Sprintf("fail to get product with idx %d: %v", idx, err)
		log.Println(TextMsg)
		return
	}
	TextMsg = "Запись удалена"
}

func (p *SolutionCommander) Edit(inputMsg *tgbotapi.Message){
	TextMsg := ""
	defer func() {
		p.SendMessage(inputMsg, TextMsg)
	}()
	idx, TextMsg := GetArgument(inputMsg)
	if TextMsg != "" { return}

	product, err := p.SolutionService.Describe(idx)
	if err != nil {
		TextMsg = fmt.Sprintf("fail to get product with idx %d: %v", idx, err)
		log.Println(TextMsg)
		return
	}
	servicedata.EditedChat[inputMsg.Chat.ID] = *(servicedata.GetOperationData(idx, servicedata.EditoperationData))
	TextMsg = product.String() + "\n Измененная запись должна содержать поля TaskID, Autor, Title. Все поля "+
		"должны быть в одном сообщении каждое поле в отдельной строке."
}

func (c *SolutionCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	var msg = tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		strings.Join(c.SolutionService.List(parsedData.Start, parsedData.Offset), "\n"))

	if c.SolutionService.Len() > parsedData.Start + parsedData.Offset {
		parsedData.Start += parsedData.Offset
		data, _ := json.Marshal(parsedData)
		callbackData := path.CallbackPath{
			Domain:       "education",
			Subdomain:    "solution",
			CallbackName: "list",
			CallbackData: string(data),
		}
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next", callbackData.String(),
				)))
	}
	c.bot.Send(msg)
}
func (c *SolutionCommander) Default(inputMessage *tgbotapi.Message) {
	TextMsg := ""
	defer func() {
		c.SendMessage(inputMessage, TextMsg)
	}()
	if idx, ok := servicedata.EditedChat[inputMessage.Chat.ID]; ok {
		data := strings.Split(inputMessage.Text, "\n")
		if len(data) != 3 {
			TextMsg = "В сообщение должно быть 3 строки, повторите ввод, пожалуйста"
			return
		}
		taskID, err := strconv.ParseUint(data[0], 0, 64)
		if err != nil {
			TextMsg = "Первая строка не содержит число, повторите ввод, пожалуйста"
			return
		}
		solution := education.Solution{	}
		solution.Id = idx.ProductID
		solution.TaskID = taskID
		solution.Autor = data[1]
		solution.Title = data[2]
		c.SolutionService.Update(idx.ProductID, solution)
		delete(servicedata.EditedChat, inputMessage.Chat.ID)
		sol, _ := c.SolutionService.Describe(idx.ProductID)
		TextMsg = "Запись заменена: \n " + sol.String()
	} else {
		log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)

		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)

		c.bot.Send(msg)
	}
}

func (c *SolutionCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("DemoSubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *SolutionCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}
