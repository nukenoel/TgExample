package bot

import (
	"TelegramBot/config"
	"TelegramBot/internal/division"
	"TelegramBot/internal/project"
	"TelegramBot/internal/user"
	"TelegramBot/pkg/postgresql"
	"context"
	log "github.com/sirupsen/logrus"
	"gopkg.in/telebot.v3"
	"net/http"
	"strconv"
	"time"
)

type app struct {
	cfg                *config.Config
	httpServer         *http.Server
	bot                *telebot.Bot
	userRepository     user.RepositoryUser
	divisionRepository division.RepositoryDivision
	projectRepository  project.RepositoryProject
}

type App interface {
	Run()
}

func NewApp(cfg *config.Config, postgresqlClient postgresql.Client) (App, error) {
	return &app{
		cfg:                cfg,
		userRepository:     user.NewRepository(postgresqlClient),
		projectRepository:  project.NewRepository(postgresqlClient),
		divisionRepository: division.NewRepository(postgresqlClient),
	}, nil
}

func (a *app) Run() {
	bot, err := a.createBot()
	if err != nil {
		return //error: bot functionality is not complete from createBot() function
	}
	a.bot = bot
	log.Info("bot successfully started")
	a.bot.Start()
}

func (a *app) createBot() (tbot *telebot.Bot, botErr error) {
	pref := telebot.Settings{
		Token:   a.cfg.TgToken,
		Poller:  &telebot.LongPoller{Timeout: 60 * time.Second},
		OnError: a.OnBotError,
	}

	tbot, botErr = telebot.NewBot(pref)
	if botErr != nil {
		log.Fatalf("error on create BOT, description: %v", botErr)
		return
	}

	//Create control buttons
	var (
		selector = &telebot.ReplyMarkup{}
		btsm     = selector.Data(a.cfg.Buttons.BtStart, "someButton")
		btsm2    = selector.Data(a.cfg.Buttons.BtNext, "someButton")
	)

	selector.Inline(
		selector.Row(btsm, btsm2),
	)
	//The end creating control buttons

	//Create marks buttons
	var (
		selectorMark = &telebot.ReplyMarkup{
			ResizeKeyboard: true,
		}
		btOne   = selector.Data(strconv.Itoa(oneMark), "someMark")
		btTwo   = selector.Data(strconv.Itoa(twoMark), "someMark")
		btThree = selector.Data(strconv.Itoa(threeMark), "someMark")
		btFour  = selector.Data(strconv.Itoa(fourMark), "someMark")
		btFive  = selector.Data(strconv.Itoa(fiveMark), "someMark")
		btSix   = selector.Data(strconv.Itoa(sixMark), "someMark")
		btSeven = selector.Data(strconv.Itoa(sevenMark), "someMark")
		btEight = selector.Data(strconv.Itoa(eightMark), "someMark")
		btNine  = selector.Data(strconv.Itoa(nineMark), "someMark")
		btTen   = selector.Data(strconv.Itoa(tenMark), "someMark")
	)
	selectorMark.Inline(
		selectorMark.Row(btOne, btTwo, btThree, btFour, btFive),
		selectorMark.Row(btSix, btSeven, btEight, btNine, btTen),
	)
	//The end creating marks buttons

	tbot.Handle(&btsm, func(c telebot.Context) error {
		//c.Edit("вы поставили оценку 1")
		return nil
	})

	tbot.Handle(&btsm2, func(c telebot.Context) error {

		return nil
	})

	tbot.Handle("/start", func(c telebot.Context) error {
		//c.Send("Вы начали опрос\n Выберите свой отдел", selectorMark)
		all, err := a.divisionRepository.FindAll(context.TODO())
		if err != nil {
			log.Errorf("error on getting info about all divisions descriptpion:%v", err)
			return err
		}
		log.Info(all)
		divisionsSelector := &telebot.ReplyMarkup{ResizeKeyboard: true}
		for _, fields := range all {
			but := selector.Data(fields.Name, "task", fields.Name)
			tbot.Handle(&but)
		}
		c.Send("Выберите свой отдел", divisionsSelector)
		return nil
	})

	return
}

func (a *app) OnBotError(err error, ctx telebot.Context) {
	log.Error(err)
}
