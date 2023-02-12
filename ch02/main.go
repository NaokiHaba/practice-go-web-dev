package main

import (
	"context"
	"fmt"
	"time"
)

func child(ctx context.Context) {
	// ctx.Err()でキャンセルされているかどうかを確認する
	if err := ctx.Err(); err != nil {
		return
	}
	fmt.Println("キャンセルされていない")
}

// WithTimeout タイムアウトを設定する
func WithTimeout(ctx context.Context, timeout time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, timeout)
}

// WithDeadline タイムアウトを設定する
func WithDeadline(ctx context.Context, d time.Time) (context.Context, context.CancelFunc) {
	// context.WithDeadlineでキャンセル可能なコンテキストを作成する
	return context.WithDeadline(ctx, d)
}

type TraceID string

const ZeroTraceID TraceID = ""

type TraceIDKey struct{}

func SetTraceID(ctx context.Context, traceID TraceID) context.Context {
	// context.WithValue: Contextに値を設定する
	// (ctx, TraceIDKey{}, traceID): ContextにTraceIDKey{}というキーでtraceIDという値を設定する
	return context.WithValue(ctx, TraceIDKey{}, traceID)
}

func GetTraceID(ctx context.Context) TraceID {
	// ctx.Value: Contextから値を取得する
	if traceID, ok := ctx.Value(TraceIDKey{}).(TraceID); ok {
		return traceID
	}
	return ZeroTraceID
}

func main() {
	// Section 05 キャンセルを通知する
	/**
	context.WithCancel: キャンセル可能なコンテキストを作成する
	context.Background: 親コンテキストを作成する
	ctx, cancel := context.WithCancel(context.Background())

	キャンセルされていない
	child(ctx)

	cancel(): キャンセルを通知する
	cancel()

	nilが返ってくる
	child(ctx)
	*/

	// Section 06 キャンセルされるまで待つ
	/**
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	// defer: 関数の最後に実行される
	defer cancel()

	// go: ゴルーチンを起動する
	go func() {
		fmt.Println("別ゴルーチン")
	}()

	fmt.Println("Stop")

	// <-: 受信専用チャネルを受信する
	// ctx.Done(): キャンセル通知を受信するチャネルを返す
	<-ctx.Done()
	fmt.Println("Start")
	*/

	// Section 07 キャンセル通知まで別処理を繰り返す
	/**
	ctx, cancel := context.WithCancel(context.Background())

	// make: チャネルを作成する make(chan 型)
	task := make(chan int)

	// ゴルーチンを起動する
	go func() {
		for {
			select {
			// <-ctx.Done(): キャンセル通知を受信する
			case <-ctx.Done():
				fmt.Println("キャンセルされた")
				return
				// <-task: taskから値を受信する
			case i := <-task:
				fmt.Println("get", i)
			default:
				fmt.Println("キャンセルされていない")
			}
			time.Sleep(time.Millisecond * 300)
		}
	}()

	time.Sleep(time.Second)

	for i := 0; i < 10; i++ {
		// task <- i: iをtaskに送信する
		task <- i
	}
	cancel()
	*/

	// Contextに値を設定する
	ctx := context.Background()
	fmt.Println("TraceID:", GetTraceID(ctx))
	ctx = SetTraceID(ctx, "12345")
	fmt.Println("TraceID:", GetTraceID(ctx))
}
