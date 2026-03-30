package main

import (
	"context"
	"fmt"
	"strings"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) BuildStudyPlan(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		name = "Go 同学"
	}

	return fmt.Sprintf("%s，建议按 Gin -> gRPC -> Cobra -> GORM -> Testify 的顺序学习，每个 demo 都先运行，再尝试自己改一版。", name)
}
