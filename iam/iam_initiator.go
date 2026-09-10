// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package iam

import cqrsiam "github.com/bborbe/cqrs/iam"

// InitiatorPermissions returns all permissions granted to this initiator through role bindings.
func InitiatorPermissions(i cqrsiam.Initiator) cqrsiam.Permissions {
	return AvailableRoleBindings.FindByInitiator(i).Roles().Permissions()
}

var (
	// Service Initiators - Automated services that perform operations

	// ApiInitiator represents API gateway service for external requests.
	ApiInitiator cqrsiam.Initiator = "api"
	// AgentBacktest represents the standalone backtest agent service that processes task files.
	AgentBacktest cqrsiam.Initiator = "agent-backtest"
	// CoreBacktestAgentJob represents the job execution service for backtest processing.
	CoreBacktestAgentJob cqrsiam.Initiator = "core-backtest-agent-job"
	// CoreBacktestAgentOptimize represents the optimization service for parameter tuning.
	CoreBacktestAgentOptimize cqrsiam.Initiator = "core-backtest-agent-optimize"
	// CoreBacktestAgentQueue represents the queue service that triggers backtests automatically.
	CoreBacktestAgentQueue cqrsiam.Initiator = "core-backtest-agent-queue"
	// CoreBacktestAgentRun represents the execution service that runs individual backtests.
	CoreBacktestAgentRun cqrsiam.Initiator = "core-backtest-agent-run"
	// CoreBacktestController represents the main backtest controller service.
	CoreBacktestController cqrsiam.Initiator = "core-backtest-controller"
	// CoreBacktestDummyUser represents a test user for backtest operations.
	CoreBacktestDummyUser cqrsiam.Initiator = "core-backtest-dummyuser"
	// CoreBacktestCron represents the automated strategy queuing service.
	CoreBacktestCron cqrsiam.Initiator = "core-backtest-cron"
	// CoreReview represents the review management service.
	CoreReview cqrsiam.Initiator = "core-review"
	// CoreStrategy represents the strategy management service.
	CoreStrategy cqrsiam.Initiator = "core-strategy"
	// CoreStrategyPublishStrategies represents the strategy publishing command service.
	CoreStrategyPublishStrategies cqrsiam.Initiator = "core-strategy-publish-strategies"
	// CoreStrategyWatchStrategies represents the strategy watching command service.
	CoreStrategyWatchStrategies cqrsiam.Initiator = "core-strategy-watch-strategies"
	// CoreReportCronjob represents the automated report scheduling service.
	CoreReportCronjob cqrsiam.Initiator = "core-report-cron"
	// CoreMailControllerRefreshTokenCronjob represents the daily OAuth2 token refresh CronJob.
	CoreMailControllerRefreshTokenCronjob cqrsiam.Initiator = "core-mail-controller-refresh-token-cronjob" // #nosec G101 -- false positive: this is a service identifier, not a credential
	// Mt5TradeExecutor represents the MT5 broker integration service that handles order/position callbacks.
	Mt5TradeExecutor cqrsiam.Initiator = "mt5-trade-executor"
	// CapitalcomTradeExecutor represents the Capital.com broker integration service that handles order/position callbacks.
	CapitalcomTradeExecutor cqrsiam.Initiator = "capitalcom-trade-executor"

	// CoreNotificationController sends discord messages for alerts and notifications.
	CoreNotificationController cqrsiam.Initiator = "core-notification-controller"
	// CoreActualTradeController represents the actual trade controller service that manages trade execution.
	CoreActualTradeController cqrsiam.Initiator = "core-actualtrade-controller"
	// Mt5AccountConverter represents the MT5 account converter service that processes raw MT5 account data.
	Mt5AccountConverter cqrsiam.Initiator = "mt5-account-converter"
	// Mt5AccountController represents the MT5 account controller service that manages MT5 account lifecycle.
	Mt5AccountController cqrsiam.Initiator = "mt5-account-controller"

	// User cqrsiam.Initiators - Human users with different roles

	// UserBenBusiness represents the business account for trading operations.
	UserBenBusiness cqrsiam.Initiator = "borbe.capital@gmail.com"
	// UserBenPrivate represents the personal administrative account.
	UserBenPrivate cqrsiam.Initiator = "benjamin.borbe@gmail.com"
	// UserBenSeibert represents the corporate development account.
	UserBenSeibert cqrsiam.Initiator = "benjamin.borbe@seibert.group"
	// UserStefan represents an additional authorized user.
	UserStefan cqrsiam.Initiator = "stefan.nicolin@googlemail.com"

	// OpenClaw cqrsiam.Initiators - AI assistant services with different access levels

	// BossOpenClaw represents the Boss AI assistant with readonly monitoring access.
	BossOpenClaw cqrsiam.Initiator = "boss-openclaw"
	// TradingOpenClaw represents the Trading AI assistant with readonly access (may expand later).
	TradingOpenClaw cqrsiam.Initiator = "trading-openclaw"

	// CapitalcomMarketDetailConverter converts market details from CapitalCom into trading symbols.
	CapitalcomMarketDetailConverter cqrsiam.Initiator = "capitalcom-marketdetail-converter"
	// AcgSymbolConverter converts symbols from ACG broker into trading symbols.
	AcgSymbolConverter cqrsiam.Initiator = "acg-symbol-converter"
	// DwxSymbolConverter converts symbols from DWX broker into trading symbols.
	DwxSymbolConverter cqrsiam.Initiator = "dwx-symbol-converter"
	// FtmoSymbolConverter converts symbols from FTMO broker into trading symbols.
	FtmoSymbolConverter cqrsiam.Initiator = "ftmo-symbol-converter"
	// CapitalcomAccountConverter converts account data from CapitalCom broker into trading accounts.
	CapitalcomAccountConverter cqrsiam.Initiator = "capitalcom-account-converter"
	// News Converter cqrsiam.Initiators - Services that create news from external sources

	// ForexfactoryDetailConverter converts forex factory details into news entries.
	ForexfactoryDetailConverter cqrsiam.Initiator = "forexfactory-detail-converter"
	// FtmoEconomicCalendarConverter converts FTMO economic calendar events into news.
	FtmoEconomicCalendarConverter cqrsiam.Initiator = "ftmo-economiccalendar-converter"
	// TagesschauConverter converts Tagesschau news articles into trading news.
	TagesschauConverter cqrsiam.Initiator = "tagesschau-converter"
	// FinancialjuiceEventConverter converts FinancialJuice events into trading news.
	FinancialjuiceEventConverter cqrsiam.Initiator = "financialjuice-event-converter"
)

// AvailableInitiators contains all registered initiators (services and users) in the system.
var AvailableInitiators = cqrsiam.Initiators{
	AgentBacktest,
	ApiInitiator,
	UserBenSeibert,
	CoreBacktestAgentJob,
	CoreBacktestAgentOptimize,
	CoreBacktestAgentQueue,
	CoreBacktestAgentRun,
	CoreBacktestController,
	CoreBacktestDummyUser,
	CoreBacktestCron,
	CoreReview,
	CoreStrategy,
	CoreStrategyPublishStrategies,
	CoreStrategyWatchStrategies,
	CoreReportCronjob,
	CoreMailControllerRefreshTokenCronjob,
	Mt5TradeExecutor,
	CapitalcomTradeExecutor,
	CoreNotificationController,
	CoreActualTradeController,
	Mt5AccountConverter,
	Mt5AccountController,
	UserStefan,
	UserBenBusiness,
	UserBenPrivate,
	BossOpenClaw,
	TradingOpenClaw,
	CapitalcomMarketDetailConverter,
	AcgSymbolConverter,
	DwxSymbolConverter,
	FtmoSymbolConverter,
	Mt5AccountConverter,
	CapitalcomAccountConverter,
	ForexfactoryDetailConverter,
	FtmoEconomicCalendarConverter,
	TagesschauConverter,
	FinancialjuiceEventConverter,
}
