package main

import (
	"CareerCenter/internal/config/database"
	account3 "CareerCenter/internal/delivery/http/account"
	application3 "CareerCenter/internal/delivery/http/application"
	company3 "CareerCenter/internal/delivery/http/company"
	jobs3 "CareerCenter/internal/delivery/http/jobs"
	profile3 "CareerCenter/internal/delivery/http/profile"
	"CareerCenter/internal/repository/account"
	"CareerCenter/internal/repository/application"
	"CareerCenter/internal/repository/company"
	"CareerCenter/internal/repository/jobs"
	"CareerCenter/internal/repository/profile"
	account2 "CareerCenter/internal/usecase/account"
	application2 "CareerCenter/internal/usecase/application"
	company2 "CareerCenter/internal/usecase/company"
	jobs2 "CareerCenter/internal/usecase/jobs"
	profile2 "CareerCenter/internal/usecase/profile"
)

var (
	mysqlConn = database.InitMysqlDB()

	repoAccount    = account.NewAccountMysqlInteractor(mysqlConn)
	useCaseAccount = account2.NewAccountUsecase(repoAccount, repoProfile)
	handlerAccount = account3.NewUseCaseAccountHandler(useCaseAccount)

	repoJobs    = jobs.NewJobsMysqlInteractor(mysqlConn)
	useCaseJobs = jobs2.NewJobsUsecase(repoJobs)
	handlerJobs = jobs3.NewUseCaseJobsHandler(useCaseJobs)

	repoProfile    = profile.NewProfileMysqlInteractor(mysqlConn)
	useCaseProfile = profile2.NewProfileUsecase(repoProfile)
	handlerProfile = profile3.NewUseCaseProfileHandler(useCaseProfile)

	repoApplication    = application.NewApplicationMysqlInteractor(mysqlConn)
	useCaseApplication = application2.NewApplicationUsecase(repoApplication, repoProfile)
	handlerApplication = application3.NewUseCaseApplicationHandler(useCaseApplication)

	repoCompany    = company.NewCompanyMysqlInteractor(mysqlConn)
	useCaseCompany = company2.NewCompanyUsecase(repoCompany, repoJobs)
	handlerCompany = company3.NewUseCaseCompanyHandler(useCaseCompany)
)
