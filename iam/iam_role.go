// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iam

import cqrsiam "github.com/bborbe/cqrs/iam"

var (
	// BacktestAdmin provides full administrative access to backtest operations.
	// Intended for: Core administrators and the backtest controller service.
	// cqrsiam.Permissions: All backtest operations including admin-level functions.
	BacktestAdmin = cqrsiam.NewRole("BacktestAdmin",
		CoreBacktestActiveStrategiesPermission,
		CoreBacktestAdminPermission,
		CoreBacktestCreatePermission,
		CoreBacktestDeletePermission,
		CoreBacktestDeleteGroupPermission,
		CoreBacktestQueuePermission,
		CoreBacktestRecreateJobPermission,
		CoreBacktestRerunFailedPermission,
		CoreBacktestFailRunningPermission,
		CoreBacktestUpdatePermission,
	)
	// BacktestCron provides minimal permissions for scheduled backtest operations.
	// Intended for: Cron jobs and automated strategy queuing services.
	// cqrsiam.Permissions: Only access to active strategies for automated queuing.
	BacktestCron = cqrsiam.NewRole("BacktestUser",
		CoreBacktestActiveStrategiesPermission,
	)
	// BacktestUser provides standard backtest operations for regular users and services.
	// Intended for: API users, business users, and backtest agent services.
	// cqrsiam.Permissions: Create, delete, update, and queue backtests (excludes admin functions).
	BacktestUser = cqrsiam.NewRole("BacktestUser",
		CoreBacktestCreatePermission,
		CoreBacktestDeletePermission,
		CoreBacktestDeleteGroupPermission,
		CoreBacktestQueuePermission,
		CoreBacktestUpdatePermission,
	)
	// BacktestViewer provides read-only access to backtest data.
	// Intended for: Users who need to view but not modify backtest results.
	// cqrsiam.Permissions: None (read-only access through other mechanisms).
	BacktestViewer = cqrsiam.NewRole("BacktestViewer")
	// BacktestCreator provides minimal permissions for services that only create backtests.
	// Intended for: Automated services like core-backtest-agent-queue that trigger backtests.
	// cqrsiam.Permissions: Only backtest creation (follows principle of least privilege).
	BacktestCreator = cqrsiam.NewRole("BacktestCreator",
		CoreBacktestCreatePermission,
	)
	// ReviewAdmin provides full administrative access to review operations.
	// Intended for: Core administrators and the review service.
	// cqrsiam.Permissions: All review operations including admin functions.
	ReviewAdmin = cqrsiam.NewRole("ReviewAdmin",
		CoreReviewAdminPermission,
		CoreReviewCreatePermission,
		CoreReviewDeletePermission,
		CoreReviewReviewPermission,
	)
	// ReviewUser provides standard review operations for regular users.
	// Intended for: Business users who can participate in review processes.
	// cqrsiam.Permissions: Review and comment on items (excludes admin functions).
	ReviewUser = cqrsiam.NewRole("ReviewUser",
		CoreReviewReviewPermission,
	)
	// ReviewViewer provides read-only access to review data.
	// Intended for: Users who need to view but not participate in reviews.
	// cqrsiam.Permissions: None (read-only access through other mechanisms).
	ReviewViewer = cqrsiam.NewRole("ReviewViewer")

	// StrategyAdmin provides full administrative access to strategy operations.
	// Intended for: Core administrators and the strategy service.
	// cqrsiam.Permissions: All strategy operations including admin-level functions.
	StrategyAdmin = cqrsiam.NewRole("StrategyAdmin",
		CoreStrategyCreatePermission,
		CoreStrategyUpdatePermission,
		CoreStrategyDeletePermission,
		CoreStrategyDisablePermission,
		CoreStrategyAdminPermission,
	)
	// StrategyUser provides standard strategy operations for regular users and services.
	// Intended for: API users and business users managing strategies.
	// cqrsiam.Permissions: Create, update, and disable strategies (excludes delete and admin functions).
	StrategyUser = cqrsiam.NewRole("StrategyUser",
		CoreStrategyCreatePermission,
		CoreStrategyUpdatePermission,
		CoreStrategyDisablePermission,
	)
	// StrategyViewer provides read-only access to strategy data.
	// Intended for: Users who need to view but not modify strategies.
	// cqrsiam.Permissions: None (read-only access through other mechanisms).
	StrategyViewer = cqrsiam.NewRole("StrategyViewer")

	// SymbolAdmin provides full administrative access to symbol operations.
	// Intended for: Core administrators and symbol services.
	// cqrsiam.Permissions: All symbol operations including admin-level functions.
	SymbolAdmin = cqrsiam.NewRole("SymbolAdmin",
		CoreSymbolFetchPermission,
		CoreSymbolAdminPermission,
		CoreSymbolCreatePermission,
		CoreSymbolUpdatePermission,
		CoreSymbolDeletePermission,
	)
	// SymbolUser provides standard symbol operations for regular users and services.
	// Intended for: API users and business users fetching market information.
	// cqrsiam.Permissions: Fetch market details (excludes admin functions).
	SymbolUser = cqrsiam.NewRole("SymbolUser",
		CoreSymbolFetchPermission,
		CoreSymbolCreatePermission,
		CoreSymbolUpdatePermission,
	)

	// ReportAdmin provides full administrative access to report operations.
	// Intended for: Core administrators who can trigger and manage reports.
	// cqrsiam.Permissions: All report operations including sending reports.
	ReportAdmin = cqrsiam.NewRole("ReportAdmin",
		CoreReportSendPermission,
		CoreReportAdminPermission,
		CoreMailSendPermission,
	)
	// ReportCron provides minimal permissions for scheduled report operations.
	// Intended for: Cron jobs that send automated trading reports.
	// cqrsiam.Permissions: Report sending and mail sending capabilities.
	ReportCron = cqrsiam.NewRole("ReportCron",
		CoreReportSendPermission,
		CoreMailSendPermission,
	)

	// MailRefresh provides permissions for OAuth2 token refresh and mail sending operations.
	// Intended for: CronJob that refreshes OAuth2 tokens and API users who can send mails.
	// cqrsiam.Permissions: OAuth2 token refresh and mail sending capabilities.
	MailRefresh = cqrsiam.NewRole("MailRefresh",
		CoreMailRefreshTokenPermission,
		CoreMailSendPermission,
	)

	// Mt5AccountAdmin provides full administrative access to MT5 account operations.
	// Intended for: Core administrators and MT5 account controller service.
	// cqrsiam.Permissions: All MT5 account operations including admin functions.
	Mt5AccountAdmin = cqrsiam.NewRole("Mt5AccountAdmin",
		Mt5AccountAdminPermission,
		Mt5AccountCreatePermission,
		Mt5AccountUpdatePermission,
		Mt5AccountDeletePermission,
	)
	// Mt5AccountUser provides standard MT5 account operations for converters and services.
	// Intended for: MT5 converters and automated services that process account data.
	// cqrsiam.Permissions: Create and update MT5 accounts (not delete - admin only).
	Mt5AccountUser = cqrsiam.NewRole("Mt5AccountUser",
		Mt5AccountCreatePermission,
		Mt5AccountUpdatePermission,
	)
	// Mt5AccountViewer provides read-only access to MT5 account data.
	// Intended for: Users who need to view but not modify MT5 accounts (e.g., monitoring services).
	// cqrsiam.Permissions: None (read-only access through other mechanisms).
	Mt5AccountViewer = cqrsiam.NewRole("Mt5AccountViewer")

	// ExpectedTradeAdmin provides full administrative access to expected trade operations.
	// Intended for: Core administrators and the expected trade controller service.
	// cqrsiam.Permissions: All expected trade operations including admin functions.
	ExpectedTradeAdmin = cqrsiam.NewRole("ExpectedTradeAdmin",
		CoreExpectedTradeAdminPermission,
		CoreExpectedTradeApprovePermission,
		CoreExpectedTradeRejectPermission,
		CoreExpectedTradeClosePermission,
		CoreExpectedTradeHedgePermission,
	)
	// ExpectedTradeUser provides standard expected trade operations for traders.
	// Intended for: Traders who can approve and reject expected trades.
	// cqrsiam.Permissions: Approve and reject expected trades.
	ExpectedTradeUser = cqrsiam.NewRole("ExpectedTradeUser",
		CoreExpectedTradeApprovePermission,
		CoreExpectedTradeRejectPermission,
		CoreExpectedTradeClosePermission,
		CoreExpectedTradeHedgePermission,
	)
	// ExpectedTradeViewer provides read-only access to expected trade data.
	// Intended for: Users who need to view but not modify expected trades (e.g., monitoring services).
	// cqrsiam.Permissions: None (read-only access through other mechanisms).
	ExpectedTradeViewer = cqrsiam.NewRole("ExpectedTradeViewer")

	// AccountAdmin provides full administrative access to account operations.
	// Intended for: Core administrators and the account controller service.
	// cqrsiam.Permissions: All account operations including admin-level functions.
	AccountAdmin = cqrsiam.NewRole("AccountAdmin",
		CoreAccountAdminPermission,
		CoreAccountCreatePermission,
		CoreAccountUpdatePermission,
		CoreAccountDeletePermission,
	)
	// AccountUser provides standard account operations for regular users and services.
	// Intended for: API users and business users managing accounts.
	// cqrsiam.Permissions: Create and update accounts (excludes delete and admin functions).
	AccountUser = cqrsiam.NewRole("AccountUser",
		CoreAccountCreatePermission,
		CoreAccountUpdatePermission,
	)

	// ClosingAdmin provides full administrative access to closing operations.
	// Intended for: Core administrators and the closing controller service.
	// cqrsiam.Permissions: All closing operations including admin-level functions.
	ClosingAdmin = cqrsiam.NewRole("ClosingAdmin",
		CoreClosingAdminPermission,
		CoreClosingCreatePermission,
		CoreClosingUpdatePermission,
		CoreClosingDeletePermission,
	)
	// ClosingUser provides standard closing operations for regular users and services.
	// Intended for: API users and business users managing closings.
	// cqrsiam.Permissions: Create and update closings (excludes delete and admin functions).
	ClosingUser = cqrsiam.NewRole("ClosingUser",
		CoreClosingCreatePermission,
		CoreClosingUpdatePermission,
	)

	// DiscordAdmin provides full administrative access to discord operations.
	// Intended for: Core administrators and the discord controller service.
	// cqrsiam.Permissions: All discord operations including admin-level functions.
	DiscordAdmin = cqrsiam.NewRole("DiscordAdmin",
		CoreDiscordAdminPermission,
		CoreDiscordSendPermission,
	)
	// DiscordUser provides standard discord operations for regular users and services.
	// Intended for: API users and services that send discord messages.
	// cqrsiam.Permissions: Send discord messages (excludes admin functions).
	DiscordUser = cqrsiam.NewRole("DiscordUser",
		CoreDiscordSendPermission,
	)

	// NewsAdmin provides full administrative access to news operations.
	// Intended for: Core administrators and the news controller service.
	// cqrsiam.Permissions: All news operations including admin functions.
	NewsAdmin = cqrsiam.NewRole("NewsAdmin",
		CoreNewsAdminPermission,
		CoreNewsUpdatePermission,
		CoreNewsRatePermission,
	)
	// NewsUser provides standard news operations for regular users and services.
	// Intended for: Services and users that create and modify news entries.
	// cqrsiam.Permissions: Update and rate news entries.
	NewsUser = cqrsiam.NewRole("NewsUser",
		CoreNewsUpdatePermission,
		CoreNewsRatePermission,
	)
	// NewsViewer provides read-only access to news data.
	// Intended for: Users who need to view but not modify news (e.g., monitoring services).
	// cqrsiam.Permissions: None (read-only access through other mechanisms).
	NewsViewer = cqrsiam.NewRole("NewsViewer")

	// SignalAdmin provides full administrative access to signal operations.
	// Intended for: Core administrators and signal services.
	// cqrsiam.Permissions: All signal operations including break-even and admin functions.
	SignalAdmin = cqrsiam.NewRole("SignalAdmin",
		CoreSignalCreatePermission,
		CoreSignalModifyPermission,
		CoreSignalClosePermission,
		CoreSignalBreakEvenPermission,
		CoreSignalAdminPermission,
	)
	// SignalUser provides standard signal operations for traders.
	// Intended for: Traders who can create, modify, and close trading signals.
	// cqrsiam.Permissions: Create, modify, and close signals (excludes break-even and admin).
	SignalUser = cqrsiam.NewRole("SignalUser",
		CoreSignalCreatePermission,
		CoreSignalModifyPermission,
		CoreSignalClosePermission,
	)
	// SignalViewer provides read-only access to signal data.
	// Intended for: Users who need to view but not modify signals (e.g., monitoring services).
	// cqrsiam.Permissions: None (read-only access through other mechanisms).
	SignalViewer = cqrsiam.NewRole("SignalViewer")

	// ActualTradeAdmin provides full administrative access to actual trade operations.
	// Intended for: Core administrators and the actual trade controller service.
	// cqrsiam.Permissions: All actual trade operations including admin functions.
	ActualTradeAdmin = cqrsiam.NewRole("ActualTradeAdmin",
		CoreActualTradeRetryPermission,
		CoreActualTradeBrokerCallbackPermission,
		CoreActualTradeAdminPermission,
	)
	// ActualTradeUser provides standard actual trade operations for traders.
	// Intended for: Traders who can retry failed trades.
	// cqrsiam.Permissions: Retry failed trades.
	ActualTradeUser = cqrsiam.NewRole("ActualTradeUser",
		CoreActualTradeRetryPermission,
	)
	// ActualTradeBrokerCallback provides broker callback permissions.
	// Intended for: Broker converter services that send order/position updates.
	// cqrsiam.Permissions: All broker callback operations.
	ActualTradeBrokerCallback = cqrsiam.NewRole("ActualTradeBrokerCallback",
		CoreActualTradeBrokerCallbackPermission,
	)
	// ActualTradeViewer provides read-only access to actual trade data.
	// Intended for: Users who need to view but not modify actual trades (e.g., monitoring services).
	// cqrsiam.Permissions: None (read-only access through other mechanisms).
	ActualTradeViewer = cqrsiam.NewRole("ActualTradeViewer")

	// Mt5SyncTrigger provides permission to trigger the two MT5 sync operations.
	// Intended for: the API gateway identity (ApiInitiator, shared with every other
	// business-tier automation via FRONTEND_GATEWAY_API_KEY in every environment, not
	// dev-scoped) and the human administrator. This deviates from the Mt5AccountAdmin
	// precedent of keeping admin-tier permissions off ApiInitiator (iam_rolebinding.go);
	// it is accepted here because both operations are idempotent refreshes (re-select
	// symbols, re-read account state) — nothing is created, deleted, or traded, and no
	// money moves. That is a materially smaller blast radius than Mt5AccountAdmin
	// (create/update/delete accounts), which is why the deviation is acceptable rather
	// than merely convenient. scripts/api-call.sh, which authenticates as this identity,
	// must be able to trigger these without a human-authenticated session.
	// cqrsiam.Permissions: MT5 symbol re-selection and MT5 account re-read triggers.
	Mt5SyncTrigger = cqrsiam.NewRole("Mt5SyncTrigger",
		Mt5TickSelectSymbolsPermission,
		Mt5AccountSyncAccountsPermission,
	)

	// SuperAdmin provides full system administrative access including gateway admin routes.
	// Intended for: System administrators who need access to all admin interfaces.
	// cqrsiam.Permissions: Gateway admin access for service management.
	SuperAdmin = cqrsiam.NewRole("SuperAdmin",
		FrontendGatewayAdminPermission,
	)
)

// AvailableRoles contains all registered roles in the IAM system.
var AvailableRoles = cqrsiam.Roles{
	BacktestAdmin,
	BacktestUser,
	BacktestViewer,
	BacktestCreator,
	BacktestCron,
	ReviewAdmin,
	ReviewUser,
	ReviewViewer,
	StrategyAdmin,
	StrategyUser,
	StrategyViewer,
	SymbolAdmin,
	SymbolUser,
	ReportAdmin,
	ReportCron,
	MailRefresh,
	Mt5AccountAdmin,
	Mt5AccountUser,
	Mt5AccountViewer,
	ExpectedTradeAdmin,
	ExpectedTradeUser,
	ExpectedTradeViewer,
	AccountAdmin,
	AccountUser,
	ClosingAdmin,
	ClosingUser,
	DiscordAdmin,
	DiscordUser,
	NewsAdmin,
	NewsUser,
	NewsViewer,
	SignalAdmin,
	SignalUser,
	SignalViewer,
	ActualTradeAdmin,
	ActualTradeUser,
	ActualTradeBrokerCallback,
	ActualTradeViewer,
	Mt5SyncTrigger,
	SuperAdmin,
}
