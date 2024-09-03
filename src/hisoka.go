package conn

import (
	"context"
	"fmt"
	"hisoka/src/handlers"
	"hisoka/src/helpers"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	_ "hisoka/src/commands"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waCompanionReg"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/protobuf/proto"
)

type Template struct {
	Nama   string
	Status bool
}

var log helpers.Logger

func init() {
	store.DeviceProps.PlatformType = waCompanionReg.DeviceProps_EDGE.Enum()
	store.DeviceProps.Os = proto.String("Linux")
}

func StartClient() {
	dbLog := waLog.Stdout("Database", "ERROR", true)
	container, err := sqlstore.New("sqlite3", "file:session.db?_foreign_keys=on", dbLog)
	if err != nil {
		panic(err)
	}
	handler := handlers.NewHandler(container)
	log.Info("Connecting Socket")
	conn := handler.Client()
	conn.PrePairCallback = func(jid types.JID, platform, businessName string) bool {
		log.Info("Connected Socket")
		return true
	}

	if conn.Store.ID == nil {
		// No ID stored, new login
		pairingNumber := os.Getenv("PAIRING_NUMBER")

		if pairingNumber != "" {
			pairingNumber = regexp.MustCompile(`\D+`).ReplaceAllString(pairingNumber, "")

			if err := conn.Connect(); err != nil {
				panic(err)
			}

			code, err := conn.PairPhone(pairingNumber, true, whatsmeow.PairClientChrome, "Edge (Linux)")
			if err != nil {
				panic(err)
			}

			fmt.Println("Code Kamu : " + code)
		} else {
			qrChan, _ := conn.GetQRChannel(context.Background())
			if err := conn.Connect(); err != nil {
				panic(err)
			}

			for evt := range qrChan {
				switch string(evt.Event) {
				case "code":
					qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
					log.Info("Qr Required")
				}
			}
		}
	} else {
		// Already logged in, just connect
		if err := conn.Connect(); err != nil {
			panic(err)
		}
		log.Info("Connected Socket")
	}

	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	conn.Disconnect()
}
