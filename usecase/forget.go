package usecase

import (
	"fmt"
	"time"

	"github.com/fillu87gyc/lambda-go/lib/zap"
	"github.com/fillu87gyc/takubo_core/domain/model"
	"github.com/fillu87gyc/takubo_core/lib"
)

func (takubo *takuboUsecase) Forget(targetWord string) error {
	ln, title, _, _ := takubo.repository.GetCurrentState()
	response, err := takubo.client.FetchWithForgetWord(ln, title, targetWord)
	if err != nil {
		zap.GetLogger().Fatal("DetectのBFFへのリクエストに失敗しました:" + err.Error())
		return err
	}
	zap.GetLogger().Info("DetectのBFFへのリクエストに成功しました:" + fmt.Sprintf("%+v", response))
	takubo.Do(response)
	if response.State == model.Forget {
		// Forgetの場合は数珠つなぎ終わりでgoroutineで処理をする
		return nil
	}
	takubo.remembered()
	return nil
}

func (takubo *takuboUsecase) forgetHint(behaviorPattern int, progressCount int) {
	zap.GetLogger().Info(fmt.Sprintf("behaviorPattern: %d, progressCount: %d", behaviorPattern, progressCount))
	switch behaviorPattern % 3 {
	case FIXEDTEXT:
		takubo.Speak("なんだっけ？")
	case INTERROGATIVE_SENTENCE:
		takubo.Speak(takubo.forgetCond.question)
	case SMALLAMOUNTS:
		count := progressCount
		ans := takubo.forgetCond.bestAnswer
		zap.GetLogger().Info(fmt.Sprintf("count: %d, ans: %s", count, ans))
		// 文字列をルーン（Unicodeコードポイント）のスライスに変換
		runes := []rune(ans)
		if count >= len(runes) {
			// 思い出す
			zap.GetLogger().Info(lib.Color("ヒントで全部でちゃった", lib.Yellow))
			takubo.Speak("思い出したー")
			takubo.remembered()
			return
		}
		// 文字列の一部分を順番に増やしながら出力
		takubo.Speak("えっとー。" + string(runes[:count]) + "。だよ")
	default:
		zap.GetLogger().Error("想定外の値が入っています: " + fmt.Sprintf("%+v", takubo.forgetCond))
	}
}
func (t *takuboUsecase) remembered() {
	close(t.forgetCond.closeChannel)
	t.SetState(model.Talking)
	time.Sleep(3 * time.Second)
	t.Talking()
}
