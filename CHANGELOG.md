
<a name="v1.0.1"></a>
## [v1.0.1](https://github.com/Bad3r/antidot-home/compare/v1.0.0...v1.0.1) (2025-08-03)

### Bug Fixes

* update goreleaser to v2 with --clean flag and explicit version


<a name="v1.0.0"></a>
## v1.0.0 (2025-08-03)

### Bug Fixes

* fix file exists check
* fallback to system default xdg values ([#37](https://github.com/Bad3r/antidot-home/issues/37))
* revert go mod tidy changes and remove toolchain directive
* go-dependency-submission hash ([#13](https://github.com/Bad3r/antidot-home/issues/13))
* add minimal build step for integration tests
* update go.mod to Go 1.18 for dependency compatibility
* handle empty braces {} in fish shell unbracketEnvVar function
* handle errors in various places
* remove eq sign in fish set command
* skip dump when nothing applied
* add missing dollar sign to bash init stub
* write tmp file to same partition in atomic write
* format code and add pre-commit framework
* handle action errors
* print help when no command is passed ([#43](https://github.com/Bad3r/antidot-home/issues/43))
* check file existance before sourcing it ([#88](https://github.com/Bad3r/antidot-home/issues/88))
* remove quotes from env fallback values
* remove init var check ([#70](https://github.com/Bad3r/antidot-home/issues/70))
* skip ignored rules
* go version format in go.mod
* don't override existing xdg values ([#42](https://github.com/Bad3r/antidot-home/issues/42))
* add message about running antidot init ([#87](https://github.com/Bad3r/antidot-home/issues/87))
* use correct data path for alias and env files
* ensure data home exists before rules update
* add proper error when missing rules file
* **config:** override old config
* **rules:** correct Simplescreenrecorder dest

### Code Refactoring

* err on unknown regex group name
* improve error handling
* dedup alias and export code
* minor ux improvements
* move shell logic to package
* warn on failed rules
* change to prompt per action ([#22](https://github.com/Bad3r/antidot-home/issues/22))
* change to a nicer prompt ([#23](https://github.com/Bad3r/antidot-home/issues/23))
* load config concurrently and log fixes
* rename and break action package
* change ignored rules text
* move get shell files to a func
* return config from load function
* use an action registry
* remove unused function
* move get home dir to utils
* **clean:** search for rules in home dir and not the opposite
* **clean:** search for rules in home dir and not the opposite
* **tui:** replace deprecated survey with bubbletea v1.3.6

### Features

* prompt before applying rules
* detect all dotfiles in homedir
* add automatic changelog generation ([#21](https://github.com/Bad3r/antidot-home/issues/21))
* respect NO_COLOR env var
* initial dotfile rules engine
* support fish and zsh
* support printing rules' notes ([#68](https://github.com/Bad3r/antidot-home/issues/68))
* add verbose flag ([#23](https://github.com/Bad3r/antidot-home/issues/23))
* add update rules file command
* init cmd for initializing antidot
* improve GoReleaser configuration ([#19](https://github.com/Bad3r/antidot-home/issues/19))
* update CI to use Go 1.24 matching local development
* rules file path flag
* add alias action
* add export action
* add version flag
* add color output
* delete file action
* rules printing and application
* load actions from rules file
* load rules from config file
* **ci:** Add GH Actions linting ([#15](https://github.com/Bad3r/antidot-home/issues/15))
* **shell:** print supported shells

