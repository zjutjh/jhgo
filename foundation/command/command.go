package command

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
	"time"

	"github.com/spf13/cobra"

	"github.com/zjutjh/mygo/config"
	"github.com/zjutjh/mygo/foundation/kernel"
	"github.com/zjutjh/mygo/nlog"
)

var cfgPath string

var boot func() kernel.BootList
var defaultRun func(cmd *cobra.Command, args []string) error
var runningCommand *cobra.Command

var once sync.Once
var root = &cobra.Command{
	Use: "app",
	Run: func(cmd *cobra.Command, args []string) {
		if runningCommand == nil || (runningCommand.Use != "app" && runningCommand.Use != "server") {
			runningCommand = cmd
		}
		kernel.Bootstrap(cfgPath, boot)
		Run(defaultRun, cmd, args)
	},
}

func init() {
	root.PersistentFlags().StringVar(&cfgPath, "config", "conf/", "config path(default is conf/)")
}

// Execute 应用程序执行主入口
func Execute(b func() kernel.BootList, rc func(*cobra.Command), dr func(cmd *cobra.Command, args []string) error) {
	once.Do(func() {
		// 设置引导器与默认运行器
		boot = b
		defaultRun = dr

		// 注册业务命令
		rc(root)

		// 执行
		if err := root.Execute(); err != nil {
			fmt.Fprintln(os.Stdout, "执行命令错误", err)
			os.Exit(1)
		}
	})
}

// Add 注册一个命令
func Add(key string, runner func(cmd *cobra.Command, args []string) error) {
	cmd := &cobra.Command{
		Use:   key,
		Short: fmt.Sprintf("运行命令[%s]", key),
		Long:  fmt.Sprintf("运行命令[%s]", key),
		Run: func(cmd *cobra.Command, args []string) {
			if runningCommand == nil || (runningCommand.Use != "app" && runningCommand.Use != "server") {
				runningCommand = cmd
			}
			kernel.Bootstrap(cfgPath, boot)
			Run(runner, cmd, args)
		},
	}
	root.AddCommand(cmd)
}

// Run 运行命令
func Run(runner func(cmd *cobra.Command, args []string) error, cmd *cobra.Command, args []string) {
	// 初始化配置和日志实例
	conf := DefaultConfig
	err := config.Pick().UnmarshalKey("command", &conf)
	if err != nil {
		fmt.Fprintln(os.Stdout, "初始化命令配置错误", err)
		os.Exit(1)
	}
	logger := nlog.Pick(conf.Logger)

	// 启动pprof
	if conf.PprofSwitch {
		for _, t := range conf.PprofType {
			switch t {
			case "cpu":
				w, err := os.OpenFile(conf.PprofOutput+cmd.Use+".run.cpu", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
				if err != nil {
					fmt.Fprintln(os.Stdout, "处理命令CPU pprof错误", err)
					os.Exit(1)
				}
				err = pprof.StartCPUProfile(w)
				if err != nil {
					fmt.Fprintln(os.Stdout, "启动命令CPU pprof错误", err)
					os.Exit(1)
				}
				defer pprof.StopCPUProfile()
			default:
				w, err := os.OpenFile(conf.PprofOutput+cmd.Use+".run."+t, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
				if err != nil {
					fmt.Fprintln(os.Stdout, fmt.Sprintf("处理命令%s pprof错误", t), err)
					os.Exit(1)
				}
				defer w.Close()
				defer pprof.Lookup(t).WriteTo(w, 0)
			}
		}
	}

	// 标记开始时间
	start := time.Now()

	// 处理panic
	defer func() {
		if pnc := recover(); pnc != nil {
			if conf.Output {
				fmt.Fprintf(os.Stdout, "命令[%s]发生panic, 耗时[%s]\n", cmd.Use, time.Since(start).String())
			}
			logger.Errorf("命令[%s]发生panic, 耗时[%s]", cmd.Use, time.Since(start).String())
			panic(pnc)
		}
	}()

	// 声明开始执行信息
	if conf.Output {
		fmt.Fprintf(os.Stdout, "命令[%s]开始执行\n", cmd.Use)
	}
	logger.WithField("args", args).Infof("命令[%s]开始执行", cmd.Use)

	// 执行命名逻辑
	err = runner(cmd, args)

	// 声明执行结果
	if err == nil {
		if conf.Output {
			fmt.Fprintf(os.Stdout, "命令[%s]执行成功, 耗时[%s]\n", cmd.Use, time.Since(start).String())
		}
		logger.Infof("命令[%s]执行成功, 耗时[%s]", cmd.Use, time.Since(start).String())
	} else {
		if conf.Output {
			fmt.Fprintf(os.Stdout, "命令[%s]发生错误, 耗时[%s], 错误: %s\n", cmd.Use, time.Since(start).String(), err.Error())
		}
		logger.WithError(err).Errorf("命令[%s]发生错误, 耗时[%s]", cmd.Use, time.Since(start).String())
	}
}

// GetRunCommand 对外暴露正在运行的 command 信息
func GetRunCommand() *cobra.Command {
	return runningCommand
}
