package metric

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"math"
	"os"
	"sync"
	"testing"
	"time"
)

const (
	ScriptExecuteQPS = "script_execute_qps_metric"

	LabelService = "Service"
	LabelRelease = "Release"
	LabelSpace   = "Space"
)

var (
	queryArray  [50]float64
	metricArray []float64
	queryLock   sync.Mutex
)

func Test_Meter(t *testing.T) {
	scriptExecuteQPS := NewMeterVec(1, MeterOpts{
		Name: ScriptExecuteQPS,
		Help: "qps of script execute",
	},
		[]string{LabelService, LabelRelease, LabelSpace},
	)
	var stopC = make(chan bool)
	var waitGroup sync.WaitGroup
	waitGroup.Add(3)

	// write routine
	go userMocker1(scriptExecuteQPS, &waitGroup)
	go userMocker2(scriptExecuteQPS, &waitGroup)
	go userMocker3(scriptExecuteQPS, &waitGroup)

	// read routine
	go printMetricInfo(stopC, &waitGroup)

	// wait write goroutines exit
	waitGroup.Wait()

	waitGroup.Add(1)
	stopC <- false
	// wait read goroutines exit
	waitGroup.Wait()

	//
	if len(metricArray) != len(queryArray) {
		fmt.Println("metricArray Length:", len(metricArray), "| queryArray Length:", len(queryArray))
		fmt.Println("metricArray:", metricArray)
		fmt.Println("queryArray:", queryArray)
		return
	}

	var isSame bool = true
	for index := 0; index < len(metricArray); index++ {
		if metricArray[index] != queryArray[index] {
			isSame = false
			break
		}
	}

	if isSame {
		fmt.Println("#############Verify Success!!!#############")
	} else {
		fmt.Println("#############Verify Failed!!!#############")
	}
}

func printMetricInfo(stopC chan bool, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	var qpsGauge *dto.MetricFamily //, curCounter, lastCounter
	var ticker = time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("******************************************************")
			metricFamilies, err := prometheus.DefaultGatherer.Gather()
			if err != nil {
				fmt.Printf("get metric family from default gatherer failed. Error: %v\n", err)
				os.Exit(10)
			}
			for _, metricFamily := range metricFamilies {
				if *metricFamily.Type == dto.MetricType_GAUGE && metricFamily.GetName() == ScriptExecuteQPS {
					qpsGauge = metricFamily
				}
			}
			//
			if qpsGauge == nil {
				fmt.Println("metric not found")
				continue
			}
			for _, metricItem := range qpsGauge.Metric {
				var qpsValue = metricItem.GetGauge().GetValue()
				// var curCounter = fmt.Printf("[%s]: %v\n", metricItem.Label, curValue)
				fmt.Printf("%v ---> QPS[%v]\n", metricItem.Label, qpsValue)
				metricArray = append(metricArray, qpsValue)
			}
		case <-stopC:
			return
		}
	}
}

//
func userMocker1(meter Meter, waitGroup *sync.WaitGroup) {
	for times := 0; times < 50; times++ {
		if math.Mod(float64(times), 5) == 0 {
			for index := 0; index < 100; index++ {
				meter.Inc(prometheus.Labels{LabelService: "jvessel", LabelRelease: "0.0.1", LabelSpace: "dfhciaxe"})
			}
			queryLock.Lock()
			queryArray[times] += 100
			queryLock.Unlock()
		}
		time.Sleep(time.Second)
	}
	waitGroup.Done()
}

func userMocker2(meter Meter, waitGroup *sync.WaitGroup) {
	for times := 0; times < 50; times++ {
		for index := 0; index < 10; index++ {
			meter.Inc(prometheus.Labels{LabelService: "jvessel", LabelRelease: "0.0.1", LabelSpace: "dfhciaxe"})
		}
		queryLock.Lock()
		queryArray[times] += 10
		queryLock.Unlock()

		time.Sleep(time.Second)
	}
	waitGroup.Done()
}

func userMocker3(meter Meter, waitGroup *sync.WaitGroup) {
	for times := 0; times < 50; times++ {
		if times&1 == 0 {
			meter.Add(prometheus.Labels{LabelService: "jvessel", LabelRelease: "0.0.1", LabelSpace: "dfhciaxe"}, 10.5)

			queryLock.Lock()
			queryArray[times] += 10.5
			queryLock.Unlock()
		}
		time.Sleep(time.Second)
	}
	waitGroup.Done()
}

func GetMetricFamilyOverview() {
	fmt.Println("****************************************")
	metricFamilies, err := prometheus.DefaultGatherer.Gather()
	if err != nil {
		fmt.Printf("get metric family from default gatherer failed. Error: %v\n", err)
	}
	for _, metricFamily := range metricFamilies {
		fmt.Println("MetricType: ", *metricFamily.Type, "MetricName:", metricFamily.GetName())
	}
	fmt.Println()
}
