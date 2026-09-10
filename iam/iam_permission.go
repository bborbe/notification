// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iam

import (
	"context"

	cqrsiam "github.com/bborbe/cqrs/iam"
	"github.com/bborbe/errors"
	"github.com/bborbe/validation"
)

// ValidatePermission checks if the permission is non-empty and exists in the available permissions.
func ValidatePermission(ctx context.Context, p cqrsiam.Permission) error {
	if p == "" {
		return errors.Wrap(ctx, validation.Error, "permission is empty")
	}
	if !AvailablePermissions.Contains(p) {
		return errors.Wrap(ctx, validation.Error, "unknown permission")
	}
	return nil
}

var (
	// Backtest cqrsiam.Permissions
	CoreBacktestActiveStrategiesPermission cqrsiam.Permission = "backtest.active-strategies"
	CoreBacktestAdminPermission            cqrsiam.Permission = "backtest.admin"
	CoreBacktestCreatePermission           cqrsiam.Permission = "backtest.create"
	CoreBacktestDeleteGroupPermission      cqrsiam.Permission = "backtest.delete-group"
	CoreBacktestDeletePermission           cqrsiam.Permission = "backtest.delete"
	CoreBacktestQueuePermission            cqrsiam.Permission = "backtest.queue"
	CoreBacktestRecreateJobPermission      cqrsiam.Permission = "backtest.recreate-job"
	CoreBacktestRerunFailedPermission      cqrsiam.Permission = "backtest.rerun-failed"
	CoreBacktestFailRunningPermission      cqrsiam.Permission = "backtest.fail-running"
	CoreBacktestUpdatePermission           cqrsiam.Permission = "backtest.update"

	// Review cqrsiam.Permissions
	CoreReviewAdminPermission  cqrsiam.Permission = "review.admin"
	CoreReviewCreatePermission cqrsiam.Permission = "review.create"
	CoreReviewDeletePermission cqrsiam.Permission = "review.delete"
	CoreReviewReviewPermission cqrsiam.Permission = "review.review"

	// Strategy cqrsiam.Permissions
	CoreStrategyCreatePermission  cqrsiam.Permission = "strategy.create"
	CoreStrategyUpdatePermission  cqrsiam.Permission = "strategy.update"
	CoreStrategyDeletePermission  cqrsiam.Permission = "strategy.delete"
	CoreStrategyDisablePermission cqrsiam.Permission = "strategy.disable"
	CoreStrategyAdminPermission   cqrsiam.Permission = "strategy.admin"

	// Report cqrsiam.Permissions
	CoreReportSendPermission  cqrsiam.Permission = "report.send"
	CoreReportAdminPermission cqrsiam.Permission = "report.admin"

	// Mail cqrsiam.Permissions
	CoreMailSendPermission         cqrsiam.Permission = "mail.send"
	CoreMailRefreshTokenPermission cqrsiam.Permission = "mail.refresh-token"
	CoreMailAdminPermission        cqrsiam.Permission = "mail.admin"

	// Symbol cqrsiam.Permissions
	CoreSymbolFetchPermission  cqrsiam.Permission = "symbol.fetch"
	CoreSymbolAdminPermission  cqrsiam.Permission = "symbol.admin"
	CoreSymbolCreatePermission cqrsiam.Permission = "symbol.create"
	CoreSymbolUpdatePermission cqrsiam.Permission = "symbol.update"
	CoreSymbolDeletePermission cqrsiam.Permission = "symbol.delete"

	// MT5Account cqrsiam.Permissions
	Mt5AccountCreatePermission       cqrsiam.Permission = "mt5account.create"
	Mt5AccountUpdatePermission       cqrsiam.Permission = "mt5account.update"
	Mt5AccountDeletePermission       cqrsiam.Permission = "mt5account.delete"
	Mt5AccountAdminPermission        cqrsiam.Permission = "mt5account.admin"
	Mt5AccountSyncAccountsPermission cqrsiam.Permission = "mt5account.sync-accounts"

	// MT5Tick cqrsiam.Permissions
	Mt5TickSelectSymbolsPermission cqrsiam.Permission = "mt5tick.select-symbols"

	// ExpectedTrade cqrsiam.Permissions
	CoreExpectedTradeApprovePermission cqrsiam.Permission = "expectedtrade.approve"
	CoreExpectedTradeRejectPermission  cqrsiam.Permission = "expectedtrade.reject"
	CoreExpectedTradeClosePermission   cqrsiam.Permission = "expectedtrade.close"
	CoreExpectedTradeHedgePermission   cqrsiam.Permission = "expectedtrade.hedge"
	CoreExpectedTradeAdminPermission   cqrsiam.Permission = "expectedtrade.admin"

	// Account cqrsiam.Permissions
	CoreAccountAdminPermission  cqrsiam.Permission = "account.admin"
	CoreAccountCreatePermission cqrsiam.Permission = "account.create"
	CoreAccountUpdatePermission cqrsiam.Permission = "account.update"
	CoreAccountDeletePermission cqrsiam.Permission = "account.delete"

	// Closing cqrsiam.Permissions
	CoreClosingAdminPermission  cqrsiam.Permission = "closing.admin"
	CoreClosingCreatePermission cqrsiam.Permission = "closing.create"
	CoreClosingUpdatePermission cqrsiam.Permission = "closing.update"
	CoreClosingDeletePermission cqrsiam.Permission = "closing.delete"

	// Discord cqrsiam.Permissions
	CoreDiscordSendPermission  cqrsiam.Permission = "discord.send"
	CoreDiscordAdminPermission cqrsiam.Permission = "discord.admin"

	// News cqrsiam.Permissions
	CoreNewsUpdatePermission cqrsiam.Permission = "news.update"
	CoreNewsRatePermission   cqrsiam.Permission = "news.rate"
	CoreNewsAdminPermission  cqrsiam.Permission = "news.admin"

	// Signal cqrsiam.Permissions
	CoreSignalCreatePermission    cqrsiam.Permission = "signal.create"
	CoreSignalModifyPermission    cqrsiam.Permission = "signal.modify"
	CoreSignalClosePermission     cqrsiam.Permission = "signal.close"
	CoreSignalBreakEvenPermission cqrsiam.Permission = "signal.breakeven"
	CoreSignalAdminPermission     cqrsiam.Permission = "signal.admin"

	// ActualTrade cqrsiam.Permissions
	CoreActualTradeRetryPermission          cqrsiam.Permission = "actualtrade.retry"
	CoreActualTradeBrokerCallbackPermission cqrsiam.Permission = "actualtrade.broker-callback"
	CoreActualTradeAdminPermission          cqrsiam.Permission = "actualtrade.admin"

	// Frontend cqrsiam.Permissions
	FrontendGatewayAdminPermission cqrsiam.Permission = "frontend.gateway.admin"
)

// AvailablePermissions contains all registered permissions in the IAM system.
var AvailablePermissions = cqrsiam.Permissions{
	CoreBacktestActiveStrategiesPermission,
	CoreBacktestAdminPermission,
	CoreBacktestCreatePermission,
	CoreBacktestDeleteGroupPermission,
	CoreBacktestDeletePermission,
	CoreBacktestQueuePermission,
	CoreBacktestRecreateJobPermission,
	CoreBacktestRerunFailedPermission,
	CoreBacktestFailRunningPermission,
	CoreBacktestUpdatePermission,
	CoreReviewAdminPermission,
	CoreReviewCreatePermission,
	CoreReviewDeletePermission,
	CoreReviewReviewPermission,
	CoreStrategyCreatePermission,
	CoreStrategyUpdatePermission,
	CoreStrategyDeletePermission,
	CoreStrategyDisablePermission,
	CoreStrategyAdminPermission,
	CoreReportSendPermission,
	CoreReportAdminPermission,
	CoreMailSendPermission,
	CoreMailRefreshTokenPermission,
	CoreMailAdminPermission,
	CoreSymbolFetchPermission,
	CoreSymbolAdminPermission,
	CoreSymbolCreatePermission,
	CoreSymbolUpdatePermission,
	CoreSymbolDeletePermission,
	Mt5AccountCreatePermission,
	Mt5AccountUpdatePermission,
	Mt5AccountDeletePermission,
	Mt5AccountAdminPermission,
	Mt5AccountSyncAccountsPermission,
	Mt5TickSelectSymbolsPermission,
	CoreExpectedTradeApprovePermission,
	CoreExpectedTradeRejectPermission,
	CoreExpectedTradeClosePermission,
	CoreExpectedTradeHedgePermission,
	CoreExpectedTradeAdminPermission,
	CoreAccountAdminPermission,
	CoreAccountCreatePermission,
	CoreAccountUpdatePermission,
	CoreAccountDeletePermission,
	CoreClosingAdminPermission,
	CoreClosingCreatePermission,
	CoreClosingUpdatePermission,
	CoreClosingDeletePermission,
	CoreDiscordSendPermission,
	CoreDiscordAdminPermission,
	CoreNewsUpdatePermission,
	CoreNewsRatePermission,
	CoreNewsAdminPermission,
	CoreSignalCreatePermission,
	CoreSignalModifyPermission,
	CoreSignalClosePermission,
	CoreSignalBreakEvenPermission,
	CoreSignalAdminPermission,
	CoreActualTradeRetryPermission,
	CoreActualTradeBrokerCallbackPermission,
	CoreActualTradeAdminPermission,
	FrontendGatewayAdminPermission,
}
