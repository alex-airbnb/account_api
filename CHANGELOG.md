# [1.3.0](https://github.com/alex-airbnb/account_api/compare/v1.2.0...v1.3.0) (2021-03-09)


### Bug Fixes

* Pass the article as pointer in the Account REST Use Case. ([ae3b61a](https://github.com/alex-airbnb/account_api/commit/ae3b61a533e0dfb5653a9d5f405753a73746875f))


### Features

* Add and run the REST Handler. ([7286d3d](https://github.com/alex-airbnb/account_api/commit/7286d3d118d4a7a387869b33bca77b88ace7e3b1))
* Add docker-compose-dependencies file. ([6302012](https://github.com/alex-airbnb/account_api/commit/630201279c9219b3291d4eec6373b5cacc949035))


### Performance Improvements

* Refactor SetUpPostgres function to return a db instance. ([dcdc569](https://github.com/alex-airbnb/account_api/commit/dcdc5693dbd1a415d22ef4ffb4250f1a25f7b6bb))
* Update the route for the account handlers. ([db85deb](https://github.com/alex-airbnb/account_api/commit/db85deb0a81b129c89b62384023e280bec66b4b2))

# [1.2.0](https://github.com/alex-airbnb/account_api/compare/v1.1.0...v1.2.0) (2021-03-05)


### Features

* Add NewPostgres to the adapter package. ([e6c5f2d](https://github.com/alex-airbnb/account_api/commit/e6c5f2d83eb6d2ac5db2006ae626420eb5f969f4))
* Add Postgres and SetUpPostgres to client package. ([be00630](https://github.com/alex-airbnb/account_api/commit/be0063095e973f0f05811a48c62670f5d52ac349))

# [1.1.0](https://github.com/alex-airbnb/account_api/compare/v1.0.0...v1.1.0) (2021-03-03)


### Features

* Add Account and CreateAccount to the package entity. ([afe9f5c](https://github.com/alex-airbnb/account_api/commit/afe9f5c4d6b7e8e5fb79def2494dfb058ae4941d))
* Add AccountREST, ErrInvalidJSONFormat, ErrMissingRequiredProperty, CreateAccountRequest, and CreateAccountResponse to the usecase package. ([86b9090](https://github.com/alex-airbnb/account_api/commit/86b909074fa6dc8c2d6f92c505fbb1674e984cbf))
* Add RepositoryPort to the adapter package. ([becacdb](https://github.com/alex-airbnb/account_api/commit/becacdbf6f6578b85061d13856c13eced57226bd))
* Add Setup, response, createAccount, and AccountRouter. ([ee3abeb](https://github.com/alex-airbnb/account_api/commit/ee3abeb82d16868bb10d5c8e77dc71601f03d983))
* Add tests for the POST /account endpoint. ([5d1d6f9](https://github.com/alex-airbnb/account_api/commit/5d1d6f9759c130fb314b2883eed37d8dd082fe5a))
* Add UseCase and NewAccountREST to the usecase package. ([a37d684](https://github.com/alex-airbnb/account_api/commit/a37d6848ec13b8d2a124e7672252115d8d6496b7))
* Rename TestAccount to TestAccountEntity ([7926b5d](https://github.com/alex-airbnb/account_api/commit/7926b5d39f02b045b93c61b7acc9ff7de55ce15f))
* Update codecov badge. ([3ffec65](https://github.com/alex-airbnb/account_api/commit/3ffec655358e0fb5b7284c25191e4f2921fbe059))
* Use gorm.Model struct in the Account entity. ([0e2ff3d](https://github.com/alex-airbnb/account_api/commit/0e2ff3decac93347d58d57e35399da4ba30b79ee))


### Performance Improvements

* Replace Setup by NewRESTHandler function. ([a423a0e](https://github.com/alex-airbnb/account_api/commit/a423a0efd96783dfdfc5696864d2cbc979815521))

# 1.0.0 (2021-02-17)


### Features

* Add CircleCI config. ([b550b4e](https://github.com/alex-airbnb/account_api/commit/b550b4e74fe3d7488bbd622f74afa52b3cc59213))
* Add Dockerfile. ([bd23c31](https://github.com/alex-airbnb/account_api/commit/bd23c31429f1924c295b5faf17bbc18a51257b98))
* Add gitignore file. ([354c084](https://github.com/alex-airbnb/account_api/commit/354c0842ba9ab0fde8e2087621b25ed6ac7353a4))
* Add main package. ([52fab11](https://github.com/alex-airbnb/account_api/commit/52fab1160ff805f77d9bb41f32cb07d68ea5dbd4))
* Add Makefile. ([e57ec22](https://github.com/alex-airbnb/account_api/commit/e57ec226bf15add6c19a51cf38fd65cd264c9c75))
* Add semantic release config. ([faff80b](https://github.com/alex-airbnb/account_api/commit/faff80b1c2e4cf6ab91cb4800a812c5243b67671))
* Initialize github.com/alex-airbnb/account_api project. ([5ae93f7](https://github.com/alex-airbnb/account_api/commit/5ae93f77c01f4557bc0562feb02d7ec0bc5f174d))
* Update README file. ([e43fc3e](https://github.com/alex-airbnb/account_api/commit/e43fc3e40442fad2143808a55ec9375afc8ee66c))
* Update the name of the default branch. ([5758613](https://github.com/alex-airbnb/account_api/commit/5758613947bfd2f6c5d2a147d411f1992138f7e0))
