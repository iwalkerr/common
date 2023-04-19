package toolx

import "time"

type TimerFunc func(any) bool

/**
 * 定时调用
 * @delay 首次延时
 * @tick  间隔
 * @fun   定时执行function
 * @param fun参数
 */
func Timer(tick time.Duration, fun TimerFunc, param any, funcDefer TimerFunc, paramDefer any) {
	go func() {
		defer func() {
			if funcDefer != nil {
				funcDefer(paramDefer)
			}
		}()

		if fun == nil {
			return
		}
		t := time.NewTimer(2 * time.Second)
		for {
			select {
			case <-t.C:
				if !fun(param) {
					return
				}
				t.Reset(tick)
			}
		}

		// for range t.C {
		// 	if !fun(param) {
		// 		return
		// 	}
		// 	t.Reset(tick)
		// }
	}()
}
