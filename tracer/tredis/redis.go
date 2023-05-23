package tredis

import (
	"context"
	"strings"

	"github.com/go-redis/redis"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// WrapRedisClient adds opentracing measurements for commands and returns cloned client
func WrapRedisClient(ctx context.Context, client *redis.Client) *redis.Client {
	if ctx == nil {
		return client
	}

	// 获取ctx携带的span, 找不到返回nil
	parentSpan := opentracing.SpanFromContext(ctx)
	if parentSpan == nil {
		return client.WithContext(ctx)
	}

	// ctxClient 复制一个携带ctx的redis客户端
	ctxClient := client.WithContext(ctx)
	// opts 只读属性
	opts := ctxClient.Options()
	// WrapProcess 包装处理 Redis 命令的函数。
	ctxClient.WrapProcess(process(parentSpan, opts))
	// WrapProcessPipeline 批量包装处理 Redis 命令的函数。
	ctxClient.WrapProcessPipeline(processPipeline(parentSpan, opts))
	return ctxClient
}

// process 扩展 Redis 命令, 添加可观测性功能
func process(parentSpan opentracing.Span, opts *redis.Options) func(oldProcess func(cmd redis.Cmder) error) func(cmd redis.Cmder) error {
	return func(oldProcess func(cmd redis.Cmder) error) func(cmd redis.Cmder) error {
		return func(cmd redis.Cmder) error {
			dbMethod := formatCommandAsDbMethod(cmd)
			span := getSpan(parentSpan, opts, "redis-cmd", dbMethod)
			defer span.Finish()
			return oldProcess(cmd)
		}
	}
}

// processPipeline 批量扩展 Redis 命令, 添加可观测性功能
func processPipeline(parentSpan opentracing.Span, opts *redis.Options) func(oldProcess func(cmds []redis.Cmder) error) func(cmds []redis.Cmder) error {
	return func(oldProcess func(cmds []redis.Cmder) error) func(cmds []redis.Cmder) error {
		return func(cmds []redis.Cmder) error {
			dbMethod := formatCommandsAsDbMethods(cmds)
			span := getSpan(parentSpan, opts, "redis-pipeline-cmd", dbMethod)
			defer span.Finish()
			return oldProcess(cmds)
		}
	}
}

// formatCommandAsDbMethod 返回 Redis 命令的名称
func formatCommandAsDbMethod(cmd redis.Cmder) string {
	return cmd.Name()
}

// formatCommandsAsDbMethods 返回 Redis 命令的名称列表
func formatCommandsAsDbMethods(cmds []redis.Cmder) string {
	cmdsAsDbMethods := make([]string, len(cmds))
	for i, cmd := range cmds {
		dbMethod := formatCommandAsDbMethod(cmd)
		cmdsAsDbMethods[i] = dbMethod
	}
	return strings.Join(cmdsAsDbMethods, " -> ")
}

// getSpan 根据父span创建一个span, 并添加额外标签
func getSpan(parentSpan opentracing.Span, opts *redis.Options, operationName, dbMethod string) opentracing.Span {
	// tracer 提供对创建此 Span 的 Tracer 的访问
	tracer := parentSpan.Tracer()
	// 创建一个span
	span := tracer.StartSpan(operationName, opentracing.ChildOf(parentSpan.Context()))
	// 扩展 span 属性, 设置数据库tag
	ext.DBType.Set(span, "redis")
	// 扩展 span 属性, 设置数据库连接串
	ext.PeerAddress.Set(span, opts.Addr)
	// 扩展 span 属性, 设置类型为客户端
	ext.SpanKind.Set(span, ext.SpanKindEnum("client"))
	// 扩展 span 属性, 设置数据库方法的名称
	span.SetTag("db.method", dbMethod)
	return span
}
