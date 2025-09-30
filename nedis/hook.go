package nedis

import (
	"context"
	"errors"
	"net"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type hook struct {
	logger         *logrus.Logger
	InfoRecordTime time.Duration
	WarnRecordTime time.Duration
}

func (h *hook) DialHook(next redis.DialHook) redis.DialHook {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		return next(ctx, network, addr)
	}
}

func (h *hook) ProcessHook(next redis.ProcessHook) redis.ProcessHook {
	return func(ctx context.Context, cmd redis.Cmder) error {
		if h.logger == nil {
			return next(ctx, cmd)
		}

		// 执行
		start := time.Now()
		err := next(ctx, cmd)

		// 错误记录
		if err != nil && !ignoreError(err) {
			h.logger.WithContext(ctx).WithError(err).WithFields(logrus.Fields{
				"cmd": cmd.Name(),
			}).Error("执行Redis命令错误")
			return err
		}

		// 计算操作耗时
		finish := time.Now()
		cost := finish.Sub(start)

		// 记录操作耗时
		if h.WarnRecordTime >= 0 && cost >= h.WarnRecordTime {
			h.logger.WithContext(ctx).WithFields(logrus.Fields{
				"cmd":       cmd.Name(),
				"cost":      cost.String(),
				"threshold": h.WarnRecordTime.String(),
			}).Warn("执行Redis命令时长超过期望")
		} else if h.InfoRecordTime >= 0 && cost >= h.InfoRecordTime {
			h.logger.WithContext(ctx).WithFields(logrus.Fields{
				"cmd":  cmd.Name(),
				"cost": cost.String(),
			}).Info("执行Redis命令时长记录")
		}

		return err
	}
}

func (h *hook) ProcessPipelineHook(next redis.ProcessPipelineHook) redis.ProcessPipelineHook {
	return func(ctx context.Context, cmds []redis.Cmder) error {
		if h.logger == nil {
			return next(ctx, cmds)
		}

		// 执行
		start := time.Now()
		errPipeline := next(ctx, cmds)

		// 错误记录
		if errPipeline != nil && !ignoreError(errPipeline) {
			h.logger.WithContext(ctx).WithError(errPipeline).WithFields(logrus.Fields{
				"cmds": h.getCmdStringList(cmds),
			}).Error("执行Redis Pipeline命令错误")
			return errPipeline
		}
		ok := true
		for _, cmd := range cmds {
			err := cmd.Err()
			if err != nil && !ignoreError(err) {
				h.logger.WithContext(ctx).WithError(err).WithFields(logrus.Fields{
					"cmd": cmd.Name(),
				}).Error("执行Redis Pipeline命令错误")
				ok = false
			}
		}
		if !ok {
			return errPipeline
		}

		// 计算操作耗时
		finish := time.Now()
		cost := finish.Sub(start)

		// 记录操作耗时
		if h.WarnRecordTime >= 0 && cost >= h.WarnRecordTime {
			h.logger.WithContext(ctx).WithFields(logrus.Fields{
				"cmds":      h.getCmdStringList(cmds),
				"cost":      cost.String(),
				"threshold": h.WarnRecordTime.String(),
			}).Warn("执行Redis Pipeline命令时长超过期望")
		} else if h.InfoRecordTime >= 0 && cost >= h.InfoRecordTime {
			h.logger.WithContext(ctx).WithFields(logrus.Fields{
				"cmds": h.getCmdStringList(cmds),
				"cost": cost.String(),
			}).Info("执行Redis Pipeline命令时长记录")
		}

		return errPipeline
	}
}

func (h *hook) getCmdStringList(cmds []redis.Cmder) []string {
	res := make([]string, len(cmds))
	for k, cmd := range cmds {
		res[k] = cmd.String()
	}
	return res
}

func ignoreError(err error) bool {
	if errors.Is(err, redis.Nil) {
		return true
	}
	if redis.HasErrorPrefix(err, "NOSCRIPT") {
		return true
	}
	if errors.Is(err, context.Canceled) {
		return true
	}
	return false
}
