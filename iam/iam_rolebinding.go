// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iam

import cqrsiam "github.com/bborbe/cqrs/iam"

// cqrsiam.Role bindings define which initiators (users/services) have which roles.
//
// Example: Adding a new service with minimal permissions
//
//	1. Add the service as an initiator constant:
//	   CoreMyNewService cqrsiam.Initiator = "core-my-new-service"
//
//	2. Add to AvailableInitiators slice
//
//	3. Create role binding (use existing role or create new one):
//	   NewRoleBinding(BacktestCreator, CoreMyNewService)
//
//	4. For new permissions, create a new role:
//	   MyServicecqrsiam.Role = cqrsiam.NewRole("MyServicecqrsiam.Role", MyNewPermission)
//
// Follow the principle of least privilege - give services only the minimum
// permissions they need to function.

// AvailableRoleBindings defines the complete access control matrix for the trading system.
// Each binding grants specific roles to users and services.
//
// Current role assignments:
//   - BacktestAdmin: Core administrators and controller service
//   - BacktestCron: Automated strategy queuing services
//   - BacktestUser: API users and standard backtest agent services
//   - BacktestCreator: Services that only create backtests (principle of least privilege)
//   - ReviewAdmin: Core administrators and review service
//   - ReviewUser: Business users for review participation
//   - StrategyAdmin: Core administrators and strategy service
//   - StrategyUser: API users and business users managing strategies
var AvailableRoleBindings = cqrsiam.RoleBindings{
	cqrsiam.NewRoleBinding(BacktestAdmin, UserBenPrivate, CoreBacktestController),
	cqrsiam.NewRoleBinding(BacktestCron, CoreBacktestCron),
	cqrsiam.NewRoleBinding(
		BacktestUser,
		ApiInitiator,
		UserBenBusiness,
		CoreBacktestAgentJob,
		CoreBacktestAgentOptimize,
		CoreBacktestAgentRun,
		AgentBacktest,
	),
	cqrsiam.NewRoleBinding(BacktestCreator, CoreBacktestAgentQueue),
	cqrsiam.NewRoleBinding(BacktestViewer, BossOpenClaw, TradingOpenClaw),
	cqrsiam.NewRoleBinding(ActualTradeAdmin, UserBenPrivate, CoreActualTradeController),
	cqrsiam.NewRoleBinding(ActualTradeBrokerCallback, Mt5TradeExecutor, CapitalcomTradeExecutor),
	cqrsiam.NewRoleBinding(ReviewAdmin, UserBenPrivate, CoreReview),
	cqrsiam.NewRoleBinding(ReviewUser, UserBenBusiness),
	cqrsiam.NewRoleBinding(ReviewViewer, BossOpenClaw, TradingOpenClaw),
	cqrsiam.NewRoleBinding(
		StrategyAdmin,
		UserBenPrivate,
		CoreStrategy,
		CoreStrategyPublishStrategies,
		CoreStrategyWatchStrategies,
	),
	cqrsiam.NewRoleBinding(StrategyUser, ApiInitiator, UserBenBusiness),
	cqrsiam.NewRoleBinding(StrategyViewer, BossOpenClaw, TradingOpenClaw),
	cqrsiam.NewRoleBinding(SymbolAdmin, UserBenPrivate),
	cqrsiam.NewRoleBinding(
		SymbolUser,
		ApiInitiator,
		CapitalcomMarketDetailConverter,
		AcgSymbolConverter,
		DwxSymbolConverter,
		FtmoSymbolConverter,
	),
	cqrsiam.NewRoleBinding(ClosingAdmin, UserBenPrivate),
	cqrsiam.NewRoleBinding(ClosingUser, ApiInitiator),
	cqrsiam.NewRoleBinding(Mt5AccountAdmin, UserBenPrivate, Mt5AccountController),
	cqrsiam.NewRoleBinding(Mt5AccountUser, Mt5AccountConverter, ApiInitiator),
	cqrsiam.NewRoleBinding(Mt5AccountViewer, BossOpenClaw, TradingOpenClaw),
	cqrsiam.NewRoleBinding(ExpectedTradeAdmin, UserBenPrivate),
	cqrsiam.NewRoleBinding(ExpectedTradeUser, UserBenBusiness, ApiInitiator),
	cqrsiam.NewRoleBinding(ExpectedTradeViewer, BossOpenClaw, TradingOpenClaw),
	cqrsiam.NewRoleBinding(SignalAdmin, UserBenPrivate),
	cqrsiam.NewRoleBinding(SignalUser, UserBenBusiness, ApiInitiator),
	cqrsiam.NewRoleBinding(SignalViewer, BossOpenClaw, TradingOpenClaw),
	cqrsiam.NewRoleBinding(ActualTradeUser, UserBenBusiness, ApiInitiator),
	cqrsiam.NewRoleBinding(ActualTradeViewer, BossOpenClaw, TradingOpenClaw),
	cqrsiam.NewRoleBinding(ReportAdmin, UserBenPrivate),
	cqrsiam.NewRoleBinding(ReportCron, CoreReportCronjob),
	cqrsiam.NewRoleBinding(MailRefresh, CoreMailControllerRefreshTokenCronjob, ApiInitiator),
	cqrsiam.NewRoleBinding(AccountAdmin, UserBenPrivate),
	cqrsiam.NewRoleBinding(
		AccountUser,
		ApiInitiator,
		Mt5AccountConverter,
		CapitalcomAccountConverter,
	),
	cqrsiam.NewRoleBinding(DiscordAdmin, UserBenPrivate),
	cqrsiam.NewRoleBinding(DiscordUser, ApiInitiator, CoreNotificationController),
	cqrsiam.NewRoleBinding(NewsAdmin, UserBenPrivate),
	cqrsiam.NewRoleBinding(
		NewsUser,
		UserBenBusiness,
		ApiInitiator,
		FinancialjuiceEventConverter,
		ForexfactoryDetailConverter,
		FtmoEconomicCalendarConverter,
		TagesschauConverter,
	),
	cqrsiam.NewRoleBinding(NewsViewer, BossOpenClaw, TradingOpenClaw),
	cqrsiam.NewRoleBinding(Mt5SyncTrigger, ApiInitiator, UserBenPrivate),
	cqrsiam.NewRoleBinding(SuperAdmin, UserBenPrivate),
}
