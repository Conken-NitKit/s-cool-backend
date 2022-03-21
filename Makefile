# export PATH := bin:${PATH}

# .DEFAULT_GOAL := help
# .PHONY: help
# help: ## コマンド一覧を表示（実行方法: make help ）
# 	@echo ""
# 	@echo "------------------------------------------------------------------"
# 	@echo " [使い方] ルートディレクトリで \033[36mmake [コマンド名]\033[0m で実行可能"
# 	@echo " [利用例] \033[36mmake help\033[0m"
# 	@echo "------------------------------------------------------------------"
# 	@echo ""
# 	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' ${MAKEFILE_LIST} | awk 'BEGIN {FS = ":.*?## "}; { \
# 		printf "- \033[36m%-20s\033[0m %s\n", $$1, $$2 \
# 	}'
# 	@echo ""

# .PHONY: setup
# setup: ## セットアップ
# 	go get github.com/google/wire/cmd/wire

# .PHONY: gen
# gen: ## 各種生成系コマン度を実行
# 	gen-wire

# .PHONY: gen-wire
# gen-wire: ## DI関数を生成
# 	cd api/internal/di && wire

# .PHONY: up
# up: ## サービス用のコンテナの構築、作成、起動、アタッチを行う（バックグラウンド）
# 	docker compose up -d

# .PHONY: up-show-log
# up-show-log: ## サービス用のコンテナの構築、作成、起動、アタッチを行う
# 	docker compose up

# .PHONY: stop
# stop: ## サービス用のコンテナを停止
# 	docker compose stop

# .PHONY: down
# down: ## サービス用のコンテナを停止し、そのコンテナとネットワークを削除
# 	docker compose down