package cases

import (
	"time"

	"github.com/mgobench"
	"github.com/roshanraj/goRandString/goRand"

	"gopkg.in/mgo.v2/bson"
)

type FlatT1Doc struct {
	ID    bson.ObjectId `bson:"_id,omitempty`
	StrF  string        `bson:"strf"`
	IntF  int64         `bson:"intf"`
	BoolF bool          `bson:"boolf"`
	TimeF time.Time     `bson:"timef"`
}

func FlatT1DocTest(t time.Duration, r *mgobench.ResultWorker, wm mgobench.WorkerManager, mt mgobench.MongoTask) {

	killTime := time.After(t)
Loop:
	for {

		select {

		case <-killTime:
			// send to influxdb
			break Loop
		default:
			var data = make([]interface{}, 0)
			data = append(data, &FlatT1Doc{
				ID:    bson.NewObjectId(),
				StrF:  goRand.RandString(8),
				IntF:  26,
				BoolF: true,
				TimeF: time.Now(),
			})
			ch := mgobench.InsertTask{
				MongoTask: mt,
				Docs:      data,
				Name:      "Oorder",
			}

			wm.Send(ch)
		}
	}

}
