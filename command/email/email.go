package email

import (
	"log"

	c "github.com/AdilHoumadi/background/config"
	"github.com/remind101/mq-go"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   c.Email,
	Short: c.EmailShort,
	Run: func(cmd *cobra.Command, args []string) {
		c.Listen(Processor, c.Demo)
	},
}

func Processor(message *mq.Message) error {
	log.Println("Hello", *message.SQSMessage.MessageId)
	return nil
}
