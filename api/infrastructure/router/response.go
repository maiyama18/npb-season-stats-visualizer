package router

import (
	"fmt"

	"github.com/mui87/npb-season-stats-visualizer/domain"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type PlayerSearchResponse struct {
	Players []Player `json:"players"`
}

type PitcherStatsResponse struct {
	Player Player       `json:"player"`
	Stats  PitcherStats `json:"stats"`
}

type BatterStatsResponse struct {
	Player Player      `json:"player"`
	Stats  BatterStats `json:"stats"`
}

type Player struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PitcherStats struct {
	Game             IntStat   `json:"game"`
	Era              FloatStat `json:"era"`
	GameStart        IntStat   `json:"gameStart"`
	Complete         IntStat   `json:"complete"`
	ShutOut          IntStat   `json:"shutOut"`
	QualityStart     IntStat   `json:"qualityStart"`
	Win              IntStat   `json:"win"`
	Lose             IntStat   `json:"lose"`
	Hold             IntStat   `json:"hold"`
	HoldPoint        IntStat   `json:"holdPoint"`
	Save             IntStat   `json:"save"`
	WinPercent       FloatStat `json:"winPercent"`
	Inning           FloatStat `json:"inning"`
	Hit              IntStat   `json:"hit"`
	HomeRun          IntStat   `json:"homeRun"`
	StrikeOut        IntStat   `json:"strikeOut"`
	StrikeOutPercent FloatStat `json:"strikeOutPercent"`
	Walk             IntStat   `json:"walk"`
	HitByPitch       IntStat   `json:"hitByPitch"`
	WildPitch        IntStat   `json:"wildPitch"`
	Balk             IntStat   `json:"balk"`
	Run              IntStat   `json:"run"`
	EarnedRun        IntStat   `json:"earnedRun"`
	Average          FloatStat `json:"average"`
	Kbb              FloatStat `json:"kbb"`
	Whip             FloatStat `json:"whip"`
}

type BatterStats struct {
	Game                       IntStat   `json:"game"`
	Average                    FloatStat `json:"average"`
	PlateAppearance            IntStat   `json:"plateAppearance"`
	AtBat                      IntStat   `json:"atBat"`
	Hit                        IntStat   `json:"hit"`
	Double                     IntStat   `json:"double"`
	Triple                     IntStat   `json:"triple"`
	HomeRun                    IntStat   `json:"homeRun"`
	TotalBase                  IntStat   `json:"totalBase"`
	RunBattedIn                IntStat   `json:"runBattedIn"`
	Run                        IntStat   `json:"run"`
	StrikeOut                  IntStat   `json:"strikeOut"`
	Walk                       IntStat   `json:"walk"`
	HitByPitch                 IntStat   `json:"hitByPitch"`
	Sacrifice                  IntStat   `json:"sacrifice"`
	SacrificeFly               IntStat   `json:"sacrificeFly"`
	StolenBase                 IntStat   `json:"stolenBase"`
	CaughtStealing             IntStat   `json:"caughtStealing"`
	DoublePlay                 IntStat   `json:"doublePlay"`
	OnBasePercent              FloatStat `json:"onBasePercent"`
	SluggingPercent            FloatStat `json:"sluggingPercent"`
	Ops                        FloatStat `json:"ops"`
	AverageWithScoringPosition FloatStat `json:"averageWithScoringPosition"`
	Error                      IntStat   `json:"error"`
}

type IntStat struct {
	Dates  []string `json:"dates"`
	Values []int    `json:"values"`
}

func newIntStat() IntStat {
	return IntStat{
		Dates:  make([]string, 0),
		Values: make([]int, 0),
	}
}

func (is *IntStat) addDataPoint(date string, value int) {
	is.Dates = append(is.Dates, date)
	is.Values = append(is.Values, value)
}

type FloatStat struct {
	Dates  []string  `json:"dates"`
	Values []float64 `json:"values"`
}

func newFloatStat() FloatStat {
	return FloatStat{
		Dates:  make([]string, 0),
		Values: make([]float64, 0),
	}
}

func (fs *FloatStat) addDataPoint(date string, value float64) {
	fs.Dates = append(fs.Dates, date)
	fs.Values = append(fs.Values, value)
}

func constructErrorResponse(format string, a ...interface{}) ErrorResponse {
	errMsg := fmt.Sprintf(format, a...)
	return ErrorResponse{Error: errMsg}
}

func constructPlayerSearchResponseFromPitchers(pitchers []domain.Pitcher) PlayerSearchResponse {
	players := make([]Player, 0)
	for _, p := range pitchers {
		players = append(players, Player{ID: p.ID, Name: p.Name})
	}

	return PlayerSearchResponse{Players: players}
}

func constructPlayerSearchResponseFromBatters(batters []domain.Batter) PlayerSearchResponse {
	players := make([]Player, 0)
	for _, p := range batters {
		players = append(players, Player{ID: p.ID, Name: p.Name})
	}

	return PlayerSearchResponse{Players: players}
}

func constructPitcherStatsResponse(pitcher domain.Pitcher) PitcherStatsResponse {
	player := Player{ID: pitcher.ID, Name: pitcher.Name}

	game := newIntStat()
	gameStart := newIntStat()
	complete := newIntStat()
	shutOut := newIntStat()
	qualityStart := newIntStat()
	win := newIntStat()
	lose := newIntStat()
	hold := newIntStat()
	holdPoint := newIntStat()
	save := newIntStat()
	hit := newIntStat()
	homeRun := newIntStat()
	strikeOut := newIntStat()
	walk := newIntStat()
	hitByPitch := newIntStat()
	wildPitch := newIntStat()
	balk := newIntStat()
	run := newIntStat()
	earnedRun := newIntStat()

	era := newFloatStat()
	winPercent := newFloatStat()
	inning := newFloatStat()
	strikeOutPercent := newFloatStat()
	average := newFloatStat()
	kbb := newFloatStat()
	whip := newFloatStat()

	for _, pStats := range pitcher.PitcherStatsList {
		date := pStats.Date.Format("2006-01-02")

		if pStats.Game != nil {
			game.addDataPoint(date, *pStats.Game)
		}
		if pStats.Era != nil {
			era.addDataPoint(date, *pStats.Era)
		}
		if pStats.GameStart != nil {
			gameStart.addDataPoint(date, *pStats.GameStart)
		}
		if pStats.Complete != nil {
			complete.addDataPoint(date, *pStats.Complete)
		}
		if pStats.ShutOut != nil {
			shutOut.addDataPoint(date, *pStats.ShutOut)
		}
		if pStats.QualityStart != nil {
			qualityStart.addDataPoint(date, *pStats.QualityStart)
		}
		if pStats.Win != nil {
			win.addDataPoint(date, *pStats.Win)
		}
		if pStats.Lose != nil {
			lose.addDataPoint(date, *pStats.Lose)
		}
		if pStats.Hold != nil {
			hold.addDataPoint(date, *pStats.Hold)
		}
		if pStats.HoldPoint != nil {
			holdPoint.addDataPoint(date, *pStats.HoldPoint)
		}
		if pStats.Save != nil {
			save.addDataPoint(date, *pStats.Save)
		}
		if pStats.WinPercent != nil {
			winPercent.addDataPoint(date, *pStats.WinPercent)
		}
		if pStats.Inning != nil {
			inning.addDataPoint(date, *pStats.Inning)
		}
		if pStats.Hit != nil {
			hit.addDataPoint(date, *pStats.Hit)
		}
		if pStats.HomeRun != nil {
			homeRun.addDataPoint(date, *pStats.HomeRun)
		}
		if pStats.StrikeOut != nil {
			strikeOut.addDataPoint(date, *pStats.StrikeOut)
		}
		if pStats.StrikeOutPercent != nil {
			strikeOutPercent.addDataPoint(date, *pStats.StrikeOutPercent)
		}
		if pStats.Walk != nil {
			walk.addDataPoint(date, *pStats.Walk)
		}
		if pStats.HitByPitch != nil {
			hitByPitch.addDataPoint(date, *pStats.HitByPitch)
		}
		if pStats.WildPitch != nil {
			wildPitch.addDataPoint(date, *pStats.WildPitch)
		}
		if pStats.Balk != nil {
			balk.addDataPoint(date, *pStats.Balk)
		}
		if pStats.Run != nil {
			run.addDataPoint(date, *pStats.Run)
		}
		if pStats.EarnedRun != nil {
			earnedRun.addDataPoint(date, *pStats.EarnedRun)
		}
		if pStats.Average != nil {
			average.addDataPoint(date, *pStats.Average)
		}
		if pStats.Kbb != nil {
			kbb.addDataPoint(date, *pStats.Kbb)
		}
		if pStats.Whip != nil {
			whip.addDataPoint(date, *pStats.Whip)
		}
	}

	stats := PitcherStats{
		Game:             game,
		Era:              era,
		GameStart:        gameStart,
		Complete:         complete,
		ShutOut:          shutOut,
		QualityStart:     qualityStart,
		Win:              win,
		Lose:             lose,
		Hold:             hold,
		HoldPoint:        holdPoint,
		Save:             save,
		WinPercent:       winPercent,
		Inning:           inning,
		Hit:              hit,
		HomeRun:          homeRun,
		StrikeOut:        strikeOut,
		StrikeOutPercent: strikeOutPercent,
		Walk:             walk,
		HitByPitch:       hitByPitch,
		WildPitch:        wildPitch,
		Balk:             balk,
		Run:              run,
		EarnedRun:        earnedRun,
		Average:          average,
		Kbb:              kbb,
		Whip:             whip,
	}

	return PitcherStatsResponse{Player: player, Stats: stats}
}

func constructBatterStatsResponse(batter domain.Batter) BatterStatsResponse {
	player := Player{ID: batter.ID, Name: batter.Name}

	//var game, plateAppearance, atBat, hit, double, triple, homeRun, totalBase, runBattedIn, run, strikeOut, walk, hitByPitch, sacrifice, sacrificeFly,
	//	stolenBase, caughtStealing, doublePlay, errorCount IntStat
	//var average, onBasePercent, sluggingPercent, ops, averageWithScoringPosition FloatStat

	game := newIntStat()
	plateAppearance := newIntStat()
	atBat := newIntStat()
	hit := newIntStat()
	double := newIntStat()
	triple := newIntStat()
	homeRun := newIntStat()
	totalBase := newIntStat()
	runBattedIn := newIntStat()
	run := newIntStat()
	strikeOut := newIntStat()
	walk := newIntStat()
	hitByPitch := newIntStat()
	sacrifice := newIntStat()
	sacrificeFly := newIntStat()
	stolenBase := newIntStat()
	caughtStealing := newIntStat()
	doublePlay := newIntStat()
	errorCount := newIntStat()

	average := newFloatStat()
	onBasePercent := newFloatStat()
	sluggingPercent := newFloatStat()
	ops := newFloatStat()
	averageWithScoringPosition := newFloatStat()

	for _, bStats := range batter.BatterStatsList {
		date := bStats.Date.Format("2006-01-02")

		if bStats.Game != nil {
			game.addDataPoint(date, *bStats.Game)
		}
		if bStats.Average != nil {
			average.addDataPoint(date, *bStats.Average)
		}
		if bStats.PlateAppearance != nil {
			plateAppearance.addDataPoint(date, *bStats.PlateAppearance)
		}
		if bStats.AtBat != nil {
			atBat.addDataPoint(date, *bStats.AtBat)
		}
		if bStats.Hit != nil {
			hit.addDataPoint(date, *bStats.Hit)
		}
		if bStats.Double != nil {
			double.addDataPoint(date, *bStats.Double)
		}
		if bStats.Triple != nil {
			triple.addDataPoint(date, *bStats.Triple)
		}
		if bStats.HomeRun != nil {
			homeRun.addDataPoint(date, *bStats.HomeRun)
		}
		if bStats.TotalBase != nil {
			totalBase.addDataPoint(date, *bStats.TotalBase)
		}
		if bStats.RunBattedIn != nil {
			runBattedIn.addDataPoint(date, *bStats.RunBattedIn)
		}
		if bStats.Run != nil {
			run.addDataPoint(date, *bStats.Run)
		}
		if bStats.StrikeOut != nil {
			strikeOut.addDataPoint(date, *bStats.StrikeOut)
		}
		if bStats.Walk != nil {
			walk.addDataPoint(date, *bStats.Walk)
		}
		if bStats.HitByPitch != nil {
			hitByPitch.addDataPoint(date, *bStats.HitByPitch)
		}
		if bStats.Sacrifice != nil {
			sacrifice.addDataPoint(date, *bStats.Sacrifice)
		}
		if bStats.SacrificeFly != nil {
			sacrificeFly.addDataPoint(date, *bStats.SacrificeFly)
		}
		if bStats.StolenBase != nil {
			stolenBase.addDataPoint(date, *bStats.StolenBase)
		}
		if bStats.CaughtStealing != nil {
			caughtStealing.addDataPoint(date, *bStats.CaughtStealing)
		}
		if bStats.DoublePlay != nil {
			doublePlay.addDataPoint(date, *bStats.DoublePlay)
		}
		if bStats.OnBasePercent != nil {
			onBasePercent.addDataPoint(date, *bStats.OnBasePercent)
		}
		if bStats.SluggingPercent != nil {
			sluggingPercent.addDataPoint(date, *bStats.SluggingPercent)
		}
		if bStats.Ops != nil {
			ops.addDataPoint(date, *bStats.Ops)
		}
		if bStats.AverageWithScoringPosition != nil {
			averageWithScoringPosition.addDataPoint(date, *bStats.AverageWithScoringPosition)
		}
		if bStats.Error != nil {
			errorCount.addDataPoint(date, *bStats.Error)
		}
	}

	stats := BatterStats{
		Game:                       game,
		Average:                    average,
		PlateAppearance:            plateAppearance,
		AtBat:                      atBat,
		Hit:                        hit,
		Double:                     double,
		Triple:                     triple,
		HomeRun:                    homeRun,
		TotalBase:                  totalBase,
		RunBattedIn:                runBattedIn,
		Run:                        run,
		StrikeOut:                  strikeOut,
		Walk:                       walk,
		HitByPitch:                 hitByPitch,
		Sacrifice:                  sacrifice,
		SacrificeFly:               sacrificeFly,
		StolenBase:                 stolenBase,
		CaughtStealing:             caughtStealing,
		DoublePlay:                 doublePlay,
		OnBasePercent:              onBasePercent,
		SluggingPercent:            sluggingPercent,
		Ops:                        ops,
		AverageWithScoringPosition: averageWithScoringPosition,
		Error:                      errorCount,
	}

	return BatterStatsResponse{Player: player, Stats: stats}
}
